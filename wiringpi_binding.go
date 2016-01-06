// Package pigpiogo provides GPIO support on the raspberry pi (2).
package pigpiogo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib_wiringpi.sh

/*
//#cgo CFLAGS: -I${SRCDIR}/ext/wiringpi
//#cgo LDFLAGS: -L${SRCDIR}/ext/wiringpi -Wall -pthread -lwiringpi -lrt
//#include <wiringpi.h>
*/
import "C"

// ... more stuff here...
