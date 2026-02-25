package config

import "github.com/spf13/cobra"

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "Manage VITO config.",
	Long:    "Default, repair or create configuration on VITO",
	GroupID: "core",
}

func init() {
	ConfigCmd.AddCommand(initCmd)
}
