package gorch

import (
	"gorch/internal/client"
	"gorch/internal/request"
)

type Gorch struct {
	baseUrl        string
	currentRequest request.Request
	client         *client.Client
}

func New(baseUrl string) *Gorch {
	gorch := new(Gorch)
	gorch.client = client.New()
	gorch.baseUrl = baseUrl
	return gorch
}

func (gorch *Gorch) Get(endpoint string) *Gorch {
	gorch.currentRequest = request.New("GET", gorch.baseUrl+endpoint)
	return gorch
}

func (gorch *Gorch) WithParams(params map[string]string) *Gorch {
	gorch.currentRequest.SetParams(params)
	return gorch
}

func (gorch *Gorch) WithHeaders(headers map[string]string) *Gorch {
	gorch.currentRequest.SetHeaders(headers)
	return gorch
}

func (gorch *Gorch) Execute() (map[string]interface{}, error) {
	return gorch.client.Execute(gorch.currentRequest)
}
