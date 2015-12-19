package gpigo // code.bitsetter.de/tk/gpigo
//go:generate ./ext/makelib.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/wiringPi/wiringPi
#cgo LDFLAGS: -L${SRCDIR}/ext/wiringPi/wiringPi -lwiringPi
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

func Initialize() (err error) {
	_, err = C.wiringPiSetup()
	return
}

// pin modes – make them private so they cannot be used on other
// functions accidentially
type pMode int

const (
	INPUT pMode = iota
	OUTPUT
	PWM_OUTPUT
	GPIO_CLOCK
)

// export "enum" as interface
type PMode interface {
	PMode() pMode
}

// ... and wrap the enum
func (m pMode) PMode() pMode { return m }

// void pinMode(int pin, int mode);
//  INPUT 0
//  OUTPUT 1
//  PWM_OUTPUT 2
//  GPIO_CLOCK 3
func PinMode(pin int, mode PMode) (err error) {
	_, err = C.pinMode(C.int(pin), C.int(mode.PMode()))
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

// void digitalWrite(int pin, int value); HIGH, LOW
func dWrite(pin, value int) error {
	_, err := C.digitalWrite(C.int(pin), C.int(value))
	return err
}

func delay(ms uint) (err error) {
	_, err = C.delay(C.uint(ms))
	return
}

func TestGPIGO() {
	Initialize()
	PinMode(1, OUTPUT) //output
	for i := 0; i < 10; i++ {
		dWrite(1, 1)
		delay(100)
		dWrite(1, 0)
		delay(300)
	}
	//	_, err = C.wpitest()
}
