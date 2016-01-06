package gpigo

/*
 * interface
 */

// This configuration is wiringPi-specific

/*
 * pigpio
 */

// Nothing to put here

/*
 * wiringPi
 */
type wp_pwmMode int

const (
	WP_PWM_MODE_MS wp_pwmMode = iota
	WP_PWM_MODE_BAL
)

type PWM interface {
	Base() wp_pwmMode
}

func (pwm wp_pwmMode) Base() wp_pwmMode {
	return pwm
}
