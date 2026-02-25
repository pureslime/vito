package cmd

import (
	"fmt"
	"libs/ui"
	"os"
	"runtime"

	"github.com/pureslime/vito/internal/config"
	"github.com/pureslime/vito/internal/utils"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade VITO to the latest version",
	GroupID: "core",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintInfo("Checking for updates...")

		baseURL := config.GetUpgradeURL()

		targetOS := runtime.GOOS
		targetArch := runtime.GOARCH
		if targetArch == "x86_64" {
			targetArch = "amd64"
		}

		finalURL := fmt.Sprintf("%s/vito-%s-%s", baseURL, targetOS, targetArch)

		ui.PrintInfo(fmt.Sprintf("Downloading from: %s", finalURL))

		currentBinary, _ := os.Executable()
		tmpBinary := currentBinary + ".tmp"

		if err := utils.DownloadFile(tmpBinary, finalURL); err != nil {
			utils.HandleError(err, "upgrade")
			return
		}

		os.Chmod(tmpBinary, 0755)
		if err := os.Rename(tmpBinary, currentBinary); err != nil {
			utils.HandleError(fmt.Errorf("Could not replace binary: %v", err), "upgrade")
			return
		}
		ui.PrintSuccess("VITO has been updgraded to the latest version!")
	},
}

func init() {
	upgradeCmd.GroupID = "core"
}
