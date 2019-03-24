package device

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Types.h>
import "C"

// GetDriverKey returns a key that uniquely identifies the device driver.
//
// This key identifies the underlying implementation. Several variants of a product may share a driver.
func (dev *SDRDevice) GetDriverKey() (driverKey string) {

	return C.GoString(C.SoapySDRDevice_getDriverKey(dev.device))
}

// GetHardwareKey returns a key that uniquely identifies the hardware.
//
// This key should be meaningful to the user to optimize for the underlying hardware.
func (dev *SDRDevice) GetHardwareKey() (hardwareKey string) {

	return C.GoString(C.SoapySDRDevice_getHardwareKey(dev.device))
}

// GetHardwareInfo queries a dictionary of available device information.
//
// This dictionary can any number of values like vendor name, product name, revisions, serials...
// This information can be displayed to the user to help identify the instantiated device.
func (dev *SDRDevice) GetHardwareInfo() (hardwareInfo map[string]string) {

	info := (C.SoapySDRKwargs)(C.SoapySDRDevice_getHardwareInfo(dev.device))
	defer ArgsClear(info)

	return Args2Go(info)
}
