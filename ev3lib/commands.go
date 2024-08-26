package ev3lib

import (
	"fmt"
	"slices"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Command Interface                                                          //
////////////////////////////////////////////////////////////////////////////////

// Command defines the interface for all commands to implement.
type Command interface {
	Init()
	Run()
	End(interrupted bool)
	IsDone() bool
}

////////////////////////////////////////////////////////////////////////////////
// Base Command                                                               //
////////////////////////////////////////////////////////////////////////////////

// CommandBase provides a default implementation for all commands.
type CommandBase struct{}

func (b *CommandBase) Init() {}

func (b *CommandBase) Run() {}

func (b *CommandBase) End(interrupted bool) {}

func (b *CommandBase) IsDone() bool {
	return true
}

////////////////////////////////////////////////////////////////////////////////
// Command Decorators                                                         //
////////////////////////////////////////////////////////////////////////////////

// WithTimeout adds a timeout to a command specified by `dur`.
func WithTimeout(c Command, dur time.Duration) Command {
	return NewParallelRace(c, NewWaitCommand(dur))
}

////////////////////////////////////////////////////////////////////////////////

type untilCommandDecorator struct {
	c Command
	p func() bool

	interrupted bool
}

func (u *untilCommandDecorator) Init() {
	u.interrupted = false
	u.c.Init()
}

func (u *untilCommandDecorator) Run() {
	u.c.Run()
}

func (u *untilCommandDecorator) End(interrupted bool) {
	if u.interrupted {
		u.c.End(true)
		return
	}
	u.c.End(interrupted)
}

func (u *untilCommandDecorator) IsDone() bool {
	if u.p() {
		u.interrupted = true
		return true
	}
	return u.c.IsDone()
}

// Until will run a command but will interrupt if a predicate is satisfied.
func Until(c Command, predicate func() bool) Command {
	return &untilCommandDecorator{c: c, p: predicate}
}

////////////////////////////////////////////////////////////////////////////////
// Blocking Command Runner                                                    //
////////////////////////////////////////////////////////////////////////////////

// RunCommand will run a command in a blocking fashion.
func RunCommand(c Command) {
	c.Init()
	for !c.IsDone() {
		c.Run()
	}
	c.End(false)
}

////////////////////////////////////////////////////////////////////////////////
// Sequence Command                                                           //
////////////////////////////////////////////////////////////////////////////////

type sequence struct {
	CommandBase

	current  int
	commands []Command
}

// NewSequence will run several commands one after another.
func NewSequence(commands ...Command) Command {
	return &sequence{current: 0, commands: commands}
}

func (s *sequence) Init() {
	s.commands[0].Init()
}

func (s *sequence) Run() {
	currCommand := s.commands[s.current]
	currCommand.Run()

	if currCommand.IsDone() {
		currCommand.End(false)

		s.current++
		if !s.IsDone() {
			s.commands[s.current].Init()
		}
	}
}

func (s *sequence) IsDone() bool {
	return s.current >= len(s.commands)
}

////////////////////////////////////////////////////////////////////////////////
// Parallel Command                                                           //
////////////////////////////////////////////////////////////////////////////////

type parallel struct {
	CommandBase

	commands   []Command
	incomplete int
}

// NewParallel will run several commands at the same time waiting for all commands to complete.
func NewParallel(commands ...Command) Command {
	return &parallel{commands: commands, incomplete: len(commands)}
}

func (p *parallel) Init() {
	for _, c := range p.commands {
		c.Init()
	}
}

func (p *parallel) Run() {
	toRemove := make([]Command, 0)
	for _, c := range p.commands {
		c.Run()

		if c.IsDone() {
			c.End(false)
			toRemove = append(toRemove, c)
			p.incomplete--
		}
	}

	if len(toRemove) > 0 {
		p.commands = slices.DeleteFunc(p.commands, func(c Command) bool { return slices.Contains(toRemove, c) })
	}
}

func (p *parallel) IsDone() bool {
	return p.incomplete == 0
}

////////////////////////////////////////////////////////////////////////////////
// Parallel Race Command                                                      //
////////////////////////////////////////////////////////////////////////////////

type parallelRace struct {
	CommandBase

	commands []Command
	finished Command
	done     bool
}

// NewParallelRace will run several commands at the same time.
// It will finish when the first command finishes and will interrupt the rest.
func NewParallelRace(commands ...Command) Command {
	return &parallelRace{commands: commands}
}

func (p *parallelRace) Init() {
	p.done = false
	for _, c := range p.commands {
		c.Init()
	}
}

func (p *parallelRace) Run() {
	for _, c := range p.commands {
		c.Run()

		if c.IsDone() {
			p.finished = c
			p.done = true
		}
	}
}

func (p *parallelRace) End(interrupted bool) {
	for _, c := range p.commands {
		if c != p.finished {
			c.End(true)
		} else {
			c.End(false)
		}
	}
}

func (p *parallelRace) IsDone() bool {
	return p.done
}

////////////////////////////////////////////////////////////////////////////////
// Utility Commands                                                           //
////////////////////////////////////////////////////////////////////////////////

type funcCommand struct {
	CommandBase

	f func()
}

// NewFuncCommand runs a function once.
func NewFuncCommand(f func()) Command {
	return &funcCommand{f: f}
}

func (f *funcCommand) Init() {
	f.f()
}

func (f *funcCommand) IsDone() bool {
	return true
}

////////////////////////////////////////////////////////////////////////////////

type waitCommand struct {
	CommandBase

	init time.Time
	dur  time.Duration
}

// NewWaitCommand waits for a time duration.
func NewWaitCommand(time time.Duration) Command {
	return &waitCommand{dur: time}
}

func (w *waitCommand) Init() {
	w.init = time.Now()
}

func (w *waitCommand) IsDone() bool {
	return time.Since(w.init) > w.dur
}

////////////////////////////////////////////////////////////////////////////////

type printCommand struct {
	CommandBase

	text string
}

// NewPrintCommand prints out some text.
func NewPrintCommand(text string) Command {
	return &printCommand{text: text}
}

func NewPrintlnCommand(text string) Command {
	return &printCommand{text: text + "\n"}
}

func (p *printCommand) Init() {
	fmt.Print(p.text)
}

func (p *printCommand) IsDone() bool {
	return true
}
