[![Travis](https://travis-ci.org/manios/sma-webbox-go.svg?branch=master)](https://travis-ci.org/manios/sma-webbox-go) [![project status badge](https://img.shields.io/badge/status-active-green.svg)](https://img.shields.io/badge/status-active-green.svg) [![Project Goreport card analysis badge](https://goreportcard.com/badge/github.com/manios/sma-webbox-go?status.svg)](https://goreportcard.com/report/github.com/manios/sma-webbox-go)&nbsp;[![GoDoc](https://godoc.org/github.com/manios/sma-webbox-go?status.svg)](https://godoc.org/github.com/manios/sma-webbox-go)

# Go API client for SMA Sunny WebBox

The data loggers Sunny WebBox and Sunny WebBox with Bluetooth continuously record all the data of a PV plant. This is then averaged over a configurable interval and cached. The data can be transmitted at regular intervals to the Sunny Portal for analysis and visualization. 

Via the Sunny WebBox and Sunny WebBox with Bluetooth RPC interface, selected data from the PV plant can be transmitted to a remote terminal by means of an RPC protocol (Remote Procedure Call protocol).  Related documents: 

* [SUNNY WEBBOX RPC User Manual v1.4](http://files.sma.de/dl/2585/SWebBoxRPC-BA-en-14.pdf) 
* [SUNNY WEBBOX RPC User Manual v1.3](http://files.sma.de/dl/4253/SWebBoxRPC-eng-BUS112713.pdf)

# Overview

- API version: 1.4.0

This is an unofficial [Go](http://golang.org/) API client which currently supports [Sunny WebBox with Bluetooth](http://files.sma.de/dl/11567/SWebBox20-BA-en-13.pdf).

# Requirements

Run:

```
go get golang.org/x/oauth2
go get github.com/manios/sma-webbox-go
```

# Usage

Currently this client supports the following procedures:

* RPC_GET_PLANT_OVERVIEW
* RPC_GET_DEVICES
* RPC_GET_PROCESS_DATA_CHANNEL
* RPC_GET_PROCESS_DATA

To run **sample code** you can have a look on ```./examples``` directory.

## Important Note

According to the official docs the interval between two queries <u>should not be less than 30 seconds</u>.

## Create a new client

To create a new client object:

```go
// Set up a new client
smaClient := sma.NewWebboxClient("http://mysolarpark:82")
```

where ```http://mysolarpark:82``` is the HTTP address and port of your Sunny WebBox in the internet. HTTP Basic Authentication is not supported yet.

## Get plant overview

To get a general overview of the solar plant you can execute the following:

```go
// Get plant overview
plantResponse, err := smaClient.GetPlantOverview()

if err != nil {
    panic(err)
}

fmt.Printf("Plant response:%+v\nPlant result:%+v\n", plantResponse, plantResponse.Result)
```


This will output a response that may look like this (it may vary depending on your solar plant setup):

```
TODO add response
```

## Get devices

To list all the devices attached to Sunny WebBox execute:

```go
// Get a list of all devices in the park
devicesResponse, err := smaClient.GetDevices()
    
if err != nil {
    panic(err)
} 

fmt.Printf("Device response:%+v\nDevice result:%+v\n", devicesResponse, devicesResponse.Result)
```

This will output a response that may look like this (it may vary depending on your solar plant setup):

```
TODO add response
```

## Get process data channels

To get the data channels offered by a specific given device :

```go
// Get the data channels offered by the given device
processDataChannels, err := smaClient.GetProcessDataChannels("0145:ad221af2")

if err != nil {
    panic(err)
} 

fmt.Printf("Process data channels response:%+v\nData channels result:%+v\n", processDataChannels, processDataChannels.Result)
```

where ```"0145:ad221af2"``` is the ```uid``` of a device returned by ```GetDevices()``` function.

This will output a response that may look like this (it may vary depending on your solar plant setup):

```
TODO add response
```
## Get process data

To return the current data of a specific device , you can use the following:

```go
// Get the current data for the given device
processData, err := smaClient.GetProcessData("0145:ad221af2")

if err != nil {
    panic(err)
} 

fmt.Printf("Process data response:%+v\nData result:%+v\n", processData, processData.Result.Devices[0])

```

# Acknowledgements

This project would never come into life without the help and inspiration from the following resources:

* [SMA-WebBox (smawb)](https://github.com/smarthomeNG/plugins/tree/master/smawb)
* [jraedler/SunnyWebBox](https://github.com/jraedler/SunnyWebBox)
* Gist [juhaautioniemi/0850bcda0d246dd7d3141308522d00a7](https://gist.github.com/juhaautioniemi/0850bcda0d246dd7d3141308522d00a7)

# Feature requests / support

This is a free time project so, any feature requests should be done as new issues.