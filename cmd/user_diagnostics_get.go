package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var userDiagnosticsGetCmd = &cobra.Command{
	Use:     userDiagnosticsGetUse,
	Example: userDiagnosticsGetExample,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)

		linkId = args[0]
		svc.UserDiagnosticsGet(linkId, mtr, dig, curl)

	},
}

func init() {

	userDiagnosticsCmd.AddCommand(userDiagnosticsGetCmd)
	userDiagnosticsGetCmd.Flags().SortFlags = false

	userDiagnosticsGetCmd.Short = internal.GetMessageForKey(userDiagnosticsGetCmd, internal.Short)
	userDiagnosticsGetCmd.Long = internal.GetMessageForKey(userDiagnosticsGetCmd, internal.Long)

	userDiagnosticsGetCmd.Flags().BoolVar(&mtr, "mtr", false, internal.GetMessageForKey(userDiagnosticsGetCmd, "mtr"))
	userDiagnosticsGetCmd.Flags().BoolVar(&curl, "curl", false, internal.GetMessageForKey(userDiagnosticsGetCmd, "curl"))
	userDiagnosticsGetCmd.Flags().BoolVar(&dig, "dig", false, internal.GetMessageForKey(userDiagnosticsGetCmd, "dig"))

}
