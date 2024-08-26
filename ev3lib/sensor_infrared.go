package ev3lib

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// Infrared Sensor Interface                                                  //
////////////////////////////////////////////////////////////////////////////////

type InfraredSensor interface {
	Distance() float64
	Buttons(channel int) []BeaconButton
}

////////////////////////////////////////////////////////////////////////////////
// Test Infrared Sensor                                                       //
////////////////////////////////////////////////////////////////////////////////

var _ InfraredSensor = &testInfraredSensor{}

type testInfraredSensor struct{}

func NewTestInfraredSensor() InfraredSensor {
	return &testInfraredSensor{}
}

func (s *testInfraredSensor) Distance() float64 {
	return 0
}

func (s *testInfraredSensor) Buttons(channel int) []BeaconButton {
	return []BeaconButton{}
}

////////////////////////////////////////////////////////////////////////////////
// EV3 Infrared Sensor                                                        //
////////////////////////////////////////////////////////////////////////////////

const infraredSensorDriverName string = "lego-ev3-ir"

type infraredSensorMode string

const (
	infraredSensorModeProximity infraredSensorMode = "IR-PROX"
	infraredSensorModeRemote    infraredSensorMode = "IR-REMOTE"
)

var _ InfraredSensor = &ev3InfraredSensor{}

// Provides access to the EV3 infrared sensor.
type ev3InfraredSensor struct {
	sensor *ev3dev.Sensor

	currentMode infraredSensorMode
}

// Creates a new infrared sensor from the provided port.
func NewInfraredSensor(port EV3Port) (InfraredSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), infraredSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(infraredSensorModeProximity))
	return &ev3InfraredSensor{sensor: sensor, currentMode: infraredSensorModeProximity}, nil
}

type BeaconButton int

const (
	LeftUp BeaconButton = iota
	LeftDown
	RightUp
	RightDown
	Beacon
)

// Returns the distance measured by the sensor from 0 to 1.
func (s *ev3InfraredSensor) Distance() float64 {
	if s.currentMode != infraredSensorModeProximity {
		s.sensor.SetMode(string(infraredSensorModeProximity))
	}

	val, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	distance := float64(valInt) / 100

	return distance
}

var buttonMap = map[int][]BeaconButton{
	0:  {},
	1:  {LeftUp},
	2:  {LeftDown},
	3:  {RightUp},
	4:  {RightDown},
	5:  {LeftUp, RightUp},
	6:  {LeftUp, RightDown},
	7:  {LeftDown, RightUp},
	8:  {LeftDown, RightDown},
	9:  {Beacon},
	10: {LeftUp, LeftDown},
	11: {RightUp, RightDown},
}

// Returns a slice of BeaconButton's containing all the buttons that are currently being pressed.
// Checks the buttons on the provided channel.
func (s *ev3InfraredSensor) Buttons(channel int) []BeaconButton {
	if s.currentMode != infraredSensorModeRemote {
		s.sensor.SetMode(string(infraredSensorModeRemote))
	}

	if channel < 0 || channel > 4 {
		fmt.Println("Channel does not exist, only 0 to 4 are accepted")
		return []BeaconButton{}
	}

	val, err := s.sensor.Value(channel)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	return buttonMap[valInt]
}
