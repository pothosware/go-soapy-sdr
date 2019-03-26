package device

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Formats.h>
// #include <SoapySDR/Types.h>
import "C"
import "unsafe"

// ListSensors gets a list of the available global readable sensors.
//
// Return a list of available sensor string names
func (dev *SDRDevice) ListSensors() []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listSensors(dev.device, &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// GetSensorInfo gets meta-information about a sensor.
//
// Params:
//  - key: the ID name of an available sensor
//
// Return meta-information about a sensor
func (dev *SDRDevice) GetSensorInfo(key string) SDRArgInfo {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	info := C.SoapySDRDevice_getSensorInfo(dev.device, cKey)
	defer ArgInfoClear(info)

	return ArgInfo2Go(&info)
}

// ReadSensor reads a global sensor given the name. The value returned is a string which can represent
// a boolean ("true"/"false"), an integer, or float.
//
// Params:
//  - key: the ID name of an available sensor
//
// Return the current value of the sensor
func (dev *SDRDevice) ReadSensor(key string) string {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	val := (*C.char)(C.SoapySDRDevice_readSensor(dev.device, cKey))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// ListChannelSensors gets a list of the available channel readable sensors.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of available sensor string names
func (dev *SDRDevice) ListChannelSensors(direction Direction, channel uint) []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listChannelSensors(dev.device, C.int(direction), C.size_t(channel), &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// GetChannelSensorInfo gets meta-information about a channel sensor.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - key: the ID name of an available sensor
//
// Return meta-information about a sensor
func (dev *SDRDevice) GetChannelSensorInfo(direction Direction, channel uint, key string) SDRArgInfo {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	info := C.SoapySDRDevice_getChannelSensorInfo(dev.device, C.int(direction), C.size_t(channel), cKey)
	defer ArgInfoClear(info)

	return ArgInfo2Go(&info)
}

// ReadChannelSensor reads a channel sensor given the name. The value returned is a string which can represent
// a boolean ("true"/"false"), an integer, or float.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - key: the ID name of an available sensor
//
// Return the current value of the sensor
func (dev *SDRDevice) ReadChannelSensor(direction Direction, channel uint, key string) string {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	val := (*C.char)(C.SoapySDRDevice_readChannelSensor(dev.device, C.int(direction), C.size_t(channel), cKey))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
