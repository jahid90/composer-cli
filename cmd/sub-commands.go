package cmd

import (
	"github.com/jahid90/composer/cmd/compose"
	"github.com/urfave/cli/v2"
)

// GetSubCommands returns the subcommands for the app.
func GetSubCommands() []*cli.Command {

	composeCmd := compose.Cmd()

	return []*cli.Command{
		composeCmd,
	}
}
