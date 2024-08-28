package ev3lib

import "math"

////////////////////////////////////////////////////////////////////////////////
// MotorInterface Base                                                                 //
////////////////////////////////////////////////////////////////////////////////

type Motor struct {
	MotorInterface
}

func NewMotorBase(m MotorInterface) *Motor {
	return &Motor{m}
}

////////////////////////////////////////////////////////////////////////////////
// MotorInterface Commands                                                             //
////////////////////////////////////////////////////////////////////////////////

func (m *Motor) SetCommand(power float64) *Command {
	return NewFuncCommand(func() { m.Set(power) }).Repeatedly().WhenDone(func(_ bool) { m.Set(0) })
}

////////////////////////////////////////////////////////////////////////////////

type runToRelPosCommand struct {
	DefaultCommand

	pos float64

	targetPos float64

	done      bool
	tolerance float64

	pid PIDController
	m   MotorInterface
}

func (r *runToRelPosCommand) Init() {
	r.targetPos = r.pos + r.m.Position()
}

func (r *runToRelPosCommand) Run() {
	pow := r.pid.Get(r.m.Position(), r.targetPos)

	if math.Abs(pow) < r.tolerance {
		r.done = true
	}

	r.m.Set(pow)
}

func (r *runToRelPosCommand) End(interrupted bool) {
	r.m.Stop()
}

func (r *runToRelPosCommand) IsDone() bool {
	return r.done
}

func (m *Motor) RunToRelPos(pos float64, tolerance float64, pid PIDController) *Command {
	return NewCommand(&runToRelPosCommand{pos: pos, pid: pid, m: m.MotorInterface})
}

////////////////////////////////////////////////////////////////////////////////

type runToAbsPosCommand struct {
	DefaultCommand

	pos float64

	done      bool
	tolerance float64

	pid PIDController
	m   MotorInterface
}

func (r *runToAbsPosCommand) Run() {
	pow := r.pid.Get(r.m.Position(), r.pos)

	if math.Abs(pow) < r.tolerance {
		r.done = true
	}

	r.m.Set(pow)
}

func (r *runToAbsPosCommand) End(interrupted bool) {
	r.m.Stop()
}

func (r *runToAbsPosCommand) IsDone() bool {
	return r.done
}

func (m *Motor) RunToAbsPos(pos float64, tolerance float64, pid PIDController) *Command {
	return NewCommand(&runToAbsPosCommand{pos: pos, pid: pid, m: m.MotorInterface})
}
