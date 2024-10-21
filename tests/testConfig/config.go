package bopConfig

import (
	"fmt"
	"time"

	"github.com/Alanlu217/ev3lib/tests/commands"

	"github.com/Alanlu217/ev3lib/ev3lib"
)

type Config struct {
	Ev3 *ev3lib.EV3Brick

	Gyro *ev3lib.GyroSensor

	LeftColor, CentreColor, RightColor *ev3lib.ColorSensor

	LeftDrive, RightDrive *ev3lib.Motor
}

func (b *Config) Run1() *ev3lib.Command {
	return ev3lib.NewSequence(
		ev3lib.NewFuncCommand(func() { fmt.Printf("b.gyro.Angle(): %v\n", b.Gyro.Angle()) }),
		ev3lib.NewWaitCommand(10*time.Second).WithTimeout(1*time.Second),
		commands.NewCounterCommand(10).Repeatedly().WithTimeout(1*time.Second),
		// b.LeftDrive.SetCommand(1).WithTimeout(5*time.Second),
		ev3lib.NewPrintlnCommand("Finished Run1").WithTimeout(5*time.Second),
	)
}

func (b *Config) GetCommandPages() *ev3lib.Menu {
	m := ev3lib.NewCommandMenu()

	m.AddPage("runs").
		AddCommand("test", ev3lib.NewPrintlnCommand("Testing")).
		AddCommand("count", ev3lib.NewSequence(
			ev3lib.NewPrintlnCommand("Starting"),
			commands.NewCounterCommand(100).While(commands.NewCounterCommand(50)),
		)).
		AddCommand("run2", b.Run1()).
		AddCommand("run3", b.Run1()).
		AddCommand("run4", b.Run1()).
		AddCommand("run5", b.Run1()).
		AddCommand("run6", b.Run1()).
		AddCommand("run7", b.Run1()).
		AddCommand("motor", ev3lib.NewSequence(
			ev3lib.NewPrintlnCommand("Started Motor"),
			b.LeftDrive.SetCommand(1).WithTimeout(5*time.Second),
			ev3lib.NewPrintlnCommand("Finished Motor"),
		)).
		Add()

	m.AddPage("util").AddCommand("gyroAng",
		ev3lib.NewFuncCommand(func() { fmt.Printf("b.gyro.Angle(): %v\n", b.Gyro.Angle()) })).Add()

	return m
}
