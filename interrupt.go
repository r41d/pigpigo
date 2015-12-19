package gpigo //code.bitsetter.de/tk/gpigo

type intType int

const (
	INT_SETUP   intType = iota // Interrupts already set up
	INT_FALLING                // Interrupt on falling edge of signal
	INT_RAISING                // Interrupt on raising edge of signal
	INT_BOTH                   // Interrupt on both edges
)

type INT interface {
	Base() intType
}

func (i intType) Base() intType { return i }
