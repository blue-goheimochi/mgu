package commands

import (
	"strings"
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/pkg/config"
)

// We need to mock the survey.Ask function
func mockSurveyAsk(answers interface{}) SurveyAskFunc {
	return func(qs []*survey.Question, response interface{}) error {
		// Type assert response to get the struct we expect
		if ptr, ok := response.(*struct {
			Name  string
			Email string
		}); ok {
			// Set the mock answers
			mockedAnswers := answers.(struct {
				Name  string
				Email string
			})
			ptr.Name = mockedAnswers.Name
			ptr.Email = mockedAnswers.Email
		}
		return nil
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name            string
		initialized     bool
		userAnswers     interface{}
		expectedOutputs []string
	}{
		{
			name:            "not initialized",
			initialized:     false,
			expectedOutputs: []string{"You need to initialize", "Please execute", "mgu init"},
		},
		{
			name:        "add user successfully",
			initialized: true,
			userAnswers: struct {
				Name  string
				Email string
			}{
				Name:  "Test User",
				Email: "test@example.com",
			},
			expectedOutputs: []string{"Test User <test@example.com> has been added"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test helper
			helper := NewTestCommandHelper(t)

			// Save original path and functions
			origPath := config.SettingFilePath
			defer func() {
				config.SettingFilePath = origPath
			}()

			// Setup config
			if tt.initialized {
				helper.SetupConfig([]config.User{})
			}

			// Set path to test file
			config.SettingFilePath = helper.SettingFile

			// Mock survey.Ask
			originalAskFunc := askFunc
			askFunc = mockSurveyAsk(tt.userAnswers)
			defer func() {
				askFunc = originalAskFunc
			}()

			// Capture output
			output := CaptureOutput(func() {
				err := Add(helper.GetContext())
				if err != nil {
					t.Errorf("Add() error = %v", err)
				}
			})

			// Check expected outputs
			for _, expected := range tt.expectedOutputs {
				if !strings.Contains(output, expected) {
					t.Errorf("Add() output = %q, should contain %q", output, expected)
				}
			}

			// Verify user was added if initialized
			if tt.initialized {
				// Check that the user was actually added to the config
				manager := config.NewManager(helper.SettingFile)
				users, err := manager.GetUsers()
				if err != nil {
					t.Errorf("Failed to get users: %v", err)
				}

				if len(users) != 1 {
					t.Errorf("Expected 1 user, got %d", len(users))
				}

				answers := tt.userAnswers.(struct {
					Name  string
					Email string
				})

				if len(users) > 0 {
					user := users[0]
					if user.Name != answers.Name || user.Email != answers.Email {
						t.Errorf("User data mismatch. Expected %s <%s>, got %s <%s>",
							answers.Name, answers.Email, user.Name, user.Email)
					}
				}
			}
		})
	}
}