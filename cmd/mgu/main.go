package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/cmd/mgu/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	// Custom survey settings
	survey.ErrorTemplate = survey.ErrorTemplate + "X"
	
	app := cli.NewApp()
	app.Name = "mgu"
	app.Usage = "Manage git local users"
	app.Version = "0.0.2"
	
	// Default action when no command is specified
	app.Action = func(c *cli.Context) error {
		if c.Args().Get(0) == "" {
			return commands.Show(c)
		} else {
			return cli.ShowCommandHelp(c, "")
		}
	}

	// Define all available commands
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create settings file",
			Action:  commands.Init,
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show current Git user",
			Action:  commands.Show,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a Git user profile",
			Action:  commands.Add,
		},
		{
			Name:   "set",
			Usage:  "Set a Git user as the local repository user",
			Action: commands.Set,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Display all saved Git users",
			Action:  commands.List,
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove a Git user profile",
			Action:  commands.Remove,
		},
	}

	// Run the application
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}