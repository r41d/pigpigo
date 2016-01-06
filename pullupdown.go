package gpigo

// The values here actually 100% compatible across the native libraries!

/*
 * interface
 */

type Pullupdown int

const (
	PUD_OFF Pullupdown = iota
	PUD_DOWN
	PUD_UP
)

/*
 * pigpio
 */

type pg_Pullupdown int

const ( // from pigpio.h
	PG_PUD_OFF pg_Pullupdown = iota
	PG_PUD_DOWN
	PG_PUD_UP
)

/*
 * wiringPi
 */

type wp_Pullupdown int

const ( // from wiringPi.h
	WP_PUD_OFF wp_Pullupdown = iota
	WP_PUD_DOWN
	WP_PUD_UP
)
