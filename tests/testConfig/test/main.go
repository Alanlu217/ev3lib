package main

import (
	"time"

	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/Alanlu217/ev3lib/ev3lib/testUtils"
	"github.com/Alanlu217/ev3lib/tests/bopConfig"
)

func main() {
	config := &bopConfig.BopConfig{}

	config.Ev3 = testUtils.NewTestEV3Brick()

	config.Gyro = testUtils.NewTestGyroSensor()

	config.LeftColor = testUtils.NewTestColorSensor()
	config.CentreColor = testUtils.NewTestColorSensor()
	config.RightColor = testUtils.NewTestColorSensor()

	config.LeftDrive = testUtils.NewTestMotor("Left Drive")
	config.RightDrive = testUtils.NewTestMotor("Right Drive Motor")

	ev3lib.RunTimedCommand(config.GetCommandPages().Pages[0].Commands[8], 20*time.Millisecond)
}
