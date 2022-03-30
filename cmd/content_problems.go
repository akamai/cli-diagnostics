package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var (
	contentProblemsRequest internal.ContentProblemsRequest
)

var contentProblemsCmd = &cobra.Command{
	Use:     contentProblemsUse,
	Example: contentProblemsExample,
	Aliases: []string{"cp"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateContentProblemsFields(args, &contentProblemsRequest)

		svc.ContentProblems(contentProblemsRequest)
	},
}

func init() {

	rootCmd.AddCommand(contentProblemsCmd)
	contentProblemsCmd.Flags().SortFlags = false

	contentProblemsCmd.Short = internal.GetMessageForKey(contentProblemsCmd, internal.Short)
	contentProblemsCmd.Long = internal.GetMessageForKey(contentProblemsCmd, internal.Long)

	contentProblemsCmd.Flags().StringVarP(&contentProblemsRequest.EdgeLocationId, "client-location", "l", "", internal.GetMessageForKey(contentProblemsCmd, "client-location"))
	contentProblemsCmd.Flags().StringVarP(&contentProblemsRequest.EdgeIp, "edge-server-ip", "e", "", internal.GetMessageForKey(contentProblemsCmd, "edge-server-ip"))
	contentProblemsCmd.Flags().StringVarP(&contentProblemsRequest.IpVersion, "ip-version", "i", "", internal.GetMessageForKey(contentProblemsCmd, "ip-version"))
	contentProblemsCmd.Flags().StringArrayVarP(&contentProblemsRequest.RequestHeaders, "request-header", "H", []string{}, internal.GetMessageForKey(contentProblemsCmd, "requestHeader"))
	contentProblemsCmd.Flags().BoolVarP(&contentProblemsRequest.RunFromSiteShield, "run-from-site-shield-map", "r", false, internal.GetMessageForKey(contentProblemsCmd, "run-from-site-shield-map"))

}
