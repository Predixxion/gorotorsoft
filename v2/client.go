package v2

import (
	"net/http"
)

type RotorSoftClient struct {
	HTTPClient *http.Client
	URL        string
	Username   string
	Password   string
}
