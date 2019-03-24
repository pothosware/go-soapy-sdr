package device

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <stdlib.h>
// #include <stddef.h>
// #include <SoapySDR/Device.h>
// #include <SoapySDR/Formats.h>
// #include <SoapySDR/Types.h>
import "C"

// TransactSPI performs a SPI transaction and return the result.
//
// Its up to the implementation to set the clock rate, and read edge, and the write edge of the SPI core. SPI slaves
// without a readback pin will return 0.
//
// If the device contains multiple SPI masters, the address bits can encode which master.
//
// Params:
//  - addr: an address of an available SPI slave
//  - data: the SPI data, numBits-1 is first out
//  - numBits: the number of bits to clock out
//
// Return the readback data, numBits-1 is first in
func (dev *SDRDevice) TransactSPI(addr int32, data uint32, numBits uint32) uint32 {

	return uint32(C.SoapySDRDevice_transactSPI(dev.device, C.int(addr), C.uint(data), C.size_t(numBits)))
}
