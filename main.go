package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// PathName = Default pathname where CS would be initialized
const PathName = ".configswitcher"

func initializeApp() {
	if _, err := os.Stat("./" + PathName); os.IsNotExist(err) {
		os.Mkdir(".configswitcher", 0755)
	}
}

func main() {
	initializeApp()
	app := cli.NewApp()

	app.Name = "ConfigSwitcher"
	app.Usage = "Switch between bento configs"
	app.Version = "0.0.2"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:  "config",
			Usage: "Manage config for this project",
			Subcommands: cli.Commands{
				cli.Command{
					Name: "list",
					Action: func(c *cli.Context) error {
						ListConfig()
						return nil
					},
				},
				cli.Command{
					Name: "add",
					Action: func(c *cli.Context) error {
						AddConfig(c)
						return nil
					},
				},
				cli.Command{
					Name: "apply",
					Action: func(c *cli.Context) error {
						ApplyConfig(c)
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
