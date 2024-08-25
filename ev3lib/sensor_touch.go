package ev3lib

import (
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// Touch Sensor Interface                                                     //
////////////////////////////////////////////////////////////////////////////////

type TouchSensor interface {
	IsPressed() bool
}

////////////////////////////////////////////////////////////////////////////////
// Test Touch Sensor                                                          //
////////////////////////////////////////////////////////////////////////////////

var _ TouchSensor = &TestTouchSensor{}

type TestTouchSensor struct{}

func NewTestTouchSensor() *TestTouchSensor {
	return &TestTouchSensor{}
}

func (s *TestTouchSensor) IsPressed() bool {
	return false
}

////////////////////////////////////////////////////////////////////////////////
// EV3 Touch Sensor                                                           //
////////////////////////////////////////////////////////////////////////////////

const touchSensorDriverName string = "lego-ev3-touch"

var _ TouchSensor = &EV3TouchSensor{}

// Provides access to a EV3 touch sensor.
type EV3TouchSensor struct {
	sensor *ev3dev.Sensor
}

// Creates a new touch sensor on the provided port.
func NewTouchSensor(port EV3Port) (TouchSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), touchSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode("TOUCH")
	return &EV3TouchSensor{sensor: sensor}, nil
}

// Returns whether the button is currently being pressed.
func (s *EV3TouchSensor) IsPressed() bool {
	val, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	return valInt == 1
}
