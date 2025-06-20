package commands

import (
	"fmt"
	"os"

	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/urfave/cli/v2"
)

// Init initializes the application by creating necessary directories and files
func Init(c *cli.Context) error {
	if !config.FileExists(config.ConfigDirPath) {
		if err := config.CreateDirectory(config.ConfigDirPath); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	if !config.FileExists(config.AppConfigDirPath) {
		if err := config.CreateDirectory(config.AppConfigDirPath); err != nil {
			return fmt.Errorf("failed to create app config directory: %w", err)
		}
	}

	if config.FileExists(config.SettingFilePath) {
		fmt.Println(config.SettingFilePath + " already exists.")
		return nil
	}

	file, err := os.Create(config.SettingFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	_, err = file.Write(([]byte)("[]"))
	if err != nil {
		return fmt.Errorf("failed to write initial content: %w", err)
	}
	
	fmt.Println(config.SettingFilePath + " has been created.")
	fmt.Println("Initialization successful.")
	return nil
}
