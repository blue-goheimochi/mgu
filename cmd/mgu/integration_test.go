package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIIntegration(t *testing.T) {
	// Skip in CI environment since git commands would fail
	if os.Getenv("CI") == "true" {
		t.Skip("Skipping integration test in CI environment")
	}

	// Create a temporary directory for this test
	tempDir := t.TempDir()
	
	// Create a temporary git repository
	repoDir := filepath.Join(tempDir, "testrepo")
	err := os.Mkdir(repoDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	
	// Override settings path using environment variable 
	// since the executable will be a separate process
	os.Setenv("HOME", tempDir)
	defer func() {
		os.Unsetenv("HOME")
	}()
	
	// Create settings directory structure
	settingsDir := filepath.Join(tempDir, ".config", "mgu")
	os.MkdirAll(settingsDir, 0755)
	
	// Build the current package into a temporary executable
	executable := filepath.Join(tempDir, "mgu")
	cmd := exec.Command("go", "build", "-o", executable, ".")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build executable: %v", err)
	}
	
	// Initialize git repository
	runCmd(t, "git", []string{"init"}, repoDir)
	
	// Test init command
	out := runCmd(t, executable, []string{"init"}, "")
	if !strings.Contains(out, "Initialization successful") {
		t.Errorf("Expected init command to mention initialization success, got: %s", out)
	}
	
	// Test add command
	// Note: we'd need to mock the survey prompts here for real integration testing,
	// but that's beyond the scope of this simple test
	
	// Instead, let's manually create a user in the settings file
	userData := `[{"name":"Test User","email":"test@example.com"}]`
	settingFile := filepath.Join(settingsDir, "setting.json")
	err = os.WriteFile(settingFile, []byte(userData), 0644)
	if err != nil {
		t.Fatalf("Failed to write test user data: %v", err)
	}
	
	// Test list command
	out = runCmd(t, executable, []string{"list"}, "")
	if !strings.Contains(out, "Test User <test@example.com>") {
		t.Errorf("Expected list command to show user, got: %s", out)
	}
	
	// Test set command in git repo
	// Note: we'd need to mock the survey prompts here too
	// Instead, let's just verify the show command
	
	// Test show command
	out = runCmd(t, executable, []string{"show"}, repoDir)
	if !strings.Contains(out, "not set") {
		t.Errorf("Expected show command to mention user not set, got: %s", out)
	}
}

// runCmd runs a command and returns its output
func runCmd(t *testing.T, command string, args []string, dir string) string {
	// Get the temporary home directory from the HOME environment variable
	home := os.Getenv("HOME")
	cmd := exec.Command(command, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	
	// Set HOME environment for the command to our temp dir
	cmd.Env = append(os.Environ(), "HOME="+home)
	
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	
	err := cmd.Run()
	if err != nil {
		t.Logf("Command failed: %s %v", command, args)
		t.Logf("Output: %s", out.String())
		t.Fatalf("Command execution error: %v", err)
	}
	
	return out.String()
}