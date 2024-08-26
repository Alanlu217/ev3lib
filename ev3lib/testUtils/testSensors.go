package testUtils

import "github.com/Alanlu217/ev3lib/ev3lib"

////////////////////////////////////////////////////////////////////////////////
// Test Color Sensor                                                          //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.ColorSensor = &testColorSensor{}

type testColorSensor struct{}

func NewTestColorSensor() ev3lib.ColorSensor {
	return &testColorSensor{}
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

var _ ev3lib.GyroSensor = &testGyroSensor{}

type testGyroSensor struct{}

func NewTestGyroSensor() ev3lib.GyroSensor {
	return &testGyroSensor{}
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

func (s *testGyroSensor) ResetAngle(angle float64) {}

func (s *testGyroSensor) Calibrate() {}

////////////////////////////////////////////////////////////////////////////////
// Test Infrared Sensor                                                       //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.InfraredSensor = &testInfraredSensor{}

type testInfraredSensor struct{}

func NewTestInfraredSensor() ev3lib.InfraredSensor {
	return &testInfraredSensor{}
}

func (s *testInfraredSensor) Distance() float64 {
	return 0
}

func (s *testInfraredSensor) Buttons(channel int) []ev3lib.BeaconButton {
	return []ev3lib.BeaconButton{}
}

////////////////////////////////////////////////////////////////////////////////
// Test Touch Sensor                                                          //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.TouchSensor = &testTouchSensor{}

type testTouchSensor struct{}

func NewTestTouchSensor() ev3lib.TouchSensor {
	return &testTouchSensor{}
}

func (s *testTouchSensor) IsPressed() bool {
	return false
}

////////////////////////////////////////////////////////////////////////////////
// Test Ultrasonic Sensor                                                     //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.UltrasonicSensor = &testUltrasonicSensor{}

type testUltrasonicSensor struct{}

func NewTestUltrasonicSensor() ev3lib.UltrasonicSensor {
	return &testUltrasonicSensor{}
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
