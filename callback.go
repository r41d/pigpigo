package gpigo // code.bitsetter.de/tk/gpigo

// I will never write redundant code again ...
// I will never write redundant code again ...
// I will never write redundant code again ...
// I will never write redundant code again ...

/*
#cgo CFLAGS: -I${SRCDIR}/ext/wiringPi/wiringPi
#cgo LDFLAGS: -L${SRCDIR}/ext/wiringPi/wiringPi -lwiringPi
#include <stdio.h>
#include "wiringPi.h"

extern void callback00();
extern void callback01();
extern void callback02();
extern void callback03();
extern void callback04();
extern void callback05();
extern void callback06();
extern void callback07();
extern void callback08();
extern void callback09();

// only static inline functions are allowed in cgo
// when exporting go functions.

static inline void unsetintr(int pin, int mode) {
	printf("C: unsetting isr\n");
	wiringPiISR(pin, mode, NULL);
	printf("C: unsetting isr done\n");
}

static inline void setintr(int pin, int mode) {
	printf("C: setting up isr\n");
	switch(pin) {
		case 0: wiringPiISR(pin, mode, &callback00); break;
		case 1: wiringPiISR(pin, mode, &callback01); break;
		case 2: wiringPiISR(pin, mode, &callback02); break;
		case 3: wiringPiISR(pin, mode, &callback03); break;
		case 4: wiringPiISR(pin, mode, &callback04); break;
		case 5: wiringPiISR(pin, mode, &callback05); break;
		case 6: wiringPiISR(pin, mode, &callback06); break;
		case 7: wiringPiISR(pin, mode, &callback07); break;
		case 8: wiringPiISR(pin, mode, &callback08); break;
		case 9: wiringPiISR(pin, mode, &callback09); break;
		default: unsetintr(pin, mode); break;
	}
	printf("C: setup isr done\n");
}

*/
import "C"

//export callback00
func callback00() { cb(0) }

//export callback01
func callback01() { cb(1) }

//export callback02
func callback02() { cb(2) }

//export callback03
func callback03() { cb(3) }

//export callback04
func callback04() { cb(4) }

//export callback05
func callback05() { cb(5) }

//export callback06
func callback06() { cb(6) }

//export callback07
func callback07() { cb(7) }

//export callback08
func callback08() { cb(8) }

//export callback09
func callback09() { cb(9) }

func cb(pin int) {
	myMap[pin].Handle(pin)
}

var myMap = make(map[int]InterruptHandler)

// Registers a func to be called on an Interrupt.
func OnEvent(pin int, mode INT, i InterruptHandler) (err error) {
	if i == nil {
		delete(myMap, pin)
		_, err = C.unsetintr(C.int(pin), C.int(mode.Base()))
	} else {
		myMap[pin] = i
		_, err = C.setintr(C.int(pin), C.int(mode.Base()))
	}
	return
}
