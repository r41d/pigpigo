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
type wp_intType int

const ( // from wiringPi.h
	WP_INT_EDGE_SETUP   wp_intType = iota // Interrupts already set up
	WP_INT_EDGE_FALLING                   // Interrupt on falling edge of signal
	WP_INT_EDGE_RISING                    // Interrupt on raising edge of signal
	WP_INT_EDGE_BOTH                      // Interrupt on both edges
)

type INT interface {
	Base() wp_intType
}

func (i wp_intType) Base() wp_intType {
	return i
}
