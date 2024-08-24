package ev3lib

import (
	"log"
	"strconv"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// Color Sensor Interface                                                     //
////////////////////////////////////////////////////////////////////////////////

type ColorSensor interface {
	Ambient() float64
	Reflection() float64
	GetRGB() (float64, float64, float64)
}

////////////////////////////////////////////////////////////////////////////////
// Test Color Sensor                                                          //
////////////////////////////////////////////////////////////////////////////////

var _ ColorSensor = &TestColorSensor{}

type TestColorSensor struct{}

func (s *TestColorSensor) Ambient() float64 {
	return 0
}

func (s *TestColorSensor) Reflection() float64 {
	return 0
}

func (s *TestColorSensor) GetRGB() (float64, float64, float64) {
	return 0, 0, 0
}

////////////////////////////////////////////////////////////////////////////////
// EV3 Color Sensor                                                           //
////////////////////////////////////////////////////////////////////////////////

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

var _ ColorSensor = &EV3ColorSensor{}

// Provides access to the EV3 color sensor
type EV3ColorSensor struct {
	sensor *ev3dev.Sensor

	minReflect, maxReflect float64
	currentMode            colorSensorMode
}

// Creates a new color sensor with the provided port.
// Defaults calibration values of minReflect to 0, and maxReflect to 1.
func NewColorSensor(port EV3Port) (ColorSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), colorSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(colorSensorModeReflect))
	return &EV3ColorSensor{sensor: sensor, minReflect: 0, maxReflect: 1, currentMode: colorSensorModeReflect}, nil
}

// Returns the ambient light intensity from 0 to 1.
func (s *EV3ColorSensor) Ambient() float64 {
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
func (s *EV3ColorSensor) Reflection() float64 {
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
func (s *EV3ColorSensor) GetRGB() (float64, float64, float64) {
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
