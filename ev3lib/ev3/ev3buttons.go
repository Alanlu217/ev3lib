package ev3

import (
	"log"
	"sync"
	"time"

	"github.com/Alanlu217/ev3lib/ev3lib"
	"github.com/ev3go/ev3dev"
)

var buttons = [...]ev3lib.EV3Button{ev3lib.Back, ev3lib.Left, ev3lib.Middle, ev3lib.Right, ev3lib.Up, ev3lib.Down}

////////////////////////////////////////////////////////////////////////////////
// ButtonStates                                                               //
////////////////////////////////////////////////////////////////////////////////

type buttonState int

const (
	up buttonState = iota
	down
	released
	pressed
)

type button struct {
	state buttonState

	m sync.Mutex
}

func newButton() *button {
	return &button{up, sync.Mutex{}}
}

func (b *button) get() buttonState {
	b.m.Lock()

	val := b.state

	b.m.Unlock()

	return val
}

func (b *button) set(state buttonState) {
	b.m.Lock()

	b.state = state

	b.m.Unlock()
}

////////////////////////////////////////////////////////////////////////////////
// EV3 Button Handler                                                         //
////////////////////////////////////////////////////////////////////////////////

type ev3ButtonHandler struct {
	buttons map[ev3lib.EV3Button]*button

	poller ev3dev.ButtonPoller
}

func newEv3ButtonHandler() *ev3ButtonHandler {
	b := make(map[ev3lib.EV3Button]*button)

	b[ev3lib.Back] = newButton()
	b[ev3lib.Up] = newButton()
	b[ev3lib.Down] = newButton()
	b[ev3lib.Middle] = newButton()
	b[ev3lib.Left] = newButton()
	b[ev3lib.Right] = newButton()

	return &ev3ButtonHandler{b, ev3dev.ButtonPoller{}}
}

func (b *ev3ButtonHandler) updateButton(isUp bool, button ev3lib.EV3Button) {
	last := b.buttons[button].get()

	if isUp {
		if last == pressed || last == down {
			b.buttons[button].set(released)
		} else {
			b.buttons[button].set(up)
		}
	} else {
		// Previously released
		if last == released || last == up {
			b.buttons[button].set(pressed)
		} else {
			b.buttons[button].set(down)
		}
	}
}

func (b *ev3ButtonHandler) run() {
	t := time.NewTicker(50 * time.Millisecond)

	for {
		val, err := b.poller.Poll()
		if err != nil {
			log.Fatal(err)
		}

		b.updateButton(val&ev3dev.Back == 0, ev3lib.Back)
		b.updateButton(val&ev3dev.Left == 0, ev3lib.Left)
		b.updateButton(val&ev3dev.Right == 0, ev3lib.Right)
		b.updateButton(val&ev3dev.Middle == 0, ev3lib.Middle)
		b.updateButton(val&ev3dev.Up == 0, ev3lib.Up)
		b.updateButton(val&ev3dev.Down == 0, ev3lib.Down)

		<-t.C
	}
}

func (b *ev3ButtonHandler) get(button ev3lib.EV3Button) buttonState {
	return b.buttons[button].get()
}

func (b *ev3ButtonHandler) getDown() []ev3lib.EV3Button {
	result := make([]ev3lib.EV3Button, 0)
	for _, button := range buttons {
		if b.buttons[button].get() == down {
			result = append(result, button)
		}
	}

	return result
}
