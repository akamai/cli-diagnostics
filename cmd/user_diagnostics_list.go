package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var userDiagnosticsListCmd = &cobra.Command{
	Use:     userDiagnosticsListUse,
	Example: userDiagnosticsListExample,
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)

		svc.UserDiagnosticsList(urlFlag, user, active)

	},
}

func init() {

	userDiagnosticsCmd.AddCommand(userDiagnosticsListCmd)
	userDiagnosticsListCmd.Flags().SortFlags = false

	userDiagnosticsListCmd.Short = internal.GetMessageForKey(userDiagnosticsListCmd, internal.Short)
	userDiagnosticsListCmd.Long = internal.GetMessageForKey(userDiagnosticsListCmd, internal.Long)

	userDiagnosticsListCmd.Flags().StringVarP(&urlFlag, "url", "u", "", internal.GetMessageForKey(userDiagnosticsListCmd, "url"))
	userDiagnosticsListCmd.Flags().StringVar(&user, "user", "", internal.GetMessageForKey(userDiagnosticsListCmd, "user"))
	userDiagnosticsListCmd.Flags().BoolVarP(&active, "active", "a", false, internal.GetMessageForKey(userDiagnosticsListCmd, "active"))

}
