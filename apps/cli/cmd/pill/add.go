package pill

import (
	"fmt"
	"io"
	"libs/ui"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pureslime/vito/internal/config"
	"github.com/pureslime/vito/internal/utils"
	"github.com/spf13/cobra"
)

var addPill = &cobra.Command{
	Use:   "add [path or url]",
	Short: "Add a pill from a local path or remote URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]

		if err := config.EnsureDirs(); err != nil {
			utils.HandleError(err, "pill add")
			return
		}

		if utils.IsURL(source) {
			addRemotePill(source)
		} else {
			addLocalPill(source)
		}
	},
}

func addLocalPill(path string) {
	info, err := os.Stat(path)

	if os.IsNotExist(err) || info.IsDir() {
		utils.HandleError(fmt.Errorf("local file not found: %s", path), "pill add")
		return
	}

	fileName := filepath.Base(path)
	if !strings.HasPrefix(fileName, "vito-") {
		fileName = "vito-" + fileName
	}

	dest := filepath.Join(config.GetPillsDir(), fileName)

	if err := utils.CopyFile(path, dest); err != nil {
		utils.HandleError(err, "pill add")
		return
	}

	os.Chmod(dest, 0755)
	ui.PrintSuccess(fmt.Sprintf("Pill added from local path: %s", fileName))
}

func addRemotePill(url string) {
	ui.PrintInfo("Downloading pill...")

	resp, err := http.Get(url)
	if err != nil {
		utils.HandleError(err, "pill add")
		return
	}
	defer resp.Body.Close()

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	if !strings.HasPrefix(fileName, "vito-") {
		fileName = "vito-" + fileName
	}

	dest := filepath.Join(config.GetPillsDir(), fileName)

	out, err := os.Create(dest)
	if err != nil {
		utils.HandleError(err, "pill add")
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		utils.HandleError(err, "pill add")
		return
	}

	os.Chmod(dest, 0755)
	ui.PrintSuccess(fmt.Sprintf("Pill downloaded and installed: %s", fileName))
}
