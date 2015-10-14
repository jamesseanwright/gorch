package deserialiser

import (
	"encoding/json"
	"io"
	"reflect"
)

func InstanceFromReadCloser(reader io.ReadCloser, target interface{}) (interface{}, error) {
	instancePtr := reflect.New(reflect.TypeOf(target))
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(instancePtr)
	return instancePtr.Interface(), err
}
