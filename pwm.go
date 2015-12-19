package gpigo //code.bitsetter.de/tk/gpigo

type pwmMode int

const (
	PWM_MS pwmMode = iota
	PWM_BAL
)

type PWM interface {
	Base() pwmMode
}

func (m pwmMode) Base() pwmMode { return m }
