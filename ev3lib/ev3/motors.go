//go:build linux && arm

package ev3

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// EV3 Motor                                                                  //
////////////////////////////////////////////////////////////////////////////////

const (
	mediumMotorDriverName string = "lego-ev3-m-motor"
	largeMotorDriverName  string = "lego-ev3-l-motor"
)

func StopAllMotors() {
	m, err := ev3dev.TachoMotorFor("", mediumMotorDriverName)
	for err == nil {
		m.Command("stop")
		temp, e := m.Next()

		m = temp
		err = e
	}

	m, err = ev3dev.TachoMotorFor("", largeMotorDriverName)
	for err == nil {
		m.Command("stop")
		temp, e := m.Next()

		m = temp
		err = e
	}
}

var _ ev3lib.Motor = &ev3Motor{}

type ev3Motor struct {
	motor *ev3dev.TachoMotor

	scale float64
}

func NewMediumMotor(port ev3lib.EV3Port) (ev3lib.Motor, error) {
	m, err := ev3dev.TachoMotorFor(string(port), mediumMotorDriverName)
	if err != nil {
		return nil, err
	}

	return &ev3Motor{motor: m, scale: 1}, nil
}

func NewLargeMotor(port ev3lib.EV3Port) (ev3lib.Motor, error) {
	m, err := ev3dev.TachoMotorFor(string(port), largeMotorDriverName)
	if err != nil {
		return nil, err
	}

	return &ev3Motor{motor: m, scale: 1}, nil
}

func (m *ev3Motor) CountPerRot() int {
	return m.motor.CountPerRot()
}

func (m *ev3Motor) State() ev3lib.MotorState {
	s, _ := m.motor.State()
	return ev3lib.MotorState(s)
}

func (m *ev3Motor) Inverted() bool {
	p, _ := m.motor.Polarity()
	return p == ev3dev.Inversed
}

func (m *ev3Motor) SetInverted(inverted bool) {
	if inverted {
		m.motor.SetPolarity(ev3dev.Inversed)
	} else {
		m.motor.SetPolarity(ev3dev.Normal)
	}
}

func (m *ev3Motor) Scale() float64 {
	return m.scale
}

func (m *ev3Motor) SetScale(scale float64) {
	m.scale = scale
}

func (m *ev3Motor) Position() float64 {
	p, _ := m.motor.Position()
	return float64(p) * m.scale
}

func (m *ev3Motor) ResetPosition(pos float64) {
	m.motor.SetPosition(int(pos / m.scale))
}

func (m *ev3Motor) Speed() float64 {
	s, _ := m.motor.Speed()
	return float64(s) * m.scale
}

func (m *ev3Motor) Set(power float64) {
	m.motor.SetDutyCycleSetpoint(int(power * 100)).Command("run-direct")
}

func (m *ev3Motor) Stop() {
	m.motor.Command("stop")
}

func (m *ev3Motor) StopAction() ev3lib.MotorStopAction {
	s, _ := m.motor.StopAction()
	return ev3lib.MotorStopAction(s)
}

func (m *ev3Motor) SetStopAction(s ev3lib.MotorStopAction) {
	m.motor.SetStopAction(string(s))
}
