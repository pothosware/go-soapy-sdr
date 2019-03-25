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

// ListRegisterInterfaces gets a list of available register interfaces by name.
//
// Return a list of available register interfaces
func (dev *SDRDevice) ListRegisterInterfaces() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listRegisterInterfaces(dev.device, &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// WriteRegister writes a register on the device given the interface name. This can represent a register on a soft CPU,
// FPGA, IC; the interpretation is up the implementation to decide.
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//  - value: the register value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteRegister(name string, addr uint32, value uint32) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)
	cValue := C.uint(value)

	return sdrerror.Err(int(C.SoapySDRDevice_writeRegister(dev.device, cName, cAddr, cValue)))
}

// ReadRegister reads a register on the device given the interface name.
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//
// Return an error or nil in case of success
func (dev *SDRDevice) ReadRegister(name string, addr uint32) uint32 {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)

	return uint32(C.SoapySDRDevice_readRegister(dev.device, cName, cAddr))
}

// WriteRegisters writes a memory block on the device given the interface name. This can represent a memory block on a
// soft CPU, FPGA, IC; the interpretation is up the implementation to decide.
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//  - value: the register value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteRegisters(name string, addr uint32, value []uint32) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)
	cValue := (*C.uint)(unsafe.Pointer(&value[0]))
	cLength := C.size_t(len(value))

	return sdrerror.Err(int(C.SoapySDRDevice_writeRegisters(dev.device, cName, cAddr, cValue, cLength)))
}

// ReadRegisters reads a a memory block on the device given the interface name. Pass the number of words to be read
// in via length;
//
// Params:
//  - name: the name of a available register interface
//  - addr: the register address
//
// Return an error or nil in case of success
func (dev *SDRDevice) ReadRegisters(name string, addr uint32, length uint) []uint32 {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cAddr := C.uint(addr)
	cLength := C.size_t(length)

	cValue := C.SoapySDRDevice_readRegisters(dev.device, cName, cAddr, &cLength)

	var uintTemplate *C.uint

	results := make([]uint32, int(cLength))

	// Read all the strings
	for i := 0; i < int(cLength); i++ {
		val := (*C.uint)(unsafe.Pointer(uintptr(unsafe.Pointer(cValue)) + uintptr(i)*unsafe.Sizeof(uintTemplate)))
		results[i] = uint32(*val)
	}

	return results
}
