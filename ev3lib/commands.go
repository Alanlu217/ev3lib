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
// Default Command                                                            //
////////////////////////////////////////////////////////////////////////////////

// DefaultCommand provides a default implementation for all commands.
type DefaultCommand struct{}

func (b *DefaultCommand) Init() {}

func (b *DefaultCommand) Run() {}

func (b *DefaultCommand) End(_ bool) {}

func (b *DefaultCommand) IsDone() bool {
	return true
}

////////////////////////////////////////////////////////////////////////////////
// Base Command                                                               //
////////////////////////////////////////////////////////////////////////////////

// CommandBase provides decorator functions on commands.
type CommandBase struct {
	c Command
}

func NewCommand(c Command) *CommandBase {
	return &CommandBase{c}
}

func (c *CommandBase) Init() {
	c.c.Init()
}

func (c *CommandBase) Run() {
	c.c.Run()
}

func (c *CommandBase) End(interrupted bool) {
	c.c.End(interrupted)
}

func (c *CommandBase) IsDone() bool {
	return c.c.IsDone()
}

////////////////////////////////////////////////////////////////////////////////
// Command Decorators                                                         //
////////////////////////////////////////////////////////////////////////////////

// WithTimeout adds a timeout to a command specified by `dur`.
func (c *CommandBase) WithTimeout(dur time.Duration) *CommandBase {
	return NewCommand(NewParallelRace(c.c, NewWaitCommand(dur)))
}

////////////////////////////////////////////////////////////////////////////////

type untilCommandDecorator struct {
	c Command
	p func() bool

	interrupted bool
}

// Until will run a command but will interrupt if a predicate is satisfied.
func (c *CommandBase) Until(predicate func() bool) *CommandBase {
	return NewCommand(&untilCommandDecorator{c: c.c, p: predicate})
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

////////////////////////////////////////////////////////////////////////////////

type repeatCommandDecorator struct {
	c Command
}

// Repeatedly will run a command forever reinitialising it if it ends.
func (c *CommandBase) Repeatedly() *CommandBase {
	return NewCommand(&repeatCommandDecorator{c: c.c})
}

func (r *repeatCommandDecorator) Init() {
	r.c.Init()
}

func (r *repeatCommandDecorator) Run() {
	r.c.Run()
}

func (r *repeatCommandDecorator) End(interrupted bool) {
	r.c.End(interrupted)
}

func (r *repeatCommandDecorator) IsDone() bool {
	if r.c.IsDone() {
		r.c.End(false)
		r.c.Init()
	}
	return false
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

// RunTimedCommand will run a command in a blocking fashion with a target interval time.
func RunTimedCommand(c Command, intervalTime time.Duration) {
	t := time.NewTicker(intervalTime)

	c.Init()
	for !c.IsDone() {
		c.Run()
		<-t.C
	}
	c.End(false)

	t.Stop()
}

////////////////////////////////////////////////////////////////////////////////
// Sequence Command                                                           //
////////////////////////////////////////////////////////////////////////////////

type sequence struct {
	DefaultCommand

	current  int
	commands []Command
}

// NewSequence will run several commands one after another.
func NewSequence(commands ...Command) *CommandBase {
	return NewCommand(&sequence{current: 0, commands: commands})
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
	DefaultCommand

	commands   []Command
	incomplete int
}

// NewParallel will run several commands at the same time waiting for all commands to complete.
func NewParallel(commands ...Command) *CommandBase {
	return NewCommand(&parallel{commands: commands, incomplete: len(commands)})
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
	DefaultCommand

	commands []Command
	finished Command
	done     bool
}

// NewParallelRace will run several commands at the same time.
// It will finish when the first command finishes and will interrupt the rest.
func NewParallelRace(commands ...Command) *CommandBase {
	return NewCommand(&parallelRace{commands: commands})
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

func (p *parallelRace) End(_ bool) {
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
	DefaultCommand

	f func()
}

// NewFuncCommand runs a function once.
func NewFuncCommand(f func()) *CommandBase {
	return NewCommand(&funcCommand{f: f})
}

func (f *funcCommand) Init() {
	f.f()
}

func (f *funcCommand) IsDone() bool {
	return true
}

////////////////////////////////////////////////////////////////////////////////

type waitCommand struct {
	DefaultCommand

	init time.Time
	dur  time.Duration
}

// NewWaitCommand waits for a time duration.
func NewWaitCommand(time time.Duration) *CommandBase {
	return NewCommand(&waitCommand{dur: time})
}

func (w *waitCommand) Init() {
	w.init = time.Now()
}

func (w *waitCommand) IsDone() bool {
	return time.Since(w.init) > w.dur
}

////////////////////////////////////////////////////////////////////////////////

type printCommand struct {
	DefaultCommand

	text string
}

// NewPrintCommand prints out some text.
func NewPrintCommand(text string) *CommandBase {
	return NewCommand(&printCommand{text: text})
}

func NewPrintlnCommand(text string) *CommandBase {
	return NewCommand(&printCommand{text: text + "\n"})
}

func (p *printCommand) Init() {
	fmt.Print(p.text)
}

func (p *printCommand) IsDone() bool {
	return true
}

////////////////////////////////////////////////////////////////////////////////

type ifCommand struct {
	DefaultCommand

	isA  bool
	pred func() bool
	a, b Command
}

// NewIfCommand returns a command that will run a command depending on a predicate.
// If the predicate returns true when the command is initialised, then command a will be run.
func NewIfCommand(runA func() bool, a, b Command) *CommandBase {
	return NewCommand(&ifCommand{isA: true, pred: runA, a: a, b: b})
}

func (f *ifCommand) Init() {
	f.isA = f.pred()

	if f.isA {
		f.a.Init()
	} else {
		f.b.Init()
	}
}

func (f *ifCommand) Run() {
	if f.isA {
		f.a.Run()
	} else {
		f.b.Run()
	}
}

func (f *ifCommand) End(interrupted bool) {
	if f.isA {
		f.a.End(interrupted)
	} else {
		f.b.End(interrupted)
	}
}

func (f *ifCommand) IsDone() bool {
	if f.isA {
		return f.a.IsDone()
	} else {
		return f.b.IsDone()
	}
}
