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

/* ******************************************************************************* */
/*                                                                                 */
/*                           SOAPY LIST OF STRINGS                                 */
/*                                                                                 */
/* ******************************************************************************* */

// stringArray2Go converts an array of C string to an array of Go String
func stringArray2Go(strings **C.char, length C.size_t) []string {

	results := make([]string, int(length))
	var charPtrTemplate *C.char

	// Read all the strings
	for i := 0; i < int(length); i++ {
		val := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(strings)) + uintptr(i)*unsafe.Sizeof(charPtrTemplate)))
		results[i] = C.GoString(*val)
	}

	return results
}

// stringArrayClear frees an array of C strings
func stringArrayClear(strings **C.char, length C.size_t) {

	ptrStrings := &strings

	C.SoapySDRStrings_clear(ptrStrings, length)
}

/* ******************************************************************************* */
/*                                                                                 */
/*                            SOAPY LIST OF RANGES                                 */
/*                                                                                 */
/* ******************************************************************************* */

// rangeArray2Go converts an array of C Range to an array of Go SDRRange
func rangeArray2Go(ranges *C.SoapySDRRange, length C.size_t) []SDRRange {

	results := make([]SDRRange, int(length))

	var rangeTemplate C.SoapySDRRange

	// Read all the ranges
	for i := 0; i < int(length); i++ {
		val := (*C.SoapySDRRange)(unsafe.Pointer(uintptr(unsafe.Pointer(ranges)) + uintptr(i)*unsafe.Sizeof(rangeTemplate)))
		results[i] = SDRRange{
			Minimum: float64(val.minimum),
			Maximum: float64(val.maximum),
			Step:    float64(val.step),
		}
	}

	return results
}

// rangeArrayClear frees an array of C Range
func rangeArrayClear(ranges *C.SoapySDRRange) {

	// Free the array
	C.free(unsafe.Pointer(ranges))
}

/* ******************************************************************************* */
/*                                                                                 */
/*                                   SOAPY ARGINFO                                 */
/*                                                                                 */
/* ******************************************************************************* */

// argInfo2Go converts the type of a C ArgInfo to Go type
func argInfo2Go(argInfo *C.SoapySDRArgInfo) SDRArgInfo {

	var argType SDRArgInfoType
	switch argInfo._type {
	case 0:
		argType = ArgInfoBool
	case 1:
		argType = ArgInfoInt
	case 2:
		argType = ArgInfoFloat
	case 3:
		argType = ArgInfoString
	}

	return SDRArgInfo{
		Key:         C.GoString(argInfo.key),
		Value:       C.GoString(argInfo.value),
		Name:        C.GoString(argInfo.name),
		Description: C.GoString(argInfo.description),
		Unit:        C.GoString(argInfo.units),
		Type:        argType,
		Range: SDRRange{
			Minimum: float64(argInfo._range.minimum),
			Maximum: float64(argInfo._range.maximum),
			Step:    float64(argInfo._range.step),
		},
		NumOptions:  int(argInfo.numOptions),
		Options:     stringArray2Go(argInfo.options, argInfo.numOptions),
		OptionNames: stringArray2Go(argInfo.optionNames, argInfo.numOptions),
	}
}

// argInfoClear frees a single C ArgInfo
func argInfoClear(argInfo C.SoapySDRArgInfo) {

	// SoapySDRArgInfo_clear take a pointer, but does not free the object itself, only its content.
	// So it is safe to use for freeing a stack allocated object
	C.SoapySDRArgInfo_clear(&argInfo)
}

/* ******************************************************************************* */
/*                                                                                 */
/*                             SOAPY ARGINFO LIST                                  */
/*                                                                                 */
/* ******************************************************************************* */

// argInfoList2Go converts an array of C ArgInfo to an array of Go SDRArgInfo
func argInfoList2Go(argInfos *C.SoapySDRArgInfo, length C.size_t) []SDRArgInfo {

	results := make([]SDRArgInfo, int(length))
	var argInfoTemplate C.SoapySDRArgInfo

	// Read all the arg infos
	for i := 0; i < int(length); i++ {
		val := (*C.SoapySDRArgInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(argInfos)) + uintptr(i)*unsafe.Sizeof(argInfoTemplate)))
		results[i] = argInfo2Go(val)
	}

	return results
}

// argInfoListClear frees a list of C ArgInfo
func argInfoListClear(args *C.SoapySDRArgInfo, length C.size_t) {
	if args != nil && int(length) > 0 {
		C.SoapySDRArgInfoList_clear(args, length)
	}
}

/* ******************************************************************************* */
/*                                                                                 */
/*                                SOAPY KWARGS                                     */
/*                                                                                 */
/* ******************************************************************************* */

// args2Go converts a single C Args to Go Arg
func args2Go(args C.SoapySDRKwargs) map[string]string {

	results := make(map[string]string, args.size)

	keys := (**C.char)(unsafe.Pointer(args.keys))
	vals := (**C.char)(unsafe.Pointer(args.vals))

	// Read all the strings
	for i := 0; i < int(args.size); i++ {
		key := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(keys)) + uintptr(i)*unsafe.Sizeof(*keys)))
		val := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(vals)) + uintptr(i)*unsafe.Sizeof(*vals)))
		results[C.GoString(*key)] = C.GoString(*val)
	}

	return results

}

// argsClear frees a single C Args
func argsClear(args C.SoapySDRKwargs) {

	// SoapySDRKwargs_clear take a pointer, but does not free the object itself, only its content.
	// So it is safe to use for freeing a stack allocated object
	C.SoapySDRKwargs_clear(&args)
}

// go2Args converts a single Go args to a C Args
func go2Args(args map[string]string) (*C.SoapySDRKwargs, C.size_t) {

	if len(args) == 0 {
		return nil, C.size_t(0)
	}

	var charPtrTemplate *C.char
	keys := (**C.char)(C.malloc(C.size_t(len(args) * int(unsafe.Sizeof(charPtrTemplate)))))
	vals := (**C.char)(C.malloc(C.size_t(len(args) * int(unsafe.Sizeof(charPtrTemplate)))))

	idx := 0
	for k, v := range args {

		key := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(keys)) + uintptr(idx)*unsafe.Sizeof(*keys)))
		val := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(vals)) + uintptr(idx)*unsafe.Sizeof(*vals)))

		*key = C.CString(k)
		*val = C.CString(v)

		idx++
	}

	var result *C.SoapySDRKwargs
	size := unsafe.Sizeof(*result)

	// Allocate the result with malloc because it could be freed by `SoapySDRKwargs_clear`
	result = (*C.SoapySDRKwargs)(C.malloc(C.size_t(size)))
	result.size = C.size_t(len(args))
	result.keys = keys
	result.vals = vals

	return result, C.size_t(1)
}

/* ******************************************************************************* */
/*                                                                                 */
/*                             SOAPY KWARGS LIST                                   */
/*                                                                                 */
/* ******************************************************************************* */

// argsList2Go converts a list of C Args to a lis of Go Arg
func argsList2Go(args *C.SoapySDRKwargs, length C.size_t) []map[string]string {

	results := make([]map[string]string, 0, length)

	if args == nil || length == 0 {
		return results
	}

	// For all args (as args is not an object, but an array...)
	for i := 0; i < int(length); i++ {

		// Get the current argument
		currentArgs := (*C.SoapySDRKwargs)(unsafe.Pointer(uintptr(unsafe.Pointer(args)) + uintptr(i)*unsafe.Sizeof(*args)))
		size := int(currentArgs.size)
		keys := (**C.char)(unsafe.Pointer(currentArgs.keys))
		vals := (**C.char)(unsafe.Pointer(currentArgs.vals))

		// make the map that will receive the args
		argsData := make(map[string]string, size)

		// Read all the strings
		for i := 0; i < size; i++ {
			key := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(keys)) + uintptr(i)*unsafe.Sizeof(*keys)))
			val := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(vals)) + uintptr(i)*unsafe.Sizeof(*vals)))
			argsData[C.GoString(*key)] = C.GoString(*val)
		}

		results = append(results, argsData)
	}

	return results
}

// argsListClear frees a list of C Args
func argsListClear(args *C.SoapySDRKwargs, length C.size_t) {
	if args != nil && int(length) > 0 {
		C.SoapySDRKwargsList_clear(args, length)
	}
}

/* ******************************************************************************* */
/*                                                                                 */
/*                             OTHER FUNCTIONS                                     */
/*                                                                                 */
/* ******************************************************************************* */

// go2SizeTList converts a list of Go uint to a list of C size_t
func go2SizeTList(integers []uint) (*C.size_t, C.size_t) {

	if len(integers) == 0 {
		return nil, C.size_t(0)
	}

	var sizeTemplate C.size_t
	results := (*C.size_t)(C.malloc(C.size_t(len(integers) * int(unsafe.Sizeof(sizeTemplate)))))

	for i, v := range integers {
		val := (*C.size_t)(unsafe.Pointer(uintptr(unsafe.Pointer(results)) + uintptr(i)*unsafe.Sizeof(sizeTemplate)))
		*val = C.size_t(v)
	}

	return results, C.size_t(len(integers))
}
