package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

func cmdList(c *cli.Context) {
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
		fmt.Println("  mugu init")
		fmt.Println("")
		return
	}

	var uc []User
	json.Unmarshal(raw, &uc)

	for _, u := range uc {
		fmt.Println(u.Name + " <" + u.Email + ">")
	}
}
