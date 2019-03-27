package sdrlogger

/*

#include <stdio.h>
#include <SoapySDR/Logger.h>

// C code can call exported Go functions with their explicit name. But if a C-program wants a function pointer, a
// gateway function has to be written. This is because we can't take the address of a Go function and give that to
// C-code since the cgo tool will generate a stub in C that should be called. The following example shows how to
// integrate with C code wanting a function pointer of a give type.
// In logger.go:
//   - forward declaration of logHandlerBridge_cgo, so it can be used as a parameter to SoapySDR_registerLogHandler
//   - export of function logHandlerBridge, that is receiving the log data ultimately
// In cfuncs.go:
//   - implementation of logHandlerBridge_cgo, that is calling logHandlerBridge

// The gateway function, written in C. this function be used as a call back by SoapySDR_registerLogHandler. The
// Gateway function is simply calling the Go function(logHandlerBridge) with the same parameters
void logHandlerBridge_cgo(const SoapySDRLogLevel logLevel, const char *message)
{
	// printf("C.logHandlerBridge_cgo(): called with arg = %d, str = %s\n", logLevel, message);

	// Declare locally the function exported from logger.go
	void logHandlerBridge(const SoapySDRLogLevel, const char *);

	// Call the function in logger.go
	return logHandlerBridge(logLevel, message);
}
*/
import "C"
