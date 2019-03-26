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
	"unsafe"
)

// SetFrontendMapping sets the frontend mapping of available DSP units to RF frontends.
//
// This mapping controls channel mapping and channel availability.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//  - mapping: a vendor-specific mapping string
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetFrontendMapping(direction Direction, mapping string) (err sdrerror.SDRError) {

	cMapping := C.CString(mapping)
	defer C.free(unsafe.Pointer(cMapping))

	return sdrerror.Err(int(C.SoapySDRDevice_setFrontendMapping(dev.device, C.int(direction), cMapping)))
}

// GetFrontendMapping gets the mapping configuration string.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//
// Return the vendor-specific mapping string
func (dev *SDRDevice) GetFrontendMapping(direction Direction) string {

	val := (*C.char)(C.SoapySDRDevice_getFrontendMapping(dev.device, C.int(direction)))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetNumChannels gets a number of channels given the streaming direction.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//
// Return the number of channels
func (dev *SDRDevice) GetNumChannels(direction Direction) uint {

	return uint(C.SoapySDRDevice_getNumChannels(dev.device, C.int(direction)))
}

// GetChannelInfo gets channel info given the streaming direction.
//
// Params:
//  - direction: the channel direction DirectionRX or DIRECTION_TX
//  - channel: the channel number to get info for
//
// Return channel information
func (dev *SDRDevice) GetChannelInfo(direction Direction, channel uint) map[string]string {

	info := C.SoapySDRDevice_getChannelInfo(dev.device, C.int(direction), C.size_t(channel))
	defer argsClear(info)

	return args2Go(info)
}

// GetFullDuplex finds out if the specified channel is full or half duplex.
//
// Params:
//  - direction the channel direction DirectionRX or DIRECTION_TX
//  - channel an available channel on the device
//
// Return true for full duplex, false for half duplex
func (dev *SDRDevice) GetFullDuplex(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_getFullDuplex(dev.device, C.int(direction), C.size_t(channel)))
}
