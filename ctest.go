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
    pinMode(0, INPUT);
    printf("waiting...\n");
    delay(3000);
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
