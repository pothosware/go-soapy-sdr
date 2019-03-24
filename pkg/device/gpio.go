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

// ListGPIOBanks a list of available GPIO banks by name.
//
// Return a list of available GPIO banks
func (dev *SDRDevice) ListGPIOBanks() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listGPIOBanks(dev.device, &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// WriteGPIO writes the value of a GPIO bank.
//
// Params:
//  - bank: the name of an available bank
//  - value: an integer representing GPIO bits
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteGPIO(bank string, value uint32) (err sdrerror.SDRError) {

	cBank := C.CString(bank)
	defer C.free(unsafe.Pointer(cBank))

	cValue := C.uint(value)

	return sdrerror.Err(int(C.SoapySDRDevice_writeGPIO(dev.device, cBank, cValue)))
}

// WriteGPIOMasked writes the value of a GPIO bank with modification mask.
//
// Params:
//  - bank: the name of an available bank
//  - value: an integer representing GPIO bits
//  - mask: a modification mask where 1 = modify
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteGPIOMasked(bank string, value uint32, mask uint32) (err sdrerror.SDRError) {

	cBank := C.CString(bank)
	defer C.free(unsafe.Pointer(cBank))

	cValue := C.uint(value)
	cMask := C.uint(mask)

	return sdrerror.Err(int(C.SoapySDRDevice_writeGPIOMasked(dev.device, cBank, cValue, cMask)))
}

// ReadGPIO reads the value of a GPIO bank.
//
// Params:
//  - bank: the name of an available bank
//
// Return an integer representing GPIO bits
func (dev *SDRDevice) ReadGPIO(bank string) uint32 {

	cBank := C.CString(bank)
	defer C.free(unsafe.Pointer(cBank))

	return uint32(C.SoapySDRDevice_readGPIO(dev.device, cBank))
}

// WriteGPIODir writes the data direction of a GPIO bank. 1 bits represent outputs, 0 bits represent inputs.
//
// Params:
//  - bank: the name of an available bank
//  - dir: an integer representing data direction bits
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteGPIODir(bank string, dir uint32) (err sdrerror.SDRError) {

	cBank := C.CString(bank)
	defer C.free(unsafe.Pointer(cBank))

	cDir := C.uint(dir)

	return sdrerror.Err(int(C.SoapySDRDevice_writeGPIODir(dev.device, cBank, cDir)))
}

// WriteGPIODirMasked writes the data direction of a GPIO bank with modification mask.  1 bits represent outputs,
// 0 bits represent inputs.
//
// Params:
//  - bank: the name of an available bank
//  - dir: an integer representing data direction bits
//  - mask: a modification mask where 1 = modify
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteGPIODirMasked(bank string, dir uint32, mask uint32) (err sdrerror.SDRError) {

	cBank := C.CString(bank)
	defer C.free(unsafe.Pointer(cBank))

	cDir := C.uint(dir)
	cMask := C.uint(mask)

	return sdrerror.Err(int(C.SoapySDRDevice_writeGPIODirMasked(dev.device, cBank, cDir, cMask)))
}

// ReadGPIODir read the data direction of a GPIO bank. 1 bits represent outputs, 0 bits represent inputs.
//
// Params:
//  - bank: the name of an available bank
//
// Return an integer representing data direction bits
func (dev *SDRDevice) ReadGPIODir(bank string) uint32 {

	cBank := C.CString(bank)
	defer C.free(unsafe.Pointer(cBank))

	return uint32(C.SoapySDRDevice_readGPIODir(dev.device, cBank))
}
