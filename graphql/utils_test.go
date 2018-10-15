package graphql

import (
	"github.com/buger/jsonparser"
)

func TypeInSchema(in []byte, name, description string, extra ...string) bool {
	var found bool
	jsonparser.ArrayEach(in, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if found {
			return
		}
		var in string

		// Break if name is different
		if in, err = jsonparser.GetString(value, "name"); err != nil || in != name {
			return
		}
		// Break if description is different
		if in, err = jsonparser.GetString(value, "description"); err != nil || in != description {
			return
		}
		// Break if any of the extra fields can't be found in the type structure
		for _, f := range extra {
			if _, _, _, err := jsonparser.Get(value, f); err != nil {
				return
			}
		}
		// Otherwise set to true
		found = true
	}, "__schema", "types")
	return found
}
