//go:build !ev3test

package ev3

import (
	"log"

	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// Actual EV3Brick                                                            //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.EV3BrickInterface = &ev3{}

type ev3 struct {
	p ev3dev.PowerSupply
	b ev3dev.ButtonPoller
}

func NewEV3() *ev3lib.EV3Brick {
	return ev3lib.NewEV3BrickBase(&ev3{p: "", b: ev3dev.ButtonPoller{}})
}

func (e *ev3) ButtonsPressed() []ev3lib.EV3Button {
	val, err := e.b.Poll()
	if err != nil {
		log.Fatal(err)
	}

	result := make([]ev3lib.EV3Button, 0)

	if val&ev3dev.Back != 0 {
		result = append(result, ev3lib.Back)
	}
	if val&ev3dev.Left != 0 {
		result = append(result, ev3lib.Left)
	}
	if val&ev3dev.Middle != 0 {
		result = append(result, ev3lib.Middle)
	}
	if val&ev3dev.Right != 0 {
		result = append(result, ev3lib.Right)
	}
	if val&ev3dev.Up != 0 {
		result = append(result, ev3lib.Up)
	}
	if val&ev3dev.Down != 0 {
		result = append(result, ev3lib.Down)
	}

	return result
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
	LCD.Clear()
	// for i := 0; i < LCDByteLength; i++ {
	// 	LCD.Data[i] = 255
	// }
}

func (e *ev3) DrawText(x int, y int, text string) {
	for i, char := range text {
		values := FontMap[char]

		for _, coord := range values {
			new_x := x + coord.x + i*CharWidth
			if new_x <= LCDWidth {
				e.DrawPixel(new_x, y+coord.y, true)
			}
		}
	}
}

func (e *ev3) PrintScreen(text ...string) {
	panic("not implemented") // TODO: Implement
}

func (e *ev3) DrawPixel(x int, y int, black bool) {
	if black {
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)] = 0
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)+1] = 0
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)+2] = 0
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)+3] = 0
	} else {
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)] = 255
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)+1] = 255
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)+2] = 255
		LCD.Data[ev3lib.LCDPixelToIndex(x, y)+3] = 255
	}
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
