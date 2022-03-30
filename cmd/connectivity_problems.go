package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var connectivityProblemsRequest internal.ConnectivityProblemsRequest
var portStr string

var connectivityProblemsCmd = &cobra.Command{
	Use:     connectivityProblemsUse,
	Example: connectivityProblemsExample,
	Aliases: []string{"cvp"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateConnectivityProblemsFields(args, portStr, &connectivityProblemsRequest)
		svc.ConnectivityProblems(connectivityProblemsRequest)
	},
}

func init() {
	rootCmd.AddCommand(connectivityProblemsCmd)
	translateErrorCmd.Flags().SortFlags = false

	connectivityProblemsCmd.Flags().StringVarP(&connectivityProblemsRequest.EdgeLocationId, "client-location", "l", "", internal.GetMessageForKey(connectivityProblemsCmd, "client-location"))
	connectivityProblemsCmd.Flags().StringVarP(&connectivityProblemsRequest.SpoofEdgeIp, "edge-server-ip", "e", "", internal.GetMessageForKey(connectivityProblemsCmd, "edge-server-ip"))
	connectivityProblemsCmd.Flags().StringVar(&connectivityProblemsRequest.ClientIp, "client-ip", "", internal.GetMessageForKey(connectivityProblemsCmd, "clientIp"))
	connectivityProblemsCmd.Flags().StringArrayVarP(&connectivityProblemsRequest.RequestHeaders, "request-header", "H", []string{}, internal.GetMessageForKey(connectivityProblemsCmd, "requestHeader"))
	connectivityProblemsCmd.Flags().StringVarP(&connectivityProblemsRequest.IpVersion, "ip-version", "i", "", internal.GetMessageForKey(connectivityProblemsCmd, "ipVersion"))
	connectivityProblemsCmd.Flags().StringVar(&connectivityProblemsRequest.PacketType, "packet-type", "", internal.GetMessageForKey(connectivityProblemsCmd, "packetType"))
	connectivityProblemsCmd.Flags().StringVarP(&portStr, "port", "p", "", internal.GetMessageForKey(connectivityProblemsCmd, "port"))
	connectivityProblemsCmd.Flags().BoolVarP(&connectivityProblemsRequest.RunFromSiteShield, "run-from-site-shield-map", "r", false, internal.GetMessageForKey(connectivityProblemsCmd, "run-from-site-shield-map"))

	connectivityProblemsCmd.Short = internal.GetMessageForKey(connectivityProblemsCmd, internal.Short)
	connectivityProblemsCmd.Long = internal.GetMessageForKey(connectivityProblemsCmd, internal.Long)
}
