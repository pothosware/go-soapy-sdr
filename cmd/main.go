package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lSoapySDR
import "C"
import (
	"fmt"
	"github.com/pothosware/go-soapy-sdr/pkg/device"
	"github.com/pothosware/go-soapy-sdr/pkg/version"
	"log"
)

func main() {

	// Display the version
	fmt.Printf("Soapy SDR\n")
	fmt.Printf("ABI version: %v\n", version.GetABIVersion())
	fmt.Printf("API version: %v\n", version.GetAPIVersion())
	fmt.Printf("Lib version: %v\n", version.GetLibVersion())

	// List all devices
	devices := device.Enumerate(nil)
	for i, dev := range devices {
		fmt.Printf("Found device #%v: ", i)
		for k, v := range dev {
			fmt.Printf("%v=%v, ", k, v)
		}
		fmt.Printf("\n")
	}

	if len(devices) == 0 {
		fmt.Printf("No device found!!\n")
	}

	// Take the first device
	args := map[string]string{
		"driver": devices[0]["driver"],
	}

	dev, err := device.Make(args)
	if err != nil {
		log.Panic(err)
	}

	// Display information about the device
	displayDetails(dev)

	// Apply settings
	if err := dev.SetSampleRate(device.DirectionRX, 0, 1e6); err != nil {
		log.Fatal(fmt.Printf("setSampleRate fail: error: %v\n", err))
	}
	if err := dev.SetFrequency(device.DirectionRX, 0, 912.3e6, nil); err != nil {
		log.Fatal(fmt.Printf("setFrequency fail: error: %v\n", err))
	}

	stream, err := dev.SetupSDRStreamCS8(device.DirectionRX, []uint{0}, nil)
	if err != nil {
		log.Fatal(fmt.Printf("SetupStream fail: error: %v\n", err))
	}

	if err := stream.Activate(0, 0, 0); err != nil {
		log.Fatal(fmt.Printf("Activate fail: error: %v\n", err))
	}

	fmt.Printf("Stream MTU: %v\n", stream.GetMTU())
	fmt.Printf("NumDirectAccessBuffers: %v\n", stream.GetNumDirectAccessBuffers())

	buffers := make([][]int8, 1)
	buffers[0] = make([]int8, 1024)
	flags := make([]int, 1)

	for i := 0; i < 10; i++ {
		timeNs, numElemsRead, err := stream.Read(buffers, 511, flags, 100000)
		fmt.Printf("flags=%v, numElemsRead=%v, timeNs=%v, err=%v\n", flags, numElemsRead, timeNs, err)
	}

	if err := stream.Deactivate(0, 0); err != nil {
		log.Fatal(fmt.Printf("Deactivate fail: error: %v\n", err))
	}

	if err := stream.Close(); err != nil {
		log.Fatal(fmt.Printf("Close fail: error: %v\n", err))
	}

	err = dev.Unmake()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Done\n")
}

func displayDetails(dev *device.SDRDevice) {

	// Function from identification API
	fmt.Printf("Identification / DriverKey: %v\n", dev.GetDriverKey())
	fmt.Printf("Identification / HardwareKey: %v\n", dev.GetHardwareKey())

	hardwareInfo := dev.GetHardwareInfo()
	if len(hardwareInfo) > 0 {
		for k, v := range hardwareInfo {
			fmt.Printf("Identification / HardwareInfo: {%v:%v}\n", k, v)
		}
	} else {
		fmt.Printf("Identification / HardwareInfo: [none]\n")
	}

	//
	// GPIO
	//
	banks := dev.ListGPIOBanks()
	if len(banks) > 0 {
		for i, bank := range banks {
			fmt.Printf("GPIO / Bank #%d: %v\n", i, bank)
		}
	} else {
		fmt.Printf("GPIO / Banks: [none]\n")
	}

	//
	// Settings
	//
	settings := dev.GetSettingInfo()
	if len(settings) > 0 {
		for i, setting := range settings {
			fmt.Printf("Settings / Setting #%d: %v\n", i, setting.ToString())
		}
	} else {
		fmt.Printf("Settings: [none]\n")
	}

	//
	// UARTs
	//
	uarts := dev.ListUARTs()
	if len(settings) > 0 {
		for i, uart := range uarts {
			fmt.Printf("UARTs #%d: / UART: %v\n", i, uart)
		}
	} else {
		fmt.Printf("UARTs: [none]\n")
	}

	//
	// Clocking
	//
	fmt.Printf("MasterClockRate: %v\n", dev.GetMasterClockRate())
	clockRanges := dev.GetMasterClockRates()
	if len(clockRanges) > 0 {
		for i, clockRange := range clockRanges {
			fmt.Printf("MasterClockRate range #%d: %v\n", i, clockRange)
		}
	} else {
		fmt.Printf("MasterClockRate ranges: [none]\n")
	}
	clockSources := dev.ListClockSources()
	if len(clockSources) > 0 {
		for i, clockSource := range clockSources {
			fmt.Printf("Clock source #%d: %v\n", i, clockSource)
		}
	} else {
		fmt.Printf("Clock sources: [none]\n")
	}

	//
	// Register
	//
	registers := dev.ListRegisterInterfaces()
	if len(registers) > 0 {
		for i, register := range registers {
			fmt.Printf("Register #%d: %v\n", i, register)
		}
	} else {
		fmt.Printf("Registers: [none]\n")
	}

	//
	// Device Sensor
	//
	sensors := dev.ListSensors()
	if len(sensors) > 0 {
		for i, sensor := range sensors {
			fmt.Printf("Sensor #%d: %v\n", i, sensor)
		}
	} else {
		fmt.Printf("Sensors: [none]\n")
	}

	//
	// TimeSource
	//
	timeSources := dev.ListTimeSources()
	if len(timeSources) > 0 {
		for i, timeSource := range timeSources {
			fmt.Printf("Time source #%d: %v\n", i, timeSource)
		}
	} else {
		fmt.Printf("Time sources: [none]\n")
	}
	hasHardwareTime := dev.HasHardwareTime("")
	fmt.Printf("Time source / Has hardware time: %v\n", hasHardwareTime)
	if hasHardwareTime {
		fmt.Printf("Time source / Hardware time: %v\n", dev.GetHardwareTime(""))
	}

	displayDirectionDetails(dev, device.DirectionTX)
	displayDirectionDetails(dev, device.DirectionRX)
}

func displayDirectionDetails(dev *device.SDRDevice, direction device.Direction) {

	if direction == device.DirectionTX {
		fmt.Printf("Direction TX\n")
	} else {
		fmt.Printf("Direction RX\n")
	}

	frontEndMapping := dev.GetFrontendMapping(direction)
	if len(frontEndMapping) > 0 {
		fmt.Printf("FrontendMapping: %v\n", frontEndMapping)
	} else {
		fmt.Printf("FrontendMapping: [none]\n")
	}

	numChannels := dev.GetNumChannels(direction)
	fmt.Printf("NumChannel: %v\n", numChannels)

	for channel := uint(0); channel < numChannels; channel++ {
		displayDirectionChannelDetails(dev, direction, channel)
	}
}

func displayDirectionChannelDetails(dev *device.SDRDevice, direction device.Direction, channel uint) {

	// Settings
	settings := dev.GetChannelSettingInfo(direction, channel)
	if len(settings) > 0 {
		for i, setting := range settings {
			fmt.Printf("Channel #%d / Setting #%d: / Banks: %v\n", channel, i, setting)
		}
	} else {
		fmt.Printf("Channel #%d / Settings: [none]\n", channel)
	}

	//
	// Channel
	//

	channelInfo := dev.GetChannelInfo(direction, channel)
	if len(channelInfo) > 0 {
		for k, v := range channelInfo {
			fmt.Printf("Channel #%d / ChannelInfo: {%v:%v}\n", channel, k, v)
		}
	} else {
		fmt.Printf("Channel #%d / ChannelInfo: [none]\n", channel)
	}

	fmt.Printf("Channel #%d / FullDuplex: %v\n", channel, dev.GetFullDuplex(direction, channel))

	//
	// Antenna
	//

	antennas := dev.ListAntennas(direction, channel)
	fmt.Printf("Channel #%d / NumAntennas: %v\n", channel, len(antennas))

	for i, antenna := range antennas {
		fmt.Printf("Channel #%d / Antenna #%d: %v\n", channel, i, antenna)
	}

	//
	// Bandwidth
	//

	fmt.Printf("Channel #%d / Baseband filter width Hz: %v Hz\n", channel, dev.GetBandwidth(direction, channel))

	bandwidthRanges := dev.GetBandwidthRanges(direction, channel)
	for i, bandwidthRange := range bandwidthRanges {
		fmt.Printf("Channel #%d / Baseband filter #%d: %v\n", channel, i, bandwidthRange)
	}

	//
	// Gain
	//

	fmt.Printf("Channel #%d / HasGainMode (Automatic gain possible): %v\n", channel, dev.HasGainMode(direction, channel))
	fmt.Printf("Channel #%d / GainMode (Automatic gain enabled): %v\n", channel, dev.GetGainMode(direction, channel))
	fmt.Printf("Channel #%d / Gain: %v\n", channel, dev.GetGain(direction, channel))
	fmt.Printf("Channel #%d / GainRange: %v\n", channel, dev.GetGainRange(direction, channel).ToString())

	gainElements := dev.ListGains(direction, channel)
	fmt.Printf("Channel #%d / NumGainElements: %v\n", channel, len(gainElements))

	for i, gainElement := range gainElements {
		fmt.Printf("Channel #%d / Gain Element #%d / Name: %v\n", channel, i, gainElement)
		fmt.Printf("Channel #%d / Gain Element #%d / Value: %v\n", channel, i, dev.GetGainElement(direction, channel, gainElement))
		fmt.Printf("Channel #%d / Gain Element #%d / Range: %v\n", channel, i, dev.GetGainElementRange(direction, channel, gainElement).ToString())
	}

	//
	// SampleRate
	//

	fmt.Printf("Channel #%d / Sample Rate: %v\n", channel, dev.GetSampleRate(direction, channel))
	for i, sampleRateRange := range dev.GetSampleRateRange(direction, channel) {
		fmt.Printf("Channel #%d / Sample Rate Range #%d: %v\n", channel, i, sampleRateRange.ToString())
	}

	//
	// Frequencies
	//

	fmt.Printf("Channel #%d / Frequency: %v\n", channel, dev.GetFrequency(direction, channel))
	for i, frequencyRange := range dev.GetFrequencyRange(direction, channel) {
		fmt.Printf("Channel #%d / Frequency Range #%d: %v\n", channel, i, frequencyRange.ToString())
	}

	frequencyArgsInfos := dev.GetFrequencyArgsInfo(direction, channel)
	if len(frequencyArgsInfos) > 0 {
		for i, argInfo := range frequencyArgsInfos {
			fmt.Printf("Channel #%d / Frequency ArgInfo #%d: %v\n", channel, i, argInfo.ToString())
		}
	} else {
		fmt.Printf("Channel #%d / Frequency ArgInfo: [none]\n", channel)
	}

	frequencyComponents := dev.ListFrequencies(direction, channel)
	fmt.Printf("Channel #%d / NumFrequencyComponents: %v\n", channel, len(frequencyComponents))

	for i, frequencyComponent := range frequencyComponents {
		fmt.Printf("Channel #%d / Frequency Component #%d / Name: %v\n", channel, i, frequencyComponent)
		fmt.Printf("Channel #%d / Frequency Component #%d / Frequency: %v\n", channel, i, dev.GetFrequencyComponent(direction, channel, frequencyComponent))

		frequencyRanges := dev.GetFrequencyRangeComponent(direction, channel, frequencyComponent)
		for j, frequencyRange := range frequencyRanges {
			fmt.Printf("Channel #%d / Frequency Component #%d / Frequency Range #%d: %v\n", channel, i, j, frequencyRange.ToString())
		}
	}

	//
	// Stream
	//

	fmt.Printf("Channel #%d / Stream / Formats: %v\n", channel, dev.GetStreamFormats(direction, channel))
	nativeStreamFormat, nativeStreamFullScale := dev.GetNativeStreamFormat(direction, channel)
	fmt.Printf("Channel #%d / Stream / NativeFormat: %v (fullScale: %v)\n", channel, nativeStreamFormat, nativeStreamFullScale)

	streamArgsInfos := dev.GetStreamArgsInfo(direction, channel)
	if len(streamArgsInfos) > 0 {
		for i, argInfo := range streamArgsInfos {
			fmt.Printf("Channel #%d / Stream ArgInfo #%d: %v\n", channel, i, argInfo.ToString())
		}
	} else {
		fmt.Printf("Channel #%d / Stream ArgInfo: [none]\n", channel)
	}

	//
	// Front-end correction
	//
	available := dev.HasDCOffsetMode(direction, channel)
	fmt.Printf("Channel #%d / Stream / Correction / Auto DC correction available: %v\n", channel, available)
	if available {
		fmt.Printf("Channel #%d / Stream / Correction / Auto DC correction: %v\n", channel, dev.GetDCOffsetMode(direction, channel))
	}
	available = dev.HasDCOffset(direction, channel)
	fmt.Printf("Channel #%d / Stream / Correction / DC correction available: %v\n", channel, available)
	if available {
		I, Q, err := dev.GetDCOffset(direction, channel)
		fmt.Printf("Channel #%d / Stream / Correction / DC correction I: %v, Q: %v, err :%v\n", channel, I, Q, err)
	}
	available = dev.HasIQBalance(direction, channel)
	fmt.Printf("Channel #%d / Stream / Correction / IQ Balance available: %v\n", channel, available)
	if available {
		I, Q, err := dev.GetIQBalance(direction, channel)
		fmt.Printf("Channel #%d / Stream / Correction / IQ Balance I: %v, Q: %v, err :%v\n", channel, I, Q, err)
	}
	available = dev.HasFrequencyCorrection(direction, channel)
	fmt.Printf("Channel #%d / Stream / Correction / Frequency correction available: %v\n", channel, available)
	if available {
		fmt.Printf("Channel #%d / Stream / Correction / Frequency correction: %v PPM\n", channel, dev.GetFrequencyCorrection(direction, channel))
	}

	//
	// Channel Sensor
	//
	sensors := dev.ListChannelSensors(direction, channel)
	if len(sensors) > 0 {
		for i, sensor := range sensors {
			fmt.Printf("Channel #%d / Sensor #%d: %v\n", channel, i, sensor)
		}
	} else {
		fmt.Printf("Channel #%d / Sensors: [none]\n", channel)
	}
}
