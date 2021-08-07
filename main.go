package main

import (
	"log"
	"os"

	"github.com/jahid90/composer/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	// Disable timestamps in log messages
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	app := cli.NewApp()
	app.Name = "composer"
	app.Usage = "Compose files"
	app.Description = "Composes files by interpolating variables in template files."

	app.Version = "0.1.0"
	app.Flags = []cli.Flag{}
	app.Commands = cmd.GetSubCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
