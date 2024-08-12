package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var subCmdList = []*cli.Command{}

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:        "lgtm-generator",
				Usage:       "CLI for lgtm image",
				Subcommands: subCmdList,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
