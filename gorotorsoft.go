package gorotorsoft

import (
	"net/http"

	v0 "github.com/Predixxion/gorotorsoft/v0"
	v1 "github.com/Predixxion/gorotorsoft/v1"
	v2 "github.com/Predixxion/gorotorsoft/v2"
	v3 "github.com/Predixxion/gorotorsoft/v3"
)

type Client struct {
	V0 *v0.RotorSoftClient
	V1 *v1.RotorSoftClient
	V2 *v2.RotorSoftClient
	V3 *v3.RotorSoftClient
}

func NewClient(httpClient *http.Client, url, username, password string) *Client {
	rotorSoftClientV3 := &v3.RotorSoftClient{
		HTTPClient: httpClient,
		URL:        url,
		Username:   username,
		Password:   password,
		Helper:     &v3.Helper{},
	}
	rotorSoftClientV3.Helper.Client = rotorSoftClientV3

	return &Client{
		V0: &v0.RotorSoftClient{HTTPClient: httpClient, URL: url, Username: username, Password: password},
		V1: &v1.RotorSoftClient{HTTPClient: httpClient, URL: url, Username: username, Password: password},
		V2: &v2.RotorSoftClient{HTTPClient: httpClient, URL: url, Username: username, Password: password},
		V3: rotorSoftClientV3,
	}
}
