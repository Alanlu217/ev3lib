package ev3lib

import (
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

const touchSensorDriverName string = "lego-ev3-touch"

// Provides access to a EV3 touch sensor.
type TouchSensor struct {
	sensor *ev3dev.Sensor
}

// Creates a new touch sensor on the provided port.
func NewTouchSensor(port EV3Port) (*TouchSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), touchSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode("TOUCH")
	return &TouchSensor{sensor: sensor}, nil
}

// Returns whether the button is currently being pressed.
func (s *TouchSensor) IsPressed() bool {
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
