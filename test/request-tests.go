package gorch

import (
	"github.com/stretchr/testify/assert"
	"gorch/internal"
	"testing"
)

func TestSetQueryString(t *testing.T) {
	const testName = "It should correctly format query strings"

	var gorchRequest *internal.GorchRequest = internal.NewRequest("http://url", "GET")
	queryParams := make(map[string]string)
	queryParams["key1"] = "value1"
	queryParams["key2"] = "value2"
	queryParams["key3"] = "value3"

	gorchRequest.SetQueryString(queryParams)
	var actual string = gorchRequest.GetQueryString()
	const expected string = "?key1=value1&key2=value2&key3=value3"

	assert.EqualValues(t, expected, actual, testName)
}
