package graphql

import (
	"github.com/buger/jsonparser"
)

func TypeInSchema(in []byte, name, description string, extra ...string) bool {
	var found bool
	hasExtra := len(extra) > 0

	jsonparser.ArrayEach(in, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if found {
			return
		}
		in, err := jsonparser.GetString(value, "name")
		if err != nil || in != name {
			return
		}
		if in, err = jsonparser.GetString(value, "description"); err != nil || in != description {
			return
		}
		if hasExtra {
			for _, f := range extra {
				if _, _, _, err := jsonparser.Get(value, f); err != nil {
					return
				}
			}
		}
		found = true
	}, "__schema", "types")

	return found
}
