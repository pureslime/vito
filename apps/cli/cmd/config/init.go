package config

import (
	"fmt"
	"libs/ui"
	"os"
	"path/filepath"

	"github.com/pureslime/vito/internal/config"
	"github.com/pureslime/vito/internal/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize default configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := filepath.Join(config.GetConfigDir(), "config.yaml")

		if _, err := os.Stat(configPath); err == nil {
			ui.PrintInfo("Config file already exists.")
			return
		}

		yamlContent := []byte(`remote:
  upgrade_url: "https://github.com/pureslime/vito/releases/latest/download"
  pills_url: "https://raw.githubusercontent.com/pureslime/vito-pills/main"
`)

		if err := os.WriteFile(configPath, yamlContent, 0644); err != nil {
			utils.HandleError(err, "config init")
			return
		}

		ui.PrintSuccess(fmt.Sprintf("Config initialized at: %s", configPath))
	},
}
