package ev3

import (
	"log"
	"os"
	"reflect"
	"sync"
	"syscall"
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

const (
	key_backspace = 14
	key_enter     = 28
	key_up        = 103
	key_down      = 108
	key_left      = 105
	key_right     = 106

	key_max = 0x2ff

	keyBufLen = (key_max + 7) / 8
)

var ev3devButtons = [...]uint{
	key_backspace,
	key_left,
	key_enter,
	key_right,
	key_up,
	key_down,
}

type ev3ButtonHandler struct {
	buttons map[ev3lib.EV3Button]*button

	buf []byte
	ev  *os.File

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

	buf := make([]byte, keyBufLen)

	ev, err := os.Open(ev3dev.ButtonPath)
	if err != nil {
		log.Fatal(err)
	}

	return &ev3ButtonHandler{b, buf, ev, ev3dev.ButtonPoller{}}
}

func isSet(bit uint, buf []byte) bool {
	return buf[bit>>3]&(1<<(bit&7)) != 0
}

func getButton(buf []byte) ev3dev.Button {
	var pressed ev3dev.Button
	for i, bit := range &ev3devButtons {
		if isSet(bit, buf) {
			pressed |= 1 << uint(i)
		}
	}
	return pressed
}

func (b *ev3ButtonHandler) updateButton(isUp bool, button ev3lib.EV3Button) {
	last := b.buttons[button].get()

	if isUp {
		if last == pressed || last == down {
			b.buttons[button].set(released)
		}
	} else {
		// Previously released
		if last == released || last == up {
			b.buttons[button].set(pressed)
		}
	}
}

const (
	_ioc_read = 2

	_ioc_nrbits   = 8
	_ioc_typebits = 8
	_ioc_sizebits = 14
	_ioc_dirbits  = 2

	_ioc_nrmask   = 1<<_ioc_nrbits - 1
	_ioc_typemask = 1<<_ioc_typebits - 1
	_ioc_sizemask = 1<<_ioc_sizebits - 1
	_ioc_dirmask  = 1<<_ioc_dirbits - 1

	_ioc_nrshift   = 0
	_ioc_typeshift = _ioc_nrshift + _ioc_nrbits
	_ioc_sizeshift = _ioc_typeshift + _ioc_typebits
	_ioc_dirshift  = _ioc_sizeshift + _ioc_sizebits
)

func eviocgkey(buf []byte) uintptr {
	return _ioc_read<<_ioc_dirshift | uintptr(len(buf))<<_ioc_sizeshift | 'E'<<_ioc_typeshift | 0x18<<_ioc_nrshift
}

func ioctl(fd, cmd, ptr uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr)
	if errno != 0 {
		return errno
	}
	return nil
}

func (b *ev3ButtonHandler) run() {
	t := time.NewTicker(60 * time.Millisecond)

	for {
		err := ioctl(b.ev.Fd(), eviocgkey(b.buf), reflect.ValueOf(b.buf).Index(0).Addr().Pointer())
		if err != nil {
			log.Fatalf("ev3dev: failed to set ioctl command for button event device: %v\n", err)
		}
    val := getButton(b.buf)

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
	val := b.buttons[button].get()

	if val == pressed {
		b.buttons[button].set(down)
	} else if val == released {
		b.buttons[button].set(up)
	}

	return val
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
