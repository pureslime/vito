package utils

import (
	"fmt"
	"libs/ui"
	"os"
	"path/filepath"
)

func CheckPath(binDir string) {
	pathEnv := os.Getenv("PATH")
	if !contains(pathEnv, binDir) {
		fmt.Println()
		ui.PrintInfo("Note: VITO's bin directory is not in your PATH.")
		ui.PrintInfo(fmt.Sprintf("Add this to your shell profile: export PATH=\"%s:$PATH\"", binDir))
	}
}

func contains(pathEnv, binDir string) bool {
	binDir = filepath.Clean(binDir)

	dirs := filepath.SplitList(pathEnv)

	for _, d := range dirs {
		if filepath.Clean(d) == binDir {
			return true
		}
	}
	return false
}
