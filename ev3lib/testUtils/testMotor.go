package testUtils

import (
	"log"

	"github.com/Alanlu217/ev3lib/ev3lib"
)

////////////////////////////////////////////////////////////////////////////////
// Test MotorInterface                                                                 //
////////////////////////////////////////////////////////////////////////////////

type testMotor struct {
	name string

	inverted   bool
	scale      float64
	stopAction ev3lib.MotorStopAction
}

func NewTestMotor(name string) *ev3lib.Motor {
	return ev3lib.NewMotorBase(&testMotor{name: name, inverted: false, scale: 1, stopAction: ev3lib.Brake})
}

func (m *testMotor) CountPerRot() int {
	return 0
}

func (m *testMotor) State() ev3lib.MotorState {
	return ev3lib.Holding
}

func (m *testMotor) Inverted() bool {
	return m.inverted
}

func (m *testMotor) SetInverted(inverted bool) {
	m.inverted = inverted
}

func (m *testMotor) Scale() float64 {
	return m.scale
}

func (m *testMotor) SetScale(scale float64) {
	m.scale = scale
}

func (m *testMotor) Position() float64 {
	return 0
}

func (m *testMotor) ResetPosition(pos float64) {
	log.Printf("%v reset position\n", m.name)
}

func (m *testMotor) Speed() float64 {
	return 0
}

func (m *testMotor) Set(power float64) {
	log.Printf("%v set to %v\n", m.name, power)
}

func (m *testMotor) Stop() {
	log.Printf("%v stopped\n", m.name)
}

func (m *testMotor) StopAction() ev3lib.MotorStopAction {
	return m.stopAction
}

func (m *testMotor) SetStopAction(s ev3lib.MotorStopAction) {
	m.stopAction = s
}
