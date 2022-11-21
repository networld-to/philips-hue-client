package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLights() (lights *Lights, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/light", apiEndpointPrefix)
	var responseBody []byte
	responseBody, resp, err = getCall(endpoint)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &lights)
	return
}

func GetLight(id string) (lights *Lights, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/light/%s", apiEndpointPrefix, id)
	var responseBody []byte
	responseBody, resp, err = getCall(endpoint)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &lights)
	return
}

func SwitchOn(id string) (respBody []byte, resp *http.Response, err error) {
	var control = &OnOff{
		On: On{On: true},
	}
	var endpoint = fmt.Sprintf("%s/resource/light/%s", apiEndpointPrefix, id)

	var jsonBody []byte
	jsonBody, err = json.Marshal(control)
	if err != nil {
		return
	}
	respBody, resp, err = putCall(endpoint, jsonBody)
	return
}

func SwitchOff(id string) (respBody []byte, resp *http.Response, err error) {
	var control = &OnOff{
		On: On{On: false},
	}
	var endpoint = fmt.Sprintf("%s/resource/light/%s", apiEndpointPrefix, id)

	var jsonBody []byte
	jsonBody, err = json.Marshal(control)
	if err != nil {
		return
	}
	respBody, resp, err = putCall(endpoint, jsonBody)
	return
}

func ChangeLightState(id string, control *LightControl) (respBody []byte, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/light/%s", apiEndpointPrefix, id)

	var jsonBody []byte
	jsonBody, err = json.Marshal(control)
	if err != nil {
		return
	}
	respBody, resp, err = putCall(endpoint, jsonBody)
	return
}
