/*
 * SMA WebBox RPC over HTTP REST API
 *
 * The data loggers Sunny WebBox and Sunny WebBox with Bluetooth continuously record all the data of a PV plant. This is then averaged over a configurable interval and cached. The data can be transmitted at regular intervals to the Sunny Portal for analysis and visualization.  Via the Sunny WebBox and Sunny WebBox with Bluetooth RPC interface, selected data from the PV plant can be transmitted to a remote terminal by means of an RPC protocol (Remote Procedure Call protocol).  Related documents: * [SUNNY WEBBOX RPC User Manual v1.4](http://files.sma.de/dl/2585/SWebBoxRPC-BA-en-14.pdf) * [SUNNY WEBBOX RPC User Manual v1.3](http://files.sma.de/dl/4253/SWebBoxRPC-eng-BUS112713.pdf)
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package smawebboxgo

type PlantOverviewResponse struct {
	Result *PlantOverview `json:"result,omitempty"`
	// Format of the response. In current implementation it always returns: JSON
	Format string `json:"format"`
	// Name of the RPC call. For this call it is: PlantOverview
	Proc string `json:"proc"`
	// Returns the API version of this response
	Version string `json:"version"`
	// A random
	Id string `json:"id"`
	// 
	Error_ string `json:"error,omitempty"`
}
