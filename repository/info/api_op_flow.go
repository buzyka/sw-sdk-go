package info

import (
	"net/http"

	"github.com/buzyka/sw-sdk-go/shopware"
)

const (
	flowsPath = "/_info/flow-actions.json"
)

type FlowsResponse []Flow

type Flow struct {
    Name        string   `json:"name"`
    Requirements []string `json:"requirements"`
    Extensions  []string `json:"extensions"`
}

func (r *InfoRepository) GetFlows() (*FlowsResponse, error) {
	flowsResponse := &FlowsResponse{}
	err := r.client.InvokeRequest(&shopware.SwAPIRequest{
		Method: http.MethodGet,
		URL:    flowsPath,
		Body:   nil,
	}, flowsResponse)
	return flowsResponse, err
}