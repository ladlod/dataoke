package dataoke

import (
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

var dJson = jsoniter.ConfigCompatibleWithStandardLibrary

type Params map[string]interface{}

func NewParams() Params {
	p := make(Params)
	return p
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

func (p Params) SetParams(params map[string]interface{}) {
	for key, value := range params {
		p[key] = value
	}
}

func (p Params) GetQuery() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, toString(v))
	}
	return u.Encode()
}

func (p Params) GetJsonBody() []byte {
	body, _ := dJson.Marshal(p)
	return body
}

func (p Params) Clear() {
	p = make(Params)
}
