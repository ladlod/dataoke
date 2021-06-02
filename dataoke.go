package dataoke

import (
	"sync"

	"github.com/parnurzeal/gorequest"
)

type DaTaoKeApp struct {
	appKey     string
	apiVersion string

	clients *sync.Pool
}

// NewDaTaoKeApp
func NewDaTaoKeApp(appKey string, apiVersion string) *DaTaoKeApp {
	return &DaTaoKeApp{
		appKey:     appKey,
		apiVersion: apiVersion,
		clients: &sync.Pool{
			New: func() interface{} {
				return gorequest.New()
			},
		},
	}
}

// postQuery
func (d *DaTaoKeApp) postQuery(params Params, url string) (err error) {
	client := d.clients.Get().(*gorequest.SuperAgent)
	defer d.clients.Put(client)
	params.Set("appKey", d.appKey)
	params.Set("version", d.apiVersion)

	_, b, errors := client.Get().End()

	return
}
