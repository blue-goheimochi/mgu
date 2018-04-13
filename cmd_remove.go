package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
)

func cmdRemove(c *cli.Context) {
	if !isInitialize() {
		fmt.Println("You need to initialize.")
		fmt.Println("Please execute the following command.")
		fmt.Println("")
		fmt.Println("  mgu init")
		fmt.Println("")
		return
	}

	raw, err := ioutil.ReadFile(appConfigFilePath)
	if err != nil {
		fmt.Println("You need to initialize.")
		fmt.Println("Please execute the following command.")
		fmt.Println("")
		fmt.Println("  mgu init")
		fmt.Println("")
		return
	}

	var uc []User
	if err := json.Unmarshal(raw, &uc); err != nil {
		panic(err)
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
	err = survey.AskOne(prompt, &selected, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := strings.Split(selected, " ")
	r := strings.NewReplacer("<", "", ">", "")
	name := s[0]
	email := r.Replace(s[1])

	flg := false
	prompt2 := &survey.Confirm{
		Message: "Do you want to remove?",
	}
	survey.AskOne(prompt2, &flg, nil)

	if flg {

		var nuc []User
		for _, u := range uc {
			if u.Name != name && u.Email != email {
				nuc = append(nuc, u)
			}
		}

		bytes, _ := json.Marshal(&nuc)
		ioutil.WriteFile(appConfigFilePath, bytes, os.ModePerm)
		fmt.Println(name + " <" + email + "> is removed.")
	}
}
