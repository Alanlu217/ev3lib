package ev3lib

import "github.com/ev3go/ev3dev"

type MotorStopAction string

const (
	Coast MotorStopAction = "coast"
	Brake MotorStopAction = "brake"
	Hold  MotorStopAction = "hold"
)

type MotorState ev3dev.MotorState

const (
	Running MotorState = 1 << iota
	Ramping
	Holding
	Overloaded
	Stalled
)

////////////////////////////////////////////////////////////////////////////////
// Motor Interface                                                            //
////////////////////////////////////////////////////////////////////////////////

type Motor interface {
	CountPerRot() int
	State() MotorState

	Inverted() bool
	SetInverted(inverted bool)

	Scale() float64
	SetScale(scale float64)
	Position() float64
	ResetPosition(pos float64)
	Speed() float64

	Set(power float64)
	Stop()

	StopAction() MotorStopAction
	SetStopAction(s MotorStopAction)
}

////////////////////////////////////////////////////////////////////////////////
// Test Motor                                                                 //
////////////////////////////////////////////////////////////////////////////////

var _ Motor = &TestMotor{}

type TestMotor struct{}

func NewTestMotor() Motor {
	return &TestMotor{}
}

func (m *TestMotor) CountPerRot() int {
	return 0
}

func (m *TestMotor) State() MotorState {
	return Holding
}

func (m *TestMotor) Inverted() bool {
	return false
}

func (m *TestMotor) SetInverted(inverted bool) {}

func (m *TestMotor) Scale() float64 {
	return 0
}

func (m *TestMotor) SetScale(scale float64) {}

func (m *TestMotor) Position() float64 {
	return 0
}

func (m *TestMotor) ResetPosition(pos float64) {}

func (m *TestMotor) Speed() float64 {
	return 0
}

func (m *TestMotor) Set(power float64) {}

func (m *TestMotor) Stop() {}

func (m *TestMotor) StopAction() MotorStopAction {
	return Coast
}

func (m *TestMotor) SetStopAction(s MotorStopAction) {}

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

var _ Motor = &EV3Motor{}

type EV3Motor struct {
	motor *ev3dev.TachoMotor

	scale float64
}

func NewMediumMotor(port EV3Port) (Motor, error) {
	m, err := ev3dev.TachoMotorFor(string(port), mediumMotorDriverName)
	if err != nil {
		return nil, err
	}

	return &EV3Motor{motor: m, scale: 1}, nil
}

func NewLargeMotor(port EV3Port) (Motor, error) {
	m, err := ev3dev.TachoMotorFor(string(port), largeMotorDriverName)
	if err != nil {
		return nil, err
	}

	return &EV3Motor{motor: m, scale: 1}, nil
}

func (m *EV3Motor) CountPerRot() int {
	return m.motor.CountPerRot()
}

func (m *EV3Motor) State() MotorState {
	s, _ := m.motor.State()
	return MotorState(s)
}

func (m *EV3Motor) Inverted() bool {
	p, _ := m.motor.Polarity()
	return p == ev3dev.Inversed
}

func (m *EV3Motor) SetInverted(inverted bool) {
	if inverted {
		m.motor.SetPolarity(ev3dev.Inversed)
	} else {
		m.motor.SetPolarity(ev3dev.Normal)
	}
}

func (m *EV3Motor) Scale() float64 {
	return m.scale
}

func (m *EV3Motor) SetScale(scale float64) {
	m.scale = scale
}

func (m *EV3Motor) Position() float64 {
	p, _ := m.motor.Position()
	return float64(p) * m.scale
}

func (m *EV3Motor) ResetPosition(pos float64) {
	m.motor.SetPosition(int(pos / m.scale))
}

func (m *EV3Motor) Speed() float64 {
	s, _ := m.motor.Speed()
	return float64(s) * m.scale
}

func (m *EV3Motor) Set(power float64) {
	m.motor.SetDutyCycleSetpoint(int(power * 100)).Command("run-direct")
}

func (m *EV3Motor) Stop() {
	m.motor.Command("stop")
}

func (m *EV3Motor) StopAction() MotorStopAction {
	s, _ := m.motor.StopAction()
	return MotorStopAction(s)
}

func (m *EV3Motor) SetStopAction(s MotorStopAction) {
	m.motor.SetStopAction(string(s))
}
