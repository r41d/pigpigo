// Package gpigo provides GPIO support on the raspberry pi (2).
// It encapsulates the wiringPi C library (see http://wiringpi.com).
// At the current stage, only basic functions of the library are implemented.
package gpigo //code.bitsetter.de/tk/gpigo
//go:generate ./ext/makelib.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/wiringPi/wiringPi
#cgo LDFLAGS: -L${SRCDIR}/ext/wiringPi/wiringPi -lwiringPi
#include "wiringPi.h"
// no glue code needed at this point.
// we use the librarys functions directly.
*/
import "C"

// Initialize the GPIO interface
func Initialize() (err error) {
	_, err = C.wiringPiSetup()
	return
}

// pin modes – make them private so they cannot be used on other
// functions accidentially
type pinMode int

const (
	INPUT pinMode = iota
	OUTPUT
	PWM_OUTPUT
	GPIO_CLOCK
)

// export "enum" as interface
type PinMode interface {
	Base() pinMode
}

// ... and wrap the enum
func (m pinMode) Base() pinMode { return m }

// Pin Values
type pinValue int

const (
	LOW pinValue = iota
	HIGH
)

type PinValue interface {
	Base() pinValue
}

func (v pinValue) Base() pinValue { return v }

// Sets the configuration of a pin to be one of
//  INPUT
//  OUTPUT
//  PWM_OUTPUT
//	GPIO_CLOCK
// Note: PWN is not usable on every pin. Check gpio docs.
// void pinMode(int pin, int mode);
func SetPinMode(pin int, mode PinMode) (err error) {
	_, err = C.pinMode(C.int(pin), C.int(mode.Base()))
	return
}

type pud int

const (
	OFF = iota
	DOWN
	UP
)

// void pullUpDnControl(int pin, int pud);
//  PUD_OFF 0
//  PUD_DOWN 1
//  PUD_UP 2
func pullUpDnControl(pin int, p pud) (err error) {
	_, err = C.pullUpDnControl(C.int(pin), C.int(p))
	return
}

// Sets the value of a pins output to the given value,
// assuming the pin is configured as an output accordingly.
func WritePin(pin int, value PinValue) error {
	// [void digitalWrite(int pin, int value);]
	_, err := C.digitalWrite(C.int(pin), C.int(value.Base()))
	return err
}

// Delays thread execution for the given amount of milliseconds.
// This is done in a native way using the wiringPi library.
// TODO Check if this is a CPU hog!
func Delay(ms uint) (err error) {
	_, err = C.delay(C.uint(ms))
	return
}

// void pwmWrite(int pin, int value);
//   0-1024
//   not when in Sys mode (wiringPiSetup*)
//   pin1 (BCM_GPIO18, phys 12)
// int digitalRead(int pin); HIGH/LOW 1/0
// analogRead(int pin); // need special board
// analogWrite(int pin, int value); // need special board

func TestGPIGO() {
	Initialize()
	SetPinMode(1, OUTPUT) //output
	for i := 0; i < 10; i++ {
		WritePin(1, HIGH)
		Delay(100)
		WritePin(1, LOW)
		Delay(300)
	}
	//	_, err = C.wpitest()
}
