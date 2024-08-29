package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var subCmdList = []*cli.Command{}

func main() {
	app := &cli.App{
		Name:     "lgtm-generator",
		Usage:    "CLI for lgtm image",
		Commands: subCmdList,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Value:   false,
				Usage:   "output detail log",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
