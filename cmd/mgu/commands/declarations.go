package commands

import (
	"github.com/blue-goheimochi/mgu/pkg/git"
)

// defaultRepositoryFactory creates a new Git repository
func defaultRepositoryFactory() git.Repository {
	return git.NewLocalRepository()
}

// For testing - allow injecting a repository
var repositoryFactory = defaultRepositoryFactory