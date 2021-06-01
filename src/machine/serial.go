package machine

// UARTConfig is a struct with which a UART (or similar object) can be
// configured. The baud rate is usually respected, but TX and RX may be ignored
// depending on the chip and the type of object.
type UARTConfig struct {
	BaudRate uint32
	TX       Pin
	RX       Pin
}

// NullSerial is a serial version of /dev/null (or null router): it drops
// everything that is written to it.
type NullSerial struct {
}

// Configure does nothing: the null serial has no configuration.
func (ns NullSerial) Configure(config UARTConfig) error {
	return nil
}

// WriteByte is a no-op: the null serial doesn't write bytes.
func (ns NullSerial) WriteByte(b byte) error {
	return nil
}
