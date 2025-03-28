package info

import (
	"net/http"

	"github.com/buzyka/sw-sdk-go/shopware"
)

const (
	versionPath = "/_info/version"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func (r *InfoRepository) GetVersion() (*VersionResponse, error) {
	versionResponse := &VersionResponse{}
	err := r.client.InvokeRequest(&shopware.SwAPIRequest{
		Method: http.MethodGet,
		URL: versionPath,
		Body: nil,
	},  
	versionResponse,
	)
	return versionResponse, err
}
