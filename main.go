package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vanclief/compose-cli/generators"
)

var duplicatesMap = make(map[interface{}]bool)

func main() {
	app := cli.NewApp()
	app.Name = "compose-cli"
	app.Usage = "Create boilerplate for your compose based application"

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
							Name:  "new",
							Usage: "Generate a new API resource",
							Action: func(c *cli.Context) error {
								return generators.NewResourceAPI()
							},
						},
						{
							Name:  "method",
							Usage: "Generate the API resource method",
							Action: func(c *cli.Context) error {
								return generators.NewAPIMethod()
							},
						},
					},
				},
			},
			// Action: func(c *cli.Context) error {
			// reader := bufio.NewReader(os.Stdin)
			// fmt.Print("Enter something to generate: ")
			// input, _ := reader.ReadString('\n')
			// fmt.Printf("Generating something for: %s\n", input)
			// Add your generation logic here
			// return nil
			// },
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(0)
	}
}
