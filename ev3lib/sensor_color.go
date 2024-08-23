package ev3lib

import (
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

const colorSensorDriverName string = "lego-ev3-color"

type colorSensorMode string

const (
	colorSensorModeReflect      colorSensorMode = "COL-REFLECT"
	colorSensorModeAmbient      colorSensorMode = "COL-AMBIENT"
	colorSensorModeColor        colorSensorMode = "COL-COLOR"
	colorSensorModeRawReflected colorSensorMode = "REF-RAW"
	colorSensorModeRGB          colorSensorMode = "RGB-RAW"
	colorSensorModeCalibrate    colorSensorMode = "COL-CAL"
)

// Provides access to the EV3 color sensor
type ColorSensor struct {
	sensor *ev3dev.Sensor

	minReflect, maxReflect float64
	currentMode            colorSensorMode
}

// Creates a new color sensor with the provided port.
// Defaults calibration values of minReflect to 0, and maxReflect to 1.
func NewColorSensor(port Ev3Port) (*ColorSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), colorSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(colorSensorModeReflect))
	return &ColorSensor{sensor: sensor, minReflect: 0, maxReflect: 1, currentMode: colorSensorModeReflect}, nil
}

// Returns the ambient light intensity from 0 to 1.
func (s *ColorSensor) Ambient() float64 {
	if s.currentMode != colorSensorModeAmbient {
		s.sensor.SetMode(string(colorSensorModeAmbient))
	}

	val, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	ambient := float64(valInt) / 100

	return ambient
}

// Returns the reflected light intensity from 0 to 1.
func (s *ColorSensor) Reflection() float64 {
	if s.currentMode != colorSensorModeReflect {
		s.sensor.SetMode(string(colorSensorModeReflect))
	}

	val, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	reflected := float64(valInt) / 100

	return reflected
}

// Returns the measured color in RGB with each value from 0 to 1.
func (s *ColorSensor) GetRGB() (float64, float64, float64) {
	if s.currentMode != colorSensorModeRGB {
		s.sensor.SetMode(string(colorSensorModeRGB))
	}

	rgb := [3]float64{0, 0, 0}

	for i := 0; i < 3; i++ {
		val, err := s.sensor.Value(i)
		if err != nil {
			log.Fatal(err.Error())
		}
		valInt, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err.Error())
		}

		rgb[i] = float64(valInt) / 1020
	}

	return rgb[0], rgb[1], rgb[2]
}
