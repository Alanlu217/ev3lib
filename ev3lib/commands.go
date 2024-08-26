package ev3lib

import (
	"fmt"
	"slices"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Command Interface                                                          //
////////////////////////////////////////////////////////////////////////////////

type Command interface {
	Init()
	Run()
	End(interrupted bool)
	IsDone() bool
}

////////////////////////////////////////////////////////////////////////////////
// Command Decorators                                                         //
////////////////////////////////////////////////////////////////////////////////

func WithTimeout(c Command, dur time.Duration) Command {
	return NewParallelRace(c, NewWaitCommand(dur))
}

////////////////////////////////////////////////////////////////////////////////
// Blocking Command Runner                                                    //
////////////////////////////////////////////////////////////////////////////////

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
	current  int
	commands []Command
}

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

func (s *sequence) End(interrupted bool) {}

func (s *sequence) IsDone() bool {
	return s.current >= len(s.commands)
}

////////////////////////////////////////////////////////////////////////////////
// Parallel Command                                                           //
////////////////////////////////////////////////////////////////////////////////

type parallel struct {
	commands []Command

	incompleted int
}

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

func (p *parallel) End(interrupted bool) {}

func (p *parallel) IsDone() bool {
	return p.incompleted == 0
}

////////////////////////////////////////////////////////////////////////////////
// Parallel Race Command                                                      //
////////////////////////////////////////////////////////////////////////////////

type parallelRace struct {
	commands []Command
	finished Command
	done     bool
}

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
	f func()
}

// runs a function once
func NewFuncCommand(f func()) *funcCommand {
	return &funcCommand{f: f}
}

func (f *funcCommand) Init() {
	f.f()
}

func (f *funcCommand) Run() {}

func (f *funcCommand) End(interrupted bool) {}

func (f *funcCommand) IsDone() bool {
	return true
}

////////////////////////////////////////////////////////////////////////////////

type waitCommand struct {
	init time.Time

	dur time.Duration
}

// waits for a time duration
func NewWaitCommand(time time.Duration) *waitCommand {
	return &waitCommand{dur: time}
}

func (w *waitCommand) Init() {
	w.init = time.Now()
}

func (w *waitCommand) Run() {}

func (w *waitCommand) End(interrupted bool) {}

func (w *waitCommand) IsDone() bool {
	return time.Since(w.init) > w.dur
}

////////////////////////////////////////////////////////////////////////////////

type printCommand struct {
	text string
}

// prints out some text
func NewPrintCommand(text string) *printCommand {
	return &printCommand{text: text}
}

func NewPrintlnCommand(text string) *printCommand {
	return &printCommand{text: text + "\n"}
}

func (p *printCommand) Init() {
	fmt.Print(p.text)
}

func (p *printCommand) Run() {}

func (p *printCommand) End(interrupted bool) {}

func (p *printCommand) IsDone() bool {
	return true
}
