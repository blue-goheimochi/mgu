package commands

import (
	"strings"
	"testing"

	"github.com/blue-goheimochi/mgu/pkg/config"
)

func TestList(t *testing.T) {
	tests := []struct {
		name            string
		setupUsers      []config.User
		expectedOutputs []string
	}{
		{
			name:            "empty list",
			setupUsers:      []config.User{},
			expectedOutputs: []string{"No users found"},
		},
		{
			name: "list with users",
			setupUsers: []config.User{
				{Name: "User 1", Email: "user1@example.com"},
				{Name: "User 2", Email: "user2@example.com"},
			},
			expectedOutputs: []string{
				"User 1 <user1@example.com>",
				"User 2 <user2@example.com>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test helper
			helper := NewTestCommandHelper(t)
			
			// Save original path
			origPath := config.SettingFilePath
			defer func() { config.SettingFilePath = origPath }()
			
			// Setup config
			helper.SetupConfig(tt.setupUsers)
			
			// Set path to test file
			config.SettingFilePath = helper.SettingFile
			
			// Capture output
			output := CaptureOutput(func() {
				err := List(helper.GetContext())
				if err != nil {
					t.Errorf("List() error = %v", err)
				}
			})
			
			// Check expected outputs
			for _, expected := range tt.expectedOutputs {
				if !strings.Contains(output, expected) {
					t.Errorf("List() output = %q, should contain %q", output, expected)
				}
			}
		})
	}
}