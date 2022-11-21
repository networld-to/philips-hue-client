package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetGroups() (groups *Groups, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/grouped_light", apiEndpointPrefix)
	var responseBody []byte
	responseBody, resp, err = getCall(endpoint)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &groups)
	return
}

func GetGroup(id string) (groups *Groups, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/grouped_light/%s", apiEndpointPrefix, id)
	var responseBody []byte
	responseBody, resp, err = getCall(endpoint)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &groups)
	return
}


// SwitchGroupOn accepts the group ID (uuid) as input and switches all the
// associated lights on, uses the previous light settings
func SwitchGroupOn(id string) (respBody []byte, resp *http.Response, err error) {
	var control = &OnOff{
		On: On{ On: true},
	}
	var endpoint = fmt.Sprintf("%s/resource/grouped_light/%s", apiEndpointPrefix, id)

	var jsonBody []byte
	jsonBody, err = json.Marshal(control)
	if err != nil {
		return
	}
	respBody, resp, err = putCall(endpoint, jsonBody)
	return
}

// SwitchGroupOff accepts the group ID (uuid) as input and switches all the
// associated lights off
func SwitchGroupOff(id string) (respBody []byte, resp *http.Response, err error) {
	var control = &OnOff{
		On: On{ On: false},
	}
	var endpoint = fmt.Sprintf("%s/resource/grouped_light/%s", apiEndpointPrefix, id)

	var jsonBody []byte
	jsonBody, err = json.Marshal(control)
	if err != nil {
		return
	}
	respBody, resp, err = putCall(endpoint, jsonBody)
	return
}


// ChangeGroupState allows the switching on and off of light groups by providing
// additionally color, brightness and other settings.
func ChangeGroupState(id string, control *LightControl) (respBody []byte, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/grouped_light/%s", apiEndpointPrefix, id)

	var jsonBody []byte
	jsonBody, err = json.Marshal(control)
	if err != nil {
		return
	}
	respBody, resp, err = putCall(endpoint, jsonBody)
	return
}