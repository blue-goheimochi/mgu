package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func cmdInit(c *cli.Context) error {
	if !fileExists(configDirPath) {
		createDirectory(configDirPath)
	}

	if !fileExists(appConfigDirPath) {
		createDirectory(appConfigDirPath)
	}

	if fileExists(appConfigFilePath) {
		fmt.Println(appConfigFilePath + " is already exist.")
		return nil
	}

	file, err := os.Create(appConfigFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	file.Write(([]byte)("[]"))
	fmt.Println(appConfigFilePath + " has been created.")
	fmt.Println("Successfully Initialization.")
	return nil
}
