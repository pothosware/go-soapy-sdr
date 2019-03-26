package version

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <SoapySDR/Version.h>
import "C"

// Get the ABI version string that the library was built against
//
// Return the ABI version
func GetABIVersion() string {

	// Note: The string must not be freed
	return C.GoString((*C.char)(C.SoapySDR_getABIVersion()))
}

// Get the SoapySDR library API version as a string. The format of the version string is major.minor.increment
//
// Return the API version
func GetAPIVersion() string {

	// Note: The string must not be freed
	return C.GoString((*C.char)(C.SoapySDR_getAPIVersion()))
}

// Get the library version and build information string. The format of the version string is major.minor.patch-buildInfo.
// This function is commonly used to identify the software back-end to the user for command-line utilities and graphical
// applications.
//
// Return the library version
func GetLibVersion() string {

	// Note: The string must not be freed
	return C.GoString((*C.char)(C.SoapySDR_getAPIVersion()))
}
