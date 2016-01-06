package gpigo

/*
 * interface
 */
type PinMode uint

const (
	PIN_INPUT  PinMode = 0
	PIN_OUTPUT         = 1
)

/*
 * pigpio
 */
type pg_PinMode uint

const ( // from pigpio.h
	PG_INPUT  pg_PinMode = 0
	PG_OUTPUT            = 1
	PG_ALT0              = 4
	PG_ALT1              = 5
	PG_ALT2              = 6
	PG_ALT3              = 7
	PG_ALT4              = 3
	PG_ALT5              = 2
)

/*
 * wiringPi
 */
type wp_PinMode uint

const ( // from wiringPi.h
	WP_INPUT            wp_PinMode = 0
	WP_OUTPUT                      = 1
	WP_PWM_OUTPUT                  = 2
	WP_GPIO_CLOCK                  = 3
	WP_SOFT_PWM_OUTPUT             = 4
	WP_SOFT_TONE_OUTPUT            = 5
	WP_PWM_TONE_OUTPUT             = 6
)
