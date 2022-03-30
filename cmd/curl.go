package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var curlRequest internal.CurlRequest

var curlCmd = &cobra.Command{
	Use:     curlUse,
	Example: curlExample,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateCurlFields(args, &curlRequest)
		svc.Curl(curlRequest)

	},
}

func init() {

	rootCmd.AddCommand(curlCmd)
	curlCmd.Flags().SortFlags = false

	curlCmd.Short = internal.GetMessageForKey(curlCmd, internal.Short)
	curlCmd.Long = internal.GetMessageForKey(curlCmd, internal.Long)

	curlCmd.Flags().StringVarP(&curlRequest.EdgeLocationId, "client-location", "l", "", internal.GetMessageForKey(curlCmd, "clientLocation"))
	curlCmd.Flags().StringVarP(&curlRequest.EdgeIp, "edge-server-ip", "e", "", internal.GetMessageForKey(curlCmd, "edgeServerIp"))
	curlCmd.Flags().StringVarP(&curlRequest.IpVersion, "ip-version", "i", internal.IPV4, internal.GetMessageForKey(curlCmd, "ipVersion"))
	curlCmd.Flags().StringArrayVarP(&curlRequest.RequestHeaders, "request-header", "H", []string{}, internal.GetMessageForKey(curlCmd, "requestHeader"))
	curlCmd.Flags().BoolVarP(&curlRequest.RunFromSiteShield, "run-from-site-shield-map", "r", false, internal.GetMessageForKey(curlCmd, "run-from-site-shield-map"))

}
