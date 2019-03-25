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
	"github.com/pothosware/go-soapy-sdr/pkg/sdrerror"
)

// SetBandwidth sets the baseband filter width of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - bw: the baseband filter width in Hz
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetBandwidth(direction Direction, channel uint, bw float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setBandwidth(dev.device, C.int(direction), C.size_t(channel), C.double(bw))))
}

// GetBandwidth gets the baseband filter width of the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the baseband filter width in Hz
func (dev *SDRDevice) GetBandwidth(direction Direction, channel uint) float64 {

	return float64(C.SoapySDRDevice_getBandwidth(dev.device, C.int(direction), C.size_t(channel)))
}

// GetBandwidthRanges gets the range of possible baseband filter widths.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of bandwidth ranges in Hz
func (dev *SDRDevice) GetBandwidthRanges(direction Direction, channel uint) []SDRRange {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getBandwidthRange(dev.device, C.int(direction), C.size_t(channel), &length)
	defer RangeArrayClear(info, length)

	return RangeArray2Go(info, length)
}
