package client

import (
	"gorch/internal/request"
	"net/http"
)

func execute(request request.Request) {
	httpClient := new(http.Client)
	return httpClient.Do(http.NewRequest(request.GetMethod(), request.GetUrl(), nil))
}
