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

// SetFrequency sets the center frequency of the chain.
//  - For RX, this specifies the down-conversion frequency.
//  - For TX, this specifies the up-conversion frequency.
//
// The default implementation of setFrequency() will tune the "RF"
// component as close as possible to the requested center frequency.
// Tuning inaccuracies will be compensated for with the "BB" component.
//
// The args can be used to augment the tuning algorithm.
//  - Use "OFFSET" to specify an "RF" tuning offset,
//    usually with the intention of moving the LO out of the passband.
//    The offset will be compensated for using the "BB" component.
//  - Use the name of a component for the key and a frequency in Hz
//    as the value (any format) to enforce a specific frequency.
//    The other components will be tuned with compensation
//    to achieve the specified overall frequency.
//  - Use the name of a component for the key and the value "IGNORE"
//    so that the tuning algorithm will avoid altering the component.
//  - Vendor specific implementations can also use the same args to augment
//    tuning in other ways such as specifying fractional vs integer N tuning.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - frequency: the center frequency in Hz
//  - args: optional tuner arguments
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetFrequency(direction Direction, channel uint, frequency float64, args map[string]string) (err sdrerror.SDRError) {

	cArgs, cArgsLength := go2Args(args)
	defer argsListClear(cArgs, cArgsLength)

	return sdrerror.Err(int(C.SoapySDRDevice_setFrequency(dev.device, C.int(direction), C.size_t(channel), C.double(frequency), cArgs)))
}

// SetFrequencyComponent tunes the center frequency of the specified element.
//  - For RX, this specifies the down-conversion frequency.
//  - For TX, this specifies the up-conversion frequency.
//
// Recommended names used to represent tunable components:
//  - "CORR" - freq error correction in PPM
//  - "RF" - frequency of the RF frontend
//  - "BB" - frequency of the baseband DSP
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of a tunable element
//  - frequency: the center frequency in Hz
//  - args: optional tuner arguments
//
// Return an error or nil in case of success
func (dev *SDRDevice) SetFrequencyComponent(direction Direction, channel uint, name string, frequency float64, args map[string]string) (err sdrerror.SDRError) {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cArgs, cArgsLength := go2Args(args)
	defer argsListClear(cArgs, cArgsLength)

	return sdrerror.Err(int(C.SoapySDRDevice_setFrequencyComponent(dev.device, C.int(direction), C.size_t(channel), cName, C.double(frequency), cArgs)))
}

// GetFrequency gets the overall center frequency of the chain.
//  - For RX, this specifies the down-conversion frequency.
//  - For TX, this specifies the up-conversion frequency.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the center frequency in Hz
func (dev *SDRDevice) GetFrequency(direction Direction, channel uint) float64 {

	return float64(C.SoapySDRDevice_getFrequency(dev.device, C.int(direction), C.size_t(channel)))
}

// GetFrequencyComponent gets the frequency of a tunable element in the chain.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of a tunable element
//
// Return the tunable element's frequency in Hz
func (dev *SDRDevice) GetFrequencyComponent(direction Direction, channel uint, name string) float64 {

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.SoapySDRDevice_getFrequencyComponent(dev.device, C.int(direction), C.size_t(channel), cName))
}

// ListFrequencies lists available tunable elements in the chain.
//
// Elements should be in order RF to baseband.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return a list of tunable elements by name
func (dev *SDRDevice) ListFrequencies(direction Direction, channel uint) []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_listFrequencies(dev.device, C.int(direction), C.size_t(channel), &length)
	defer stringArrayClear(info, length)

	return stringArray2Go(info, length)
}

// GetFrequencyRange gets the range of overall frequency values.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel
//
// Return a list of frequency ranges in Hz
func (dev *SDRDevice) GetFrequencyRange(direction Direction, channel uint) []SDRRange {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getFrequencyRange(dev.device, C.int(direction), C.size_t(channel), &length)
	defer rangeArrayClear(info)

	return rangeArray2Go(info, length)
}

// GetFrequencyRangeComponent gets the range of tunable values for the specified element.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//  - name: the name of a tunable element
//
// Return a list of frequency ranges in Hz
func (dev *SDRDevice) GetFrequencyRangeComponent(direction Direction, channel uint, name string) []SDRRange {

	length := C.size_t(0)

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	info := C.SoapySDRDevice_getFrequencyRangeComponent(dev.device, C.int(direction), C.size_t(channel), cName, &length)
	defer rangeArrayClear(info)

	return rangeArray2Go(info, length)
}

// GetFrequencyArgsInfo queries the argument info description for tune args.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of argument info structures
func (dev *SDRDevice) GetFrequencyArgsInfo(direction Direction, channel uint) []SDRArgInfo {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getFrequencyArgsInfo(dev.device, C.int(direction), C.size_t(channel), &length)
	defer argInfoListClear(info, length)

	return argInfoList2Go(info, length)
}
