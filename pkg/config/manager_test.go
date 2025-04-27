package config

import (
	"reflect"
	"testing"
)

func TestManager_GetUsers(t *testing.T) {
	// Setup test helper
	helper := NewTestHelper(t)
	
	// Test cases
	tests := []struct {
		name        string
		setupFunc   func()
		want        []User
		wantErrMsg  string
	}{
		{
			name: "empty users list",
			setupFunc: func() {
				helper.CreateEmptySettingsFile()
			},
			want: []User{},
		},
		{
			name: "users list with entries",
			setupFunc: func() {
				users := []User{
					{Name: "Test User 1", Email: "test1@example.com"},
					{Name: "Test User 2", Email: "test2@example.com"},
				}
				helper.CreateSettingsFileWithUsers(users)
			},
			want: []User{
				{Name: "Test User 1", Email: "test1@example.com"},
				{Name: "Test User 2", Email: "test2@example.com"},
			},
		},
		{
			name:       "file does not exist",
			setupFunc:  func() {}, // Do nothing, file won't exist
			wantErrMsg: "settings file not found",
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup for this test case
			tt.setupFunc()
			
			// Create a manager that uses the test settings file
			manager := helper.CreateManager()
			
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
	// Setup test helper
	helper := NewTestHelper(t)
	
	// Test cases
	tests := []struct {
		name        string
		setupFunc   func()
		user        User
		wantUsers   []User
		wantErrMsg  string
	}{
		{
			name: "add to empty list",
			setupFunc: func() {
				helper.CreateEmptySettingsFile()
			},
			user: User{Name: "New User", Email: "new@example.com"},
			wantUsers: []User{
				{Name: "New User", Email: "new@example.com"},
			},
		},
		{
			name: "add to existing list",
			setupFunc: func() {
				users := []User{
					{Name: "Existing User", Email: "existing@example.com"},
				}
				helper.CreateSettingsFileWithUsers(users)
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
			// Setup for this test case
			tt.setupFunc()
			
			// Create a manager that uses the test settings file
			manager := helper.CreateManager()
			
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
	// Setup test helper
	helper := NewTestHelper(t)
	
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
			// Setup fresh users for each test
			helper.CreateSettingsFileWithUsers(initialUsers)
			
			// Create a manager that uses the test settings file
			manager := helper.CreateManager()
			
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