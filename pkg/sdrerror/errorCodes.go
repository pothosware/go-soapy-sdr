package sdrerror

// SDRError is an error of the SDR layer
type SDRError interface {
	// Error returns the error message
	Error() string
	// SDRErrorCode returns the original error code for the SoapySDR
	SDRErrorCode() int
}

// Err build a new SDR error from an SDR error code. If the SDR error code is 0 then, there was no error and nil is
// returned
func Err(errorCode int) SDRError {
	switch errorCode {
	case 0:
		return nil
	case -1:
		return &Timeout{}
	case -2:
		return &StreamError{}
	case -3:
		return &Corruption{}
	case -4:
		return &Overflow{}
	case -5:
		return &NotSupported{}
	case -6:
		return &TimeError{}
	case -7:
		return &Underflow{}
	default:
		return &Unknown{}
	}
}

// Timeout denotes a Timeout error during a read operation
type Timeout struct {
}

// Error returns the error message
func (err *Timeout) Error() string {
	return "timeout error during read operation"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *Timeout) SDRErrorCode() int {
	return -1
}

// StreamError denotes a Stream error
type StreamError struct {
}

// Error returns the error message
func (err *StreamError) Error() string {
	return "non-specific stream errors"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *StreamError) SDRErrorCode() int {
	return -2
}

// Corruption denotes that read has data corruption. For example, the driver saw a malformed packet.
type Corruption struct {
}

// Error returns the error message
func (err *Corruption) Error() string {
	return "data corruption during read operation"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *Corruption) SDRErrorCode() int {
	return -3
}

// Overflow denotes that read has an overflow condition. For example, and internal buffer has filled.
type Overflow struct {
}

// Error returns the error message
func (err *Overflow) Error() string {
	return "overflow during read operation"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *Overflow) SDRErrorCode() int {
	return -4
}

// NotSupported denotes that requested operation or flag setting is not supported by the underlying implementation.
type NotSupported struct {
}

// Error returns the error message
func (err *NotSupported) Error() string {
	return "requested operation or flag setting is not supported"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *NotSupported) SDRErrorCode() int {
	return -5
}

// TimeError denotes that a the device encountered a stream time which was expired (late) or too early to process.
type TimeError struct {
}

// Error returns the error message
func (err *TimeError) Error() string {
	return "device encountered a stream time expired or too early"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *TimeError) SDRErrorCode() int {
	return -6
}

// Underflow denotes that a write caused an underflow condition. For example, a continuous stream was interrupted.
type Underflow struct {
}

// Error returns the error message
func (err *Underflow) Error() string {
	return "write operation caused an underflow condition"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *Underflow) SDRErrorCode() int {
	return -7
}

// Unknown denotes an unknown error. This should not happen.
type Unknown struct {
}

// Error returns the error message
func (err *Unknown) Error() string {
	return "unknown error"
}

// SDRErrorCode returns the original error code for the SoapySDR
func (err *Unknown) SDRErrorCode() int {
	return -255
}
