package internal

import (
	"bytes"
	"sort"
)

func BuildQueryString(queryParams map[string]string) string {
	var separator string
	var buffer bytes.Buffer
	i := 0
	var sortedKeys []string

	for key := range queryParams {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		if i > 0 {
			separator = "&"
		} else {
			separator = "?"
		}

		buffer.WriteString(separator + key + "=" + queryParams[key])
		i++
	}

	return buffer.String()
}
