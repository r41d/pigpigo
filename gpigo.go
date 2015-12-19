package main // code.bitsetter.de/tk/gpigo
//go:generate ./ext/makelib.sh


/*
#cgo CFLAGS: -Iext/wiringPi/wiringPi
#cgo LDFLAGS: -Lext/wiringPi/wiringPi -lwiringPi
#include "wiringPi.h"
*/
import "C"

// void pwmWrite(int pin, int value);
//   0-1024
//   not when in Sys mode (wiringPiSetup*)
//   pin1 (BCM_GPIO18, phys 12)
// int digitalRead(int pin); HIGH/LOW 1/0
// analogRead(int pin); // need special board
// analogWrite(int pin, int value); // need special board

func gpigoInit() (err error) {
	_, err = C.wiringPiSetup()
	return
}

type pinMode int
const (
	INPUT pinMode = iota
	OUTPUT
	PWM_OUTPUT
	GPIO_CLOCK
)
// void pinMode(int pin, int mode);
//  INPUT 0
//  OUTPUT 1
//  PWM_OUTPUT 2
//  GPIO_CLOCK 3
func pinMode(pin int, mode int) (err error) {
	_, err = C.pinMode(C.int(pin), C.int(mode))
	return
}

// void pullUpDnControl(int pin, int pud);
//  PUD_OFF 0
//  PUD_DOWN 1
//  PUD_UP 2
func pullUpDnControl(int pin, int pud) (err error) {
	_, err = C.pullUpDnControl(C.int(pin), C.int(pud)
	return
}


// void digitalWrite(int pin, int value); HIGH, LOW
func dWrite(pin, value int) error {
	_, err := C.digitalWrite(C.int(pin), C.int(value))
	return err
}

func delay(ms uint) (err error) {
	_, err = C.delay(C.uint(ms))
	return
}

func main() {
	gpigoInit()
	pinMode(1, 1) //output
	for i := 0; i<10; i++ {
		dWrite(1, 1)
		delay(100)
		dWrite(1, 0)
		delay(300)
	}
//	_, err = C.wpitest()
}

