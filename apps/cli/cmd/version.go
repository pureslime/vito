package cmd

import (
	"libs/ui"

	"github.com/spf13/cobra"
)

const LibVersion = "0.1.0"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows version of VITO",
	GroupID: "core",
	Run: func(cmd *cobra.Command, args []string) {
		ui.PrintVersion("VITO", LibVersion)
	},
}
