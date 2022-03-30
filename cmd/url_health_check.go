package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var (
	logs                  bool
	networkConnectivity   bool
	urlHealthCheckRequest internal.UrlHealthCheckRequest
)

var urlHealthCheckCmd = &cobra.Command{
	Use:     urlHealthCheckUse,
	Example: urlHealthCheckExample,
	Aliases: []string{"uhc", "debug-url"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateUrlHealthCheckFields(args, portStr, &urlHealthCheckRequest, logs, networkConnectivity)

		svc.UrlHealthCheck(urlHealthCheckRequest)
	},
}

func init() {

	rootCmd.AddCommand(urlHealthCheckCmd)
	urlHealthCheckCmd.Flags().SortFlags = false

	urlHealthCheckCmd.Short = internal.GetMessageForKey(urlHealthCheckCmd, internal.Short)
	urlHealthCheckCmd.Long = internal.GetMessageForKey(urlHealthCheckCmd, internal.Long)

	urlHealthCheckCmd.Flags().StringVarP(&urlHealthCheckRequest.EdgeLocationId, "client-location", "l", "", internal.GetMessageForKey(urlHealthCheckCmd, "client-location"))
	urlHealthCheckCmd.Flags().StringVarP(&urlHealthCheckRequest.EdgeIp, "edge-server-ip", "e", "", internal.GetMessageForKey(urlHealthCheckCmd, "edge-server-ip"))
	urlHealthCheckCmd.Flags().StringVarP(&portStr, "port", "p", "", internal.GetMessageForKey(urlHealthCheckCmd, "port"))
	urlHealthCheckCmd.Flags().StringVar(&urlHealthCheckRequest.PacketType, "packet-type", "", internal.GetMessageForKey(urlHealthCheckCmd, "packet-type"))
	urlHealthCheckCmd.Flags().StringVarP(&urlHealthCheckRequest.IpVersion, "ip-version", "i", "", internal.GetMessageForKey(urlHealthCheckCmd, "ip-version"))
	urlHealthCheckCmd.Flags().StringVarP(&urlHealthCheckRequest.QueryType, "query-type", "q", "", internal.GetMessageForKey(urlHealthCheckCmd, "query-type"))
	urlHealthCheckCmd.Flags().StringArrayVarP(&urlHealthCheckRequest.RequestHeaders, "request-header", "H", []string{}, internal.GetMessageForKey(urlHealthCheckCmd, "requestHeader"))
	urlHealthCheckCmd.Flags().BoolVar(&logs, "logs", false, internal.GetMessageForKey(urlHealthCheckCmd, "logs"))
	urlHealthCheckCmd.Flags().BoolVar(&networkConnectivity, "network-connectivity", false, internal.GetMessageForKey(urlHealthCheckCmd, "network-connectivity"))
	urlHealthCheckCmd.Flags().BoolVarP(&urlHealthCheckRequest.RunFromSiteShield, "run-from-site-shield-map", "r", false, internal.GetMessageForKey(urlHealthCheckCmd, "run-from-site-shield-map"))

}
