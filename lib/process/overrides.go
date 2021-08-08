package process

import (
	"errors"
	"strings"
)

// ProcessOverrides Processes command line overrides and returns a map
func ProcessOverrides(args []string) (map[interface{}]interface{}, error) {

	overrides := make(map[interface{}]interface{})

	for _, arg := range args {
		keyval := strings.Split(arg, "=")

		if len(keyval) != 2 {
			return nil, errors.New("error: invalid value override: " + arg + "; expected of the form <key>=<val>")
		}

		overrides[keyval[0]] = keyval[1]
	}

	// fmt.Printf("%#v", overrides)

	return overrides, nil
}
