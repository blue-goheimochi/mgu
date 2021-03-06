package main

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1/core"
)

var (
	homeDirPath, _    = homedir.Dir()
	configDirPath     = homeDirPath + "/.config"
	appConfigDirPath  = configDirPath + "/mgu"
	appConfigFilePath = appConfigDirPath + "/settign.json"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	core.ErrorIcon = "X"
	core.SelectFocusIcon = ">"
	core.MarkedOptionIcon = "[x]"
	core.UnmarkedOptionIcon = "[ ]"

	app := cli.NewApp()
	app.Name = "mgu"
	app.Usage = "Manage git local users"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		if c.Args().Get(0) == "" {
			cmdShow(c)
		} else {
			cli.ShowCommandHelp(c, "")
		}
		return nil
	}

	app.Commands = []cli.Command{
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

	app.Run(os.Args)
}
