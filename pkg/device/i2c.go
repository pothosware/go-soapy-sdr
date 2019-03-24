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

// WriteI2C writes to an available I2C slave.
//
// If the device contains multiple I2C masters, the address bits can encode which master.
//
// Params:
//  - addr: the address of the slave
//  - data: an array of bytes write out
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteI2C(addr int32, data []uint8) (err sdrerror.SDRError) {

	cAddr := C.int(addr)
	cData := (*C.char)(unsafe.Pointer(&data[0]))
	cNumBytes := C.size_t(len(data))

	return sdrerror.Err(int(C.SoapySDRDevice_writeI2C(dev.device, cAddr, cData, cNumBytes)))
}

// ReadI2C reads from an available I2C slave.
//
// If the device contains multiple I2C masters, the address bits can encode which master.
//
// Params:
//  - addr: the address of the slave
//  - numBytes: the number of bytes to read
//
// Return the bytes actually read.
func (dev *SDRDevice) ReadI2C(addr int32, numBytes uint) (data []uint8) {

	cAddr := C.int(addr)
	cNumBytes := C.size_t(len(data))

	// TODO check if need free
	cData := C.SoapySDRDevice_readI2C(dev.device, cAddr, &cNumBytes)

	data = make([]uint8, int(cNumBytes))

	for i := 0; i < int(cNumBytes); i++ {

		// Get the data from the returned array
		valPtr := (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cData)) + uintptr(i)))
		val := uint8(*valPtr)

		// Fill the slice to return
		data[i] = val
	}

	return data
}
