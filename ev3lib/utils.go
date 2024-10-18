package ev3lib

import (
	"math"

	"golang.org/x/exp/constraints"
)

////////////////////////////////////////////////////////////////////////////////
// PID Controller                                                             //
////////////////////////////////////////////////////////////////////////////////

type PIDController struct {
	kp, ki, kd float64

	derivative, integral, lastError float64
}

func NewPIDController(kp, ki, kd float64) *PIDController {
	return &PIDController{kp, ki, kd, 0, 0, 0}
}

func (p *PIDController) Get(current, setPoint float64) float64 {
	e := setPoint - current

	p.derivative = p.lastError - e
	p.integral = p.integral/2 + e
	p.lastError = e

	return (e * p.kp) + (p.integral * p.ki) + (p.derivative * p.kd)
}

func (p *PIDController) Kp() float64 {
	return p.kp
}

func (p *PIDController) SetKp(kp float64) {
	p.kp = kp
}

func (p *PIDController) Ki() float64 {
	return p.ki
}

func (p *PIDController) SetKi(ki float64) {
	p.ki = ki
}

func (p *PIDController) Kd() float64 {
	return p.kd
}

func (p *PIDController) SetKd(kd float64) {
	p.kd = kd
}

func (p *PIDController) PID() (float64, float64, float64) {
	return p.kp, p.ki, p.kd
}

func (p *PIDController) SetPID(kp, ki, kd float64) {
	p.kp = kp
	p.ki = ki
	p.kd = kd
}

////////////////////////////////////////////////////////////////////////////////
// LCD Utils                                                                  //
////////////////////////////////////////////////////////////////////////////////

func LCDIndexToPixel(idx int) (x, y int) {
	row, rem := math.Modf(float64(idx) / (128 * 4))

	col := math.Floor(rem * 128)

	return int(row), int(col)
}

func LCDPixelToIndex(x, y int) int {
	return x*4 + y*4*178
}

////////////////////////////////////////////////////////////////////////////////
// Math                                                                       //
////////////////////////////////////////////////////////////////////////////////

func Clamp[T constraints.Ordered](v, min, max T) T {
	if v > max {
		return max
	}

	if v < min {
		return min
	}

	return v
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
