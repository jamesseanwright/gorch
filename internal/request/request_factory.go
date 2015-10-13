package request

func New(method string, url string) Request {
	return NewGetRequest(url)
}
