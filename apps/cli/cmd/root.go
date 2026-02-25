package cmd

import (
	"fmt"
	"libs/ui"

	"github.com/pureslime/vito/cmd/config"
	"github.com/pureslime/vito/cmd/pill"
	internalConfig "github.com/pureslime/vito/internal/config"
	"github.com/pureslime/vito/internal/engine"
	"github.com/pureslime/vito/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vito",
	Short: "Modular CLI for modern devs.",
	Long:  "VITO is a CLI designed to be extensible and expandible for any workflow.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'vito --help' for seeing all the commands.")
	},
}

func init() {

	rootCmd.AddGroup(&cobra.Group{ID: "core", Title: "CORE COMMANDS"})
	rootCmd.AddGroup(&cobra.Group{ID: "pills", Title: "INSTALLED PILLS"})

	cobra.AddTemplateFunc("vitoBadge", func(s string, padding int) string {
		formatted := fmt.Sprintf("%-*s", padding, s)
		return ui.Badge(s) + formatted[len(s):]
	})

	cobra.AddTemplateFunc("vitoTitle", func(s string) string {
		return ui.MakeTitle(s)
	})

	cobra.OnInitialize(func() {
		if err := internalConfig.EnsureDirs(); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
		}

		if err := internalConfig.LoadConfig(); err != nil {
			fmt.Printf("Error loading config: %v\n", err)
		}
	})

	// Set internal "pills"
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(pill.PillCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(upgradeCmd)

	// Set help template
	rootCmd.SetHelpTemplate(engine.GetHelpTemplate())
	rootCmd.SetUsageTemplate(engine.GetHelpTemplate())
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	engine.LoadAndRegisterPills(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.ExitIfError(err, "Executing command")
	}
}
