package cmd

import (
	"fmt"
	"libs/ui"
	"os"
	"path/filepath"
	"runtime"

	internalConfig "github.com/pureslime/vito/internal/config"
	"github.com/pureslime/vito/internal/utils"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade VITO to the latest version",
	GroupID: "core",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintInfo("Checking for updates...")

		baseURL := internalConfig.GetUpgradeURL()

		targetOS := runtime.GOOS
		targetArch := runtime.GOARCH
		if targetArch == "x86_64" {
			targetArch = "amd64"
		}

		finalURL := fmt.Sprintf("%s/vito-%s-%s", baseURL, targetOS, targetArch)

		installDir := internalConfig.GetBinDir()
		installPath := filepath.Join(installDir, "vito")
		tmpBinary := installPath + ".tmp"
		ui.PrintInfo(fmt.Sprintf("Downloading from: %s", finalURL))

		if err := utils.DownloadFile(tmpBinary, finalURL); err != nil {
			utils.HandleError(err, "upgrade")
			return
		}

		os.Chmod(tmpBinary, 0755)
		if err := os.Rename(tmpBinary, installPath); err != nil {
			utils.HandleError(fmt.Errorf("Could not replace binary: %v", err), "upgrade")
			return
		}
		ui.PrintSuccess("VITO has been updgraded to the latest version!")

		utils.CheckPath(installDir)
	},
}

func init() {
	upgradeCmd.GroupID = "core"
}
