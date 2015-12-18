package main // code.bitsetter.de/tk/gpigo
//go:generate ./ext/makelib.sh

import (
	"fmt"
	"log"
)

/*
#cgo CFLAGS: -Iext/wiringPi/wiringPi
#cgo LDFLAGS: -Lext/wiringPi/wiringPi -lwiringPi
#include <stdio.h>
#include "wiringPi.h"

void wpitest() {
	printf("setup...\n");
    wiringPiSetup();
    pinMode(1, OUTPUT);
	int i;
    for (i=0; i<10; i++) {
		digitalWrite(1, HIGH);
		delay(200);
		digitalWrite(1, LOW);
		delay(200);
	}
    printf("done.\n");
}

 int sum(int a, int b) {
   return a+b;
 }

*/
import "C"

func main() {
	s, err := C.sum(21, 88)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hi", s )
	_, err = C.wpitest()
}
