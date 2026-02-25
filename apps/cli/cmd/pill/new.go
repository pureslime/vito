package pill

import (
	"fmt"

	"github.com/pureslime/vito/internal/templates"
	"github.com/pureslime/vito/internal/utils"
	"github.com/spf13/cobra"
)

var lang string

var newPill = &cobra.Command{
	Use:   "new [name]",
	Short: "Generate a new Pill for VITO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		switch lang {
		case "go":
			templates.GenGoTemplate(name)
		default:
			utils.HandleError(fmt.Errorf("Language not supported."), "pill new")
		}
	},
}

func init() {
	newPill.Flags().StringVarP(&lang, "lang", "l", "go", "Language of the plugin to be made (go).")
}
