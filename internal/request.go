package internal

import (
	"bytes"
)

type GorchRequest struct {
	url         string
	method      string
	queryString string
}

func Create(url string, method string) *GorchRequest {
	gorchRequest := new(GorchRequest)
	gorchRequest.url = url
	gorchRequest.method = method
	return gorchRequest
}

func (gorchRequest *GorchRequest) SetQueryString(queryParams map[string]string) {
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
	}

	return buffer.String()
}
