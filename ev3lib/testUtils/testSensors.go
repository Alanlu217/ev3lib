package testUtils

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"log"
)

////////////////////////////////////////////////////////////////////////////////
// Test Color Sensor                                                          //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.ColorSensorInterface = &testColorSensor{}

type testColorSensor struct{}

func NewTestColorSensor() *ev3lib.ColorSensor {
	return ev3lib.NewColorSensorBase(&testColorSensor{})
}

func (s *testColorSensor) Ambient() float64 {
	return 0
}

func (s *testColorSensor) Reflection() float64 {
	return 0
}

func (s *testColorSensor) GetRGB() (float64, float64, float64) {
	return 0, 0, 0
}

////////////////////////////////////////////////////////////////////////////////
// Test Gyro Sensor                                                           //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.GyroSensorInterface = &testGyroSensor{}

type testGyroSensor struct{}

func NewTestGyroSensor() *ev3lib.GyroSensor {
	return ev3lib.NewGyroSensorBase(&testGyroSensor{})
}

func (s *testGyroSensor) Rate() float64 {
	return 0
}

func (s *testGyroSensor) Angle() float64 {
	return 0
}

func (s *testGyroSensor) AngleRate() (float64, float64) {
	return 0, 0
}

func (s *testGyroSensor) ResetAngle(angle float64) {
	log.Printf("reset angle to %v\n", angle)
}

func (s *testGyroSensor) Calibrate() {}

////////////////////////////////////////////////////////////////////////////////
// Test Infrared Sensor                                                       //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.InfraredSensorInterface = &testInfraredSensor{}

type testInfraredSensor struct{}

func NewTestInfraredSensor() *ev3lib.InfraredSensor {
	return ev3lib.NewInfraredSensorBase(&testInfraredSensor{})
}

func (s *testInfraredSensor) Distance() float64 {
	return 0
}

func (s *testInfraredSensor) Buttons(_ int) []ev3lib.BeaconButton {
	return []ev3lib.BeaconButton{}
}

////////////////////////////////////////////////////////////////////////////////
// Test Touch Sensor                                                          //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.TouchSensorInterface = &testTouchSensor{}

type testTouchSensor struct{}

func NewTestTouchSensor() *ev3lib.TouchSensor {
	return ev3lib.NewTouchSensorBase(&testTouchSensor{})
}

func (s *testTouchSensor) IsPressed() bool {
	return false
}

////////////////////////////////////////////////////////////////////////////////
// Test Ultrasonic Sensor                                                     //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.UltrasonicSensorInterface = &testUltrasonicSensor{}

type testUltrasonicSensor struct{}

func NewTestUltrasonicSensor() *ev3lib.UltrasonicSensor {
	return ev3lib.NewUltrasonicSensorBase(&testUltrasonicSensor{})
}

func (s *testUltrasonicSensor) Distance() float64 {
	return 0
}

func (s *testUltrasonicSensor) DistanceSilent() float64 {
	return 0
}

func (s *testUltrasonicSensor) Presence() bool {
	return false
}
