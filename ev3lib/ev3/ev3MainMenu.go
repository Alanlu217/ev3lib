package ev3

import (
	"fmt"

	"github.com/Alanlu217/ev3lib/ev3lib"
)

////////////////////////////////////////////////////////////////////////////////
// EV3 Main Menu                                                              //
////////////////////////////////////////////////////////////////////////////////

var _ ev3lib.MainMenuInterface = &EV3MainMenu{}

type EV3MainMenu struct {
	ev3 *ev3lib.EV3Brick
}

func NewEV3MainMenu(ev3 *ev3lib.EV3Brick) *EV3MainMenu {
	return &EV3MainMenu{ev3}
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
	return e.ev3.IsButtonPressed(ev3lib.Right)

}

func (e *EV3MainMenu) PreviousPage() bool {
	return e.ev3.IsButtonPressed(ev3lib.Left)
}

func (e *EV3MainMenu) SetPage() (bool, int) {
	return false, 0
}

func (e *EV3MainMenu) Display(menu *ev3lib.Menu, command int, page int) {
	e.ev3.ClearScreen()
	e.ev3.DrawText(0, 0, fmt.Sprintf("Command %v on page %v", command, page))

	// fmt.Println(e.ev3.ButtonsPressed())
}
