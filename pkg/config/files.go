package config

import (
	"os"
)

// IsInitialized checks if the application has been initialized
func IsInitialized() bool {
	return FileExists(SettingFilePath)
}

// FileExists checks if the specified file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// CreateDirectory creates a directory with proper permissions
func CreateDirectory(path string) error {
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	return nil
}
