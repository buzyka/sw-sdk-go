package info

import (

	"github.com/buzyka/sw-sdk-go/shopware"
)


const (
	configPath  = "/_info/config"
)

type InfoRepository struct {
	client *shopware.Client
}

func (r *InfoRepository) SetClient(client *shopware.Client) {
	r.client = client
}