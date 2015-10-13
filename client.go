package gorch

import (
	"fmt"
	"gorch/internal"
)

type GorchClient struct {
	baseUrl string
	onError func(err error)
}

func New(baseUrl string) *GorchClient {
	gorchClient := new(GorchClient)
	gorchClient.baseUrl = baseUrl
	return gorchClient
}

// TODO: return some deserialised object
func (gorchClient *GorchClient) Get(endpoint string, queryParams map[string]string) string {
	var req *internal.GorchRequest = internal.CreateRequest(gorchClient.baseUrl, endpoint)
	//resp, err := http.Get(gorchClient.baseUrl + endpoint)
	fmt.Print(req.GetQueryString())
	return "deserialised to obj"
}
