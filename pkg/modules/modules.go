// Utility functions to deal with modules. These utility functions are made available for advanced usage.
// For most use cases, the API will automatically load modules.
package modules

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Modules.h>
import "C"
import (
	"errors"
	"unsafe"
)

// GetRootPath queries the root installation path
//
// Return the root installation path
func GetRootPath() string {

	// Do not deallocate - SOAPY internal string
	val := (*C.char)(C.SoapySDR_getRootPath())

	return C.GoString(val)
}

// ListSearchPaths gets a list of paths automatically searched by loadModules().
//
// Return a list of paths
func ListSearchPaths() []string {

	length := C.size_t(0)

	paths := C.SoapySDR_listSearchPaths(&length)
	defer C.SoapySDRStrings_clear(&paths, length)

	return stringArray2Go(paths, length)
}

// ListModules gets a list of all modules found in default path.
//
// Return a list of file paths to loadable modules
func ListModules() []string {

	length := C.size_t(0)

	paths := C.SoapySDR_listModules(&length)
	defer C.SoapySDRStrings_clear(&paths, length)

	return stringArray2Go(paths, length)
}

// ListModulesPath gets a list of all modules found in the given path
//
// Params:
//  - path: a directory on the system
//
// Return a list of file paths to loadable modules
func ListModulesPath(path string) []string {

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	length := C.size_t(0)

	paths := C.SoapySDR_listModulesPath(cPath, &length)
	defer C.SoapySDRStrings_clear(&paths, length)

	return stringArray2Go(paths, length)
}

// LoadModule gets a list of all modules found in default path.
//
// Params:
//  - path: the path to a specific module file
//
// Return an error or nil on success
func LoadModule(path string) error {

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	val := C.SoapySDR_loadModule(cPath)
	defer C.free(unsafe.Pointer(val))

	result := C.GoString(val)
	if len(result) > 0 {
		return errors.New(result)
	}

	return nil
}

// GetLoaderResult lists all registration loader errors for a given module path.
//
// The resulting dictionary contains all registry entry names provided by the specified module. The value of each entry
// is an error message string or empty on successful load.
//
// Params:
//  - path: the path to a specific module file
//
// Return a dictionary of registry names to error messages
func GetLoaderResult(path string) map[string]string {

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	val := C.SoapySDR_getLoaderResult(cPath)
	defer C.SoapySDRKwargs_clear(&val)

	return args2Go(val)
}

// GetModuleVersion gets a version string for the specified module. Modules may optionally provide version strings.
//
// Params:
//  - path: the path to a specific module file
//
// Return a version string or empty if no version provided
func GetModuleVersion(path string) string {

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	val := C.SoapySDR_getModuleVersion(cPath)
	defer C.free(unsafe.Pointer(val))

	return C.GoString(val)
}

// UnloadModule unloads a module that was loaded with loadModule().
//
// Params:
//  - path: the path to a specific module file
//
// Return an error or nil on success
func UnloadModule(path string) error {

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	val := C.SoapySDR_unloadModule(cPath)
	defer C.free(unsafe.Pointer(val))

	result := C.GoString(val)
	if len(result) > 0 {
		return errors.New(result)
	}

	return nil
}

// LoadModules loads the support modules installed on this system. This call will only actually perform the load once.
// Subsequent calls are a NOP.
func LoadModules() {

	C.SoapySDR_loadModules()
}
