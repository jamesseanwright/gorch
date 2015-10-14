package deserialiser

import (
	"encoding/json"
	"io"
)

func InstanceFromReader(reader io.Reader, targetPtr interface{}) error {
	decoder := json.NewDecoder(reader)

	for decoder.More() {
		err := decoder.Decode(targetPtr)

		if err != nil {
			return err
		}
	}

	return nil
}
