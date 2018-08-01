package main

import (
	"fmt"
	"time"

	sma "gitlab.com/manios/sma-webbox-go"
)

func main() {

	// Set up a new client
	smaClient := sma.NewWebboxClient("http://mysolarpark:82")

	// Get plant overview
	plantResponse, err := smaClient.GetPlantOverview()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Plant response:%+v\nPlant result:%+v\n", plantResponse, plantResponse.Result)
	}

	pauseFor(10)

	// Get a list of all devices in the park
	devicesResponse, err := smaClient.GetDevices()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Device response:%+v\nDevice result:%+v\n", devicesResponse, devicesResponse.Result)
	}

	pauseFor(10)

	// Get the data channels offered by the given device
	processDataChannels, err := smaClient.GetProcessDataChannels("0145:ad221af2")

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Process data channels response:%+v\nData channels result:%+v\n", processDataChannels, processDataChannels.Result)
	}

	pauseFor(10)

	// Get the current data for the given device
	processData, err := smaClient.GetProcessData("0145:ad221af2")

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Process data response:%+v\nData result:%+v\n", processData, processData.Result.Devices[0])
	}

}

func pauseFor(seconds time.Duration) {

	sleepSeconds := seconds * time.Second
	fmt.Printf("Sleep for %d seconds\n", seconds)
	time.Sleep(sleepSeconds)
}
