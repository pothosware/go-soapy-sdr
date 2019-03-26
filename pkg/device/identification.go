package device

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Types.h>
import "C"
import "unsafe"

// GetDriverKey returns a key that uniquely identifies the device driver.
//
// This key identifies the underlying implementation. Several variants of a product may share a driver.
func (dev *SDRDevice) GetDriverKey() (driverKey string) {

	val := (*C.char)(C.SoapySDRDevice_getDriverKey(dev.device))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetHardwareKey returns a key that uniquely identifies the hardware.
//
// This key should be meaningful to the user to optimize for the underlying hardware.
func (dev *SDRDevice) GetHardwareKey() (hardwareKey string) {

	val := (*C.char)(C.SoapySDRDevice_getHardwareKey(dev.device))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetHardwareInfo queries a dictionary of available device information.
//
// This dictionary can any number of values like vendor name, product name, revisions, serials...
// This information can be displayed to the user to help identify the instantiated device.
func (dev *SDRDevice) GetHardwareInfo() (hardwareInfo map[string]string) {

	info := (C.SoapySDRKwargs)(C.SoapySDRDevice_getHardwareInfo(dev.device))
	defer argsClear(info)

	return args2Go(info)
}
