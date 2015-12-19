package gpigo //code.bitsetter.de/tk/gpigo

// Pin Values
type pinValue int

const (
	PIN_LOW pinValue = iota
	PIN_HIGH
)

type PIN interface {
	Base() pinValue
}

func (v pinValue) Base() pinValue { return v }
