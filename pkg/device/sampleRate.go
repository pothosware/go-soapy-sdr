package device

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Formats.h>
// #include <SoapySDR/Types.h>
import "C"
import (
	"pothosware/go-soapy-sdr/go-soapy-sdr/pkg/sdrerror"
)

// SetSampleRate sets the baseband sample rate of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - rate: the sample rate in samples per second
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetSampleRate(direction Direction, channel uint, rate float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setSampleRate(dev.device, C.int(direction), C.size_t(channel), C.double(rate))))
}

// GetSampleRate gets the baseband sample rate of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the sample rate in samples per second
func (dev *SDRDevice) GetSampleRate(direction Direction, channel uint) float64 {

	return float64(C.SoapySDRDevice_getSampleRate(dev.device, C.int(direction), C.size_t(channel)))
}

// GetSampleRateRange gets the range of possible baseband sample rates.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of sample rate ranges in samples per second
func (dev *SDRDevice) GetSampleRateRange(direction Direction, channel uint) []SDRRange {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getSampleRateRange(dev.device, C.int(direction), C.size_t(channel), &length)
	defer RangeArrayClear(info, length)

	return RangeArray2Go(info, length)
}
