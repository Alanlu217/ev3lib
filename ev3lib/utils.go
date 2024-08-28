package ev3lib

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
