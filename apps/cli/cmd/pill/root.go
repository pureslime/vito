package pill

import (
	"github.com/spf13/cobra"
)

var PillCmd = &cobra.Command{
	Use:   "pill",
	Short: "Manage VITO pills (plugins)",
	Long:  `Install, create, and manage the lifecyle of VITO pills.`,
}

func init() {
	PillCmd.AddCommand(NewPill)
}
