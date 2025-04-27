package commands

import (
	"fmt"

	"github.com/blue-goheimochi/mgu/pkg/git"
	"github.com/urfave/cli/v2"
)

// Show displays the current Git user information
func Show(c *cli.Context) error {
	repo := git.NewLocalRepository()
	
	if !repo.IsGitRepository() {
		fmt.Println("Your current directory is not a git repository.")
		return nil
	}
	
	globalName := repo.GetGlobalName()
	globalEmail := repo.GetGlobalEmail()
	localName := repo.GetLocalName()
	localEmail := repo.GetLocalEmail()

	hasLocalGitUserSetting := true
	if localName == "" && localEmail == "" {
		fmt.Println("Your Git's local user name and email are not set.")
		hasLocalGitUserSetting = false
	} else if localName == "" {
		fmt.Println("Your Git's local user name is not set.")
		hasLocalGitUserSetting = false
	} else if localEmail == "" {
		fmt.Println("Your Git's local email is not set.")
		hasLocalGitUserSetting = false
	}
	
	if !hasLocalGitUserSetting {
		fmt.Println("Currently the following Git's global user is in use:")
		fmt.Printf("%s <%s>\n", globalName, globalEmail)
		return nil
	}
	
	fmt.Printf("%s <%s>\n", localName, localEmail)
	return nil
}