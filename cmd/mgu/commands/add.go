package commands

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/urfave/cli/v2"
)

// Add adds a new Git user to the configuration
func Add(c *cli.Context) error {
	if !config.IsInitialized() {
		printInitRequiredMessage()
		return nil
	}

	var qs = []*survey.Question{
		{
			Name:   "name",
			Prompt: &survey.Input{Message: "user.name"},
		},
		{
			Name:   "email",
			Prompt: &survey.Input{Message: "user.email"},
		},
	}

	answers := struct {
		Name  string
		Email string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return fmt.Errorf("survey error: %w", err)
	}

	user := config.User{
		Name:  answers.Name,
		Email: answers.Email,
	}

	manager := config.DefaultManager()
	if err := manager.AddUser(user); err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	fmt.Printf("%s <%s> has been added.\n", user.Name, user.Email)
	return nil
}
