//go:build !ev3test

package ev3

import (
	"reflect"
	"unsafe"

	ev3go "github.com/ev3go/ev3"
)

const LCDByteLength = 91136
const LCDWidth = 178
const LCDHeight = 128

var clearScreen []byte

func init() {
	clearScreen = make([]byte, LCDByteLength)
	for i := range clearScreen {
		clearScreen[i] = 255
	}
}

var LCD = newLcd()

type lcd struct {
	Data []byte
}

func getUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func newLcd() *lcd {
	err := ev3go.LCD.Init(false)
	if err != nil {
		return nil
	}

	b := getUnexportedField(reflect.ValueOf(ev3go.LCD).Elem().FieldByName("fbdev")).([]byte)
	l := &lcd{Data: b}

	return l
}

func (l *lcd) Set(i int, b byte) {
	l.Data[i] = b
}

func (l *lcd) Clear() {
	copy(l.Data, clearScreen[:])
}
