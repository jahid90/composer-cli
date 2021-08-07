package file

import (
	"errors"
	"io/ioutil"
)

// ReadFile Reads filename and returns its contents
func ReadFile(filename string) ([]byte, error) {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, errors.New("error: file not found")
	}

	return content, nil
}
