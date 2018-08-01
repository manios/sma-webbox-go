package smawebboxgo

import (
	"strconv"
	"time"
)

// https://stackoverflow.com/a/42872183/1411901
// https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
// https://stackoverflow.com/questions/19303137/golang-read-ints-from-stdin-until-eof-while-reporting-format-errors
type WebboxClient struct {
	// The Url of SMA Sunny Webbox
	URL       string
	ApiClient *APIClient
}

func NewWebboxClient(url string) WebboxClient {

	var apiConf *Configuration
	apiConf = NewConfiguration()
	apiConf.BasePath = url
	apiConf.AddDefaultHeader("Connection", "close")
	apiConf.UserAgent = "sma-webbox-go/1.4"

	var apiClientTmp = NewAPIClient(apiConf)

	return WebboxClient{
		URL:       url,
		ApiClient: apiClientTmp,
	}
}

func (w *WebboxClient) newGenericRequest(processID string) *GenericRequest {

	var nowID = generateCallID()

	return &GenericRequest{
		Id:      nowID,
		Proc:    processID,
		Format:  "JSON",
		Version: "1.0",
	}
}

func (w *WebboxClient) GetPlantOverview() (PlantOverviewResponse, error) {

	var plantRequest = w.newGenericRequest("GetPlantOverview")

	response, _, err := w.ApiClient.PlantApi.GetPlantOverview(nil, *plantRequest)

	return response, err
}

// GetDevices returns an object which contains a list of the devices which are
// attached to SMA Webbox device.
func (w *WebboxClient) GetDevices() (DevicesResponse, error) {

	var deviceRequest = w.newGenericRequest("GetDevices")

	response, _, err := w.ApiClient.DeviceApi.GetDevices(nil, *deviceRequest)

	return response, err
}

func (w *WebboxClient) GetProcessDataChannels(deviceID string) (ProcessDataChannelsResponse, error) {

	var nowID = generateCallID()

	deviceParam := DeviceDataChannelsDeviceRequest{
		Device: deviceID,
	}

	var apiRequest = DeviceDataChannelsRequest{
		Id:      nowID,
		Proc:    "GetProcessDataChannels",
		Format:  "JSON",
		Version: "1.0",
		Params:  &deviceParam,
	}

	response, _, err := w.ApiClient.ProcessDataApi.GetProcessDataChannels(nil, apiRequest)

	return response, err
}

func (w *WebboxClient) GetProcessData(deviceUid string) (ProcessDataResponse, error) {

	devRequest := ProcessDataRequestDeviceObject{
		Key:      deviceUid,
		Channels: nil,
	}

	devArray := []ProcessDataRequestDeviceObject{
		devRequest,
	}

	return w.GetProcessDataForDevices(devArray)
}

func (w *WebboxClient) GetProcessDataForDevices(devList []ProcessDataRequestDeviceObject) (ProcessDataResponse, error) {

	var nowID = generateCallID()

	procDataReqContainer := ProcessDataRequestDevicesContainer{
		Devices: devList,
	}

	var apiRequest = ProcessDataRequest{
		Id:      nowID,
		Proc:    "GetProcessData",
		Format:  "JSON",
		Version: "1.0",
		Params:  &procDataReqContainer,
	}

	response, _, err := w.ApiClient.ProcessDataApi.GetProcessData(nil, apiRequest)

	return response, err
}

func generateCallID() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
