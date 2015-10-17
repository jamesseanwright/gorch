package deserialiser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToMap(t *testing.T) {
	const testName = "It should deserialise JSON data to a map"
	const json = `{ "name": "Bob", "age": 6, "pets": [ "Trevor", "Toby" ] }`

	reader := strings.NewReader(json)

	data, err := ToMap(reader)
	assert.Empty(t, err)
	assert.EqualValues(t, "Bob", data["name"])
	assert.EqualValues(t, 6, data["age"])
	assert.EqualValues(t, []interface{}{"Trevor", "Toby"}, data["pets"])
}
