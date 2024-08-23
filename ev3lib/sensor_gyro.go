package ev3lib

import (
	"log"
	"strconv"
	"time"

	"github.com/ev3go/ev3dev"
)

const gyroSensorDriverName string = "lego-ev3-gyro"

type gyroSensorMode string

const (
	gyroSensorModeAngle     gyroSensorMode = "GYRO-ANG"
	gyroSensorModeRate      gyroSensorMode = "GYRO-RATE"
	gyroSensorModeFAS       gyroSensorMode = "GYRO-FAS"
	gyroSensorModeAngleRate gyroSensorMode = "GYRO-G&A"
	gyroSensorModeCalibrate gyroSensorMode = "GYRO_CAL"
)

// Provides access to the EV3 gyro sensor
type GyroSensor struct {
	sensor *ev3dev.Sensor

	initAngle   float64
	inverted    int
	currentMode gyroSensorMode
}

// Creates a new gyro sensor with the provided port.
// Set inverted to true if arrow markings on the gyro are facing down.
func NewGyroSensor(port Ev3Port, inverted bool) (*GyroSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), gyroSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(gyroSensorModeAngle))
	s := &GyroSensor{sensor: sensor, inverted: 1, initAngle: 0, currentMode: gyroSensorModeAngle}
	if inverted {
		s.inverted = -1
	}

	return s, nil
}

// Returns the gyro's rotational speed in degrees per second.
// Will max out at -440 and 440.
func (s *GyroSensor) Rate() float64 {
	if s.currentMode != gyroSensorModeRate {
		s.sensor.SetMode(string(gyroSensorModeRate))
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

// Returns the current angle of the gyro in degrees.
// The angle has a max cap from -32768 to 32767 degrees, depending on the manufacturer it will either freeze or overflow.
func (s *GyroSensor) Angle() float64 {
	if s.currentMode != gyroSensorModeAngle {
		s.sensor.SetMode(string(gyroSensorModeAngle))
	}

	val, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err.Error())
	}

	return float64(valInt)*float64(s.inverted) - s.initAngle
}

// Returns both the angle and rate of the gyro, see Angle() and Rate() for more details
func (s *GyroSensor) AngleRate() (float64, float64) {
	if s.currentMode != gyroSensorModeAngleRate {
		s.sensor.SetMode(string(gyroSensorModeAngleRate))
	}

	val1, err := s.sensor.Value(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	angle, err := strconv.Atoi(val1)
	if err != nil {
		log.Fatal(err.Error())
	}

	val2, err := s.sensor.Value(1)
	if err != nil {
		log.Fatal(err.Error())
	}
	rate, err := strconv.Atoi(val2)
	if err != nil {
		log.Fatal(err.Error())
	}

	return float64(angle)*float64(s.inverted) - s.initAngle, float64(rate)
}

// Resets the current angle of the gyro.
func (s *GyroSensor) ResetAngle(angle float64) {
	s.initAngle = angle
}

// Calibrates the gyro.
// This function will block for around 200ms
// Ensure that the gyro is completely still during the calirbation.
func (s *GyroSensor) Calibrate() {
	s.sensor.SetMode(string(gyroSensorModeCalibrate))
	time.Sleep(time.Millisecond * 100)
	s.sensor.SetMode(string(gyroSensorModeAngle))
	time.Sleep(time.Millisecond * 100)
}
