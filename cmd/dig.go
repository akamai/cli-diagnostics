package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var digRequest internal.DigRequest

var digCmd = &cobra.Command{
	Use:     digUse,
	Example: digExample,
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateDigFields(args, &digRequest)
		svc.Dig(digRequest)
	},
}

func init() {

	rootCmd.AddCommand(digCmd)
	digCmd.Flags().SortFlags = false

	digCmd.Short = internal.GetMessageForKey(digCmd, internal.Short)
	digCmd.Long = internal.GetMessageForKey(digCmd, internal.Long)

	digCmd.Flags().StringVarP(&digRequest.Hostname, "hostname", "d", "", internal.GetMessageForKey(digCmd, "hostname"))
	digCmd.Flags().StringVarP(&digRequest.ClientLocation, "client-location", "l", "", internal.GetMessageForKey(digCmd, "client-location"))
	digCmd.Flags().StringVarP(&digRequest.EdgeServerIp, "edge-server-ip", "e", "", internal.GetMessageForKey(digCmd, "edge-server-ip"))
	digCmd.Flags().StringVarP(&digRequest.QueryType, "query-type", "q", "A", internal.GetMessageForKey(digCmd, "query-type"))
	digCmd.Flags().BoolVarP(&digRequest.IsGtmHostName, "gtm", "g", false, internal.GetMessageForKey(digCmd, "gtm"))

}
