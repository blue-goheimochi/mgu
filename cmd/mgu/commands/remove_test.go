package commands

import (
	"strings"
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/blue-goheimochi/mgu/pkg/git"
)

// Mock AskOne behavior specifically for Remove command
func mockRemoveAskOne(selectedUser string, confirmed bool) SurveyAskOneFunc {
	return func(prompt survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
		// Check prompt type to determine what to respond with
		switch prompt.(type) {
		case *survey.Select:
			if strPtr, ok := response.(*string); ok {
				*strPtr = selectedUser
			}
		case *survey.Confirm:
			if boolPtr, ok := response.(*bool); ok {
				*boolPtr = confirmed
			}
		}
		return nil
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name               string
		initialized        bool
		users              []config.User
		isGitRepo          bool
		selectedUser       string
		confirmRemoval     bool
		currentLocalName   string
		currentLocalEmail  string
		expectedOutputs    []string
		unexpectedOutputs  []string
		expectRemoveUser   bool
		expectedRemovedName  string
		expectedRemovedEmail string
	}{
		{
			name:            "not initialized",
			initialized:     false,
			expectedOutputs: []string{"You need to initialize", "Please execute", "mgu init"},
		},
		{
			name:            "no users found",
			initialized:     true,
			users:           []config.User{},
			expectedOutputs: []string{"No users found"},
		},
		{
			name:        "remove user - confirmed",
			initialized: true,
			users: []config.User{
				{Name: "User1", Email: "user1@example.com"},
				{Name: "User2", Email: "user2@example.com"},
			},
			isGitRepo:          true,
			selectedUser:       "User1 <user1@example.com>",
			confirmRemoval:     true,
			expectRemoveUser:   true,
			expectedRemovedName:  "User1",
			expectedRemovedEmail: "user1@example.com",
			expectedOutputs:    []string{"User1 <user1@example.com> has been removed"},
		},
		{
			name:        "remove user - canceled",
			initialized: true,
			users: []config.User{
				{Name: "User1", Email: "user1@example.com"},
				{Name: "User2", Email: "user2@example.com"},
			},
			isGitRepo:          true,
			selectedUser:       "User2 <user2@example.com>",
			confirmRemoval:     false,
			expectRemoveUser:   false,
			unexpectedOutputs:  []string{"has been removed"},
		},
		{
			name:        "shows current user",
			initialized: true,
			users: []config.User{
				{Name: "User1", Email: "user1@example.com"},
				{Name: "User2", Email: "user2@example.com"},
			},
			isGitRepo:          true,
			currentLocalName:   "Current",
			currentLocalEmail:  "current@example.com",
			selectedUser:       "User2 <user2@example.com>",
			confirmRemoval:     true,
			expectRemoveUser:   true,
			expectedRemovedName:  "User2",
			expectedRemovedEmail: "user2@example.com",
			expectedOutputs:    []string{"User2 <user2@example.com> has been removed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test helper
			helper := NewTestCommandHelper(t)

			// Save original values
			origPath := config.SettingFilePath
			origAskOneFunc := askOneFunc
			origRepoFactory := repositoryFactory

			defer func() {
				config.SettingFilePath = origPath
				askOneFunc = origAskOneFunc
				repositoryFactory = origRepoFactory
			}()

			// Setup config
			if tt.initialized {
				helper.SetupConfig(tt.users)
			}

			// Set path to test file
			config.SettingFilePath = helper.SettingFile

			// Setup mock repository
			mockRepo := helper.GetMockRepo()
			mockRepo.IsGitRepoFunc = func() bool {
				return tt.isGitRepo
			}
			mockRepo.GetLocalNameFunc = func() string {
				return tt.currentLocalName
			}
			mockRepo.GetLocalEmailFunc = func() string {
				return tt.currentLocalEmail
			}

			// Setup mock repository factory
			repositoryFactory = func() git.Repository {
				return mockRepo
			}

			// Setup mock AskOne
			askOneFunc = mockRemoveAskOne(tt.selectedUser, tt.confirmRemoval)

			// Track if user was removed
			var userRemoved bool
			var removedName, removedEmail string

			// Capture output
			output := CaptureOutput(func() {
				// We'll check the users before and after to verify if RemoveUser was called
				var beforeUsers []config.User
				if tt.initialized {
					manager := config.NewManager(helper.SettingFile)
					beforeUsers, _ = manager.GetUsers()
				}

				err := Remove(helper.GetContext())
				if err != nil {
					t.Errorf("Remove() error = %v", err)
				}

				// Check if a user was removed by comparing before and after
				if tt.initialized && tt.expectRemoveUser {
					manager := config.NewManager(helper.SettingFile)
					afterUsers, _ := manager.GetUsers()
					
					if len(beforeUsers) > len(afterUsers) {
						userRemoved = true
						
						// Find which user was removed
						for _, before := range beforeUsers {
							found := false
							for _, after := range afterUsers {
								if before.Name == after.Name && before.Email == after.Email {
									found = true
									break
								}
							}
							if !found {
								removedName = before.Name
								removedEmail = before.Email
								break
							}
						}
					}
				}
			})

			// Check expected outputs
			for _, expected := range tt.expectedOutputs {
				if !strings.Contains(output, expected) {
					t.Errorf("Remove() output = %q, should contain %q", output, expected)
				}
			}

			// Check unexpected outputs
			for _, unexpected := range tt.unexpectedOutputs {
				if strings.Contains(output, unexpected) {
					t.Errorf("Remove() output = %q, should not contain %q", output, unexpected)
				}
			}

			// Check if user was removed as expected
			if tt.expectRemoveUser && !userRemoved {
				t.Errorf("Remove() should have removed a user but did not")
			} else if !tt.expectRemoveUser && userRemoved {
				t.Errorf("Remove() should not have removed a user but did")
			}

			// Check removed user name and email
			if tt.expectRemoveUser && userRemoved {
				if removedName != tt.expectedRemovedName {
					t.Errorf("Remove() removed user name = %q, want %q", removedName, tt.expectedRemovedName)
				}
				if removedEmail != tt.expectedRemovedEmail {
					t.Errorf("Remove() removed user email = %q, want %q", removedEmail, tt.expectedRemovedEmail)
				}
			}
		})
	}
}