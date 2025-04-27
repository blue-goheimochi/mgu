package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattn/go-pipeline"
)

// LocalRepository implements the Repository interface for a local git repository
type LocalRepository struct{}

// IsGitRepository checks if the current directory is a git repository
func (r *LocalRepository) IsGitRepository() bool {
	err := exec.Command("git", "config", "--local", "--list").Run()
	if err != nil {
		return false
	}
	return true
}

// GetGlobalName returns the global git user.name
func (r *LocalRepository) GetGlobalName() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--list"},
		[]string{"grep", "user.name"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.name=", "", 1), "\n")
}

// GetGlobalEmail returns the global git user.email
func (r *LocalRepository) GetGlobalEmail() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--list"},
		[]string{"grep", "user.email"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.email=", "", 1), "\n")
}

// GetLocalName returns the local git user.name
func (r *LocalRepository) GetLocalName() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--local", "--list"},
		[]string{"grep", "user.name"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.name=", "", 1), "\n")
}

// GetLocalEmail returns the local git user.email
func (r *LocalRepository) GetLocalEmail() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--local", "--list"},
		[]string{"grep", "user.email"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.email=", "", 1), "\n")
}

// SetLocalUser sets the local git user.name and user.email
func (r *LocalRepository) SetLocalUser(name string, email string) error {
	if err := exec.Command("git", "config", "--local", "user.name", name).Run(); err != nil {
		return fmt.Errorf("failed to set git user.name: %w", err)
	}
	
	if err := exec.Command("git", "config", "--local", "user.email", email).Run(); err != nil {
		return fmt.Errorf("failed to set git user.email: %w", err)
	}
	
	return nil
}