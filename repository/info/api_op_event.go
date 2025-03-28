package info

import (
	"net/http"

	"github.com/buzyka/sw-sdk-go/shopware"
)

const (
	eventsPath = "/_info/events.json"
)

type EventsResponse []Event

type Event struct {
    Name       string                 `json:"name"`
    Class      string                 `json:"class"`
    Data       map[string]interface{} `json:"data"`
    Aware      []string               `json:"aware"`
    Extensions []string               `json:"extensions"`
}

func (r *InfoRepository) GetEvents() (*EventsResponse, error) {
	eventsResponse := &EventsResponse{}
	err := r.client.InvokeRequest(&shopware.SwAPIRequest{
		Method: http.MethodGet,
		URL:    eventsPath,
		Body:   nil,
	}, eventsResponse)
	return eventsResponse, err
}