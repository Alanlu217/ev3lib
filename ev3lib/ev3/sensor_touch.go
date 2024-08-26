//go:build !ev3test

package ev3

import (
	"log"
	"strconv"

	"github.com/Alanlu217/ev3lib/ev3lib"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// EV3 Touch Sensor                                                           //
////////////////////////////////////////////////////////////////////////////////

const touchSensorDriverName string = "lego-ev3-touch"

var _ ev3lib.TouchSensor = &ev3TouchSensor{}

// Provides access to a EV3 touch sensor.
type ev3TouchSensor struct {
	sensor *ev3dev.Sensor
}

// NewTouchSensor creates a new touch sensor on the provided port.
func NewTouchSensor(port ev3lib.EV3Port) (ev3lib.TouchSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), touchSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode("TOUCH")
	return &ev3TouchSensor{sensor: sensor}, nil
}

// IsPressed returns whether the button is currently being pressed.
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
