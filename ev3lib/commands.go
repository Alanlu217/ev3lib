package ev3lib

import (
	"fmt"
	"log"
	"slices"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// CommandInterface Interface                                                          //
////////////////////////////////////////////////////////////////////////////////

// CommandInterface defines the interface for all commands to implement.
type CommandInterface interface {
	Init()
	Run()
	End(interrupted bool)
	IsDone() bool
}

////////////////////////////////////////////////////////////////////////////////
// Default CommandInterface                                                            //
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
// Base CommandInterface                                                               //
////////////////////////////////////////////////////////////////////////////////

// Command provides decorator functions on commands.
type Command struct {
	CommandInterface
}

func NewCommand(c CommandInterface) *Command {
	return &Command{c}
}

////////////////////////////////////////////////////////////////////////////////
// CommandInterface Decorators                                                         //
////////////////////////////////////////////////////////////////////////////////

// WithTimeout adds a timeout to a command specified by `dur`.
func (c *Command) WithTimeout(dur time.Duration) *Command {
	return NewCommand(NewParallelRace(c.CommandInterface, NewWaitCommand(dur)))
}

////////////////////////////////////////////////////////////////////////////////

type untilCommandDecorator struct {
	c CommandInterface
	p func() bool

	interrupted bool
}

// Until will run a command but will interrupt if a predicate is satisfied.
func (c *Command) Until(predicate func() bool) *Command {
	return NewCommand(&untilCommandDecorator{c: c.CommandInterface, p: predicate})
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

// OnlyIf will run a command only if a predicate returns true.
func (c *Command) OnlyIf(pred func() bool) *Command {
	return NewIfCommand(pred, c.CommandInterface, NewFuncCommand(func() {}))
}

////////////////////////////////////////////////////////////////////////////////

func (c *Command) Then(cc ...CommandInterface) *Command {
	return NewSequence(slices.Insert(cc, 0, c.CommandInterface)...)
}

////////////////////////////////////////////////////////////////////////////////

func (c *Command) While(cc ...CommandInterface) *Command {
	return NewParallel(append(cc, c.CommandInterface)...)
}

////////////////////////////////////////////////////////////////////////////////

func (c *Command) RaceWith(cc ...CommandInterface) *Command {
	return NewParallelRace(append(cc, c.CommandInterface)...)
}

////////////////////////////////////////////////////////////////////////////////

type repeatCommandDecorator struct {
	c CommandInterface
}

// Repeatedly will run a command forever reinitialising it if it ends.
func (c *Command) Repeatedly() *Command {
	return NewCommand(&repeatCommandDecorator{c: c.CommandInterface})
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

type whenDoneCommandDecorator struct {
	c CommandInterface

	f func(bool)
}

// WhenDone will run a command then a function when it's done.
func (c *Command) WhenDone(f func(bool)) *Command {
	return NewCommand(&whenDoneCommandDecorator{c: c.CommandInterface, f: f})
}

func (r *whenDoneCommandDecorator) Init() {
	r.c.Init()
}

func (r *whenDoneCommandDecorator) Run() {
	r.c.Run()
}

func (r *whenDoneCommandDecorator) End(interrupted bool) {
	r.c.End(interrupted)
	r.f(interrupted)
}

func (r *whenDoneCommandDecorator) IsDone() bool {
	return r.c.IsDone()
}

////////////////////////////////////////////////////////////////////////////////
// Blocking CommandInterface Runner                                                    //
////////////////////////////////////////////////////////////////////////////////

// RunCommand will run a command in a blocking fashion.
func RunCommand(c CommandInterface) {
	c.Init()
	for !c.IsDone() {
		c.Run()
	}
	c.End(false)
}

// RunTimedCommand will run a command in a blocking fashion with a target interval time.
func RunTimedCommand(c CommandInterface, intervalTime time.Duration) {
	t := time.NewTicker(intervalTime)

	c.Init()
	for !c.IsDone() {
		start := time.Now()

		c.Run()

		delta := time.Since(start)

		if delta > intervalTime {
			log.Printf("Loop time overrun, took: %v\n", delta)
		}

		<-t.C
	}
	c.End(false)

	t.Stop()
}

////////////////////////////////////////////////////////////////////////////////
// Sequence CommandInterface                                                           //
////////////////////////////////////////////////////////////////////////////////

type sequence struct {
	DefaultCommand

	current  int
	commands []CommandInterface
}

// NewSequence will run several commands one after another.
func NewSequence(commands ...CommandInterface) *Command {
	return NewCommand(&sequence{current: 0, commands: commands})
}

func (s *sequence) Init() {
	s.current = 0
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
// Parallel CommandInterface                                                           //
////////////////////////////////////////////////////////////////////////////////

type parallel struct {
	DefaultCommand

	commands   []CommandInterface
	incomplete int
}

// NewParallel will run several commands at the same time waiting for all commands to complete.
func NewParallel(commands ...CommandInterface) *Command {
	return NewCommand(&parallel{commands: commands, incomplete: len(commands)})
}

func (p *parallel) Init() {
	p.incomplete = len(p.commands)
	for _, c := range p.commands {
		c.Init()
	}
}

func (p *parallel) Run() {
	toRemove := make([]CommandInterface, 0)
	for _, c := range p.commands {
		c.Run()

		if c.IsDone() {
			c.End(false)
			toRemove = append(toRemove, c)
			p.incomplete--
		}
	}

	if len(toRemove) > 0 {
		p.commands = slices.DeleteFunc(p.commands, func(c CommandInterface) bool { return slices.Contains(toRemove, c) })
	}
}

func (p *parallel) IsDone() bool {
	return p.incomplete == 0
}

////////////////////////////////////////////////////////////////////////////////
// Parallel Race CommandInterface                                                      //
////////////////////////////////////////////////////////////////////////////////

type parallelRace struct {
	DefaultCommand

	commands []CommandInterface
	finished CommandInterface
	done     bool
}

// NewParallelRace will run several commands at the same time.
// It will finish when the first command finishes and will interrupt the rest.
func NewParallelRace(commands ...CommandInterface) *Command {
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
func NewFuncCommand(f func()) *Command {
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
func NewWaitCommand(time time.Duration) *Command {
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
func NewPrintCommand(text string) *Command {
	return NewCommand(&printCommand{text: text})
}

func NewPrintlnCommand(text string) *Command {
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
	a, b CommandInterface
}

// NewIfCommand returns a command that will run a command depending on a predicate.
// If the predicate returns true when the command is initialised, then command a will be run.
func NewIfCommand(runA func() bool, a, b CommandInterface) *Command {
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
