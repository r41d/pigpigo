// Package gpigo provides GPIO support on the raspberry pi (2).
// It encapsulates the wiringPi C library (see http://wiringpi.com).
// At the current stage, only basic functions of the library are implemented.
package gpigo //code.bitsetter.de/tk/gpigo

// Automatically fetch and build wiringPi:
//go:generate ./ext/makelib.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/wiringPi/wiringPi
#cgo LDFLAGS: -L${SRCDIR}/ext/wiringPi/wiringPi -lwiringPi
#include "wiringPi.h"
*/
import "C"

// Initialize the GPIO interface
func Initialize() (err error) {
	_, err = C.wiringPiSetup()
	return
}

// Sets the configuration of a pin to be one of
//  INPUT
//  OUTPUT
//  PWM_OUTPUT
//	GPIO_CLOCK
// Note: PWN is not usable on every pin. Check gpio docs.
// void pinMode(int pin, int mode);
func SetPinMode(pin int, mode MODE) (err error) {
	_, err = C.pinMode(C.int(pin), C.int(mode.Base()))
	return
}

// Configures the pull up / pull down resistors for a given pin.
func PullUpDnControl(pin int, p PUD) (err error) {
	// void pullUpDnControl(int pin, int pud);
	_, err = C.pullUpDnControl(C.int(pin), C.int(p.Base()))
	return
}

// Gets the value of a given pins input, assuming the
// pin is configured as an input.
func ReadPin(pin int) (int, error) {
	// int digitalRead(int pin); HIGH/LOW 1/0
	v, err := C.digitalRead(C.int(pin))
	if err != nil {
		return -1, err
	}
	return int(v), err
}

// Sets the value of a pins output to the given value,
// assuming the pin is configured as an output accordingly.
func WritePin(pin int, value PIN) (err error) {
	// void digitalWrite(int pin, int value);
	_, err = C.digitalWrite(C.int(pin), C.int(value.Base()))
	return
}

// Delays thread execution for the given amount of milliseconds.
// This is done in a native way using the wiringPi library.
// TODO Check if this is a CPU hog!
func Delay(ms uint) (err error) {
	// void delay(unsigned int howLong);
	_, err = C.delay(C.uint(ms))
	return
}

// TODO
func SetPwm(pin int, val int) (err error) {
	_, err = C.pwmWrite(C.int(pin), C.int(val))
	return
}

type InterruptHandler interface {
	Handle(value int)
}

//__________________________________________
// TODO
// Implement the following basic functions:
//
// void pwmWrite(int pin, int value);
//   0-1024
//   not when in Sys mode (wiringPiSetup*)
//   pin1 (BCM_GPIO18, phys 12)
// analogRead(int pin); // need special board
// analogWrite(int pin, int value); // need special board
