package gpigo

// The values here actually 100% compatible across the native libraries!

/*
 * interface
 */
type PinValue uint

const (
	PIN_LOW  PinValue = 0
	PIN_HIGH          = 1
)

/*
 * pigpio
 */
type pg_PinValue uint

const ( // from pigpio.h
	PG_OFF   pg_PinValue = 0
	PG_ON                = 1
	PG_CLEAR             = 0
	PG_SET               = 1
	PG_LOW               = 0
	PG_HIGH              = 1
)

/*
 * wiringPi
 */
type wp_PinValue uint

const ( // from wiringPi.h
	WP_LOW  wp_PinValue = 0
	WP_HIGH             = 1
)
