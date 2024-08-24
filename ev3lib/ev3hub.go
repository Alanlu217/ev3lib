package ev3lib

type EV3Button int

type EV3Color struct {
	r, g, b float64
}

func NewColor(r, g, b float64) EV3Color {
	return EV3Color{r: r, g: g, b: b}
}

type EV3Note string

////////////////////////////////////////////////////////////////////////////////
// EV3Brick interface                                                         //
////////////////////////////////////////////////////////////////////////////////

type EV3Brick interface {
	ButtonsPressed() []EV3Button

	SetLight(color EV3Color)

	Beep(frequency, duration float64)

	PlayNotes(notes []EV3Note, tempo float64)

	SetVolume(volume float64)

	ClearScreen()

	DrawText(x, y int, text string)

	PrintScreen(text ...string)

	DrawPixel(x, y int)

	Voltage() float64

	Current() float64
}

////////////////////////////////////////////////////////////////////////////////
// TestEV3Brick                                                               //
////////////////////////////////////////////////////////////////////////////////

var _ EV3Brick = &TestEV3Brick{}

type TestEV3Brick struct{}

func (*TestEV3Brick) ButtonsPressed() []EV3Button {
	return []EV3Button{}
}

func (*TestEV3Brick) SetLight(color EV3Color) {}

func (*TestEV3Brick) Beep(frequency float64, duration float64) {}

func (*TestEV3Brick) PlayNotes(notes []EV3Note, tempo float64) {}

func (*TestEV3Brick) SetVolume(volume float64) {}

func (*TestEV3Brick) ClearScreen() {}

func (*TestEV3Brick) DrawText(x int, y int, text string) {}

func (*TestEV3Brick) PrintScreen(text ...string) {}

func (*TestEV3Brick) DrawPixel(x int, y int) {}

func (*TestEV3Brick) Voltage() float64 {
	return 0
}

func (*TestEV3Brick) Current() float64 {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// Actual EV3Brick                                                            //
////////////////////////////////////////////////////////////////////////////////

var _ EV3Brick = &EV3{}

type EV3 struct {
}

func NewEV3() *EV3 {
	return &EV3{}
}

func (e *EV3) ButtonsPressed() []EV3Button {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) SetLight(color EV3Color) {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) Beep(frequency float64, duration float64) {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) PlayNotes(notes []EV3Note, tempo float64) {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) SetVolume(volume float64) {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) ClearScreen() {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) DrawText(x int, y int, text string) {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) PrintScreen(text ...string) {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) DrawPixel(x int, y int) {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) Voltage() float64 {
	panic("not implemented") // TODO: Implement
}

func (e *EV3) Current() float64 {
	panic("not implemented") // TODO: Implement
}
