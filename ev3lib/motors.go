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

var _ Motor = &testMotor{}

type testMotor struct{}

func NewTestMotor() Motor {
	return &testMotor{}
}

func (m *testMotor) CountPerRot() int {
	return 0
}

func (m *testMotor) State() MotorState {
	return Holding
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

func (m *testMotor) Set(power float64) {}

func (m *testMotor) Stop() {}

func (m *testMotor) StopAction() MotorStopAction {
	return Coast
}

func (m *testMotor) SetStopAction(s MotorStopAction) {}

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

var _ Motor = &ev3Motor{}

type ev3Motor struct {
	motor *ev3dev.TachoMotor

	scale float64
}

func NewMediumMotor(port EV3Port) (Motor, error) {
	m, err := ev3dev.TachoMotorFor(string(port), mediumMotorDriverName)
	if err != nil {
		return nil, err
	}

	return &ev3Motor{motor: m, scale: 1}, nil
}

func NewLargeMotor(port EV3Port) (Motor, error) {
	m, err := ev3dev.TachoMotorFor(string(port), largeMotorDriverName)
	if err != nil {
		return nil, err
	}

	return &ev3Motor{motor: m, scale: 1}, nil
}

func (m *ev3Motor) CountPerRot() int {
	return m.motor.CountPerRot()
}

func (m *ev3Motor) State() MotorState {
	s, _ := m.motor.State()
	return MotorState(s)
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

func (m *ev3Motor) StopAction() MotorStopAction {
	s, _ := m.motor.StopAction()
	return MotorStopAction(s)
}

func (m *ev3Motor) SetStopAction(s MotorStopAction) {
	m.motor.SetStopAction(string(s))
}
