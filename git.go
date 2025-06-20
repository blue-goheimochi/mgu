package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func isGitRepositoryDir() bool {
	err := exec.Command("git", "config", "--local", "--list").Run()
	if err != nil {
		return false
	}
	return true
}

func getGlobalName() string {
	out, err := exec.Command("git", "config", "--get", "user.name").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func getGlobalEmail() string {
	out, err := exec.Command("git", "config", "--get", "user.email").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func getName() string {
	out, err := exec.Command("git", "config", "--local", "--get", "user.name").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func getEmail() string {
	out, err := exec.Command("git", "config", "--local", "--get", "user.email").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func setLocalUser(name string, email string) error {
	if err := exec.Command("git", "config", "--local", "user.name", name).Run(); err != nil {
		return fmt.Errorf("failed to set git user.name: %w", err)
	}
	
	if err := exec.Command("git", "config", "--local", "user.email", email).Run(); err != nil {
		return fmt.Errorf("failed to set git user.email: %w", err)
	}
	
	return nil
}
