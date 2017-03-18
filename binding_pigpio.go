// Package gpigo provides GPIO support on the raspberry pi (2).
package pigpigo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib_pigpio.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/pigpio
#cgo LDFLAGS: -L${SRCDIR}/ext/pigpio -Wall -pthread -lpigpio -lrt
#include <pigpio.h>
*/
import "C"

// Nothing implemented yet...
