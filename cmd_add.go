package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
)

func cmdAdd(c *cli.Context) {
	if !isInitialize() {
		fmt.Println("You need to initialize.")
		fmt.Println("Please execute the following command.")
		fmt.Println("")
		fmt.Println("  mgu init")
		fmt.Println("")
		return
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
		return
	}

	user := User{
		Name:  answers.Name,
		Email: answers.Email,
	}

	raw, err := ioutil.ReadFile(appConfigFilePath)
	if err != nil {
		fmt.Println("You need to initialize.")
		fmt.Println("Please execute the following command.")
		fmt.Println("")
		fmt.Println("  mugu init")
		fmt.Println("")
		return
	}

	var uc []User
	json.Unmarshal(raw, &uc)

	uc = append(uc, user)

	bytes, _ := json.Marshal(&uc)
	ioutil.WriteFile(appConfigFilePath, bytes, os.ModePerm)
	fmt.Println(user.Name + " <" + user.Email + "> is added.")
}
