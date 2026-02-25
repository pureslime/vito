package utils

import (
	"libs/ui"
	"os"
)

func HandleError(err error, context string) {
	if err == nil {
		return
	}

	ui.PrintErr(err, context)

}

func ExitIfError(err error, context string) {
	if err != nil {
		HandleError(err, context)
		os.Exit(1)
	}
}
