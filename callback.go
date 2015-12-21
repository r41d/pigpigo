package gpigo // code.bitsetter.de/tk/gpigo

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
static inline void glue(int pin, int mode, void* func) {
	printf("C: setting up isr\n");
	wiringPiISR(pin, mode, func);
	printf("C: setup isr done\n");
}
*/
import "C"

import (
	"log"
	"unsafe"
)

// i will never write redundant code again ...
// i will never write redundant code again ...
// i will never write redundant code again ...
// i will never write redundant code again ...

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
	log.Println("interrupt on pin ", pin)
}

// keep stack pointer in order to avoid reference to be gc'ed
var myCallback00 = callback00
var myCallback01 = callback01
var myCallback02 = callback02
var myCallback03 = callback03
var myCallback04 = callback04
var myCallback05 = callback05
var myCallback06 = callback06
var myCallback07 = callback07
var myCallback08 = callback08
var myCallback09 = callback09

var myMap = make(map[int]InterruptHandler)

// Registers a func to be called on an Interrupt.
func OnEvent(pin int, mode INT, i InterruptHandler) (err error) {
	if i == nil {
		delete(myMap, pin)
		_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(nil))
	} else {
		myMap[pin] = i
		switch pin {
		case 0:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback00))
		case 1:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback01))
		case 2:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback02))
		case 3:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback03))
		case 4:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback04))
		case 5:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback05))
		case 6:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback06))
		case 7:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback07))
		case 8:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback08))
		case 9:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(&myCallback09))
		default:
			_, err = C.glue(C.int(pin), C.int(mode.Base()), unsafe.Pointer(nil))
		}
	}
	return
}
