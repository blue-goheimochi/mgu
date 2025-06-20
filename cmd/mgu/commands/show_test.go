package commands

import (
	"strings"
	"testing"
	
	"github.com/blue-goheimochi/mgu/pkg/git"
)

func TestShow(t *testing.T) {
	tests := []struct {
		name              string
		isGitRepo         bool
		localName         string
		localEmail        string
		globalName        string
		globalEmail       string
		expectedOutputs   []string
		unexpectedOutputs []string
	}{
		{
			name:            "not a git repository",
			isGitRepo:       false,
			expectedOutputs: []string{"not a git repository"},
		},
		{
			name:              "no local user set",
			isGitRepo:         true,
			localName:         "",
			localEmail:        "",
			globalName:        "Global User",
			globalEmail:       "global@example.com",
			expectedOutputs:   []string{"not set", "global user", "Global User <global@example.com>"},
			unexpectedOutputs: []string{"Local User"},
		},
		{
			name:              "local name not set",
			isGitRepo:         true,
			localName:         "",
			localEmail:        "local@example.com",
			globalName:        "Global User",
			globalEmail:       "global@example.com",
			expectedOutputs:   []string{"name is not set", "global user", "Global User <global@example.com>"},
			unexpectedOutputs: []string{"Local User"},
		},
		{
			name:              "local email not set",
			isGitRepo:         true,
			localName:         "Local User",
			localEmail:        "",
			globalName:        "Global User",
			globalEmail:       "global@example.com",
			expectedOutputs:   []string{"email is not set", "global user", "Global User <global@example.com>"},
			unexpectedOutputs: []string{"local@example.com"},
		},
		{
			name:              "local user set",
			isGitRepo:         true,
			localName:         "Local User",
			localEmail:        "local@example.com",
			globalName:        "Global User",
			globalEmail:       "global@example.com",
			expectedOutputs:   []string{"Local User <local@example.com>"},
			unexpectedOutputs: []string{"not set", "global user", "Global User"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test helper
			helper := NewTestCommandHelper(t)
			
			// Setup mock repository
			mockRepo := helper.GetMockRepo()
			mockRepo.IsGitRepoFunc = func() bool {
				return tt.isGitRepo
			}
			mockRepo.GetLocalNameFunc = func() string {
				return tt.localName
			}
			mockRepo.GetLocalEmailFunc = func() string {
				return tt.localEmail
			}
			mockRepo.GetGlobalNameFunc = func() string {
				return tt.globalName
			}
			mockRepo.GetGlobalEmailFunc = func() string {
				return tt.globalEmail
			}
			
			// Replace the factory function with one that returns our mock
			originalFactory := repositoryFactory
			repositoryFactory = func() git.Repository {
				return mockRepo
			}
			defer func() { repositoryFactory = originalFactory }()
			
			// Capture output
			output := CaptureOutput(func() {
				err := Show(helper.GetContext())
				if err != nil {
					t.Errorf("Show() error = %v", err)
				}
			})
			
			// Check for expected outputs
			for _, expected := range tt.expectedOutputs {
				if !strings.Contains(output, expected) {
					t.Errorf("Show() output = %q, should contain %q", output, expected)
				}
			}
			
			// Check for unexpected outputs
			for _, unexpected := range tt.unexpectedOutputs {
				if strings.Contains(output, unexpected) {
					t.Errorf("Show() output = %q, should not contain %q", output, unexpected)
				}
			}
		})
	}
}