package client

import (
	"gorch/internal/deserialiser"
	"gorch/internal/request"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	client := new(Client)
	client.httpClient = new(http.Client)
	return client
}

func (client *Client) Execute(request request.Request) (interface{}, error) {
	respPayload, err := http.NewRequest(request.Method(), request.Url(), nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.httpClient.Do(respPayload)

	if err != nil {
		return nil, err
	}

	target := request.DeserialisationTarget()

	err = deserialiser.InstanceFromReader(resp.Body, target)
	return target, err
}
