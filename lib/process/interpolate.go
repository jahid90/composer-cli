package process

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/jahid90/composer/lib/file"
	"gopkg.in/yaml.v2"
)

// Variables A struct that stores all the variables to interpolate a template with
type Variables struct {
	Values map[interface{}]interface{} `yaml:"values"`
}

// Interpolate Interpolates variables into a template and returns the composed output
func Interpolate(templateData []byte, variables *Variables) ([]byte, error) {

	fmt.Printf("%#v", variables)

	var buffer bytes.Buffer

	parsed, err := template.New("t").Parse(string(templateData))
	if err != nil {
		return nil, err
	}

	err = parsed.Execute(&buffer, variables)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// InterpolateFile Interpolates a template file with values from a values file and returns the composed output
func InterpolateFile(templateFile string, valuesFile string) ([]byte, error) {

	templateData, err := file.ReadFile(templateFile)
	if err != nil {
		return nil, err
	}

	valuesData, err := file.ReadFile(valuesFile)
	if err != nil {
		return nil, err
	}

	var variables Variables
	yaml.Unmarshal(valuesData, &variables)

	return Interpolate(templateData, &variables)
}
