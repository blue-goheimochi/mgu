package commands

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/blue-goheimochi/mgu/pkg/git"
	"github.com/urfave/cli/v2"
)

// Remove removes a Git user from the configuration
func Remove(c *cli.Context) error {
	if !config.IsInitialized() {
		printInitRequiredMessage()
		return nil
	}

	manager := config.DefaultManager()
	users, err := manager.GetUsers()
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users found.")
		return nil
	}

	var options []string
	for _, u := range users {
		options = append(options, fmt.Sprintf("%s <%s>", u.Name, u.Email))
	}

	repo := git.NewLocalRepository()
	
	selected := ""
	message := "Please select a user to remove:"
	currentName := repo.GetLocalName()
	currentEmail := repo.GetLocalEmail()
	if currentName != "" && currentEmail != "" {
		message = fmt.Sprintf("%s (current: %s <%s>)", message, currentName, currentEmail)
	}
	
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	
	err = survey.AskOne(prompt, &selected)
	if err != nil {
		return fmt.Errorf("survey error: %w", err)
	}

	// Parse the selected user
	s := strings.Split(selected, " ")
	r := strings.NewReplacer("<", "", ">", "")
	name := s[0]
	email := r.Replace(s[1])

	confirmed := false
	confirmPrompt := &survey.Confirm{
		Message: "Do you want to remove this user?",
	}
	
	if err := survey.AskOne(confirmPrompt, &confirmed); err != nil {
		return fmt.Errorf("survey error: %w", err)
	}

	if confirmed {
		if err := manager.RemoveUser(name, email); err != nil {
			return fmt.Errorf("failed to remove user: %w", err)
		}
		
		fmt.Printf("%s <%s> has been removed.\n", name, email)
	}
	
	return nil
}