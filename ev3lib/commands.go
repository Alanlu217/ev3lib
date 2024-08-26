package ev3lib

import (
	"fmt"
	"slices"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Command Interface                                                          //
////////////////////////////////////////////////////////////////////////////////

// defines the interface for all commands to implement.
type Command interface {
	Init()
	Run()
	End(interrupted bool)
	IsDone() bool
}

////////////////////////////////////////////////////////////////////////////////
// Base Command                                                               //
////////////////////////////////////////////////////////////////////////////////

// provides a default implementation for all commands.
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

// causes the command to be interrupted if a duration of time has passed.
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

// causes command to be interrupted if a predicate returns true.
func Until(c Command, predicate func() bool) Command {
	return &untilCommandDecorator{c: c, p: predicate}
}

////////////////////////////////////////////////////////////////////////////////
// Blocking Command Runner                                                    //
////////////////////////////////////////////////////////////////////////////////

// runs a command in a blocking manner. Used for debugging.
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

// creates a sequence of commands that execute one after another.
func NewSequence(commands ...Command) *sequence {
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

	commands    []Command
	incompleted int
}

// creates a parallel command group that executes all commands at the same time. It will finish when all internal commands finish.
func NewParallel(commands ...Command) *parallel {
	return &parallel{commands: commands, incompleted: len(commands)}
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
			p.incompleted--
		}
	}

	if len(toRemove) > 0 {
		p.commands = slices.DeleteFunc(p.commands, func(c Command) bool { return slices.Contains(toRemove, c) })
	}
}

func (p *parallel) IsDone() bool {
	return p.incompleted == 0
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

// creates a parallel race group that runs all commands at the same time. Finishes when the first command finishes and interuupts the rest.
func NewParallelRace(commands ...Command) *parallelRace {
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

// runs a function once.
func NewFuncCommand(f func()) *funcCommand {
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

// waits for a time duration.
func NewWaitCommand(time time.Duration) *waitCommand {
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

// prints out some text.
func NewPrintCommand(text string) *printCommand {
	return &printCommand{text: text}
}

func NewPrintlnCommand(text string) *printCommand {
	return &printCommand{text: text + "\n"}
}

func (p *printCommand) Init() {
	fmt.Print(p.text)
}

func (p *printCommand) IsDone() bool {
	return true
}
