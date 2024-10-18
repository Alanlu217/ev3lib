package ev3lib

////////////////////////////////////////////////////////////////////////////////
// EV3Brick interface                                                         //
////////////////////////////////////////////////////////////////////////////////

type EV3BrickInterface interface {
	IsButtonPressed(button EV3Button) bool
	IsButtonDown(button EV3Button) bool
	IsButtonReleased(button EV3Button) bool
	IsButtonUp(button EV3Button) bool
	ButtonsPressed() []EV3Button

	SetLight(color EV3Color)

	Beep(frequency, duration float64)

	PlayNotes(notes []EV3Note, tempo float64)

	SetVolume(volume float64)

	ClearScreen()

	DrawText(x, y int, text string)

	PrintScreen(text ...string)

	DrawPixel(x, y int, black bool)

	Voltage() float64

	Current() float64
}

////////////////////////////////////////////////////////////////////////////////
// Main Menu interface                                                        //
////////////////////////////////////////////////////////////////////////////////

type MainMenuInterface interface {
	Exit() bool

	RunSelected() bool

	NextCommand() bool
	PreviousCommand() bool

	SetCommand() (bool, int)

	NextPage() bool
	PreviousPage() bool

	SetPage() (bool, int)

	Display(menu *Menu, command, page int)
}

////////////////////////////////////////////////////////////////////////////////
// MotorInterface Interface                                                   //
////////////////////////////////////////////////////////////////////////////////

type MotorInterface interface {
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
// Color Sensor Interface                                                     //
////////////////////////////////////////////////////////////////////////////////

type ColorSensorInterface interface {
	Ambient() float64
	Reflection() float64
	GetRGB() (float64, float64, float64)
}

////////////////////////////////////////////////////////////////////////////////
// Gyro Sensor Interface                                                      //
////////////////////////////////////////////////////////////////////////////////

type GyroSensorInterface interface {
	Rate() float64
	Angle() float64
	AngleRate() (float64, float64)
	ResetAngle(angle float64)
	Calibrate()
}

////////////////////////////////////////////////////////////////////////////////
// Infrared Sensor Interface                                                  //
////////////////////////////////////////////////////////////////////////////////

type InfraredSensorInterface interface {
	Distance() float64
	Buttons(channel int) []BeaconButton
}

////////////////////////////////////////////////////////////////////////////////
// Touch Sensor Interface                                                     //
////////////////////////////////////////////////////////////////////////////////

type TouchSensorInterface interface {
	IsPressed() bool
}

////////////////////////////////////////////////////////////////////////////////
// Ultrasonic Sensor Interface                                                //
////////////////////////////////////////////////////////////////////////////////

type UltrasonicSensorInterface interface {
	Distance() float64
	DistanceSilent() float64
	Presence() bool
}
