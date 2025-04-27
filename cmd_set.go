package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func cmdSet(c *cli.Context) error {
	if !isInitialize() {
		fmt.Println("You need to initialize.")
		fmt.Println("Please execute the following command.")
		fmt.Println("")
		fmt.Println("  mgu init")
		fmt.Println("")
		return nil
	}

	raw, err := os.ReadFile(appConfigFilePath)
	if err != nil {
		fmt.Println("You need to initialize.")
		fmt.Println("Please execute the following command.")
		fmt.Println("")
		fmt.Println("  mgu init")
		fmt.Println("")
		return nil
	}

	var uc []User
	if err := json.Unmarshal(raw, &uc); err != nil {
		return fmt.Errorf("failed to unmarshal settings: %w", err)
	}

	var list []string
	for _, u := range uc {
		list = append(list, u.Name+" <"+u.Email+">")
	}

	selected := ""
	message := "Please select a user:"
	currentName := getName()
	currentEmail := getEmail()
	if currentName != "" && currentEmail != "" {
		message = message + " (current: " + currentName + " <" + currentEmail + ">) "
	}
	prompt := &survey.Select{
		Message: message,
		Options: list,
	}
	err = survey.AskOne(prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	s := strings.Split(selected, " ")
	r := strings.NewReplacer("<", "", ">", "")
	name := s[0]
	email := r.Replace(s[1])

	if err := setLocalUser(name, email); err != nil {
		return fmt.Errorf("failed to set local user: %w", err)
	}

	fmt.Println(name + " <" + email + "> has been set as a Git' local user.")
	return nil
}
