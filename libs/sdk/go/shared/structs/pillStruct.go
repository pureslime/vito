package structs

type PillManifest struct {
	Name     string        `json:"name"`
	Commands []CommandInfo `json:"commands"`
}

type Pill struct {
	Name     string
	Handlers map[string]PillHandler
}

type PillHandler struct {
	Description string
	Action      func(args []string) error
}

type CommandInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
