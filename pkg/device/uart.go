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

// ListUARTs enumerate the available UART devices.
//
// Return a list of names of available UARTs
func (dev *SDRDevice) ListUARTs() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listUARTs(dev.device, &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// WriteUART writes data to a UART device.
//
// Its up to the implementation to set the baud rate, carriage return settings, flushing on newline.
//
// Params:
//  - which: the name of an available UART
//  - data: an array of byte to send (packed as a string for convenience)
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteUART(which string, data string) (err sdrerror.SDRError) {

	cWhich := C.CString(which)
	defer C.free(unsafe.Pointer(cWhich))

	cData := C.CString(data)
	defer C.free(unsafe.Pointer(cData))

	return sdrerror.Err(int(C.SoapySDRDevice_writeUART(dev.device, cWhich, cData)))
}

// ReadUART read bytes from a UART until timeout or newline.
//
// Its up to the implementation to set the baud rate, carriage return settings, flushing on newline.
//
// Params:
//  - which: the name of an available UART
//  - timeoutUs: a timeout in microseconds
//
// Return an array of byte packed as a string fdr convenience
func (dev *SDRDevice) ReadUART(which string, timeoutUs uint) string {

	cWhich := C.CString(which)
	defer C.free(unsafe.Pointer(cWhich))

	val := (*C.char)(C.SoapySDRDevice_readUART(dev.device, cWhich, C.long(timeoutUs)))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
