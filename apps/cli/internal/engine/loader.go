package engine

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/pureslime/vito/internal/shared/structs"
	"github.com/pureslime/vito/internal/utils"
	"github.com/spf13/cobra"
)

func GetManifest(path string) (*structs.PillManifest, error) {
	var output []byte
	var err error

	if filepath.Ext(path) == ".wasm" {
		output, err = GetWasmOutput(path, "--vito-info")
	} else {
		cmd := exec.Command(path, "--vito-info")
		output, err = cmd.Output()
	}

	if err != nil {
		return nil, err
	}

	var m structs.PillManifest
	if err := json.Unmarshal(output, &m); err != nil {
		return nil, fmt.Errorf("Error parsing json: %w", err)
	}
	return &m, nil
}

func LoadAndRegisterPills(rootCmd *cobra.Command) {
	pills, err := DiscoverPills("pills")
	utils.ExitIfError(err, "Could not load plugins properly.")

	for _, pill := range pills {
		err := RegisterPill(rootCmd, pill)
		utils.HandleError(err, "Loading the pill"+pill.Name)
		continue
	}

}
