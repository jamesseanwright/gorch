package deserialiser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TestStruct struct {
	name string
	age  int
	pets []string
}

func TestDeserialiseToInstance(t *testing.T) {
	const testName = "It should deserialise JSON data to an instance of the specified type"
	const json = `{ "name": "Bob", "age": 6, "pets": [ "Trevor", "Toby" ] }`

	reader := strings.NewReader(json)

	testInstance := new(TestStruct)
	fmt.Println(&testInstance)
	InstanceFromReader(reader, &testInstance)
	fmt.Println(*testInstance)
	assert.EqualValues(t, "Barb", testInstance.name)
}
