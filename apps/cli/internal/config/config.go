package internalConfig

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pureslime/vito/internal/shared/structs"
	"github.com/pureslime/vito/internal/utils"
	"gopkg.in/yaml.v3"
)

var currentConfig structs.VitoConfig

const (
	DefaultUpgradeURL = "https://github.com/pureslime/VITO/releases/latest/download"
	DefaultPillsURL   = "https://raw.githubusercontent.com/pureslime/VITO/pills"
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
func GetBinDir() string    { return filepath.Join(GetVitoHome(), "bin") }

func EnsureDirs() error {
	dirs := []string{GetPillsDir(), GetConfigDir(), GetBinDir()}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}
	return nil
}

func LoadConfig() error {
	configFilePath := filepath.Join(GetConfigDir(), "config.yaml")

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		currentConfig.Remote.UpgradeURL = DefaultUpgradeURL
		currentConfig.Remote.PillsURL = DefaultPillsURL

		data, err := yaml.Marshal(&currentConfig)
		if err != nil {
			return err
		}

		return os.WriteFile(configFilePath, data, 0644)
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		utils.ExitIfError(fmt.Errorf("Couldn't find config on: %s", configFilePath), "config")
	}
	return yaml.Unmarshal(data, &currentConfig)
}

func GetUpgradeURL() string {
	if env := os.Getenv("VITO_UPGRADE_URL"); env != "" {
		return env
	}
	return currentConfig.Remote.UpgradeURL
}
