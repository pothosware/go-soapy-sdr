package device

// #include <SoapySDR/Device.h>
// #include <SoapySDR/Formats.h>
// #include <SoapySDR/Types.h>
import "C"
import (
	"fmt"
	"github.com/pothosware/go-soapy-sdr/pkg/sdrerror"
)

// Direction is the direction of the Data in the device TX and RX
type Direction int

const (
	// DirectionTX represents the transmit direction
	DirectionTX Direction = 0
	// DirectionRX represents the receive direction
	DirectionRX Direction = 1
)

// SDRArgInfoType is the type of data of an ArgInfo structure
type SDRArgInfoType int

// Possible data types for argument info
const (
	// ArgInfoBool represents an ArgInfo of type boolean
	ArgInfoBool SDRArgInfoType = 0
	// ArgInfoInt represents an ArgInfo of type integer
	ArgInfoInt SDRArgInfoType = 1
	// ArgInfoFloat represents an ArgInfo of type float
	ArgInfoFloat SDRArgInfoType = 2
	// ArgInfoString represents an ArgInfo of type string
	ArgInfoString SDRArgInfoType = 3
)

// StreamFlag is the type of data for defining the flags for a R/W operations on a stream. Flags can be summed (or or-ed
// individually to make the full flags.
type StreamFlag int

const (
	// StreamFlagEndBurst indicates end of burst for transmit or receive.
	// For write, end of burst if set by the caller.
	// For read, end of burst is set by the driver.
	StreamFlagEndBurst StreamFlag = 1 << 1

	// StreamFlagHasTime indicates that the time stamp is valid.
	// For write, the caller must set has time when timeNs is provided.
	// For read, the driver sets has time when timeNs is provided.
	StreamFlagHasTime StreamFlag = 1 << 2

	// StreamFlagEndAbrupt indicates that stream terminated prematurely.
	// This is the flag version of an overflow error
	// that indicates an overflow with the end samples.
	StreamFlagEndAbrupt StreamFlag = 1 << 3

	// StreamFlagOnePacket indicates transmit or receive only a single packet.
	// Applicable when the driver fragments samples into packets.
	// For write, the user sets this flag to only send a single packet.
	// For read, the user sets this flag to only receive a single packet.
	StreamFlagOnePacket StreamFlag = 1 << 4

	// StreamFlagMoreFragments indicate that this read call and the next results in a fragment.
	// Used when the implementation has an underlying packet interface.
	// The caller can use this indicator and the SOAPY_SDR_ONE_PACKET flag
	// on subsequent read stream calls to re-align with packet boundaries.
	StreamFlagMoreFragments StreamFlag = 1 << 5

	// StreamFlagWaitTrigger indicate that the stream should wait for an external trigger event.
	// This flag might be used with the flags argument in any of the
	// stream API calls. The trigger implementation is hardware-specific.
	StreamFlagWaitTrigger StreamFlag = 1 << 6
)

// SDRDevice is the opaque structure allowing to access device functions
type SDRDevice struct {
	device *C.SoapySDRDevice
}

// SDRStream is the opaque structure allowing to access stream functions
type SDRStream interface {
	// Close closes an open stream created by setupStream
	//
	// Params:
	//  - stream: the opaque pointer to a stream handle
	//
	// Return an error or nil in case of success
	Close() (err sdrerror.SDRError)

	// GetMTU gets the stream's maximum transmission unit (MTU) in number of elements.
	//
	// The MTU specifies the maximum payload transfer in a stream operation. This value can be used as a stream buffer
	// allocation size that can best optimize throughput given the underlying stream implementation.
	//
	// Return the MTU in number of stream elements (never zero)
	GetMTU() int

	// Activate activates a stream.
	//
	// Call activate to prepare a stream before using read/write(). The implementation control switches or stimulate data
	// flow.
	//
	// Params:
	//  - flags: optional flag indicators about the stream. The StreamFlagEndBurst flag can signal end on the finite burst.
	//    Not all implementations will support the full range of options. In this case, the implementation returns
	//    ErrorNotSupported.
	//  - timeNs: optional activation time in nanoseconds. The timeNs is only valid when the flags have StreamFlagHasTime.
	//  - numElems: optional element count for burst control. The numElems count can be used to request a finite burst size.
	//
	// Return an error or nil in case of success
	Activate(flags StreamFlag, timeNs int, numElems int) (err sdrerror.SDRError)

	// Deactivate deactivates a stream.
	//
	// Call deactivate when not using using read/write(). The implementation control switches or halt data flow.
	//
	// Params:
	//  - flags: optional flag indicators about the stream. Not all implementations will support the full range of options.
	//    In this case, the implementation returns ErrorNotSupported.
	//  - timeNs: optional deactivation time in nanoseconds. The timeNs is only valid when the flags have StreamFlagHasTime.
	//
	// Return an error or nil in case of success
	Deactivate(flags StreamFlag, timeNs int) (err sdrerror.SDRError)

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
	//  - chanMask to which channels this status applies
	//  - flags optional input flags and output flags
	//  - timeNs the buffer's timestamp in nanoseconds
	//  - timeoutUs the timeout in microseconds
	//
	// Return the buffer's timestamp in nanoseconds in case of success, an error otherwise
	ReadStreamStatus(chanMask []uint, flags []int, timeoutUs uint) (timeNs uint, err error)

	// GetNumDirectAccessBuffers returns how many direct access buffers can the stream provide.
	//
	// This is the number of times the user can call acquire() on a stream without making subsequent calls to
	// release(). A return value of 0 means that direct access is not supported.
	//
	// Return the number of direct access buffers or 0
	GetNumDirectAccessBuffers() uint

	// getDevice returns the internal device
	getDevice() *C.SoapySDRDevice
	// getStream returns the internal stream
	getStream() *C.SoapySDRStream
	// getNbChannels returns the number of channels used by the stream
	getNbChannels() uint
}

// SDRRange is the definition for a min/max numeric range with a step information
type SDRRange struct {
	Minimum float64
	Maximum float64
	Step    float64
}

// ToString returns a human string with the details of the range
func (r SDRRange) ToString() string {

	return fmt.Sprintf("%v->%v(/%v)", r.Minimum, r.Maximum, r.Step)
}

// SDRArgInfo is the definition for argument info
type SDRArgInfo struct {
	// Key is the key used to identify the argument (required)
	Key string

	// Value is the default value of the argument when not specified (required)
	// Numbers should use standard floating point and integer formats.
	// Boolean values should be represented as "true" and  "false".
	Value string

	// Name is the displayable name of the argument (optional, use key if empty)
	Name string

	// Description is a brief description about the argument (optional)
	Description string

	// Unit is the unit of the argument: dB, Hz, etc (optional)
	Unit string

	// Type is the data type of the argument (required)
	Type SDRArgInfoType

	// Range is the range of possible numeric values (optional)
	// When specified, the argument should be restricted to this range.
	// The range is only applicable to numeric argument types.
	Range SDRRange

	// NumOptions is the size of the options set, or 0 when not used.
	NumOptions int

	// A discrete list of possible values (optional)
	// When specified, the argument should be restricted to this options set.
	Options []string

	// A discrete list of displayable names for the enumerated options (optional)
	// When not specified, the option value itself can be used as a display name.
	OptionNames []string
}

// ToString returns a human string with the details of the ArgInfo
func (argInfo SDRArgInfo) ToString() string {

	nameStr := argInfo.Key
	if len(argInfo.Name) > 0 {
		nameStr = argInfo.Name
	}

	descriptionStr := "-"
	if len(argInfo.Description) > 0 {
		descriptionStr = argInfo.Description
	}

	unitStr := "-"
	if len(argInfo.Unit) > 0 {
		descriptionStr = argInfo.Unit
	}

	typeStr := "-"
	switch argInfo.Type {
	case ArgInfoBool:
		typeStr = "bool"
	case ArgInfoInt:
		typeStr = "int"
	case ArgInfoFloat:
		typeStr = "float"
	case ArgInfoString:
		typeStr = "string"
	}

	optionsStr := "None"
	if argInfo.NumOptions > 0 {
		optionsStr = "{"
		for i := 0; i < argInfo.NumOptions; i++ {
			optionsStr = optionsStr + fmt.Sprintf("%v->%v,", argInfo.OptionNames[i], argInfo.Options[i])
		}
		optionsStr = optionsStr[:argInfo.NumOptions-1] + "}"
	}

	return fmt.Sprintf("key: \"%v\", value: \"%v\", name: \"%v\", description: \"%v\", unit: \"%v\", type: \"%v\", range: %v, options: %v",
		argInfo.Key,
		argInfo.Value,
		nameStr,
		descriptionStr,
		unitStr,
		typeStr,
		argInfo.Range.ToString(),
		optionsStr)
}
