package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestLocalRepository_IsGitRepository(t *testing.T) {
	// Create a test directory
	tempDir := t.TempDir()
	
	// Create a non-git directory
	nonGitDir := filepath.Join(tempDir, "non-git")
	if err := os.Mkdir(nonGitDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	
	// Create a git repository
	gitDir := filepath.Join(tempDir, "git-repo")
	if err := os.Mkdir(gitDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	
	// Initialize git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = gitDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repository: %v", err)
	}
	
	// Test cases
	tests := []struct {
		name string
		dir  string
		want bool
	}{
		{
			name: "non-git directory",
			dir:  nonGitDir,
			want: false,
		},
		{
			name: "git directory",
			dir:  gitDir,
			want: true,
		},
	}
	
	// Store current directory
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(currentDir) // Restore directory when test finishes
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Change to test directory
			if err := os.Chdir(tt.dir); err != nil {
				t.Fatalf("Failed to change directory: %v", err)
			}
			
			repo := NewLocalRepository()
			got := repo.IsGitRepository()
			if got != tt.want {
				t.Errorf("IsGitRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalRepository_SetLocalUser(t *testing.T) {
	// Create a git repository
	tempDir := t.TempDir()
	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to initialize git repository: %v", err)
	}
	
	// Store current directory
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(currentDir) // Restore directory when test finishes
	
	// Change to test directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	
	// Test setting user
	repo := NewLocalRepository()
	testName := "Test User"
	testEmail := "test@example.com"
	
	err = repo.SetLocalUser(testName, testEmail)
	if err != nil {
		t.Fatalf("SetLocalUser() error = %v", err)
	}
	
	// Verify the user was set correctly
	cmd = exec.Command("git", "config", "--local", "user.name")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get git config: %v", err)
	}
	
	gotName := string(output)
	// Trim newline
	if len(gotName) > 0 && gotName[len(gotName)-1] == '\n' {
		gotName = gotName[:len(gotName)-1]
	}
	
	if gotName != testName {
		t.Errorf("SetLocalUser() name = %q, want %q", gotName, testName)
	}
	
	cmd = exec.Command("git", "config", "--local", "user.email")
	output, err = cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get git config: %v", err)
	}
	
	gotEmail := string(output)
	// Trim newline
	if len(gotEmail) > 0 && gotEmail[len(gotEmail)-1] == '\n' {
		gotEmail = gotEmail[:len(gotEmail)-1]
	}
	
	if gotEmail != testEmail {
		t.Errorf("SetLocalUser() email = %q, want %q", gotEmail, testEmail)
	}
}