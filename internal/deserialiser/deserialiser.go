package deserialiser

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ToMap(reader io.Reader) (map[string]interface{}, error) {
	buffer, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(buffer, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
