package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestHelper provides utility functions for testing config package
type TestHelper struct {
	TempDir    string
	SettingFile string
	t          *testing.T
}

// NewTestHelper creates a new TestHelper with a temporary directory
func NewTestHelper(t *testing.T) *TestHelper {
	tempDir := t.TempDir()
	settingFile := filepath.Join(tempDir, "setting.json")
	
	return &TestHelper{
		TempDir:    tempDir,
		SettingFile: settingFile,
		t:          t,
	}
}

// CreateEmptySettingsFile creates an empty settings file with an empty array
func (h *TestHelper) CreateEmptySettingsFile() {
	err := os.WriteFile(h.SettingFile, []byte("[]"), 0644)
	if err != nil {
		h.t.Fatalf("Failed to create empty settings file: %v", err)
	}
}

// CreateSettingsFileWithUsers creates a settings file with the given users
func (h *TestHelper) CreateSettingsFileWithUsers(users []User) {
	data, err := json.Marshal(users)
	if err != nil {
		h.t.Fatalf("Failed to marshal users: %v", err)
	}
	
	err = os.WriteFile(h.SettingFile, data, 0644)
	if err != nil {
		h.t.Fatalf("Failed to create settings file with users: %v", err)
	}
}

// ReadSettingsFile reads the settings file and returns the users
func (h *TestHelper) ReadSettingsFile() []User {
	data, err := os.ReadFile(h.SettingFile)
	if err != nil {
		h.t.Fatalf("Failed to read settings file: %v", err)
	}
	
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		h.t.Fatalf("Failed to unmarshal users: %v", err)
	}
	
	return users
}

// CreateManager creates a new Manager that uses the test settings file
func (h *TestHelper) CreateManager() *Manager {
	return NewManager(h.SettingFile)
}
