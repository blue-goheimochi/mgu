package commands

import (
	"strings"
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/blue-goheimochi/mgu/pkg/git"
)

// Mock for survey.AskOne
func mockAskOne(selectedValue string) SurveyAskOneFunc {
	return func(prompt survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
		if strPtr, ok := response.(*string); ok {
			*strPtr = selectedValue
		}
		return nil
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name               string
		initialized        bool
		users              []config.User
		isGitRepo          bool
		selectedUser       string
		currentLocalName   string
		currentLocalEmail  string
		expectedOutputs    []string
		unexpectedOutputs  []string
		expectSetLocalUser bool
		expectedSetName    string
		expectedSetEmail   string
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
			expectedOutputs: []string{"No users found", "add a user first"},
		},
		{
			name:        "select user successfully",
			initialized: true,
			users: []config.User{
				{Name: "User1", Email: "user1@example.com"},
				{Name: "User2", Email: "user2@example.com"},
			},
			isGitRepo:          true,
			selectedUser:       "User1 <user1@example.com>",
			expectSetLocalUser: true,
			expectedSetName:    "User1",
			expectedSetEmail:   "user1@example.com",
			expectedOutputs:    []string{"User1 <user1@example.com> has been set"},
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
			expectSetLocalUser: true,
			expectedSetName:    "User2",
			expectedSetEmail:   "user2@example.com",
			expectedOutputs:    []string{"User2 <user2@example.com> has been set"},
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

			var setLocalUserCalled bool
			var setLocalUserName, setLocalUserEmail string
			mockRepo.SetLocalUserFunc = func(name, email string) error {
				setLocalUserCalled = true
				setLocalUserName = name
				setLocalUserEmail = email
				return nil
			}

			// Setup mock repository factory
			repositoryFactory = func() git.Repository {
				return mockRepo
			}

			// Setup mock AskOne
			askOneFunc = mockAskOne(tt.selectedUser)

			// Capture output
			output := CaptureOutput(func() {
				err := Set(helper.GetContext())
				if err != nil {
					t.Errorf("Set() error = %v", err)
				}
			})

			// Check expected outputs
			for _, expected := range tt.expectedOutputs {
				if !strings.Contains(output, expected) {
					t.Errorf("Set() output = %q, should contain %q", output, expected)
				}
			}

			// Check unexpected outputs
			for _, unexpected := range tt.unexpectedOutputs {
				if strings.Contains(output, unexpected) {
					t.Errorf("Set() output = %q, should not contain %q", output, unexpected)
				}
			}

			// Check if SetLocalUser was called
			if tt.expectSetLocalUser && !setLocalUserCalled {
				t.Errorf("Set() should have called SetLocalUser but did not")
			} else if !tt.expectSetLocalUser && setLocalUserCalled {
				t.Errorf("Set() should not have called SetLocalUser but did")
			}

			// Check set user values
			if tt.expectSetLocalUser {
				if setLocalUserName != tt.expectedSetName {
					t.Errorf("Set() set name = %q, want %q", setLocalUserName, tt.expectedSetName)
				}
				if setLocalUserEmail != tt.expectedSetEmail {
					t.Errorf("Set() set email = %q, want %q", setLocalUserEmail, tt.expectedSetEmail)
				}
			}
		})
	}
}