//go:build !ev3test

package ev3

import (
	ev3go "github.com/ev3go/ev3"
	"reflect"
	"unsafe"
)

const LCDByteLength = 91136

var LCD = newLcd()

type lcd struct {
	data []byte
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
	l := &lcd{data: b}

	return l
}

func (l *lcd) Set(i int, b byte) {
	l.data[i] = b
}
