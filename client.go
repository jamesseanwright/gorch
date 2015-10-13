package gorch

import (
	"gorch/internal"
	"net/http"
)

type GorchClient struct {
	baseUrl string
	onError func(err error)
}

func Create(baseUrl string, onError func(err error)) *GorchClient {
	gorchClient := new(GorchClient)
	gorchClient.baseUrl = baseUrl
	gorchClient.onError = onError
	return gorchClient
}

func (gorchClient *GorchClient) Get(url endpoint, queryParams map[string]string) {
	req := GorchRequest.create(url, endpoint)
	resp, err := http.Get(gorchClient.baseUrl + endpoint)
}
