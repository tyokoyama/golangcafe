package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		&cli.StringFlag {
			Name: "config",
			Aliases: []string{"c"},
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Name = "Score"
	app.Usage = "Show student's score"
	app.Run(os.Args)
}
