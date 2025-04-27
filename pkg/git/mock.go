package git

// MockRepository is a mock implementation of Repository for testing
type MockRepository struct {
	IsGitRepoFunc     func() bool
	GetGlobalNameFunc func() string
	GetGlobalEmailFunc func() string
	GetLocalNameFunc  func() string
	GetLocalEmailFunc func() string
	SetLocalUserFunc  func(name, email string) error
}

// NewMockRepository creates a new MockRepository with default values
func NewMockRepository() *MockRepository {
	return &MockRepository{
		IsGitRepoFunc: func() bool {
			return true
		},
		GetGlobalNameFunc: func() string {
			return "Global User"
		},
		GetGlobalEmailFunc: func() string {
			return "global@example.com"
		},
		GetLocalNameFunc: func() string {
			return "Local User"
		},
		GetLocalEmailFunc: func() string {
			return "local@example.com"
		},
		SetLocalUserFunc: func(name, email string) error {
			return nil
		},
	}
}

// IsGitRepository delegates to the mock function
func (m *MockRepository) IsGitRepository() bool {
	return m.IsGitRepoFunc()
}

// GetGlobalName delegates to the mock function
func (m *MockRepository) GetGlobalName() string {
	return m.GetGlobalNameFunc()
}

// GetGlobalEmail delegates to the mock function
func (m *MockRepository) GetGlobalEmail() string {
	return m.GetGlobalEmailFunc()
}

// GetLocalName delegates to the mock function
func (m *MockRepository) GetLocalName() string {
	return m.GetLocalNameFunc()
}

// GetLocalEmail delegates to the mock function
func (m *MockRepository) GetLocalEmail() string {
	return m.GetLocalEmailFunc()
}

// SetLocalUser delegates to the mock function
func (m *MockRepository) SetLocalUser(name, email string) error {
	return m.SetLocalUserFunc(name, email)
}
