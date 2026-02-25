package templates

import (
	"fmt"
	"libs/ui"
	"os"
	"path/filepath"
)

func GenGoTemplate(name string) {
	dir := filepath.Join("pills", name)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		ui.PrintErr(err, "Error creating directory for the new pill.")
		return
	}

	mainContent := fmt.Sprintf(`package main
import (
		"fmt"
		"github.com/pureslime/vito/sdk/go"
)

func main() {
	// Create the Pill with the command name
	p := sdk.NewPill("%s")

	// Register the subcommand (example)
	p.Handle("hello", func(args []string) error {
		fmt.Println("✨ ¡Hello from native pill of Go!")
		return nil
	})

	// Execute lifecycle (Handshake + Run)
	p.Run()
}
`, name)

	err = os.WriteFile(filepath.Join(dir, "main.go"), []byte(mainContent), 0644)
	if err != nil {
		ui.PrintErr(err, "Error writting in the file 'main.go' for generating the pill.")
	}
	ui.PrintSuccess(fmt.Sprintf("Go project created on: %s", dir))
	ui.PrintInfo(fmt.Sprintf("To compile: go build -o ../vito-%s %s/main.go\n", name, dir))
}
