package engine

import (
	"path/filepath"

	"github.com/pureslime/vito/internal/shared/structs"
	"github.com/spf13/cobra"
)

func RegisterPill(rootCmd *cobra.Command, pill structs.PillInfo) error {
	manifest, err := GetManifest(pill.Path)

	if err != nil {
		simpleCmd := &cobra.Command{
			Use:                pill.Name,
			DisableFlagParsing: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				return smartExecute(pill.Path, args)
			},
		}
		rootCmd.AddCommand(simpleCmd)
		return nil
	}

	pillGroup := &cobra.Command{
		Use:   pill.Name,
		Short: "Pill: " + pill.Name,
	}

	for _, subName := range manifest.Commands {
		name := subName
		subCmd := &cobra.Command{
			Use:                name,
			DisableFlagParsing: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				return smartExecute(pill.Path, append([]string{name}, args...))
			},
		}
		pillGroup.AddCommand(subCmd)
	}
	rootCmd.AddCommand(pillGroup)
	return nil
}

func smartExecute(path string, args []string) error {
	if filepath.Ext(path) == ".wasm" {
		return ExecuteWasmPill(path, args)
	}
	return ExecutePill(path, args)
}
