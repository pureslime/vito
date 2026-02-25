package engine

func GetHelpTemplate() string {
	return `
{{vitoTitle "USAGE"}}
  {{.UseLine}}

{{- if .HasAvailableSubCommands}}
{{- range .Groups}}

{{vitoTitle .Title}}
{{- $groupID := .ID}}
{{- range $.Commands}}
{{- if (and (eq .GroupID $groupID) (.IsAvailableCommand))}}
  {{vitoBadge .Name 12}} {{.Short}}
{{- end}}
{{- end}}
{{- end}}

{{- /* Lógica para OTHER COMMANDS */ -}}
{{- $hasOrphans := false -}}
{{- range .Commands -}}
    {{- if (and (not .GroupID) (.IsAvailableCommand) (ne .Name "help") (ne .Name "completion")) -}}
        {{- $hasOrphans = true -}}
    {{- end -}}
{{- end}}

{{- if $hasOrphans}}

{{vitoTitle "OTHER COMMANDS"}}
{{- range .Commands}}
{{- if (and (not .GroupID) (.IsAvailableCommand) (ne .Name "help") (ne .Name "completion"))}}
  {{vitoBadge .Name 12}} {{.Short}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}

{{if .HasAvailableLocalFlags}}
{{vitoTitle "FLAGS"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}

{{if .HasExample}}
{{vitoTitle "EXAMPLES"}}
{{.Example}}
{{else if not .HasParent}}
{{vitoTitle "EXAMPLES"}}
  vito version
  vito pill list
{{end}}
`
}
