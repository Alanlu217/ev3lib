package commands

import (
	"fmt"
	"github.com/Alanlu217/ev3lib/ev3lib"
)

type counterCommand struct {
	ev3lib.DefaultCommand

	target, current int
}

func NewCounterCommand(target int) *ev3lib.CommandBase {
	return ev3lib.NewCommand(&counterCommand{target: target})
}

func (c *counterCommand) Init() {
	c.current = 0
}

func (c *counterCommand) Run() {
	c.current++
	fmt.Println(c.current)
}

func (c *counterCommand) End(interrupted bool) {
	if interrupted {
		fmt.Println("Interrupted at", c.current, "/", c.target)
	}
	fmt.Println("Counted to", c.target)
}

func (c *counterCommand) IsDone() bool {
	return c.current >= c.target
}
