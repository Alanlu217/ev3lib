package testUtils

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"log"
)

////////////////////////////////////////////////////////////////////////////////
// Test MotorInterface                                                                 //
////////////////////////////////////////////////////////////////////////////////

type testMotor struct{}

func NewTestMotor() *ev3lib.Motor {
	return ev3lib.NewMotorBase(&testMotor{})
}

func (m *testMotor) CountPerRot() int {
	return 0
}

func (m *testMotor) State() ev3lib.MotorState {
	return ev3lib.Holding
}

func (m *testMotor) Inverted() bool {
	return false
}

func (m *testMotor) SetInverted(inverted bool) {}

func (m *testMotor) Scale() float64 {
	return 0
}

func (m *testMotor) SetScale(scale float64) {}

func (m *testMotor) Position() float64 {
	return 0
}

func (m *testMotor) ResetPosition(pos float64) {}

func (m *testMotor) Speed() float64 {
	return 0
}

func (m *testMotor) Set(power float64) {
	log.Println("Set Motor to", power)
}

func (m *testMotor) Stop() {}

func (m *testMotor) StopAction() ev3lib.MotorStopAction {
	return ev3lib.Coast
}

func (m *testMotor) SetStopAction(s ev3lib.MotorStopAction) {}
