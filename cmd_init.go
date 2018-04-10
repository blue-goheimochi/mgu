package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func cmdInit(c *cli.Context) {
	if !fileExists(configDirPath) {
		createDirectory(configDirPath)
	}

	if !fileExists(appConfigDirPath) {
		createDirectory(appConfigDirPath)
	}

	if fileExists(appConfigFilePath) {
		fmt.Println(appConfigFilePath + " is already exist.")
		return
	}

	file, err := os.Create(appConfigFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(([]byte)("[]"))
	fmt.Println(appConfigFilePath + " has been created.")
}
