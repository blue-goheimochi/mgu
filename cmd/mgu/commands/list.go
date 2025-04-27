package commands

import (
	"fmt"

	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/urfave/cli/v2"
)

// List displays all saved Git users
func List(c *cli.Context) error {
	if !config.IsInitialized() {
		printInitRequiredMessage()
		return nil
	}

	manager := config.DefaultManager()
	users, err := manager.GetUsers()
	if err != nil {
		if !config.FileExists(config.SettingFilePath) {
			printInitRequiredMessage()
			return nil
		}
		return fmt.Errorf("failed to get users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users found.")
		return nil
	}

	for _, u := range users {
		fmt.Printf("%s <%s>\n", u.Name, u.Email)
	}
	return nil
}

// Helper function to print initialization required message
func printInitRequiredMessage() {
	fmt.Println("You need to initialize.")
	fmt.Println("Please execute the following command:")
	fmt.Println("")
	fmt.Println("  mgu init")
	fmt.Println("")
}