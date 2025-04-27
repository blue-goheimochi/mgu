package commands

import (
	"bytes"
	"flag"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/blue-goheimochi/mgu/pkg/config"
	"github.com/blue-goheimochi/mgu/pkg/git"
	"github.com/urfave/cli/v2"
)

// TestContext creates a CLI context for testing
func TestContext(t *testing.T) *cli.Context {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)
	return ctx
}

// CaptureOutput captures stdout during test execution
func CaptureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// TestCommandHelper provides utility functions for testing commands
type TestCommandHelper struct {
	t         *testing.T
	ctx       *cli.Context
	configDir string
	mockRepo  *git.MockRepository
}

// NewTestCommandHelper creates a new TestCommandHelper
func NewTestCommandHelper(t *testing.T) *TestCommandHelper {
	// Create a test CLI context
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	// Create a temporary directory for config files
	tempDir := t.TempDir()

	// Create a mock repository
	mockRepo := git.NewMockRepository()

	return &TestCommandHelper{
		t:         t,
		ctx:       ctx,
		configDir: tempDir,
		mockRepo:  mockRepo,
	}
}

// SetupConfig initializes a config file with test users
func (h *TestCommandHelper) SetupConfig(users []config.User) {
	// Create the config directory
	configPath := filepath.Join(h.configDir, "setting.json")
	
	// Create parent directory
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		h.t.Fatalf("Failed to create config directory: %v", err)
	}
	
	// Save original setting path and restore it later
	origPath := config.SettingFilePath
	defer func() { config.SettingFilePath = origPath }()
	
	// Set the config path to our test path
	config.SettingFilePath = configPath
	
	// Create and initialize the config file
	mgr := config.NewManager(configPath)
	if err := mgr.SaveUsers(users); err != nil {
		h.t.Fatalf("Failed to initialize config: %v", err)
	}
}

// GetContext returns the CLI context
func (h *TestCommandHelper) GetContext() *cli.Context {
	return h.ctx
}

// GetMockRepo returns the mock repository
func (h *TestCommandHelper) GetMockRepo() *git.MockRepository {
	return h.mockRepo
}