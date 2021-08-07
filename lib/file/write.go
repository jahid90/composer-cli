package file

import (
	"errors"
	"io/fs"
	"io/ioutil"
)

var defaultMode fs.FileMode = 0644

func WriteFile(filename string, data []byte) error {

	err := ioutil.WriteFile(filename, data, defaultMode)

	if err != nil {
		return errors.New("error: failed to write")
	}

	return nil
}
