package request

import (
	"bytes"
)

type GetRequest struct {
	url         string
	queryString string
}

func (getRequest *GetRequest) GetUrl() string {
	return getRequest.url + getRequest.queryString
}

func (getRequest *GetRequest) GetMethod() string {
	return "GET"
}

func (getRequest *GetRequest) GetBody() string {
	return ""
}

func (getRequest *GetRequest) SetParams(params map[string]string) {
	getRequest.queryString = CreateQueryString(params)
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
