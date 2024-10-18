//go:build !ev3test

package main

import (
	"log"

	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/Alanlu217/ev3lib/ev3lib/ev3"
	"github.com/Alanlu217/ev3lib/tests/bopConfig"
)

func main() {
	config := &bopConfig.BopConfig{}

	var err error

	config.Ev3 = ev3.NewEV3()

	config.Gyro, err = ev3.NewGyroSensor(ev3lib.IN4, false)
	if err != nil {
		log.Fatal(err)
	}

	config.LeftColor, err = ev3.NewColorSensor(ev3lib.IN1)
	if err != nil {
		log.Fatal(err)
	}

	config.CentreColor, err = ev3.NewColorSensor(ev3lib.IN1)
	if err != nil {
		log.Fatal(err)
	}

	config.RightColor, err = ev3.NewColorSensor(ev3lib.IN1)
	if err != nil {
		log.Fatal(err)
	}

	config.LeftDrive, err = ev3.NewLargeMotor(ev3lib.OUTA)
	if err != nil {
		log.Fatal(err)
	}

	config.RightDrive, err = ev3.NewLargeMotor(ev3lib.OUTC)
	if err != nil {
		log.Fatal(err)
	}

	// ev3lib.RunTimedCommand(config.GetCommandPages().Pages[0].Commands[1], 20*time.Millisecond)

	menu := ev3lib.NewMainMenu(ev3.NewEV3MainMenu(config.Ev3), config.GetCommandPages())

	menu.Start()
}
