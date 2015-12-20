// ugly ugly glue-around:
// wiringPi cannot give pin numbers (or any other information) into callback handlers, 
// but this information is needed in order to keep interrupt handling simple on the go side
// (the user of gpigo should NOT have to implement C callbacks).
// 
// I do not want to patch wiringPi, so the only (quick and dirty) solution is, to implement
// some glue-handlers for every possible pin-interrupt and forward into go in a clean way:
//
// TODO: make this file go generatable!

// keep a map of possible interrupt handlers:
// (there are 64 possible io pins defined in wiringPi, but not used)
void* interruptMap[64];

void handleInterrupt(int pin) {
  //...
}

void handleInterrupt_00(void) { handleInterrupt(0); }
void handleInterrupt_01(void) { handleInterrupt(1); }
void handleInterrupt_02(void) { handleInterrupt(2); }
void handleInterrupt_03(void) { handleInterrupt(3); }
void handleInterrupt_04(void) { handleInterrupt(4); }
void handleInterrupt_05(void) { handleInterrupt(5); }
void handleInterrupt_06(void) { handleInterrupt(6); }
void handleInterrupt_07(void) { handleInterrupt(7); }
void handleInterrupt_08(void) { handleInterrupt(8); }
void handleInterrupt_09(void) { handleInterrupt(9); }

void initInterruptMap() {
	interruptMap[0] = (void *)&handleInterrupt_00;
	interruptMap[1] = (void *)&handleInterrupt_01;
	interruptMap[2] = (void *)&handleInterrupt_02;
	interruptMap[3] = (void *)&handleInterrupt_03;
	interruptMap[4] = (void *)&handleInterrupt_04;
	interruptMap[5] = (void *)&handleInterrupt_05;
	interruptMap[6] = (void *)&handleInterrupt_06;
	interruptMap[7] = (void *)&handleInterrupt_07;
	interruptMap[8] = (void *)&handleInterrupt_08;
	interruptMap[9] = (void *)&handleInterrupt_09;
}


