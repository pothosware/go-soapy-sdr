package device

//go:generate go run gen/gen_streams.go

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Formats.h>
// #include <SoapySDR/Types.h>
import "C"
import (
	"errors"
	"pothosware/go-soapy-sdr/go-soapy-sdr/pkg/sdrerror"
	"unsafe"
)

// GetStreamFormats queries a list of the available stream formats.
//
// Format:
// The first character selects the number type:
//   - "C" means complex
//   - "F" means floating point
//   - "S" means signed integer
//   - "U" means unsigned integer
// The type character is followed by the number of bits per number (complex is 2x this size per sample)
// Example format strings:
//   - "CF32" -  complex float32 (8 bytes per element)
//   - "CS16" -  complex int16 (4 bytes per element)
//   - "CS12" -  complex int12 (3 bytes per element)
//   - "CS4" -  complex int4 (1 byte per element)
//   - "S32" -  int32 (4 bytes per element)
//   - "U8" -  uint8 (1 byte per element)
//
// Params:
//  - direction the channel direction RX or TX
//  - channel an available channel on the device
//
// Return a list of allowed format strings.
func (dev *SDRDevice) GetStreamFormats(direction Direction, channel uint) []string {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getStreamFormats(dev.device, C.int(direction), C.size_t(channel), &length)
	defer StringArrayClear(info, length)

	return StringArray2Go(info, length)
}

// GetNativeStreamFormat gets the hardware's native stream format for this channel.
//
// This is the format used by the underlying transport layer, and the direct buffer access API calls (when available).
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return the native stream buffer format string and the maximum possible value
func (dev *SDRDevice) GetNativeStreamFormat(direction Direction, channel uint) (format string, fullScale float64) {

	scale := C.double(0.0)

	val := (*C.char)(C.SoapySDRDevice_getNativeStreamFormat(dev.device, C.int(direction), C.size_t(channel), &scale))
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val), float64(fullScale)
}

// GetStreamArgsInfo queries the argument info description for stream args.
//
// Params:
//  - direction: the channel direction RX or TX
//  - channel: an available channel on the device
//
// Return a list of argument info structures
func (dev *SDRDevice) GetStreamArgsInfo(direction Direction, channel uint) []SDRArgInfo {

	length := C.size_t(0)

	info := C.SoapySDRDevice_getStreamArgsInfo(dev.device, C.int(direction), C.size_t(channel), &length)
	defer ArgInfoListClear(info, length)

	return ArgInfoList2Go(info, length)
}

// ReadStreamStatus reads status information about a stream.
//
// This call is typically used on a transmit stream to report time errors, underflows, and burst completion.
//
// Client code may continually poll readStreamStatus() in a loop. Implementations of readStreamStatus() should wait in
// the call for a status change event or until the timeout expiration. When stream status is not implemented on a
// particular stream, readStreamStatus() should return SOAPY_SDR_NOT_SUPPORTED. Client code may use this indication to
// disable a polling loop.
//
// Params:
//  - stream the stream from which to retrieve the status
//  - chanMask to which channels this status applies
//  - flags optional input flags and output flags
//  - timeNs the buffer's timestamp in nanoseconds
//  - timeoutUs the timeout in microseconds
//
// Return the buffer's timestamp in nanoseconds in case of success, an error otherwise
func readStreamStatus(stream SDRStream, chanMask []uint, flags []int, timeoutUs uint) (timeNs uint, err error) {

	if uint(len(flags)) != stream.getNbChannels() {
		return 0, errors.New("the flags buffer must have the same number of chanMask as the stream")
	}

	cFlags := (*C.int)(unsafe.Pointer(&flags[0]))

	// Convert the requested chanMask to a list
	channelMasks, _ := Go2SizeTList(chanMask)
	defer C.free(unsafe.Pointer(channelMasks))

	cTimeNs := C.longlong(0)

	result := int(
		C.SoapySDRDevice_readStreamStatus(
			stream.getDevice(),
			stream.getStream(),
			channelMasks,
			cFlags,
			&cTimeNs,
			C.long(timeoutUs)))
	if result < 0 {
		return 0, sdrerror.Err(int(result))
	}

	return uint(cTimeNs), nil
}

// getNumDirectAccessBuffers returns How many direct access buffers can the stream provide.
//
// This is the number of times the user can call acquire() on a stream without making subsequent calls to
// release(). A return value of 0 means that direct access is not supported.
//
// Params:
//  - stream the stream from which to retrieve the status
//
// Return the number of direct access buffers or 0
func getNumDirectAccessBuffers(stream SDRStream) uint {

	return uint(C.SoapySDRDevice_getNumDirectAccessBuffers(stream.getDevice(), stream.getStream()))
}
