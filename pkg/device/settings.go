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

// GetSettingInfo describes the allowed keys and values used for settings.
//
// Return a list of argument info structures
func (dev *SDRDevice) GetSettingInfo() []SDRArgInfo {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getSettingInfo(dev.device, &length)
	defer ArgInfoListClear(info, length)

	return ArgInfoList2Go(info, length)
}

// WriteSetting writes an arbitrary setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - key: the setting identifier
//  - value: the setting value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteSetting(key string, value string) (err sdrerror.SDRError) {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	return sdrerror.Err(int(C.SoapySDRDevice_writeSetting(dev.device, cKey, cValue)))
}

// Read an arbitrary setting on the device.
// \param device a pointer to a device instance
// \param key the setting identifier
// \return the setting value
//SOAPY_SDR_API char *SoapySDRDevice_readSetting(const SoapySDRDevice *device, const char *key);

// ReadSetting reads an arbitrary setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - key: the setting identifier
//
// Return the setting value
func (dev *SDRDevice) ReadSetting(key string) string {

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	val := (*C.char)(C.SoapySDRDevice_readSetting(dev.device, cKey))
	// TODO check if it should be un allocated
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// GetChannelSettingInfo describes the allowed keys and values used for channel settings.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of argument info structures
func (dev *SDRDevice) GetChannelSettingInfo(direction Direction, channel uint) []SDRArgInfo {

	cDirection := C.int(direction)
	cChannel := C.size_t(channel)
	length := C.size_t(0)

	info := C.SoapySDRDevice_getChannelSettingInfo(dev.device, cDirection, cChannel, &length)
	defer ArgInfoListClear(info, length)

	return ArgInfoList2Go(info, length)
}

// WriteChannelSetting writes an arbitrary channel setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - key: the setting identifier
//  - value: the setting value
//
// Return an error or nil in case of success
func (dev *SDRDevice) WriteChannelSetting(direction Direction, channel uint, key string, value string) (err sdrerror.SDRError) {

	cDirection := C.int(direction)
	cChannel := C.size_t(channel)

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	return sdrerror.Err(int(C.SoapySDRDevice_writeChannelSetting(dev.device, cDirection, cChannel, cKey, cValue)))
}

// ReadChannelSetting an arbitrary channel setting on the device.
//
// The interpretation is up the implementation.
//
// Params:
//  - key: the setting identifier
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the setting value
func (dev *SDRDevice) ReadChannelSetting(direction Direction, channel uint, key string) string {

	cDirection := C.int(direction)
	cChannel := C.size_t(channel)

	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	val := (*C.char)(C.SoapySDRDevice_readChannelSetting(dev.device, cDirection, cChannel, cKey))
	// TODO check if it should be un allocated
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}
