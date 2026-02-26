package sdk

import (
	"encoding/json"
	"fmt"
	"libs/ui"
	"os"

	"github.com/pureslime/vito-sdk/shared/structs"
)

type Pill structs.Pill

func NewPill(name string) *Pill {
	return (*Pill)(&structs.Pill{
		Name:     name,
		Handlers: make(map[string]structs.PillHandler),
	})
}

func (p *Pill) Handle(cmdName string, desc string, handler func(args []string) error) {
	p.Handlers[cmdName] = structs.PillHandler{
		Description: desc,
		Action:      handler,
	}
}

func (p *Pill) Run() {
	if len(os.Args) > 1 && os.Args[1] == "--vito-info" {
		cmds := make([]structs.CommandInfo, 0, len(p.Handlers))
		for name, h := range p.Handlers {
			cmds = append(cmds, structs.CommandInfo{
				Name:        name,
				Description: h.Description,
			})
		}
		json.NewEncoder(os.Stdout).Encode(structs.PillManifest{
			Name:     p.Name,
			Commands: cmds,
		})
		return
	}

	if len(os.Args) > 1 {
		subCmd := os.Args[1]
		if h, ok := p.Handlers[subCmd]; ok {
			if err := h.Action(os.Args[2:]); err != nil {
				p.ExitWithError(p.Name, err)
				os.Exit(1)
			}
			return
		}
	}

	p.ExitWithError(p.Name, fmt.Errorf("command '%s' not found", os.Args[1]))
}

func (p *Pill) ExitWithError(context string, err error) {
	ui.PrintErr(err, context)
	os.Exit(1)
}
