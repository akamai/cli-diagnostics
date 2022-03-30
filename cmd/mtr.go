package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var mtrRequest internal.MtrRequest

var mtrCmd = &cobra.Command{
	Use:     mtrUse,
	Example: mtrExample,
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateMtrFields(args, portStr, &mtrRequest)
		svc.Mtr(mtrRequest)

	},
}

func init() {

	rootCmd.AddCommand(mtrCmd)
	mtrCmd.Flags().SortFlags = false

	mtrCmd.Short = internal.GetMessageForKey(mtrCmd, internal.Short)
	mtrCmd.Long = internal.GetMessageForKey(mtrCmd, internal.Long)

	mtrCmd.Flags().StringVarP(&mtrRequest.Source, "source", "s", "", internal.GetMessageForKey(mtrCmd, "source"))
	mtrCmd.Flags().StringVarP(&mtrRequest.Destination, "destination", "d", "", internal.GetMessageForKey(mtrCmd, "destination"))
	mtrCmd.Flags().StringVarP(&mtrRequest.GtmHostname, "gtm-hostname", "g", "", internal.GetMessageForKey(mtrCmd, "gtm-hostname"))
	mtrCmd.Flags().StringVar(&mtrRequest.PacketType, "packet-type", internal.TCP, internal.GetMessageForKey(mtrCmd, "packet-type"))
	mtrCmd.Flags().StringVarP(&portStr, "port", "p", "", internal.GetMessageForKey(mtrCmd, "port"))
	mtrCmd.Flags().StringVarP(&mtrRequest.IPVersion, "ip-version", "i", "", internal.GetMessageForKey(mtrCmd, "ip-version"))
	mtrCmd.Flags().StringVarP(&mtrRequest.SiteShieldHostname, "site-shield-hostname", "t", "", internal.GetMessageForKey(mtrCmd, "site-shield-hostname"))
	// few flags are not set to default value because they're redundant in few scenarios
}
