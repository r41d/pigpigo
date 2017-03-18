// Package gpigo provides GPIO support on the raspberry pi (2).
package pigpigo

const (
	PIGPIO int = iota
	// PIGPIOD_IF unsupported because it's decprecated
	PIGPIOD_IF2
)

// specify the lib we use
var libInUse int = -1

// Memorize the active Raspberry Pi
// Can control only one Raspberry Pi at the moment
var pi int

////////////////////////////////////////////////////////////
// INITIALIZATION
////////////////////////////////////////////////////////////

func Initialize(lib2use int) {
	if libInUse < 0 {
		switch lib2use {
		case PIGPIO:
			libInUse = lib2use
			panic("PIGPIO unsupported yet")
		case PIGPIOD_IF2:
			libInUse = lib2use
			pi = pigpio_start("", "")
			if pi < 0 {
				panic("Initializing PIGPIOD_IF2 failed! "+ string(pi))
			}
		default:
			panic("Invalid or yet unsupported Library!")
		}
	} else {
		panic("Already Initialized!")
	}
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

///////////////////////////////////////////////////////////
// BASIC PINS
///////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////
// PWM
///////////////////////////////////////////////////////////

func GetPwm(user_gpio uint) int {
	return get_PWM_dutycycle(pi, user_gpio)
}

func SetPwm(pin uint, dutycycle uint) int {
	return set_PWM_dutycycle(pi, pin, dutycycle)
}

func SetPwmRange(pin uint, rrange uint) int {
	return set_PWM_range(pi, pin, rrange)
}

///////////////////////////////////////////////////////////
// SPI
///////////////////////////////////////////////////////////

func SpiOpen(spi_channel uint, baud uint) int {
	var flags uint = 0
	return spi_open(pi, spi_channel, baud, flags)
}

func SpiClose(handle uint) int {
	return spi_close(pi, handle)
}

func SpiRead(handle uint, buf string) int {
	return spi_read(pi, handle, buf, uint(len(buf)))
}

func SpiWrite(handle uint, buf string) int {
	return spi_write(pi, handle, buf, uint(len(buf)))
}

func SpiXfer(handle uint, txBuf string, rxBuf string, count uint) int {
	return spi_xfer(pi, handle, txBuf, rxBuf, count)
}

///////////////////////////////////////////////////////////
// I2c
///////////////////////////////////////////////////////////

// ...
