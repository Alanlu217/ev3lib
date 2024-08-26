//go:build linux && arm

package ev3

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// EV3 Ultrasonic Sensor                                                      //
////////////////////////////////////////////////////////////////////////////////

const ultrasonicSensorDriverName string = "lego-ev3-us"

type ultrasonicSensorMode string

const (
	ultrasonicSensorModeProximity       ultrasonicSensorMode = "US-DIST-CM"
	ultrasonicSensorModeSilentProximity ultrasonicSensorMode = "US-SI-CM"
	ultrasonicSensorModeListen          ultrasonicSensorMode = "US-LISTEN"
)

var _ ev3lib.UltrasonicSensor = &ev3UltrasonicSensor{}

// Provides access to an EV3 ultrasonic sensor.
type ev3UltrasonicSensor struct {
	sensor *ev3dev.Sensor

	currentMode ultrasonicSensorMode
}

// NewUltrasonicSensor creates a new ultrasonic sensor on the provided port.
func NewUltrasonicSensor(port ev3lib.EV3Port) (ev3lib.UltrasonicSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), ultrasonicSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(ultrasonicSensorModeProximity))
	return &ev3UltrasonicSensor{sensor: sensor, currentMode: ultrasonicSensorModeProximity}, nil
}

// Distance returns the measured distance in centimeters from 0 to 2550.
func (s *ev3UltrasonicSensor) Distance() float64 {
	if s.currentMode != ultrasonicSensorModeProximity {
		s.sensor.SetMode(string(ultrasonicSensorModeProximity))
	}

	val, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	return float64(valInt)
}

// DistanceSilent same as Distance(), but will turn sensor off after measurement.
func (s *ev3UltrasonicSensor) DistanceSilent() float64 {
	if s.currentMode != ultrasonicSensorModeSilentProximity {
		s.sensor.SetMode(string(ultrasonicSensorModeSilentProximity))
	}

	val, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	return float64(valInt)
}

// Presence listens for the presence of other ultrasonic sensors.
func (s *ev3UltrasonicSensor) Presence() bool {
	if s.currentMode != ultrasonicSensorModeListen {
		s.sensor.SetMode(string(ultrasonicSensorModeListen))
	}

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
