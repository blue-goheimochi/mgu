package git

// Factory function type for creating repositories (used for testing)
type RepositoryFactory func() Repository

// Default factory function to create real repositories
var NewLocalRepository RepositoryFactory = func() Repository {
	return &LocalRepository{}
}

// Ensure MockRepository implements Repository
var _ Repository = (*MockRepository)(nil)