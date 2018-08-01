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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"context"
)

// Linger please
var (
	_ context.Context
)

type DeviceApiService service

/* 
DeviceApiService Devices connected to SMA Webbox

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param rpcbody A JSON request which contains an encrypted JWT token and a groupId. The encrypted token is valid for a random time in the range between  [5,10] seconds.

@return DevicesResponse
*/
func (a *DeviceApiService) GetDevices(ctx context.Context, rpcbody GenericRequest) (DevicesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DevicesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rpc";

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &rpcbody
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v DevicesResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
