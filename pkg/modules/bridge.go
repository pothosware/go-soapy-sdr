package modules

// #include <SoapySDR/Types.h>
import "C"
import "unsafe"

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
