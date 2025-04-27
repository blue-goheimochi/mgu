package config

import (
	"reflect"
	"testing"
)

func TestManager_GetUsers(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		setupFunc   func(*TestHelper)
		want        []User
		wantErrMsg  string
	}{
		{
			name: "empty users list",
			setupFunc: func(h *TestHelper) {
				h.CreateEmptySettingsFile()
			},
			want: []User{},
		},
		{
			name: "users list with entries",
			setupFunc: func(h *TestHelper) {
				users := []User{
					{Name: "Test User 1", Email: "test1@example.com"},
					{Name: "Test User 2", Email: "test2@example.com"},
				}
				h.CreateSettingsFileWithUsers(users)
			},
			want: []User{
				{Name: "Test User 1", Email: "test1@example.com"},
				{Name: "Test User 2", Email: "test2@example.com"},
			},
		},
		{
			name:       "file does not exist",
			setupFunc:  func(h *TestHelper) {}, // Do nothing, file won't exist
			wantErrMsg: "settings file not found",
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// New helper for each test to avoid interference
			helper := NewTestHelper(t)
			
			// Setup for this test case using the new helper
			tt.setupFunc(helper)
			
			// Create a manager that uses the test settings file
			manager := NewManager(helper.SettingFile)
			
			// Test GetUsers
			got, err := manager.GetUsers()
			
			// Check error
			if tt.wantErrMsg != "" {
				if err == nil {
					t.Errorf("Expected error containing %q, but got nil", tt.wantErrMsg)
				} else if !contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("Expected error containing %q, but got %q", tt.wantErrMsg, err.Error())
				}
				return
			}
			
			// If we don't expect an error, make sure we didn't get one
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			// Check result
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_AddUser(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		setupFunc   func(*TestHelper)
		user        User
		wantUsers   []User
		wantErrMsg  string
	}{
		{
			name: "add to empty list",
			setupFunc: func(h *TestHelper) {
				h.CreateEmptySettingsFile()
			},
			user: User{Name: "New User", Email: "new@example.com"},
			wantUsers: []User{
				{Name: "New User", Email: "new@example.com"},
			},
		},
		{
			name: "add to existing list",
			setupFunc: func(h *TestHelper) {
				users := []User{
					{Name: "Existing User", Email: "existing@example.com"},
				}
				h.CreateSettingsFileWithUsers(users)
			},
			user: User{Name: "New User", Email: "new@example.com"},
			wantUsers: []User{
				{Name: "Existing User", Email: "existing@example.com"},
				{Name: "New User", Email: "new@example.com"},
			},
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new helper for each test
			helper := NewTestHelper(t)
			
			// Setup for this test case
			tt.setupFunc(helper)
			
			// Create a manager that uses the test settings file
			manager := NewManager(helper.SettingFile)
			
			// Test AddUser
			err := manager.AddUser(tt.user)
			
			// Check error
			if tt.wantErrMsg != "" {
				if err == nil {
					t.Errorf("Expected error containing %q, but got nil", tt.wantErrMsg)
				} else if !contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("Expected error containing %q, but got %q", tt.wantErrMsg, err.Error())
				}
				return
			}
			
			// If we don't expect an error, make sure we didn't get one
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			// Read the file directly to check if the user was added
			users := helper.ReadSettingsFile()
			
			// Check result
			if !reflect.DeepEqual(users, tt.wantUsers) {
				t.Errorf("After AddUser(), file contains %v, want %v", users, tt.wantUsers)
			}
		})
	}
}

func TestManager_RemoveUser(t *testing.T) {
	// Common setup
	initialUsers := []User{
		{Name: "User 1", Email: "user1@example.com"},
		{Name: "User 2", Email: "user2@example.com"},
		{Name: "User 3", Email: "user3@example.com"},
	}
	
	// Test cases
	tests := []struct {
		name       string
		removeName string
		removeEmail string
		want       []User
	}{
		{
			name:        "remove existing user",
			removeName:  "User 2",
			removeEmail: "user2@example.com",
			want: []User{
				{Name: "User 1", Email: "user1@example.com"},
				{Name: "User 3", Email: "user3@example.com"},
			},
		},
		{
			name:        "remove non-existing user",
			removeName:  "Nonexistent",
			removeEmail: "nonexistent@example.com",
			want:        initialUsers, // No change
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new helper for each test
			helper := NewTestHelper(t)
			
			// Setup fresh users for each test
			helper.CreateSettingsFileWithUsers(initialUsers)
			
			// Create a manager that uses the test settings file
			manager := NewManager(helper.SettingFile)
			
			// Test RemoveUser
			err := manager.RemoveUser(tt.removeName, tt.removeEmail)
			if err != nil {
				t.Errorf("RemoveUser() error = %v", err)
				return
			}
			
			// Read the file directly to check if the user was removed
			users := helper.ReadSettingsFile()
			
			// Check result
			if !reflect.DeepEqual(users, tt.want) {
				t.Errorf("After RemoveUser(), file contains %v, want %v", users, tt.want)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr
}