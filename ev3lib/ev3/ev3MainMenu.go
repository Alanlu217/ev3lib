package ev3

import (
	"slices"

	"github.com/Alanlu217/ev3lib/ev3lib"
)

////////////////////////////////////////////////////////////////////////////////
// EV3 Main Menu                                                              //
////////////////////////////////////////////////////////////////////////////////

type EV3MainMenu struct {
	ev3 ev3lib.EV3Brick
}

func (e *EV3MainMenu) RunSelected() bool {
	return slices.Contains(e.ev3.ButtonsPressed(), ev3lib.Middle)
}

func (e *EV3MainMenu) NextCommand() bool {
	panic("not implemented") // TODO: Implement
}

func (e *EV3MainMenu) PreviousCommand() bool {
	panic("not implemented") // TODO: Implement
}

func (e *EV3MainMenu) SetCommand() int {
	panic("not implemented") // TODO: Implement
}

func (e *EV3MainMenu) NextPage() bool {
	panic("not implemented") // TODO: Implement
}

func (e *EV3MainMenu) PreviousPage() bool {
	panic("not implemented") // TODO: Implement
}

func (e *EV3MainMenu) SetPage() int {
	panic("not implemented") // TODO: Implement
}

func (e *EV3MainMenu) Display(menu *ev3lib.Menu, command int, page int) {
	panic("not implemented") // TODO: Implement
}
