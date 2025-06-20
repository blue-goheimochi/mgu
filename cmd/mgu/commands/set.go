package commands

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/blue-goheimochi/mgu/pkg/git"
	"github.com/urfave/cli/v2"
)

// Set sets the selected Git user as the local Git user
func Set(c *cli.Context) error {
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
		fmt.Println("No users found. Please add a user first with 'mgu add'.")
		return nil
	}

	var options []string
	for _, u := range users {
		options = append(options, fmt.Sprintf("%s <%s>", u.Name, u.Email))
	}

	repo := git.NewLocalRepository()
	
	selected := ""
	message := "Please select a user:"
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

	if err := repo.SetLocalUser(name, email); err != nil {
		return fmt.Errorf("failed to set local user: %w", err)
	}

	fmt.Printf("%s <%s> has been set as the local Git user.\n", name, email)
	return nil
}
