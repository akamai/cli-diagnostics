package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var chaseOriginLogs bool
var errorTranslatorRequest internal.ErrorTranslatorRequest

var translateErrorCmd = &cobra.Command{
	Use:     translateErrorUse,
	Example: translateErrorExample,
	Aliases: []string{"tes"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateTranslateErrorFields(args, &errorTranslatorRequest, chaseOriginLogs)
		svc.TranslateError(errorTranslatorRequest)

	},
}

func init() {
	rootCmd.AddCommand(translateErrorCmd)
	translateErrorCmd.Flags().SortFlags = false

	translateErrorCmd.Short = internal.GetMessageForKey(translateErrorCmd, internal.Short)
	translateErrorCmd.Long = internal.GetMessageForKey(translateErrorCmd, internal.Long)

	translateErrorCmd.Flags().BoolVarP(&chaseOriginLogs, "chase-origin-logs", "c", false, internal.GetMessageForKey(translateErrorCmd, "chase-origin-logs"))
}
