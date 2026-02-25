package engine

func GetHelpTemplate() string {
	return `
{{vitoTitle "USAGE"}}
  {{.UseLine}}

{{if .HasAvailableSubCommands}}
{{vitoTitle "AVAILABLE COMMANDS"}}
{{range .Commands}}{{if .IsAvailableCommand}}
  {{vitoBadge .Name 10}} {{.Short}}{{end}}{{end}}
{{end}}

{{if .HasAvailableLocalFlags}}
{{vitoTitle "FLAGS"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}

{{if .HasExample}}
{{vitoTitle "EXAMPLES"}}
{{.Example}}
{{end}}

{{if not .HasParent}}
{{vitoTitle "EXAMPLES"}}
  vito version
  vito pill list
{{end}}
`
}
