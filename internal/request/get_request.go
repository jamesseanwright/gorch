package request

import (
	"gorch/internal"
)

type GetRequest struct {
	url         string
	queryString string
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

func (getRequest *GetRequest) SetParams(params map[string]string) {
	getRequest.queryString = internal.BuildQueryString(params)
}

func NewGetRequest(url string) *GetRequest {
	getRequest := new(GetRequest)
	getRequest.url = url
	return getRequest
}
