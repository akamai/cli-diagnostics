package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var translateUrlRequest internal.ArlRequest

var translateUrlCmd = &cobra.Command{
	Use:     translateUrlUse,
	Example: translateUrlExample,
	Aliases: []string{"tu"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateTranslateUrlFields(args, &translateUrlRequest)
		svc.TranslateUrl(translateUrlRequest)

	},
}

func init() {
	rootCmd.AddCommand(translateUrlCmd)
	translateUrlCmd.Flags().SortFlags = false

	translateUrlCmd.Short = internal.GetMessageForKey(translateUrlCmd, internal.Short)
	translateUrlCmd.Long = internal.GetMessageForKey(translateUrlCmd, internal.Long)

}
