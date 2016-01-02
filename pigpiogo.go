// Package pigpiogo provides GPIO support on the raspberry pi (2).
package pigpiogo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/pigpio
#cgo LDFLAGS: -L${SRCDIR}/ext/pigpio -pthread -lpigpio -lrt
#include "pigpio.h"
*/
import "C"

// Initialize the GPIO interface
func Initialize() (err error) {
	//	int gpioCfgMemAlloc(unsigned memAllocMode)
	//_, err = C.gpioCfgMemAlloc(C.uint(0))
	// int gpioInitialise(void)
	_, err = C.gpioInitialise()
	return
}

func Terminate() (err error) {
	// void gpioTerminate(void)
	_, err = C.gpioTerminate()
	return
}

func SetPinMode(pin int, mode PinMode) (err error) {
	// int gpioSetMode(unsigned gpio, unsigned mode)
	_, err = C.gpioSetMode(C.uint(pin), C.uint(mode))
	return
}

// Configures the pull up / pull down resistors for a given pin.
func PullUpDnControl(pin int, pud Pullupdown) (err error) {
	// int gpioSetPullUpDown(unsigned gpio, unsigned pud)
	_, err = C.gpioSetPullUpDown(C.uint(pin), C.uint(pud))
	return
}

// Gets the value of a given pins input, assuming the pin is configured as an input.
func ReadPin(pin int) (int, error) {
	// int gpioRead(unsigned gpio)
	v, err := C.gpioRead(C.uint(pin))
	if err != nil {
		return -1, err
	}
	return int(v), err
}

// Sets the value of a pins output to the given value,
// assuming the pin is configured as an output accordingly.
func WritePin(pin int, value PinValue) (err error) {
	// int gpioWrite(unsigned gpio, unsigned level)
	_, err = C.gpioWrite(C.uint(pin), C.uint(value))
	return
}

// Delays thread execution for the given amount of milliseconds.
// This is done in a native way using the pigpio library.
// TODO Check if this is a CPU hog!
func Delay(ms uint) (err error) {
	// uint32_t gpioDelay(uint32_t micros)
	_, err = C.gpioDelay(C.uint32_t(ms * 1000))
	return
}

func SetPwm(pin int, dutycycle int) (err error) {
	// set default pwm range always
	//	int gpioSetPWMrange(unsigned user_gpio, unsigned range)
	_, err = C.gpioSetPWMrange(C.uint(pin), 255)

	// int gpioPWM(unsigned user_gpio, unsigned dutycycle)
	_, err = C.gpioPWM(C.uint(pin), C.uint(dutycycle))
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
