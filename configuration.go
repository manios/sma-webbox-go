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

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "http://localhost",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
