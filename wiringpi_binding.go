// Package gpigo provides GPIO support on the raspberry pi (2).
package gpigo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib_wiringpi.sh

//#cgo CFLAGS: -I${SRCDIR}/ext/wiringPi/wiringPi
//#cgo LDFLAGS: -L${SRCDIR}/ext/wiringPi/wiringPi -lwiringPi
//#include <stdio.h>
//#include "wiringPi.h"
import "C"

// ... more stuff here...
