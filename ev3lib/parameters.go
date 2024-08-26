package ev3lib

type EV3Port string

const (
	IN1 EV3Port = "ev3-ports:in1"
	IN2 EV3Port = "ev3-ports:in2"
	IN3 EV3Port = "ev3-ports:in3"
	IN4 EV3Port = "ev3-ports:in4"

	OUTA EV3Port = "ev3-ports:outA"
	OUTB EV3Port = "ev3-ports:outB"
	OUTC EV3Port = "ev3-ports:outC"
	OUTD EV3Port = "ev3-ports:outD"
)

type EV3Button int

type EV3Color struct {
	r, g, b float64
}

func NewColor(r, g, b float64) EV3Color {
	return EV3Color{r: r, g: g, b: b}
}

type EV3Note string

type MotorStopAction string

const (
	Coast MotorStopAction = "coast"
	Brake MotorStopAction = "brake"
	Hold  MotorStopAction = "hold"
)

type MotorState int

const (
	Running MotorState = 1 << iota
	Ramping
	Holding
	Overloaded
	Stalled
)

type BeaconButton int

const (
	LeftUp BeaconButton = iota
	LeftDown
	RightUp
	RightDown
	Beacon
)
