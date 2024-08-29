//go:build !ev3test

package bopConfig

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/Alanlu217/ev3lib/ev3lib/ev3"
	"log"
)

func NewBopConfig() *BopConfig {
	b := &BopConfig{}

	var err error

	b.gyro, err = ev3.NewGyroSensor(ev3lib.IN4, false)
	if err != nil {
		log.Fatal(err)
	}

	b.leftColor, err = ev3.NewColorSensor(ev3lib.IN1)
	if err != nil {
		log.Fatal(err)
	}

	b.centreColor, err = ev3.NewColorSensor(ev3lib.IN1)
	if err != nil {
		log.Fatal(err)
	}

	b.rightColor, err = ev3.NewColorSensor(ev3lib.IN1)
	if err != nil {
		log.Fatal(err)
	}

	b.leftDrive, err = ev3.NewLargeMotor(ev3lib.OUTA)
	if err != nil {
		log.Fatal(err)
	}

	b.rightDrive, err = ev3.NewLargeMotor(ev3lib.OUTC)
	if err != nil {
		log.Fatal(err)
	}

	return b
}
