// Package sdrtime groups utility functions to convert time and ticks.
package sdrtime

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
// #include <SoapySDR/Time.h>
import "C"

// TicksToTimeNs converts a tick count into a time in nanoseconds using the tick rate.
//
// Params:
//  - ticks: a integer tick count
//  - rate: the ticks per second
//
// Return the time in nanoseconds
func TicksToTimeNs(ticks int, rate float64) int {

	return int(C.SoapySDR_ticksToTimeNs(C.longlong(ticks), C.double(rate)))
}

// TicksToTimeNs converts a time in nanoseconds into a tick count using the tick rate.
//
// Params:
//  - timeNs: time in nanoseconds
//  - rate: the ticks per second
//
// Return the integer tick count
func TimeNsToTicks(timeNs int, rate float64) int {

	return int(C.SoapySDR_timeNsToTicks(C.longlong(timeNs), C.double(rate)))
}
