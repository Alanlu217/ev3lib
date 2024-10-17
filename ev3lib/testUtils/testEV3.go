package testUtils

import "github.com/Alanlu217/ev3lib/ev3lib"

////////////////////////////////////////////////////////////////////////////////
// TestEV3Brick                                                               //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.EV3BrickInterface = &testEV3Brick{}

type testEV3Brick struct{}

func NewTestEV3Brick() ev3lib.EV3BrickInterface {
	return &testEV3Brick{}
}

func (*testEV3Brick) ButtonsPressed() []ev3lib.EV3Button {
	return []ev3lib.EV3Button{}
}

func (*testEV3Brick) SetLight(color ev3lib.EV3Color) {}

func (*testEV3Brick) Beep(frequency float64, duration float64) {}

func (*testEV3Brick) PlayNotes(notes []ev3lib.EV3Note, tempo float64) {}

func (*testEV3Brick) SetVolume(volume float64) {}

func (*testEV3Brick) ClearScreen() {}

func (*testEV3Brick) DrawText(x int, y int, text string) {}

func (*testEV3Brick) PrintScreen(text ...string) {}

func (*testEV3Brick) DrawPixel(x int, y int, black bool) {}

func (*testEV3Brick) Voltage() float64 {
	return 0
}

func (*testEV3Brick) Current() float64 {
	return 0
}
