package config

import (
	"os"
	"path/filepath"
)

var (
	// HomeDirPath is the path to the user's home directory
	HomeDirPath, _ = os.UserHomeDir()
	
	// ConfigDirPath is the path to the user's .config directory
	ConfigDirPath = filepath.Join(HomeDirPath, ".config")
	
	// AppConfigDirPath is the path to the application's config directory
	AppConfigDirPath = filepath.Join(ConfigDirPath, "mgu")
	
	// SettingFilePath is the path to the application's settings file
	SettingFilePath = filepath.Join(AppConfigDirPath, "setting.json")
)