package gpigo //code.bitsetter.de/tk/gpigo

type pullupdown int

const (
	PUD_OFF pullupdown = iota
	PUD_DOWN
	PUD_UP
)

type PUD interface {
	Base() pullupdown
}

func (p pullupdown) Base() pullupdown { return p }
