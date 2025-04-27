package git

// Repository provides methods to interact with a Git repository
type Repository interface {
	// IsGitRepository checks if the current directory is a git repository
	IsGitRepository() bool
	
	// GetGlobalName returns the global git user.name
	GetGlobalName() string
	
	// GetGlobalEmail returns the global git user.email
	GetGlobalEmail() string
	
	// GetLocalName returns the local git user.name
	GetLocalName() string
	
	// GetLocalEmail returns the local git user.email
	GetLocalEmail() string
	
	// SetLocalUser sets the local git user.name and user.email
	SetLocalUser(name, email string) error
}