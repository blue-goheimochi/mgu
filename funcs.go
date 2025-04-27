package main

import (
	"os"
)

func isInitialize() bool {
	return fileExists(appConfigFilePath)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func createDirectory(path string) error {
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	return nil
}
