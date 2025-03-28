package info

import (
	"net/http"

	"github.com/buzyka/sw-sdk-go/shopware"
)

const (
	healthCheckPath = "/_info/system-health-check"
)

type HealthChecksResponse struct {
    Checks []HealthCheck `json:"checks"`
}

type HealthCheck struct {
    Name    string                   `json:"name"`
    Healthy bool                     `json:"healthy"`
    Status  string                   `json:"status"`
    Message string                   `json:"message"`
    Extra   []map[string]interface{}  `json:"extra"`
}

func (r *InfoRepository) GetHealthCheck() (*HealthChecksResponse, error) {
	healthCheckResponse := &HealthChecksResponse{}
	err := r.client.InvokeRequest(&shopware.SwAPIRequest{
		Method: http.MethodGet,
		URL:    healthCheckPath,
		Body:   nil,
	}, healthCheckResponse)
	return healthCheckResponse, err
}