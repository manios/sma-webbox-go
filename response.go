/*
 * SMA WebBox RPC over HTTP REST API
 *
 * The data loggers Sunny WebBox and Sunny WebBox with Bluetooth continuously record all the data of a PV plant. This is then averaged over a configurable interval and cached. The data can be transmitted at regular intervals to the Sunny Portal for analysis and visualization.  Via the Sunny WebBox and Sunny WebBox with Bluetooth RPC interface, selected data from the PV plant can be transmitted to a remote terminal by means of an RPC protocol (Remote Procedure Call protocol).  Related documents: * [SUNNY WEBBOX RPC User Manual v1.4](http://files.sma.de/dl/2585/SWebBoxRPC-BA-en-14.pdf) * [SUNNY WEBBOX RPC User Manual v1.3](http://files.sma.de/dl/4253/SWebBoxRPC-eng-BUS112713.pdf)
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package smawebboxgo

import (
	"net/http"
)

type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
