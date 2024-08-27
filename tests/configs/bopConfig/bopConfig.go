package bopConfig

import (
	"fmt"
	"github.com/Alanlu217/ev3lib/tests/commands"
	"time"

	"github.com/Alanlu217/ev3lib/ev3lib"
)

type BopConfig struct {
	gyro ev3lib.GyroSensor

	leftColor, centreColor, rightColor ev3lib.ColorSensor
}

func (b *BopConfig) Run1() ev3lib.Command {
	return ev3lib.NewSequence(
		ev3lib.NewFuncCommand(func() { fmt.Printf("b.gyro.Angle(): %v\n", b.gyro.Angle()) }),
		ev3lib.NewWaitCommand(10*time.Second).WithTimeout(1*time.Second),
		commands.NewCounterCommand(10).Repeatedly().WithTimeout(1*time.Second),
		ev3lib.NewPrintlnCommand("Finished Run1"),
	)
}

func (b *BopConfig) GetCommandPages() ev3lib.CommandMenu {
	m := ev3lib.NewCommandMenu()

	m.AddPage("runs").
		AddCommand("test", ev3lib.NewPrintlnCommand("Testing")).
		AddCommand("run1", b.Run1()).
		Add()

	return *m
}
