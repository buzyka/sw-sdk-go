package info

import (
	"net/http"

	"github.com/buzyka/sw-sdk-go/shopware"
)

type ConfigResponse struct {
    Version        string                 `json:"version"`
    VersionRevision string                 `json:"versionRevision"`
    AdminWorker    AdminWorker             `json:"adminWorker"`
    Bundles        map[string]Bundle       `json:"bundles"`
    Settings       Settings                `json:"settings"`
}

type AdminWorker struct {
    EnableAdminWorker bool     `json:"enableAdminWorker"`
    Transports        []string `json:"transports"`
}

type Bundle struct {
    CSS []string `json:"css"`
    JS  []string `json:"js"`
}

type Settings struct {
    EnableUrlFeature bool `json:"enableUrlFeature"`
}

func (r *InfoRepository) GetConfig() (*ConfigResponse, error) {
	configResponse := &ConfigResponse{}
	err := r.client.InvokeRequest(&shopware.SwAPIRequest{
		Method: http.MethodGet,
		URL: configPath,
		Body: nil,
	},
	configResponse,
	)
	if err != nil {
		return nil, err
	}

	return configResponse, nil
}

