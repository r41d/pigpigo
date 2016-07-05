// Package gpigo provides GPIO support on the raspberry pi (2).
package gpigo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib_wiringpi.sh

//#cgo CFLAGS: -I${SRCDIR}/ext/wiringPi/wiringPi
//#cgo LDFLAGS: -L${SRCDIR}/ext/wiringPi/wiringPi -lwiringPi
//#include <stdio.h>
//#include "wiringPi.h"
import "C"

/*
 * Setup Functions
 */

func wiringPiSetup() int {
	// int wiringPiSetup (void) ;
	return int(C.wiringPiSetup())
}

func wiringPiSetupGpio() int {
	// int wiringPiSetupGpio (void) ;
	return int(C.wiringPiSetupGpio())
}

func wiringPiSetupPhys() int {
	// int wiringPiSetupPhys (void) ;
	return int(C.wiringPiSetupPhys())
}

func wiringPiSetupSys() int {
	// int wiringPiSetupSys (void) ;
	return int(C.wiringPiSetupSys())
}
