package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetVitoHome() string {
	if home := os.Getenv("VITO_HOME"); home != "" {
		return home
	}

	homeDir, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "vito")
	case "darwin":
		return filepath.Join(homeDir, "Library", "Application Support", "vito")
	default:
		return filepath.Join(homeDir, ".vito")
	}
}

func GetPillsDir() string  { return filepath.Join(GetVitoHome(), "pills") }
func GetConfigDir() string { return filepath.Join(GetVitoHome(), "config") }

func EnsureDirs() error {
	dirs := []string{GetPillsDir(), GetConfigDir()}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}
	return nil
}
