package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func cmdAdd(c *cli.Context) error {
	if !isInitialize() {
		fmt.Println("You need to initialize.")
		fmt.Println("Please execute the following command.")
		fmt.Println("")
		fmt.Println("  mgu init")
		fmt.Println("")
		return nil
	}

	var qs = []*survey.Question{
		{
			Name:   "name",
			Prompt: &survey.Input{Message: "user.name"},
		},
		{
			Name:   "email",
			Prompt: &survey.Input{Message: "user.email"},
		},
	}

	answers := struct {
		Name  string
		Email string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	user := User{
		Name:  answers.Name,
		Email: answers.Email,
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

	uc = append(uc, user)

	bytes, err := json.Marshal(&uc)
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}
	
	if err := os.WriteFile(appConfigFilePath, bytes, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write settings: %w", err)
	}
	
	fmt.Println(user.Name + " <" + user.Email + "> is added.")
	return nil
}
