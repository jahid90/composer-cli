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
			&cli.StringSliceFlag{
				Name:     "set",
				Usage:    "The values overrides",
				Aliases:  []string{"s"},
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {

			overrides, err := process.ProcessOverrides(c.StringSlice("set"))
			if err != nil {
				return err
			}

			processed, err := process.InterpolateFile("template.yaml", "values.yaml", overrides)
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
