package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var (
	urlFlag string
	user    string
	active  bool
	linkId  string
	mtr     bool
	dig     bool
	curl    bool
)

var userDiagnosticsCmd = &cobra.Command{
	Use:     userDiagnosticsUse,
	Aliases: []string{"ud"},
}

func init() {
	rootCmd.AddCommand(userDiagnosticsCmd)

	userDiagnosticsCmd.Short = internal.GetMessageForKey(userDiagnosticsCmd, internal.Short)
	userDiagnosticsCmd.Long = internal.GetMessageForKey(userDiagnosticsCmd, internal.Long)

}
