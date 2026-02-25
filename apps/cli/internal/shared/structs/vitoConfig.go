package structs

type VitoConfig struct {
	Remote struct {
		UpgradeURL string `yaml:"upgrade_url"`
		PillsURL   string `yaml:"pills_url"`
	} `yaml:"remote"`
}
