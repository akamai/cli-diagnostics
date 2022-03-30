package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var userDiagnosticsDataRequest internal.UserDiagnosticsDataRequest

var userDiagnosticsCreateCmd = &cobra.Command{
	Use:     userDiagnosticsCreateUse,
	Example: userDiagnosticsCreateExample,
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateUserDiagnosticsCreateFields(args, &userDiagnosticsDataRequest)
		svc.UserDiagnosticsCreate(userDiagnosticsDataRequest)

	},
}

func init() {

	userDiagnosticsCmd.AddCommand(userDiagnosticsCreateCmd)
	userDiagnosticsCreateCmd.Flags().SortFlags = false

	userDiagnosticsCreateCmd.Short = internal.GetMessageForKey(userDiagnosticsCreateCmd, internal.Short)
	userDiagnosticsCreateCmd.Long = internal.GetMessageForKey(userDiagnosticsCreateCmd, internal.Long)

	userDiagnosticsCreateCmd.Flags().StringVarP(&userDiagnosticsDataRequest.Url, "url", "u", "", internal.GetMessageForKey(userDiagnosticsCreateCmd, "url"))
	userDiagnosticsCreateCmd.Flags().StringVarP(&userDiagnosticsDataRequest.IpaHostname, "ipa-hostname", "", "", internal.GetMessageForKey(userDiagnosticsCreateCmd, "ipa-hostname"))
	userDiagnosticsCreateCmd.Flags().StringVarP(&userDiagnosticsDataRequest.Note, "notes", "n", "", internal.GetMessageForKey(userDiagnosticsCreateCmd, "notes"))

}
