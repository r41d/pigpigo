// Package gpigo provides GPIO support on the raspberry pi (2).
package gpigo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib_pigpio.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/pigpio
#cgo LDFLAGS: -L${SRCDIR}/ext/pigpio -Wall -pthread -lpigpiod_if -lrt
#include <pigpiod_if.h>
*/
import "C"

// Nothing yet...
