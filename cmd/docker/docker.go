package docker

import (
	"github.com/jahid90/composer/lib/file"
	"github.com/jahid90/composer/lib/process"
	"github.com/urfave/cli/v2"
)

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "docker",
		Usage: "Composes a docker-compose.yaml file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "set",
				Usage: "The value overrides",
			},
		},
		Action: func(c *cli.Context) error {

			processed, err := process.InterpolateFile("template.yaml", "values.yaml")
			if err != nil {
				return err
			}

			err = file.WriteFile("docker-compose.yaml", processed)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
