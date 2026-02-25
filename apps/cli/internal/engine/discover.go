package engine

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pureslime/vito/internal/config"
	"github.com/pureslime/vito/internal/shared/structs"
	"github.com/pureslime/vito/internal/utils"
)

func DiscoverPills(path string) ([]structs.PillInfo, error) {
	if err := config.EnsureDirs(); err != nil {
		return nil, err
	}

	pillsPath := config.GetPillsDir()

	files, err := os.ReadDir(pillsPath)
	if err != nil {
		utils.HandleError(err, "Searching in system pills directory: "+pillsPath)
		return nil, err
	}

	var detectedPills []structs.PillInfo

	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "vito-") {
			nameWithoutPrefix := strings.TrimPrefix(file.Name(), "vito-")
			pillName := strings.TrimSuffix(nameWithoutPrefix, filepath.Ext(nameWithoutPrefix))
			fullPath := filepath.Join(pillsPath, file.Name())

			detectedPills = append(detectedPills, structs.PillInfo{
				Name: pillName,
				Path: fullPath,
			})
		}
	}

	return detectedPills, nil
}
