package dataoke

import (
	"net/url"
	"strconv"

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

func toString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', 10, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', 10, 32)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := dJson.Marshal(v)
		return string(bytes)
	}
}
