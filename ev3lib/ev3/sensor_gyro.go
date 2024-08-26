//go:build linux && arm

package ev3

import (
	"github.com/Alanlu217/ev3lib/ev3lib"
	"log"
	"strconv"
	"time"

	"github.com/ev3go/ev3dev"
)

////////////////////////////////////////////////////////////////////////////////
// EV3 Gyro Sensor                                                            //
////////////////////////////////////////////////////////////////////////////////

const gyroSensorDriverName string = "lego-ev3-gyro"

type gyroSensorMode string

const (
	gyroSensorModeAngle     gyroSensorMode = "GYRO-ANG"
	gyroSensorModeRate      gyroSensorMode = "GYRO-RATE"
	gyroSensorModeAngleRate gyroSensorMode = "GYRO-G&A"
	gyroSensorModeCalibrate gyroSensorMode = "GYRO-CAL"
)

var _ ev3lib.GyroSensor = &ev3GyroSensor{}

// Provides access to the EV3 gyro sensor
type ev3GyroSensor struct {
	sensor *ev3dev.Sensor

	initAngle   float64
	inverted    int
	currentMode gyroSensorMode
}

// NewGyroSensor creates a new gyro sensor with the provided port.
// Set inverted to true if arrow markings on the gyro are facing down.
func NewGyroSensor(port ev3lib.EV3Port, inverted bool) (ev3lib.GyroSensor, error) {
	sensor, err := ev3dev.SensorFor(string(port), gyroSensorDriverName)
	if err != nil {
		return nil, err
	}

	sensor.SetMode(string(gyroSensorModeAngle))
	s := &ev3GyroSensor{sensor: sensor, inverted: 1, initAngle: 0, currentMode: gyroSensorModeAngle}
	if inverted {
		s.inverted = -1
	}

	return s, nil
}

// Rate returns the gyro's rotational speed in degrees per second.
// Will max out at -440 and 440.
func (s *ev3GyroSensor) Rate() float64 {
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

// Angle returns the current angle of the gyro in degrees.
// The angle has a max cap from -32768 to 32767 degrees, depending on the manufacturer it will either freeze or overflow.
func (s *ev3GyroSensor) Angle() float64 {
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

// AngleRate returns both the angle and rate of the gyro, see Angle() and Rate() for more details
func (s *ev3GyroSensor) AngleRate() (float64, float64) {
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

// ResetAngle resets the current angle of the gyro.
func (s *ev3GyroSensor) ResetAngle(angle float64) {
	s.initAngle = angle
}

// Calibrate calibrates the gyro.
// This function will block for around 200ms
// Ensure that the gyro is completely still during the calibration.
func (s *ev3GyroSensor) Calibrate() {
	s.sensor.SetMode(string(gyroSensorModeCalibrate))
	time.Sleep(time.Millisecond * 100)
	s.sensor.SetMode(string(gyroSensorModeAngle))
	time.Sleep(time.Millisecond * 100)
}
