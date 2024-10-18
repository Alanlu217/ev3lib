package ev3

import (
	"fmt"

	"github.com/Alanlu217/ev3lib/ev3lib"
)

const maxRows int = (LCDHeight / CharHeight) - 1

////////////////////////////////////////////////////////////////////////////////
// EV3 Main Menu                                                              //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.MainMenuInterface = &EV3MainMenu{}

type EV3MainMenu struct {
	ev3 *ev3lib.EV3Brick

	idx int
}

func NewEV3MainMenu(ev3 *ev3lib.EV3Brick) *EV3MainMenu {
	return &EV3MainMenu{ev3, 0}
}

func (e *EV3MainMenu) Exit() bool {
	return false
}

func (e *EV3MainMenu) RunSelected() bool {
	return e.ev3.IsButtonPressed(ev3lib.Middle)
}

func (e *EV3MainMenu) NextCommand() bool {
	return e.ev3.IsButtonPressed(ev3lib.Down)
}

func (e *EV3MainMenu) PreviousCommand() bool {
	return e.ev3.IsButtonPressed(ev3lib.Up)
}

func (e *EV3MainMenu) SetCommand() (bool, int) {
	return false, 0
}

func (e *EV3MainMenu) NextPage() bool {
	if e.ev3.IsButtonPressed(ev3lib.Right) {
		e.idx = 0
		return true
	}
	return false
}

func (e *EV3MainMenu) PreviousPage() bool {
	if e.ev3.IsButtonPressed(ev3lib.Left) {
		e.idx = 0
		return true
	}
	return false
}

func (e *EV3MainMenu) SetPage() (bool, int) {
	return false, 0
}

func (e *EV3MainMenu) Display(menu *ev3lib.Menu, command int, page int, running bool) {
	e.ev3.ClearScreen()

	if running {
		e.ev3.DrawText(0, 0, menu.Pages[page].Commands[command].Name)

		return
	}

	start := command - 1
	start = max(start, 0)

	end := min(start+4, len(menu.Pages[page].Commands))

	idx := 0
	for i := start; i < end; i++ {
		if i == command {
			e.ev3.DrawText(0, idx*CharHeight, fmt.Sprintf("> %v", menu.Pages[page].Commands[i].Name))
		} else {
			e.ev3.DrawText(0, idx*CharHeight, fmt.Sprintf("  %v", menu.Pages[page].Commands[i].Name))
		}
		idx++
	}

	e.ev3.DrawText(0, maxRows*CharHeight, fmt.Sprintf("%.2fV", e.ev3.Voltage()))

	// fmt.Println(e.ev3.ButtonsPressed())
}
