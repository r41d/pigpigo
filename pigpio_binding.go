// Package pigpiogo provides GPIO support on the raspberry pi (2).
package pigpiogo

// Automatically fetch and build pigpio:
//go:generate ./ext/makelib_pigpio.sh

/*
#cgo CFLAGS: -I${SRCDIR}/ext/pigpio
#cgo LDFLAGS: -L${SRCDIR}/ext/pigpio -Wall -pthread -lpigpiod_if2 -lrt
#include <pigpiod_if2.h>
*/
import "C"

func time_time() float64 {
	//double time_time(void)
	return float64(C.time_time())
}

func time_sleep(seconds float64) {
	//void time_sleep(double seconds)
	C.time_sleep(C.double(seconds))
}

func pigpio_error(errnum int) string {
	//char* pigpio_error(int errnum)
	c := C.pigpio_error(C.int(errnum))
	return C.GoString(c)
}

func pigpiod_if_version() uint {
	//unsigned pigpiod_if_version(void)
	return uint(C.pigpiod_if_version())
}

//	pthread_t* start_thread(gpioThreadFunc_t thread_func, void* userdata);

//	void stop_thread(pthread_t* pth)

func pigpio_start(addrStr string, portStr string) int {
	//int pigpio_start(char* addrStr, char* portStr)

	var addr *C.char
	var port *C.char

	if addrStr == "" {
		addr = nil
	} else {
		addr = C.CString(addrStr)
	}

	if portStr == "" {
		port = nil
	} else {
		port = C.CString(portStr)
	}

	return int(C.pigpio_start(addr, port))
}

func pigpio_stop(pi int) {
	//void pigpio_stop(int pi)
	C.pigpio_stop(C.int(pi))
}

func set_mode(pi int, gpio uint, mode uint) int {
	//int set_mode(int pi, unsigned gpio, unsigned mode)
	return int(C.set_mode(C.int(pi), C.uint(gpio), C.uint(mode)))
}

func get_mode(pi int, gpio uint) int {
	//int get_mode(int pi, unsigned gpio)
	return int(C.get_mode(C.int(pi), C.uint(gpio)))
}

func set_pull_up_down(pi int, gpio uint, pud uint) int {
	//int set_pull_up_down(int pi, unsigned gpio, unsigned pud)
	return int(C.set_pull_up_down(C.int(pi), C.uint(gpio), C.uint(pud)))
}

func gpio_read(pi int, gpio uint) int {
	//int gpio_read(int pi, unsigned gpio)
	return int(C.gpio_read(C.int(pi), C.uint(gpio)))
}

func gpio_write(pi int, gpio uint, level uint) int {
	//int gpio_write(int pi, unsigned gpio, unsigned level)
	return int(C.gpio_write(C.int(pi), C.uint(gpio), C.uint(level)))
}

func set_PWM_dutycycle(pi int, user_gpio uint, dutycycle uint) int {
	//int set_PWM_dutycycle(int pi, unsigned user_gpio, unsigned dutycycle)
	return int(C.set_PWM_dutycycle(C.int(pi), C.uint(user_gpio), C.uint(dutycycle)))
}

func get_PWM_dutycycle(pi int, user_gpio uint) int {
	//int get_PWM_dutycycle(int pi, unsigned user_gpio)
	return int(C.get_PWM_dutycycle(C.int(pi), C.uint(user_gpio)))
}

func set_PWM_range(pi int, user_gpio uint, rang uint) int {
	//int set_PWM_range(int pi, unsigned user_gpio, unsigned range)
	return int(C.set_PWM_range(C.int(pi), C.uint(user_gpio), C.uint(rang)))
}

func get_PWM_range(pi int, user_gpio uint) int {
	//int get_PWM_range(int pi, unsigned user_gpio)
	return int(C.get_PWM_range(C.int(pi), C.uint(user_gpio)))
}

func get_PWM_real_range(pi int, user_gpio uint) int {
	//int get_PWM_real_range(int pi, unsigned user_gpio)
	return int(C.get_PWM_real_range(C.int(pi), C.uint(user_gpio)))
}

func set_PWM_frequency(pi int, user_gpio uint, frequency uint) int {
	//int set_PWM_frequency(int pi, unsigned user_gpio, unsigned frequency)
	return int(C.set_PWM_frequency(C.int(pi), C.uint(user_gpio), C.uint(frequency)))
}

func get_PWM_frequency(pi int, user_gpio uint) int {
	//int get_PWM_frequency(int pi, unsigned user_gpio)
	return int(C.get_PWM_frequency(C.int(pi), C.uint(user_gpio)))
}

func set_servo_pulsewidth(pi int, user_gpio uint, pulsewidth uint) int {
	//int set_servo_pulsewidth(int pi, unsigned user_gpio, unsigned pulsewidth)
	return int(C.set_servo_pulsewidth(C.int(pi), C.uint(user_gpio), C.uint(pulsewidth)))
}

func get_servo_pulsewidth(pi int, user_gpio uint) int {
	//int get_servo_pulsewidth(int pi, unsigned user_gpio)
	return int(C.get_servo_pulsewidth(C.int(pi), C.uint(user_gpio)))
}

func notify_open(pi int) int {
	//int notify_open(int pi)
	return int(C.notify_open(C.int(pi)))
}

func notify_begin(pi int, handle uint, bits uint32) int {
	//int notify_begin(int pi, unsigned handle, uint32_t bits)
	return int(C.notify_begin(C.int(pi), C.uint(handle), C.uint32_t(bits)))
}

func notify_pause(pi int, handle uint) int {
	//int notify_pause(int pi, unsigned handle)
	return int(C.notify_pause(C.int(pi), C.uint(handle)))
}

func notify_close(pi int, handle uint) int {
	//int notify_close(int pi, unsigned handle)
	return int(C.notify_close(C.int(pi), C.uint(handle)))
}

func set_watchdog(pi int, user_gpio uint, timeout uint) int {
	//int set_watchdog(int pi, unsigned user_gpio, unsigned timeout)
	return int(C.set_watchdog(C.int(pi), C.uint(user_gpio), C.uint(timeout)))
}

func set_glitch_filter(pi int, user_gpio uint, steady uint) int {
	//int set_glitch_filter(int pi, unsigned user_gpio, unsigned steady)
	return int(C.set_glitch_filter(C.int(pi), C.uint(user_gpio), C.uint(steady)))
}

func set_noise_filter(pi int, user_gpio uint, steady uint, active uint) int {
	//int set_noise_filter(int pi, unsigned user_gpio, unsigned steady, unsigned active)
	return int(C.set_noise_filter(C.int(pi), C.uint(user_gpio), C.uint(steady), C.uint(active)))
}

func read_bank_1(pi int) uint32 {
	//uint32_t read_bank_1(int pi)
	return uint32(C.read_bank_1(C.int(pi)))
}

func read_bank_2(pi int) uint32 {
	//uint32_t read_bank_2(int pi)
	return uint32(C.read_bank_2(C.int(pi)))
}

func clear_bank_1(pi int, bits uint32) int {
	//int clear_bank_1(int pi, uint32_t bits)
	return int(C.clear_bank_1(C.int(pi), C.uint32_t(bits)))
}

func clear_bank_2(pi int, bits uint32) int {
	//int clear_bank_2(int pi, uint32_t bits)
	return int(C.clear_bank_2(C.int(pi), C.uint32_t(bits)))
}

func set_bank_1(pi int, bits uint32) int {
	//int set_bank_1(int pi, uint32_t bits)
	return int(C.set_bank_1(C.int(pi), C.uint32_t(bits)))
}

func set_bank_2(pi int, bits uint32) int {
	//int set_bank_2(int pi, uint32_t bits)
	return int(C.set_bank_2(C.int(pi), C.uint32_t(bits)))
}

func hardware_clock(pi int, gpio uint, clkfreq uint) int {
	//int hardware_clock(int pi, unsigned gpio, unsigned clkfreq)
	return int(C.hardware_clock(C.int(pi), C.uint(gpio), C.uint(clkfreq)))
}

func hardware_PWM(pi int, gpio uint, freq uint, duty uint32) int {
	//int hardware_PWM(int pi, unsigned gpio, unsigned PWMfreq, uint32_t PWMduty)
	return int(C.hardware_PWM(C.int(pi), C.uint(gpio), C.uint(freq), C.uint32_t(duty)))
}

func get_current_tick(pi int) uint32 {
	//uint32_t get_current_tick(int pi)
	return uint32(C.get_current_tick(C.int(pi)))
}

func get_hardware_revision(pi int) uint32 {
	//uint32_t get_hardware_revision(int pi)
	return uint32(C.get_hardware_revision(C.int(pi)))
}

func get_pigpio_version(pi int) uint32 {
	//uint32_t get_pigpio_version(int pi)
	return uint32(C.get_pigpio_version(C.int(pi)))
}

// NOT YET PORTED
//	int wave_clear(int pi)
//	int wave_add_new(int pi)
//	int wave_add_generic(int pi, unsigned numPulses, gpioPulse_t* pulses)
//	int wave_add_serial(int pi, unsigned user_gpio, unsigned baud, unsigned data_bits, unsigned stop_bits, unsigned offset, unsigned numBytes, char* str)
//	int wave_create(int pi)
//	int wave_delete(int pi, unsigned wave_id)
//	int wave_send_once(int pi, unsigned wave_id)
//	int wave_send_repeat(int pi, unsigned wave_id)
//	int wave_chain(int pi, char* buf, unsigned bufSize)
//	int wave_tx_busy(int pi)
//	int wave_tx_stop(int pi)
//	int wave_get_micros(int pi)
//	int wave_get_high_micros(int pi)
//	int wave_get_max_micros(int pi)
//	int wave_get_pulses(int pi)
//	int wave_get_high_pulses(int pi)
//	int wave_get_max_pulses(int pi)
//	int wave_get_cbs(int pi)
//	int wave_get_high_cbs(int pi)
//	int wave_get_max_cbs(int pi)
//	int gpio_trigger(int pi, unsigned user_gpio, unsigned pulseLen, unsigned level)
//	int store_script(int pi, char* script)
//	int run_script(int pi, unsigned script_id, unsigned numPar, uint32_t* param)
//	int script_status(int pi, unsigned script_id, uint32_t* param)
//	int stop_script(int pi, unsigned script_id)
//	int delete_script(int pi, unsigned script_id)
//	int bb_serial_read_open(int pi, unsigned user_gpio, unsigned baud, unsigned data_bits)
//	int bb_serial_read(int pi, unsigned user_gpio, void* buf, size_t bufSize)
//	int bb_serial_read_close(int pi, unsigned user_gpio)
//	int bb_serial_invert(int pi, unsigned user_gpio, unsigned invert)
//	int i2c_open(int pi, unsigned i2c_bus, unsigned i2c_addr, unsigned i2c_flags)
//	int i2c_close(int pi, unsigned handle)
//	int i2c_write_quick(int pi, unsigned handle, unsigned bit)
//	int i2c_write_byte(int pi, unsigned handle, unsigned bVal)
//	int i2c_read_byte(int pi, unsigned handle)
//	int i2c_write_byte_data(int pi, unsigned handle, unsigned i2c_reg, unsigned bVal)
//	int i2c_write_word_data(int pi, unsigned handle, unsigned i2c_reg, unsigned wVal)
//	int i2c_read_byte_data(int pi, unsigned handle, unsigned i2c_reg)
//	int i2c_read_word_data(int pi, unsigned handle, unsigned i2c_reg)
//	int i2c_process_call(int pi, unsigned handle, unsigned i2c_reg, unsigned wVal)
//	int i2c_write_block_data(int pi, unsigned handle, unsigned i2c_reg, char* buf, unsigned count)
//	int i2c_read_block_data(int pi, unsigned handle, unsigned i2c_reg, char* buf)
//	int i2c_block_process_call(int pi, unsigned handle, unsigned i2c_reg, char* buf, unsigned count)
//	int i2c_read_i2c_block_data(int pi, unsigned handle, unsigned i2c_reg, char* buf, unsigned count)
//	int i2c_write_i2c_block_data(int pi, unsigned handle, unsigned i2c_reg, char* buf, unsigned count)
//	int i2c_read_device(int pi, unsigned handle, char* buf, unsigned count)
//	int i2c_write_device(int pi, unsigned handle, char* buf, unsigned count)
//	int i2c_zip(int pi, unsigned handle, char* inBuf, unsigned inLen, char* outBuf, unsigned outLen)
//	int bb_i2c_open(int pi, unsigned SDA, unsigned SCL, unsigned baud)
//	int bb_i2c_close(int pi, unsigned SDA)
//	int bb_i2c_zip(int pi, unsigned SDA, char* inBuf, unsigned inLen, char* outBuf, unsigned outLen)
//	int spi_open(int pi, unsigned spi_channel, unsigned baud, unsigned spi_flags)
//	int spi_close(int pi, unsigned handle)
//	int spi_read(int pi, unsigned handle, char *buf, unsigned count);
//	int spi_write(int pi, unsigned handle, char* buf, unsigned count)
//	int spi_xfer(int pi, unsigned handle, char* txBuf, char* rxBuf, unsigned count)
//	int serial_open(int pi, char* ser_tty, unsigned baud, unsigned ser_flags)
//	int serial_close(int pi, unsigned handle)
//	int serial_write_byte(int pi, unsigned handle, unsigned bVal)
//	int serial_read_byte(int pi, unsigned handle)
//	int serial_write(int pi, unsigned handle, char* buf, unsigned count)
//	int serial_read(int pi, unsigned handle, char* buf, unsigned count)
//	int serial_data_available(int pi, unsigned handle)
//	int custom_1(int pi, unsigned arg1, unsigned arg2, char* argx, unsigned argc)
//	int custom_2(int pi, unsigned arg1, char* argx, unsigned argc, char* retBuf, unsigned retMax)
//	int callback(int pi, unsigned user_gpio, unsigned edge, CB Func_t f)
//	int callback_ex (int pi, unsigned user_gpio, unsigned edge, CB FuncEx_t f, void* userdata)
//	int callback_cancel(unsigned callback_id)
//	int wait_for_edge(int pi, unsigned user_gpio, unsigned edge, double timeout)
