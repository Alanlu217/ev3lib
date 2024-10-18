package bopConfig

import (
	"fmt"
	"time"

	"github.com/Alanlu217/ev3lib/tests/commands"

	"github.com/Alanlu217/ev3lib/ev3lib"
)

type BopConfig struct {
	Ev3 *ev3lib.EV3Brick

	Gyro *ev3lib.GyroSensor

	LeftColor, CentreColor, RightColor *ev3lib.ColorSensor

	LeftDrive, RightDrive *ev3lib.Motor
}

func (b *BopConfig) Run1() *ev3lib.Command {
	return ev3lib.NewSequence(
		ev3lib.NewFuncCommand(func() { fmt.Printf("b.gyro.Angle(): %v\n", b.Gyro.Angle()) }),
		ev3lib.NewWaitCommand(10*time.Second).WithTimeout(1*time.Second),
		commands.NewCounterCommand(10).Repeatedly().WithTimeout(1*time.Second),
		// b.LeftDrive.SetCommand(1).WithTimeout(5*time.Second),
		ev3lib.NewPrintlnCommand("Finished Run1").WithTimeout(5*time.Second),
	)
}

func (b *BopConfig) GetCommandPages() *ev3lib.Menu {
	m := ev3lib.NewCommandMenu()

	m.AddPage("runs").
		AddCommand("test", ev3lib.NewPrintlnCommand("Testing")).
		AddCommand("run1", b.Run1()).
		AddCommand("motor", ev3lib.NewSequence(
			ev3lib.NewPrintlnCommand("Started Motor"),
			b.LeftDrive.SetCommand(1).WithTimeout(5*time.Second),
			ev3lib.NewPrintlnCommand("Finished Motor"),
		)).
		Add()

	return m
}
