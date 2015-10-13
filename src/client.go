package gorch

import (
	"net/http"
)

type Client struct {
	baseUrl string
	onError func(err error)
}

func create(baseUrl string, onError func(err error)) *Client {
	client := new(Client)
	client.baseUrl = baseUrl
	client.onError = onError
	return client
}

func (c *Client) get(url endpoint, queryParams map[string]string) {

}
