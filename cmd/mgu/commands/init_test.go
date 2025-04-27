package commands

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/blue-goheimochi/mgu/pkg/config"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name           string
		setupFunc      func(string)
		expectedOutput string
	}{
		{
			name: "successful initialization",
			setupFunc: func(tempDir string) {
				// Don't create any directories, init should create them
			},
			expectedOutput: "has been created",
		},
		{
			name: "settings file already exists",
			setupFunc: func(tempDir string) {
				// Create config directories
				configDir := filepath.Join(tempDir, ".config")
				appConfigDir := filepath.Join(configDir, "mgu")
				os.MkdirAll(appConfigDir, 0755)
				
				// Create empty settings file
				settingsFile := filepath.Join(appConfigDir, "setting.json")
				os.WriteFile(settingsFile, []byte("[]"), 0644)
			},
			expectedOutput: "already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary directory for the test
			tempDir := t.TempDir()
			
			// Save original paths
			origHomeDirPath := config.HomeDirPath
			origConfigDirPath := config.ConfigDirPath
			origAppConfigDirPath := config.AppConfigDirPath
			origSettingFilePath := config.SettingFilePath
			
			// Set paths to test directory
			config.HomeDirPath = tempDir
			config.ConfigDirPath = filepath.Join(tempDir, ".config")
			config.AppConfigDirPath = filepath.Join(config.ConfigDirPath, "mgu")
			config.SettingFilePath = filepath.Join(config.AppConfigDirPath, "setting.json")
			
			// Restore original paths after test
			defer func() {
				config.HomeDirPath = origHomeDirPath
				config.ConfigDirPath = origConfigDirPath
				config.AppConfigDirPath = origAppConfigDirPath
				config.SettingFilePath = origSettingFilePath
			}()
			
			// Setup the test case
			tt.setupFunc(tempDir)
			
			// Create a test CLI context
			ctx := TestContext(t)
			
			// Capture output
			output := CaptureOutput(func() {
				err := Init(ctx)
				if err != nil {
					t.Errorf("Init() error = %v", err)
				}
			})
			
			// Check for expected output
			if !strings.Contains(output, tt.expectedOutput) {
				t.Errorf("Init() output = %q, expected it to contain %q", output, tt.expectedOutput)
			}
			
			// Check that the settings file exists
			if _, err := os.Stat(config.SettingFilePath); os.IsNotExist(err) {
				t.Errorf("Init() did not create settings file at %s", config.SettingFilePath)
			}
			
			// Check the content of the settings file
			content, err := os.ReadFile(config.SettingFilePath)
			if err != nil {
				t.Errorf("Failed to read settings file: %v", err)
			}
			
			if string(content) != "[]" {
				t.Errorf("Settings file has wrong content, got %q, want %q", content, "[]")
			}
		})
	}
}