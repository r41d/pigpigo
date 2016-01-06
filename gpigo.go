// Package gpigo provides GPIO support on the raspberry pi (2).
package gpigo

var pi int // controls only one raspberry pi right now

// Initialize the GPIO interface
func Initialize() {
	pi = pigpio_start("", "")
	return
}

func Terminate() {
	pigpio_stop(pi)
}

func SetPinMode(pin uint, mode PinMode) int {
	return set_mode(pi, pin, uint(mode))
}

// Configures the pull up / pull down resistors for a given pin.
func PullUpDnControl(gpio uint, pud uint) int {
	return set_pull_up_down(pi, gpio, pud)
}

// Gets the value of a given pins input, assuming the pin is configured as an input.
func ReadPin(pin uint) int {
	return gpio_read(pi, pin)
}

// Sets the value of a pins output to the given value,
// assuming the pin is configured as an output accordingly.
func WritePin(pin uint, value PinValue) int {
	return gpio_write(pi, pin, uint(value))
}

// Delays thread execution for the given amount of milliseconds.
// This is done in a native way using the pigpio library.
func Delay(ms uint) {
	time_sleep(float64(ms) * 1000)
}

func GetPwm(user_gpio uint) int {
	return get_PWM_dutycycle(pi, user_gpio)
}

func SetPwm(pin uint, dutycycle uint) int {
	return set_PWM_dutycycle(pi, pin, dutycycle)
}

func SetPwmRange(pin uint, rrange uint) int {
	return set_PWM_range(pi, pin, rrange)
}

//__________________________________________
// TODO
// Implement the following basic functions:
//
// void pwmWrite(int pin, int value);
//   0-1024
//   not when in Sys mode (wiringPiSetup*)
//   pin1 (BCM_GPIO18, phys 12)
// analogRead(int pin); // need special board
// analogWrite(int pin, int value); // need special board
