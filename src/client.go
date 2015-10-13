package gorch

import (
	"gorch/internal/request"
	"net/http"
)

type GorchClient struct {
	baseUrl string
	onError func(err error)
}

func create(baseUrl string, onError func(err error)) *GorchClient {
	gorchClient := new(GorchClient)
	gorchClient.baseUrl = baseUrl
	gorchClient.onError = onError
	return gorchClient
}

func (gorchClient *GorchClient) get(url endpoint, queryParams map[string]string) {
	req := Gorch
	resp, err := http.Get(gorchClient.baseUrl + endpoint)
}
