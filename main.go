package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"github.com/vanclief/compose-cli/generators"
	"github.com/vanclief/ez"
)

func main() {
	app := cli.NewApp()
	app.Name = "compose-cli"
	app.Usage = "Create boilerplate for your compose based application"
	app.Version = "1.0.1"

	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generates a file",
			Subcommands: []*cli.Command{
				{
					Name:  "resource",
					Usage: "Generate an API resource file",
					Subcommands: []*cli.Command{
						{
							Name:  "api",
							Usage: "Generate a new API for a resource",
							Action: func(c *cli.Context) error {
								return generators.NewResourceAPI()
							},
						},
						{
							Name:  "method",
							Usage: "Generate the API resource method",
							Flags: []cli.Flag{
								&cli.BoolFlag{
									Name:    "force",
									Usage:   "Force the creation of the resource method even if it already exists",
									Aliases: []string{"f"},
								},
							},
							Action: func(c *cli.Context) error {
								forceFlag := c.Bool("force")

								return generators.NewResourceMethod(forceFlag)
							},
						},
					},
				},
				{
					Name:  "model",
					Usage: "Generate a new Model for a resource",
					Action: func(c *cli.Context) error {
						return generators.NewResourceModel()
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red(ez.ErrorMessage(err))
		os.Exit(0)
	}
}
