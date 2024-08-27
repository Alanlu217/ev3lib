//go:build ev3test

package bopConfig

import "github.com/Alanlu217/ev3lib/ev3lib/testUtils"

func NewBopConfig() *BopConfig {
	b := &BopConfig{}

	b.gyro = testUtils.NewTestGyroSensor()

	b.leftColor = testUtils.NewTestColorSensor()
	b.centreColor = testUtils.NewTestColorSensor()
	b.rightColor = testUtils.NewTestColorSensor()

	return b
}
