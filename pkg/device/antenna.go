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
	"unsafe"
)

// ListAntennas gets a list of available antennas to select on a given chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel:  an available channel on the device
//
// Return a list of available antenna names
func (dev *SDRDevice) ListAntennas(direction Direction, channel uint) []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listAntennas(dev.device, C.int(direction), C.size_t(channel), &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// SetAntennas sets the selected antenna on a chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of an available antenna
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetAntennas(direction Direction, channel uint, name string) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return sdrerror.Err(int(C.SoapySDRDevice_setAntenna(dev.device, C.int(direction), C.size_t(channel), cName)))
}

// GetAntennas gets the selected antenna on a chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the name of an available antenna
func (dev *SDRDevice) GetAntennas(direction Direction, channel uint) string {

	val := (*C.char)(C.SoapySDRDevice_getAntenna(dev.device, C.int(direction), C.size_t(channel)))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
