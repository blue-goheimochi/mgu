package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

var (
	homeDirPath, _    = os.UserHomeDir()
	configDirPath     = filepath.Join(homeDirPath, ".config")
	appConfigDirPath  = filepath.Join(configDirPath, "mgu")
	appConfigFilePath = filepath.Join(appConfigDirPath, "setting.json") // Fix typo in filename
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	survey.ErrorTemplate = survey.ErrorTemplate + "X"
	// v2 uses a different way to customize icons
	// via survey.WithIcons() option when calling AskOne/Ask

	app := cli.NewApp()
	app.Name = "mgu"
	app.Usage = "Manage git local users"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		if c.Args().Get(0) == "" {
			return cmdShow(c)
		} else {
			return cli.ShowCommandHelp(c, "")
		}
	}

	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create setting file",
			Action:  cmdInit,
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show current Git's user",
			Action:  cmdShow,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add Git's local user",
			Action:  cmdAdd,
		},
		{
			Name:   "set",
			Usage:  "Setting Git's local user",
			Action: cmdSet,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Display Git's local users",
			Action:  cmdList,
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove Git's local user",
			Action:  cmdRemove,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
