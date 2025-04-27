package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func cmdList(c *cli.Context) error {
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

	for _, u := range uc {
		fmt.Println(u.Name + " <" + u.Email + ">")
	}
	return nil
}
