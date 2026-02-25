package pill

import (
	"github.com/spf13/cobra"
)

var PillCmd = &cobra.Command{
	Use:     "pill",
	Short:   "Manage VITO pills (plugins)",
	Long:    `Install, create, and manage the lifecyle of VITO pills.`,
	GroupID: "core",
}

func init() {
	PillCmd.AddCommand(newPill)
	PillCmd.AddCommand(addPill)
}
