package sdk

import (
	"encoding/json"
	"fmt"
	"libs/ui"
	"os"
)

type PillManifest struct {
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}

type Pill struct {
	Name     string
	handlers map[string]func(args []string) error
}

func NewPill(name string) *Pill {
	return &Pill{
		Name:     name,
		handlers: make(map[string]func(args []string) error),
	}
}

func (p *Pill) Handle(cmdName string, handler func(args []string) error) {
	p.handlers[cmdName] = handler
}

func (p *Pill) Run() {
	if len(os.Args) > 1 && os.Args[1] == "--vito-info" {
		cmds := make([]string, 0, len(p.handlers))
		for name := range p.handlers {
			cmds = append(cmds, name)
		}
		json.NewEncoder(os.Stdout).Encode(PillManifest{
			Name:     p.Name,
			Commands: cmds,
		})
		return
	}

	if len(os.Args) > 1 {
		subCmd := os.Args[1]
		if handler, ok := p.handlers[subCmd]; ok {
			if err := handler(os.Args[2:]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Printf("Pill %s: Command not found\n", p.Name)
	os.Exit(1)
}

func (p *Pill) ExitWithError(context string, err error) {
	ui.PrintErr(err, context)
	os.Exit(1)
}
