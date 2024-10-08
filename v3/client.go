package v3

import (
	"net/http"
)

type RotorSoftClient struct {
	HTTPClient *http.Client
	URL        string
	Username   string
	Password   string
}