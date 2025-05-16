[![Build](https://github.com/manios/sma-webbox-go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/manios/sma-webbox-go/actions/workflows/go.yml) [![project status badge](https://img.shields.io/badge/status-active-green.svg)](https://img.shields.io/badge/status-active-green.svg) [![Project Goreport card analysis badge](https://goreportcard.com/badge/github.com/manios/sma-webbox-go?status.svg)](https://goreportcard.com/report/github.com/manios/sma-webbox-go)&nbsp;[![GoDoc](https://godoc.org/github.com/manios/sma-webbox-go?status.svg)](https://godoc.org/github.com/manios/sma-webbox-go)

# Go API client for SMA Sunny WebBox

The data loggers Sunny WebBox and Sunny WebBox with Bluetooth continuously record all the data of a PV plant. This is then averaged over a configurable interval and cached. The data can be transmitted at regular intervals to the Sunny Portal for analysis and visualization. 

Via the Sunny WebBox and Sunny WebBox with Bluetooth RPC interface, selected data from the PV plant can be transmitted to a remote terminal by means of an RPC protocol (Remote Procedure Call protocol).  Related documents: 

* [SUNNY WEBBOX RPC User Manual v1.4](https://files.sma.de/downloads/SWebBoxRPC-BA-en-14.pdf) 
* [SUNNY WEBBOX RPC User Manual v1.3](https://files.sma.de/downloads/SWebBoxRPC-eng-BUS112713.pdf)

# Overview

- API version: 1.4.0

This is an unofficial [Go](http://golang.org/) API client which currently supports [Sunny WebBox with Bluetooth](https://files.sma.de/downloads/SWebBox20-BA-en-13.pdf).

https://files.sma.de/downloads/SWebBoxRPC-BA-en-14.pdf

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
Plant response:{Result:0xc42010e260 Format:JSON Proc:GetPlantOverview Version:1.0 Id:1536393036 Error_:}
Plant result:&{Overview:[{Meta:GriPwr Name:Power Unit:W Value:20278} {Meta:GriEgyTdy Name:Day yield Unit:kWh Value:44.455} {Meta:GriEgyTot Name:Total yield Unit:kWh Value:659.318} {Meta:OpStt Name:Condition Unit: Value:Ok, Ok, Ok} {Meta:Msg Name:Message Unit: Value:}]}
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
Device response:{Result:0xc420182120 Format:JSON Proc:GetDevices Version:1.0 Id:1536393047 Error_:}
Device result:&{Devices:[{Key:0042:a1bd74d3 Name:SN: 5182157032} {Key:0042:a1bd8a1e Name:SN: 5182163426} {Key:0042:a1bda169 Name:SN: 5182169449} {Key:0042:a1bdb610 Name:SN: 5182169552} {Key:0042:a1bdb616 Name:SN: 2110169558} {Key:0080:7dcb7484 Name:SN: 5182485636} {Key:0088:00006874 Name:SN: 26740}] TotalDevicesReturned:7}
```

## Get process data channels

To get the data channels offered by a specific given device :

```go
// Get the data channels offered by the given device
processDataChannels, err := smaClient.GetProcessDataChannels("0042:a1bd74d3")

if err != nil {
    panic(err)
} 

fmt.Printf("Process data channels response:%+v\nData channels result:%+v\n", processDataChannels, processDataChannels.Result)
```

where ```"0145:ad221af2"``` is the ```uid``` of a device returned by ```GetDevices()``` function.

This will output a response that may look like this (it may vary depending on your solar plant setup):

```
Process data channels response:{Result:map[0145:ad221af2:[Ipv Upv-Ist Fac Pac Riso h-On h-Total E-Total Netz-Ein]] Format:JSON Proc:GetProcessDataChannels Version:1.0 Id:1536393199 Error_:}
Data channels result:map[0145:ad221af2:[Ipv Upv-Ist Fac Pac Riso h-On h-Total E-Total Netz-Ein]
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

This will output a response that may look like this (it may vary depending on your solar plant setup):

```
Process data response:{Result:0xc4201282e0 Format:JSON Proc:GetProcessData Version:1.0 Id:1536393257 Error_:}
Data result:{Channels:[{Meta:Ipv Name:DC current input Unit:A Value:5.367} {Meta:Upv-Ist Name:DC voltage input Unit:V Value:531.43} {Meta:WindVel m/s Name:Wind speed Unit:m/s Value:0} {Meta:TmpAmb C Name:Ambient temperature Unit:°C Value:0} {Meta:IntSolIrr Name:Insolation Unit:W/m^2 Value:0} {Meta:Fac Name:Grid frequency Unit:Hz Value:50.01} {Meta:Iac-Ist Name:Grid current Unit:A Value:0} {Meta:Pac Name:Power Unit:W Value:3574} {Meta:Riso Name:Insulation resistance Unit:Ohm Value:3000000} {Meta:TmpMdul C Name:Module temperature Unit:°C Value:0} {Meta:h-On Name:Feed-in time Unit:h Value:30138.18} {Meta:h-Total Name:Operating time Unit:h Value:30535.05} {Meta:E-Total Name:Total yield Unit:kWh Value:173777.824} {Meta:Netz-Ein Name:Number of grid connections Unit: Value:2902}] Key:0080:7dc674d3}
```

# Acknowledgements

This project would never come into life without the help and inspiration from the following resources:

* [SMA-WebBox (smawb)](https://github.com/smarthomeNG/plugins/tree/master/smawb)
* [jraedler/SunnyWebBox](https://github.com/jraedler/SunnyWebBox)
* Gist [juhaautioniemi/0850bcda0d246dd7d3141308522d00a7](https://gist.github.com/juhaautioniemi/0850bcda0d246dd7d3141308522d00a7)

# Feature requests / support

This is a free time project so, any bugs or feature requests should be done as [new issues](https://help.github.com/articles/creating-an-issue/).