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

var _ TouchSensor = &testTouchSensor{}

type testTouchSensor struct{}

func NewTestTouchSensor() TouchSensor {
	return &testTouchSensor{}
}

func (s *testTouchSensor) IsPressed() bool {
	return false
}

////////////////////////////////////////////////////////////////////////////////
// EV3 Touch Sensor                                                           //
////////////////////////////////////////////////////////////////////////////////

const touchSensorDriverName string = "lego-ev3-touch"

var _ TouchSensor = &ev3TouchSensor{}

// Provides access to a EV3 touch sensor.
type ev3TouchSensor struct {
	sensor *ev3dev.Sensor
}

// Creates a new touch sensor on the provided port.
func NewTouchSensor(port EV3Port) (TouchSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), touchSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode("TOUCH")
	return &ev3TouchSensor{sensor: sensor}, nil
}

// Returns whether the button is currently being pressed.
func (s *ev3TouchSensor) IsPressed() bool {
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
