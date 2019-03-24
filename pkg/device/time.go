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

// ListTimeSources gets the list of available time sources.
//
// Return a list of time source names
func (dev *SDRDevice) ListTimeSources() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listTimeSources(dev.device, &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// SetTimeSource set the time source on the device.
//
// Params:
//  - source: the name of a time source
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetTimeSource(source string) (err sdrerror.SDRError) {

	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))

	return sdrerror.Err(int(C.SoapySDRDevice_setTimeSource(dev.device, cSource)))
}

// GetTimeSource gets the time source of the device.
//
// Return the name of a time source
func (dev *SDRDevice) GetTimeSource() string {

	val := (*C.char)(C.SoapySDRDevice_getTimeSource(dev.device))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// HasHardwareTime checks if the device have a hardware clock
//
// Params:
//  - what: optional argument
//
// Return true if the hardware clock exists
func (dev *SDRDevice) HasHardwareTime(what string) bool {

	cWhat := C.CString(what)
	defer C.free(unsafe.Pointer(cWhat))

	return bool(C.SoapySDRDevice_hasHardwareTime(dev.device, cWhat))
}

// GetHardwareTime reads the time from the hardware clock on the device.
//
// Params:
//  - what: optional argument. The what argument can refer to a specific time counter.
//
// Return the time in nanoseconds
func (dev *SDRDevice) GetHardwareTime(what string) uint {

	cWhat := C.CString(what)
	defer C.free(unsafe.Pointer(cWhat))

	return uint(C.SoapySDRDevice_getHardwareTime(dev.device, cWhat))
}

// SetHardwareTime writes the time to the hardware clock on the device.
//
// Params:
//  - timeNs: time in nanoseconds
//  - what: optional argument. The what argument can refer to a specific time counter.
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetHardwareTime(timeNs uint, what string) (err sdrerror.SDRError) {

	cWhat := C.CString(what)
	defer C.free(unsafe.Pointer(cWhat))

	cTimeNs := C.longlong(timeNs)

	return sdrerror.Err(int(C.SoapySDRDevice_setHardwareTime(dev.device, cTimeNs, cWhat)))
}
