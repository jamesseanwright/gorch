package request

type GorchRequest struct {
	url         string
	method      string
	queryString string
}

func create(url, method) *GorchRequest {
	gorchRequest := new(GorchRequest)
	gorchRequest.url = url
	gorchRequest.method = method
}

func (gorchRequest *GorchRequest) setQueryString(queryParams *map[string]string) {
	var separator string
	queryString := ""
	i := 0

	for key, value := range queryParams {
		if i > 0 {
			separator = "&"
		} else {
			separator = "?"
		}

		queryString += separator + key + "=" + value
	}

	return queryString
}
