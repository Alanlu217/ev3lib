package ev3

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Alanlu217/ev3lib/ev3lib"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// EV3 Infrared Sensor                                                        //
////////////////////////////////////////////////////////////////////////////////

const infraredSensorDriverName string = "lego-ev3-ir"

type infraredSensorMode string

const (
	infraredSensorModeProximity infraredSensorMode = "IR-PROX"
	infraredSensorModeRemote    infraredSensorMode = "IR-REMOTE"
)

var _ ev3lib.InfraredSensorInterface = &ev3InfraredSensor{}

// Provides access to the EV3 infrared sensor.
type ev3InfraredSensor struct {
	sensor *ev3dev.Sensor

	currentMode infraredSensorMode
}

// NewInfraredSensor creates a new infrared sensor from the provided port.
func NewInfraredSensor(port ev3lib.EV3Port) (*ev3lib.InfraredSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), infraredSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(infraredSensorModeProximity))
	return ev3lib.NewInfraredSensorBase(&ev3InfraredSensor{sensor: sensor, currentMode: infraredSensorModeProximity}), nil
}

// Distance returns the distance measured by the sensor from 0 to 1.
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

var buttonMap = map[int][]ev3lib.BeaconButton{
	0:  {},
	1:  {ev3lib.LeftUp},
	2:  {ev3lib.LeftDown},
	3:  {ev3lib.RightUp},
	4:  {ev3lib.RightDown},
	5:  {ev3lib.LeftUp, ev3lib.RightUp},
	6:  {ev3lib.LeftUp, ev3lib.RightDown},
	7:  {ev3lib.LeftDown, ev3lib.RightUp},
	8:  {ev3lib.LeftDown, ev3lib.RightDown},
	9:  {ev3lib.Beacon},
	10: {ev3lib.LeftUp, ev3lib.LeftDown},
	11: {ev3lib.RightUp, ev3lib.RightDown},
}

// Buttons returns a slice of BeaconButton's containing all the buttons that are currently being pressed.
// Checks the buttons on the provided channel.
func (s *ev3InfraredSensor) Buttons(channel int) []ev3lib.BeaconButton {
	if s.currentMode != infraredSensorModeRemote {
		s.sensor.SetMode(string(infraredSensorModeRemote))
	}

	if channel < 0 || channel > 4 {
		fmt.Println("Channel does not exist, only 0 to 4 are accepted")
		return []ev3lib.BeaconButton{}
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
