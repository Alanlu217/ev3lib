//go:build !ev3test

package ev3

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/ev3go/ev3dev"
	"log"
)

////////////////////////////////////////////////////////////////////////////////
// Actual EV3Brick                                                            //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.EV3BrickInterface = &ev3{}

type ev3 struct {
	p ev3dev.PowerSupply
}

func NewEV3() ev3lib.EV3BrickInterface {
	return &ev3{p: ""}
}

func (e *ev3) ButtonsPressed() []ev3lib.EV3Button {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) SetLight(color ev3lib.EV3Color) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) Beep(frequency float64, duration float64) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) PlayNotes(notes []ev3lib.EV3Note, tempo float64) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) SetVolume(volume float64) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) ClearScreen() {
	for i := 0; i < LCDByteLength; i++ {
		LCD.Set(i, 255)
	}
}

func (e *ev3) DrawText(x int, y int, text string) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) PrintScreen(text ...string) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) DrawPixel(x int, y int) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) Voltage() float64 {
	volt, err := e.p.Voltage()
	if err != nil {
		log.Fatal(err)
	}

	return volt
}

func (e *ev3) Current() float64 {
	curr, err := e.p.Current()
	if err != nil {
		log.Fatal(err)
	}

	return curr
}
