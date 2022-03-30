package internal

import (
	"text/template"
)

var UsageFuncMap = template.FuncMap{
	"bold":        bold,
	"header":      header,
	"italic":      italic,
	"capsToTitle": CapsToTitle,
	"green":       success,
	"blue":        blue,
}

var CustomUsageTemplate = `{{header "Usage:"}}{{if .Runnable}}
{{blue "akamai " .UseLine}}{{end}}{{if .HasAvailableSubCommands}}
{{blue "akamai " .CommandPath " [command]"}}{{end}}{{if gt (len .Aliases) 0}}

{{header "Aliases:"}}
{{.NameAndAliases}}{{end}}{{if .HasExample}}

{{header "Examples:"}}
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

{{header "Available Commands:"}}{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
{{rpad .Name .NamePadding | green}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{header "Flags:"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{header "Global Flags:"}}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{header "Additional help topics:"}}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
