package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	
	// Create a test file
	testFile := filepath.Join(tmpDir, "test_file.txt")
	err := os.WriteFile(testFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Test cases
	tests := []struct {
		name     string
		filePath string
		want     bool
	}{
		{
			name:     "existing file",
			filePath: testFile,
			want:     true,
		},
		{
			name:     "non-existing file",
			filePath: filepath.Join(tmpDir, "nonexistent.txt"),
			want:     false,
		},
		{
			name:     "empty path",
			filePath: "",
			want:     false,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FileExists(tt.filePath)
			if got != tt.want {
				t.Errorf("FileExists(%q) = %v, want %v", tt.filePath, got, tt.want)
			}
		})
	}
}

func TestCreateDirectory(t *testing.T) {
	// Create a temporary directory for testing
	rootDir := t.TempDir()
	
	// Test cases
	tests := []struct {
		name      string
		dirPath   string
		wantError bool
	}{
		{
			name:      "create new directory",
			dirPath:   filepath.Join(rootDir, "new_dir"),
			wantError: false,
		},
		{
			name:      "create existing directory",
			dirPath:   rootDir, // This already exists
			wantError: true,
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateDirectory(tt.dirPath)
			
			if (err != nil) != tt.wantError {
				t.Errorf("CreateDirectory(%q) error = %v, wantError %v", 
					tt.dirPath, err, tt.wantError)
			}
			
			if err == nil {
				// Verify the directory was created
				if !FileExists(tt.dirPath) {
					t.Errorf("CreateDirectory(%q) did not create the directory", tt.dirPath)
				}
			}
		})
	}
}

func TestIsInitialized(t *testing.T) {
	// Save the original path and restore it after the test
	originalPath := SettingFilePath
	defer func() { SettingFilePath = originalPath }()
	
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "setting.json")
	
	// Set the SettingFilePath to our test file
	SettingFilePath = testFile
	
	// Test when the file doesn't exist
	if IsInitialized() {
		t.Errorf("IsInitialized() = true when file doesn't exist")
	}
	
	// Create the file
	err := os.WriteFile(testFile, []byte("[]"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Test when the file exists
	if !IsInitialized() {
		t.Errorf("IsInitialized() = false when file exists")
	}
}