# Go bindings for SoapySDR

## Documentation

* https://github.com/pothosware/go-soapy-sdr/wiki

## Status

Bindings to the SoapySDR APIs are almost complete. The main missing part being the Direct buffer access API.

Due to lack of compatible hardware, some endpoints were not tested and may not work (but may work nonetheless).

## Dependencies

* Soapy SDR  v0.7.x
* Golang 1.12 with a working CGo toolchain

## Building

Simply reference `github.com/pothosware/go-soapy-sdr v[x].[y].[z]` in the require section of your project's go mod 
file.

## Layout

Go standard layout
* The directory `cmd` contains an example program displaying information about plugged SDR
* The directory `pkg` contains the binding itself

## Licensing information

Use, modification and distribution is subject to the Boost Software
License, Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
http://www.boost.org/LICENSE_1_0.txt)
