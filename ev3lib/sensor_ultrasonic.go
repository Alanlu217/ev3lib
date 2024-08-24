package ev3lib

import (
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// Ultrasonic Sensor Interface                                                //
////////////////////////////////////////////////////////////////////////////////

type UltrasonicSensor interface {
	Distance() float64
	DistanceSilent() float64
	Presence() bool
}

////////////////////////////////////////////////////////////////////////////////
// Test Ultrasonic Sensor                                                     //
////////////////////////////////////////////////////////////////////////////////

var _ UltrasonicSensor = &TestUltrasonicSensor{}

type TestUltrasonicSensor struct{}

func NewTestUltrasonicSensor() *TestUltrasonicSensor {
	return &TestUltrasonicSensor{}
}

func (s *TestUltrasonicSensor) Distance() float64 {
	return 0
}

func (s *TestUltrasonicSensor) DistanceSilent() float64 {
	return 0
}

func (s *TestUltrasonicSensor) Presence() bool {
	return false
}

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

var _ UltrasonicSensor = &EV3UltrasonicSensor{}

// Provides access to an EV3 ultrasonic sensor.
type EV3UltrasonicSensor struct {
	sensor *ev3dev.Sensor

	currentMode ultrasonicSensorMode
}

// Creates a new ultrasonic sensor on the provided port.
func NewUltrasonicSensor(port string) (*EV3UltrasonicSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), ultrasonicSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(ultrasonicSensorModeProximity))
	return &EV3UltrasonicSensor{sensor: sensor, currentMode: ultrasonicSensorModeProximity}, nil
}

// Returns the measured distance in centimeters from 0 to 2550.
func (s *EV3UltrasonicSensor) Distance() float64 {
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

// Same as Distance(), but will turn sensor off after measurement.
func (s *EV3UltrasonicSensor) DistanceSilent() float64 {
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

// Listens for the presence of other ultrasonic sensors.
func (s *EV3UltrasonicSensor) Presence() bool {
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

	return valInt == 1
}
