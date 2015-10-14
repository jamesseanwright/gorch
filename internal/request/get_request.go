package request

import (
	"bytes"
)

type GetRequest struct {
	url                   string
	queryString           string
	deserialisationTarget interface{} // TODO: move to request base class
}

func (getRequest *GetRequest) Url() string {
	return getRequest.url + getRequest.queryString
}

func (getRequest *GetRequest) Method() string {
	return "GET"
}

func (getRequest *GetRequest) Body() string {
	return ""
}

func (getRequest *GetRequest) DeserialisationTarget() interface{} {
	return getRequest.deserialisationTarget
}

func (getRequest *GetRequest) SetParams(params map[string]string) {
	getRequest.queryString = CreateQueryString(params)
}

func (getRequest *GetRequest) SetDeserialisationTarget(target interface{}) {
	getRequest.deserialisationTarget = target
}

func NewGetRequest(url string) *GetRequest {
	getRequest := new(GetRequest)
	getRequest.url = url
	return getRequest
}

func CreateQueryString(queryParams map[string]string) string {
	var separator string
	var buffer bytes.Buffer
	i := 0

	for key, value := range queryParams {
		if i > 0 {
			separator = "&"
		} else {
			separator = "?"
		}

		buffer.WriteString(separator + key + "=" + value)
		i++
	}

	return buffer.String()
}
