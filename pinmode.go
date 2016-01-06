package pigpiogo

// pin modes – make them private so they cannot be used on other
// functions accidentially
type PinMode uint

const (
	PI_INPUT  PinMode = 0
	PI_OUTPUT         = 1
	PI_ALT0           = 4
	PI_ALT1           = 5
	PI_ALT2           = 6
	PI_ALT3           = 7
	PI_ALT4           = 3
	PI_ALT5           = 2
)
