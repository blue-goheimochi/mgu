package git

import (
	"errors"
	"testing"
)

func TestMockRepository(t *testing.T) {
	// Test with default mock
	t.Run("Default mock", func(t *testing.T) {
		mock := NewMockRepository()
		
		// Test IsGitRepository
		if !mock.IsGitRepository() {
			t.Errorf("IsGitRepository() = false, want true")
		}
		
		// Test GetGlobalName
		wantName := "Global User"
		if got := mock.GetGlobalName(); got != wantName {
			t.Errorf("GetGlobalName() = %q, want %q", got, wantName)
		}
		
		// Test GetGlobalEmail
		wantEmail := "global@example.com"
		if got := mock.GetGlobalEmail(); got != wantEmail {
			t.Errorf("GetGlobalEmail() = %q, want %q", got, wantEmail)
		}
		
		// Test GetLocalName
		wantLocalName := "Local User"
		if got := mock.GetLocalName(); got != wantLocalName {
			t.Errorf("GetLocalName() = %q, want %q", got, wantLocalName)
		}
		
		// Test GetLocalEmail
		wantLocalEmail := "local@example.com"
		if got := mock.GetLocalEmail(); got != wantLocalEmail {
			t.Errorf("GetLocalEmail() = %q, want %q", got, wantLocalEmail)
		}
		
		// Test SetLocalUser
		if err := mock.SetLocalUser("test", "test@example.com"); err != nil {
			t.Errorf("SetLocalUser() error = %v, want nil", err)
		}
	})
	
	// Test with custom mock functions
	t.Run("Custom mock functions", func(t *testing.T) {
		mock := NewMockRepository()
		
		// Customize IsGitRepository
		mock.IsGitRepoFunc = func() bool {
			return false
		}
		
		// Customize GetGlobalName
		customName := "Custom Global User"
		mock.GetGlobalNameFunc = func() string {
			return customName
		}
		
		// Customize SetLocalUser
		expectedError := errors.New("set local user error")
		mock.SetLocalUserFunc = func(name, email string) error {
			return expectedError
		}
		
		// Test customized functions
		if mock.IsGitRepository() {
			t.Errorf("IsGitRepository() = true, want false")
		}
		
		if got := mock.GetGlobalName(); got != customName {
			t.Errorf("GetGlobalName() = %q, want %q", got, customName)
		}
		
		err := mock.SetLocalUser("test", "test@example.com")
		if err != expectedError {
			t.Errorf("SetLocalUser() error = %v, want %v", err, expectedError)
		}
	})
}