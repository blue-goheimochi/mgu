package git

import (
	"fmt"
	"os/exec"
	"strings"
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
	out, err := exec.Command("git", "config", "--get", "user.name").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// GetGlobalEmail returns the global git user.email
func (r *LocalRepository) GetGlobalEmail() string {
	out, err := exec.Command("git", "config", "--get", "user.email").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// GetLocalName returns the local git user.name
func (r *LocalRepository) GetLocalName() string {
	out, err := exec.Command("git", "config", "--local", "--get", "user.name").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// GetLocalEmail returns the local git user.email
func (r *LocalRepository) GetLocalEmail() string {
	out, err := exec.Command("git", "config", "--local", "--get", "user.email").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
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
