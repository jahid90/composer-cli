package compose

import (
	"fmt"

	"github.com/jahid90/composer/lib/file"
	"github.com/jahid90/composer/lib/process"
	"github.com/urfave/cli/v2"
)

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

			processed, err := process.InterpolateFile(templateFile, valuesFile)
			if err != nil {
				return err
			}

			outFile := c.String("out")
			if len(outFile) != 0 {
				file.WriteFile(outFile, processed)
			} else {
				fmt.Println(string(processed))
			}

			return nil
		},
	}
}
