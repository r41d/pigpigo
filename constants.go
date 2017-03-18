package pigpigo

import "C"

/* gpio: 0-53 */
const PI_MIN_GPIO      =  0
const PI_MAX_GPIO      = 53

/* user_gpio: 0-31 */
const  PI_MAX_USER_GPIO  = 31

/* level: 0-1 */
type PinValue uint
const ( // from pigpio.h
	PI_OFF   PinValue = 0
	PI_ON                = 1
	PI_CLEAR             = 0
	PI_SET               = 1
	PI_LOW               = 0
	PI_HIGH              = 1
)

/* mode: 0-7 */
type PinMode uint
const ( // from pigpio.h
	PI_INPUT  PinMode = 0
	PI_OUTPUT            = 1
	PI_ALT0              = 4
	PI_ALT1              = 5
	PI_ALT2              = 6
	PI_ALT3              = 7
	PI_ALT4              = 3
	PI_ALT5              = 2
)

/* pud: 0-2 */
type PullUpDown int
const ( // from pigpio.h
	PI_PUD_OFF PullUpDown = iota
	PI_PUD_DOWN
	PI_PUD_UP
)

/* dutycycle: 0-range */
const  PI_DEFAULT_DUTYCYCLE_RANGE    =255

/* SPI */
func PI_SPI_FLAGS_BITLEN(x uint) uint { return ((x&63)<<16) }
func PI_SPI_FLAGS_RX_LSB(x uint) uint { return  ((x&1)<<15) }
func PI_SPI_FLAGS_TX_LSB(x uint) uint { return  ((x&1)<<14) }
func PI_SPI_FLAGS_3WREN(x uint) uint { return  ((x&15)<<10) }
func PI_SPI_FLAGS_3WIRE(x uint) uint { return   ((x&1)<<9) }
func PI_SPI_FLAGS_AUX_SPI(x uint) uint { return ((x&1)<<8) }
func PI_SPI_FLAGS_RESVD(x uint) uint { return   ((x&7)<<5) }
func PI_SPI_FLAGS_CSPOLS(x uint) uint { return  ((x&7)<<2) }
func PI_SPI_FLAGS_MODE(x uint) uint { return    ((x&3)) }
