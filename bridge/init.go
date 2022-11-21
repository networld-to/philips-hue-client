package bridge

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
)

var bridge Bridge
var apiEndpointPrefix string
var client *http.Client

const HUE_BRIDGE_CA = `-----BEGIN CERTIFICATE-----
MIICMjCCAdigAwIBAgIUO7FSLbaxikuXAljzVaurLXWmFw4wCgYIKoZIzj0EAwIw
OTELMAkGA1UEBhMCTkwxFDASBgNVBAoMC1BoaWxpcHMgSHVlMRQwEgYDVQQDDAty
b290LWJyaWRnZTAiGA8yMDE3MDEwMTAwMDAwMFoYDzIwMzgwMTE5MDMxNDA3WjA5
MQswCQYDVQQGEwJOTDEUMBIGA1UECgwLUGhpbGlwcyBIdWUxFDASBgNVBAMMC3Jv
b3QtYnJpZGdlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEjNw2tx2AplOf9x86
aTdvEcL1FU65QDxziKvBpW9XXSIcibAeQiKxegpq8Exbr9v6LBnYbna2VcaK0G22
jOKkTqOBuTCBtjAPBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNV
HQ4EFgQUZ2ONTFrDT6o8ItRnKfqWKnHFGmQwdAYDVR0jBG0wa4AUZ2ONTFrDT6o8
ItRnKfqWKnHFGmShPaQ7MDkxCzAJBgNVBAYTAk5MMRQwEgYDVQQKDAtQaGlsaXBz
IEh1ZTEUMBIGA1UEAwwLcm9vdC1icmlkZ2WCFDuxUi22sYpLlwJY81Wrqy11phcO
MAoGCCqGSM49BAMCA0gAMEUCIEBYYEOsa07TH7E5MJnGw557lVkORgit2Rm1h3B2
sFgDAiEA1Fj/C3AN5psFMjo0//mrQebo0eKd3aWRx+pQY08mk48=
-----END CERTIFICATE-----`

func init() {
	bridge = LoadHueBridgeConfig()
	apiEndpointPrefix = fmt.Sprintf("https://%s/clip/v2", bridge.Bridge)

	var certs = x509.NewCertPool()
	certs.AppendCertsFromPEM([]byte(HUE_BRIDGE_CA))
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// RootCAs: certs,
				// Point this domain to your Philips Hue Bridge
				// ServerName: "philips-hue.home.arpa",
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
