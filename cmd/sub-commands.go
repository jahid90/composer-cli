package cmd

import (
	"github.com/jahid90/composer/cmd/compose"
	"github.com/jahid90/composer/cmd/docker"
	"github.com/urfave/cli/v2"
)

// GetSubCommands returns the subcommands for the app.
func GetSubCommands() []*cli.Command {

	composeCmd := compose.Cmd()
	dockerCmd := docker.Cmd()

	return []*cli.Command{
		composeCmd,
		dockerCmd,
	}
}
