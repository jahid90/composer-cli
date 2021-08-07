package compose

import (
	"os"
	"text/template"

	"github.com/jahid90/composer/lib/file"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

type Overrides struct {
	Values map[string]string `yaml:"values"`
}

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "compose",
		Usage: "Interpolate values into a template to compose a file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "in",
				Usage: "The template file",
			},
			&cli.StringFlag{
				Name:  "out",
				Usage: "The output file",
			},
			&cli.StringFlag{
				Name:  "values",
				Usage: "The file with the values",
			},
			&cli.StringFlag{
				Name:  "set",
				Usage: "The value overrides",
			},
		},
		Action: func(c *cli.Context) error {

			templateFile := c.String("in")
			valuesFile := c.String("values")

			templateData, err := file.ReadFile(templateFile)
			if err != nil {
				return err
			}

			valuesData, err := file.ReadFile(valuesFile)
			if err != nil {
				return err
			}

			var overrides Overrides
			yaml.Unmarshal(valuesData, &overrides)

			tmpl, err := template.New("test").Parse(string(templateData))
			if err != nil {
				return err
			}

			err = tmpl.Execute(os.Stdout, overrides)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
