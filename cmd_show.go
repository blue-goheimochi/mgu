package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdShow(c *cli.Context) {
	if !isGitRepositoryDir() {
		fmt.Println("Your current directory is not a git repository.")
		return
	}
	globalName := getGlobalName()
	globalEmail := getGlobalEmail()
	name := getName()
	email := getEmail()

	hasLocalGitUserSetting := true
	if name == "" && email == "" {
		fmt.Println("Your Git's local user name and email are not set.")
		hasLocalGitUserSetting = false
	} else if name == "" {
		fmt.Println("Your Git's local user name is not set.")
		hasLocalGitUserSetting = false
	} else if email == "" {
		fmt.Println("Your Git's local email is not set.")
		hasLocalGitUserSetting = false
	}
	if !hasLocalGitUserSetting {
		fmt.Println("Currently the following Git's global user are in use.")
		fmt.Println(globalName + " <" + globalEmail + ">")
		return
	}
	fmt.Println(name + " <" + email + ">")
}
