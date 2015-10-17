package main

import (
	"fmt"
	"gorch"
)

func main() {
	fmt.Println("Fetching manly wisdom from Ron...")

	api := gorch.New("http://ron-swanson-quotes.herokuapp.com")
	response, err := api.Get("/quotes").Execute()

	if err != nil {
		panic(err)
	}

	fmt.Printf(`"%s" - Ron Swanson`, response["quote"])
}
