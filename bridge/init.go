package bridge

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

var bridge Bridge
var apiEndpointPrefix string
var client *http.Client

func init() {
	bridge = LoadHueBridgeConfig()
	apiEndpointPrefix = fmt.Sprintf("https://%s/clip/v2", bridge.Bridge)

	client = &http.Client{
		Transport: &http.Transport{
			/* #nosec G402 */
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func getCall(endpoint string) (respBody []byte, resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("hue-application-key", bridge.Username)

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func putCall(endpoint string, body []byte) (respBody []byte, resp *http.Response, err error) {
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPut, endpoint, bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("hue-application-key", bridge.Username)

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}
