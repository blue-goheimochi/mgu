package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattn/go-pipeline"
)

func isGitRepositoryDir() bool {
	err := exec.Command("git", "config", "--local", "--list").Run()
	if err != nil {
		return false
	}
	return true
}

func getGlobalName() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--list"},
		[]string{"grep", "user.name"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.name=", "", 1), "\n")
}

func getGlobalEmail() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--list"},
		[]string{"grep", "user.email"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.email=", "", 1), "\n")
}

func getName() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--local", "--list"},
		[]string{"grep", "user.name"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.name=", "", 1), "\n")
}

func getEmail() string {
	out, err := pipeline.Output(
		[]string{"git", "config", "--local", "--list"},
		[]string{"grep", "user.email"},
	)
	if err != nil {
		return ""
	}
	return strings.TrimRight(strings.Replace(string(out), "user.email=", "", 1), "\n")
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
