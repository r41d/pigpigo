package pigpiogo

type pwmMode int

const (
	PWM_MS pwmMode = iota
	PWM_BAL
)

type PWM interface {
	Base() pwmMode
}

func (pwm pwmMode) Base() pwmMode {
	return pwm
}
