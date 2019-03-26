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
	"errors"
	"github.com/pothosware/go-soapy-sdr/pkg/sdrerror"
	"unsafe"
)

// LastStatus returns the last status code after a Device API call.
//
// The status code is cleared on entry to each Device call. When an device API call throws, the C bindings catch
// the exception, and set a non-zero last status code. Use LastStatus() to determine success/failure for
// Device calls without integer status return codes.
func LastStatus() int {

	return int(C.SoapySDRDevice_lastStatus())
}

// LastError returns the last error message after a device call fails.
//
// When an device API call throws, the C bindings catch the exception, store its message in thread-safe storage,
// and return a non-zero status code to indicate failure. Use lastError() to access the exception's error message.
func LastError() string {

	// Do not free as it is internal string of Soapy
	return C.GoString(C.SoapySDRDevice_lastError())
}

// Enumerate returns a list of available devices on the system.
//
// Params:
//  - args: device construction key/value argument filters, for example {"driver":"hackrf"}. Can be set to nil if no
// filter is needed
//
// Return a list of information, each unique to a device
func Enumerate(args map[string]string) []map[string]string {

	length := C.size_t(0)

	cArgs, cArgsLength := go2Args(args)
	defer argsListClear(cArgs, cArgsLength)

	enumerateData := C.SoapySDRDevice_enumerate(cArgs, &length)
	defer argsListClear(enumerateData, length)

	return argsList2Go(enumerateData, length)
}

// EnumerateStrArgs returns a list of available devices on the system.
//
// Params:
//  - args: a markup string of key/value argument filters. Markup format for args: "keyA=valA, keyB=valB". Can be set to
// nil if no filter is needed
//
// Return a list of information, each unique to a device
func EnumerateStrArgs(args string) []map[string]string {

	length := C.size_t(0)

	cArgs := C.CString(args)
	defer C.free(unsafe.Pointer(cArgs))

	enumerateData := C.SoapySDRDevice_enumerateStrArgs(cArgs, &length)
	defer argsListClear(enumerateData, length)

	return argsList2Go(enumerateData, length)
}

// Make makes a new Device object given device construction args.
//
// The device pointer will be stored in a table so subsequent calls with the same arguments will produce the same
// device. For every call to make, there should be a matched call to unmake.
//
// Params:
//  - args: device construction key/value argument map
//
// Return a pointer to a new Device object or null for error
func Make(args map[string]string) (device *SDRDevice, err error) {

	cArgs, cArgsLength := go2Args(args)
	defer argsListClear(cArgs, cArgsLength)

	dev := C.SoapySDRDevice_make(cArgs)
	if dev == nil {
		return nil, errors.New(LastError())
	}

	return &SDRDevice{
		device: dev,
	}, nil
}

// MakeStrArgs makes a new Device object given device construction args.
//
// The device pointer will be stored in a table so subsequent calls with the same arguments will produce the same
// device. For every call to make, there should be a matched call to unmake.
//
// Params:
//  - args: a markup string of key/value arguments
//
// Return a pointer to a new Device object or null for error
func MakeStrArgs(args string) (device *SDRDevice, err error) {

	cArgs := C.CString(args)
	defer C.free(unsafe.Pointer(cArgs))

	dev := C.SoapySDRDevice_makeStrArgs(cArgs)
	if dev == nil {
		return nil, errors.New(LastError())
	}

	return &SDRDevice{
		device: dev,
	}, nil
}

// Unmake unmakes or releases a device object handle.
//
// Params:
//  - device: a pointer to a device object
//
// Return an error or nil in case of success
func (dev *SDRDevice) Unmake() (err sdrerror.SDRError) {

	return sdrerror.Err(int(C.SoapySDRDevice_unmake(dev.device)))
}

// MakeList creates a list of devices from a list of construction arguments.
//
// This is a convenience call to parallelize device construction,
// and is fundamentally a parallel for loop of make(Kwargs).
//
// Params:
//  - argsList: a list of device arguments per each device
//
// Return a list of device pointers per each specified argument
func MakeList(argsList []map[string]string) (devices []*SDRDevice, err error) {

	type opened struct {
		id     int
		device *SDRDevice
		err    error
	}

	messages := make(chan opened)
	defer close(messages)

	// Open all devices in parallel
	for i, args := range argsList {
		go func(id int, args map[string]string) {
			dev, err := Make(args)
			messages <- opened{id: id, device: dev, err: err}
		}(i, args)
	}

	devices = make([]*SDRDevice, len(argsList))

	// Receive the results
	for i := 0; i < len(argsList); i++ {
		result := <-messages
		if result.err != nil {
			// In case of an error, keep the error
			err = result.err
		} else {
			// Otherwise, keep the result
			devices[result.id] = result.device
		}
	}

	// In case of error, close all devices opened
	if err != nil {
		for _, dev := range devices {
			if dev != nil {
				// Ignore the error as in any case situation is already bad...
				if errUnmake := dev.Unmake(); errUnmake != nil {
					err = errUnmake
				}
			}
		}
		return nil, err
	}

	return devices, nil
}

// UnmakeList unmakes or releases a list of device handles.
//
// This is a convenience call to parallelize device destruction,
// and is fundamentally a parallel for loop of unmake(Device *).
//
// Params:
//  - devices: a list of pointer to sdr devices
//
// Return an error or nil in case of success
func UnmakeList(devices []*SDRDevice) (err sdrerror.SDRError) {

	type closed struct {
		err sdrerror.SDRError
	}

	messages := make(chan closed)
	defer close(messages)

	// Close all devices in parallel
	for _, dev := range devices {
		go func(dev *SDRDevice) {
			err := dev.Unmake()
			messages <- closed{err: err}
		}(dev)
	}

	// Receive all the results of close operations
	for i := 0; i < len(devices); i++ {
		result := <-messages
		if result.err != nil {
			err = result.err
		}
	}

	return err
}
