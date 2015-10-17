package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetQueryString(t *testing.T) {
	const testName = "It should correctly build query strings"

	queryParams := make(map[string]string)
	queryParams["key1"] = "value1"
	queryParams["key2"] = "value2"
	queryParams["key3"] = "value3"

	actual := BuildQueryString(queryParams)
	const expected = "?key1=value1&key2=value2&key3=value3"

	assert.EqualValues(t, expected, actual, testName)
}
