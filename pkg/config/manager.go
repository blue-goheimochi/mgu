package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Manager provides methods for managing user configurations
type Manager struct {
	filePath string
}

// NewManager creates a new Manager instance with the specified path
func NewManager(path string) *Manager {
	return &Manager{
		filePath: path,
	}
}

// DefaultManager creates a Manager with the default settings path
func DefaultManager() *Manager {
	return NewManager(SettingFilePath)
}

// GetUsers reads and returns all saved users
func (m *Manager) GetUsers() ([]User, error) {
	if !FileExists(m.filePath) {
		return nil, fmt.Errorf("settings file not found at %s", m.filePath)
	}

	raw, err := os.ReadFile(m.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read settings: %w", err)
	}

	var users []User
	if err := json.Unmarshal(raw, &users); err != nil {
		return nil, fmt.Errorf("failed to parse settings: %w", err)
	}

	return users, nil
}

// SaveUsers writes the provided users to the settings file
func (m *Manager) SaveUsers(users []User) error {
	bytes, err := json.Marshal(&users)
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}

	if err := os.WriteFile(m.filePath, bytes, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write settings: %w", err)
	}

	return nil
}

// AddUser adds a new user to the saved users list
func (m *Manager) AddUser(user User) error {
	users, err := m.GetUsers()
	if err != nil {
		// If the file doesn't exist, start with an empty list
		if !FileExists(m.filePath) {
			users = []User{}
		} else {
			return err
		}
	}

	users = append(users, user)
	return m.SaveUsers(users)
}

// RemoveUser removes a user from the saved users list
func (m *Manager) RemoveUser(name, email string) error {
	users, err := m.GetUsers()
	if err != nil {
		return err
	}

	var newUsers []User
	for _, u := range users {
		if u.Name != name || u.Email != email {
			newUsers = append(newUsers, u)
		}
	}

	return m.SaveUsers(newUsers)
}
