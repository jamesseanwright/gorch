package gorch

import (
	"fmt"
	"gorch/internal/request"
)

type Gorch struct {
	baseUrl string
}

// func New(baseUrl string) *Gorch {
// 	gorch := new(Gorch)
// 	gorch.client = internal.NewClient(baseUrl)
// 	return gorch
// }

// TODO: return some deserialised object
func (gorch *Gorch) Get(endpoint string) string {
	request := request.New("GET", gorch.baseUrl+endpoint)
	//resp, err := http.Get(Gorch.baseUrl + endpoint)
	fmt.Print(request.GetUrl())
	return "deserialised to obj"
}
