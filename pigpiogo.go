// Package pigpiogo provides GPIO support on the raspberry pi (2).
package pigpiogo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/pigpio
#cgo LDFLAGS: -L${SRCDIR}/ext/pigpio -Wall -pthread -lpigpiod_if2 -lrt
#include <pigpiod_if2.h>
*/
import "C"

var pi C.int

// Initialize the GPIO interface
func Initialize() (err error) {
	//int pigpio_start(char *addrStr, char *portStr)
	pi, err = C.pigpio_start(nil, nil)
	return
}

func Terminate() (err error) {
	// void pigpio_stop(int pi)
	_, err = C.pigpio_stop(pi)
	return
}

func SetPinMode(pin int, mode PinMode) (err error) {
	// int set_mode(int pi, unsigned gpio, unsigned mode)
	_, err = C.set_mode(pi, C.uint(pin), C.uint(mode))
	return
}

// Configures the pull up / pull down resistors for a given pin.
//func PullUpDnControl(pin int, pud Pullupdown) (err error) {
//	// int gpioSetPullUpDown(unsigned gpio, unsigned pud)
//	_, err = C.gpioSetPullUpDown(C.uint(pin), C.uint(pud))
//	return
//}

// Gets the value of a given pins input, assuming the pin is configured as an input.
func ReadPin(pin int) (int, error) {
	// int gpio_read(int pi, unsigned gpio)
	v, err := C.gpio_read(pi, C.uint(pin))
	if err != nil {
		return -1, err
	}
	return int(v), err
}

// Sets the value of a pins output to the given value,
// assuming the pin is configured as an output accordingly.
func WritePin(pin int, value PinValue) (err error) {
	// int gpio_write(int pi, unsigned gpio, unsigned level)
	_, err = C.gpio_write(pi, C.uint(pin), C.uint(value))
	return
}

// Delays thread execution for the given amount of milliseconds.
// This is done in a native way using the pigpio library.
// TODO Check if this is a CPU hog!
func Delay(ms uint) (err error) {
	// void time_sleep(double seconds)
	_, err = C.time_sleep(C.double(float32(ms) * 1000))
	return
}

func SetPwm(pin int, dutycycle int) (err error) {
	// int set_PWM_dutycycle(int pi, unsigned user_gpio, unsigned dutycycle)
	_, err = C.set_PWM_dutycycle(pi, C.uint(pin), C.uint(dutycycle))
	return
}

func SetPwmRange(pin int, rrange int) (err error) {
	//int set_PWM_range(int pi, unsigned user_gpio, unsigned range)
	_, err = C.set_PWM_range(pi, C.uint(pin), C.uint(rrange))
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
