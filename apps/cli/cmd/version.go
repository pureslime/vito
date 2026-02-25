package cmd

import (
	"libs/ui"

	"github.com/spf13/cobra"
)

var LibVersion = "dev-alpha"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows version of VITO",
	GroupID: "core",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintVersion("VITO", LibVersion)
	},
}
