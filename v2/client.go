package v2

import (
	"net/http"
	"time"

	"github.com/Predixxion/gosoap"
)

type RotorSoftClient struct {
	HTTPClient *http.Client
	URL        string
	Username   string
	Password   string
}
