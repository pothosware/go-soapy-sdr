package device

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Formats.h>
// #include <SoapySDR/Types.h>
import "C"
import "pothosware/go-soapy-sdr/go-soapy-sdr/pkg/sdrerror"

// HasDCOffsetMode returns if the device support automatic DC offset corrections
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return true if the device has automatic DC offset corrections, false otherwise
func (dev *SDRDevice) HasDCOffsetMode(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_hasDCOffsetMode(dev.device, C.int(direction), C.size_t(channel)))
}

// SetDCOffsetMode sets the automatic DC offset corrections mode.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//  - automatic: true for automatic offset correction
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetDCOffsetMode(direction Direction, channel uint, automatic bool) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setDCOffsetMode(dev.device, C.int(direction), C.size_t(channel), C.bool(automatic))))
}

// GetDCOffsetMode gets the automatic DC offset corrections mode.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return true for automatic offset correction
func (dev *SDRDevice) GetDCOffsetMode(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_hasDCOffsetMode(dev.device, C.int(direction), C.size_t(channel)))
}

// HasDCOffset returns if the device support frontend DC offset correction
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return true if the device supports frontend DC offset correction, false otherwise
func (dev *SDRDevice) HasDCOffset(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_hasDCOffset(dev.device, C.int(direction), C.size_t(channel)))
}

// SetDCOffset the frontend DC offset correction.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//  - offsetI: the relative correction (1.0 max)
//  - offsetQ: the relative correction (1.0 max)
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetDCOffset(direction Direction, channel uint, offsetI float64, offsetQ float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setDCOffset(dev.device, C.int(direction), C.size_t(channel), C.double(offsetI), C.double(offsetQ))))
}

// GetDCOffset gets frontend DC offset correction.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return offsetI and offsetQ the relative correction (1.0 max) and an optional error
func (dev *SDRDevice) GetDCOffset(direction Direction, channel uint) (offsetI float64, offsetQ float64, err sdrerror.SDRError) {

	cOffsetI := C.double(0)
	cOffsetQ := C.double(0)

	result := int(C.SoapySDRDevice_getDCOffset(dev.device, C.int(direction), C.size_t(channel), &cOffsetI, &cOffsetQ))

	if result < 0 {
		return 0.0, 0.0, sdrerror.Err(result)
	}

	return float64(cOffsetI), float64(cOffsetQ), nil
}

// HasIQBalance returns if the device support frontend IQ balance correction
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return true if the device supports frontend IQ balance correction, false otherwise
func (dev *SDRDevice) HasIQBalance(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_hasIQBalance(dev.device, C.int(direction), C.size_t(channel)))
}

// SetIQBalance sets the frontend IQ balance correction.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//  - balanceI: the relative correction (1.0 max)
//  - balanceQ: the relative correction (1.0 max)
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetIQBalance(direction Direction, channel uint, balanceI float64, balanceQ float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setIQBalance(dev.device, C.int(direction), C.size_t(channel), C.double(balanceI), C.double(balanceQ))))
}

// GetIQBalance gets the IQ balance correction.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return balanceI and balanceQ the relative correction (1.0 max) and an optional error
func (dev *SDRDevice) GetIQBalance(direction Direction, channel uint) (balanceI float64, balanceQ float64, err sdrerror.SDRError) {

	cBalanceI := C.double(0)
	cBalanceQ := C.double(0)

	result := int(C.SoapySDRDevice_getIQBalance(dev.device, C.int(direction), C.size_t(channel), &cBalanceI, &cBalanceQ))

	if result < 0 {
		return 0.0, 0.0, sdrerror.Err(result)
	}

	return float64(cBalanceI), float64(cBalanceQ), nil
}

// HasFrequencyCorrection returns if the device support frontend frequency correction
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return true if the device supports frontend frequency correction, false otherwise
func (dev *SDRDevice) HasFrequencyCorrection(direction Direction, channel uint) bool {

	return bool(C.SoapySDRDevice_hasFrequencyCorrection(dev.device, C.int(direction), C.size_t(channel)))
}

// SetFrequencyCorrection fine tunes the frontend frequency correction.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//  - value: the correction in PPM
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetFrequencyCorrection(direction Direction, channel uint, value float64) (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_setFrequencyCorrection(dev.device, C.int(direction), C.size_t(channel), C.double(value))))
}

// GetFrequencyCorrection gets the frontend frequency correction value.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return the correction value in PPM and an optional error
func (dev *SDRDevice) GetFrequencyCorrection(direction Direction, channel uint) (value float64) {

	return float64(C.SoapySDRDevice_getFrequencyCorrection(dev.device, C.int(direction), C.size_t(channel)))
}
