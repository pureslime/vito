package structs

type PillManifest struct {
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}
