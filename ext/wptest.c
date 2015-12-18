#include "wiringPi/wiringPi/wiringPi.h"
#include <stdio.h>

int main(int argc, char** argv) {
	printf("setup...\n");
	wiringPiSetup();
	pinMode(0, INPUT);	
    printf("waiting...\n");
    delay(3000);
	printf("done.\n");
	return 0;
}

