package config

// User represents a Git user with name and email
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}