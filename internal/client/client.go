package client

import (
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

func (client *Client) execute(request request.Request) {
	return client.httpClient.Do(http.NewRequest(request.GetMethod(), request.GetUrl(), nil))
}
