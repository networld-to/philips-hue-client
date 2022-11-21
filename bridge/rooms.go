package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRooms() (rooms *Rooms, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/room", apiEndpointPrefix)
	var responseBody []byte
	responseBody, resp, err = getCall(endpoint)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &rooms)
	return
}

func GetRoom(id string) (rooms *Rooms, resp *http.Response, err error) {
	var endpoint = fmt.Sprintf("%s/resource/room/%s", apiEndpointPrefix, id)
	var responseBody []byte
	responseBody, resp, err = getCall(endpoint)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseBody, &rooms)
	return
}