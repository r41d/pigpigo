package gpigo //code.bitsetter.de/tk/gpigo

// pin modes – make them private so they cannot be used on other
// functions accidentially
type pinMode int

const (
	MODE_INPUT pinMode = iota
	MODE_OUTPUT
	MODE_PWM_OUTPUT
	MODE_GPIO_CLOCK
)

// export "enum" as interface
type MODE interface {
	Base() pinMode
}

// ... and wrap the enum
func (m pinMode) Base() pinMode { return m }
