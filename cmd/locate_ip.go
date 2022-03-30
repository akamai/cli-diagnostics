package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var locateIpsRequest internal.VerifyLocateIpsRequest

var locateIpCmd = &cobra.Command{
	Use:     locateIpUse,
	Example: locateIpExample,
	Aliases: []string{"li"},
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateVerifyIpOrLocateIpFields(args, &locateIpsRequest)
		svc.LocateIp(locateIpsRequest)

	},
}

func init() {

	rootCmd.AddCommand(locateIpCmd)
	locateIpCmd.Flags().SortFlags = false

	locateIpCmd.Short = internal.GetMessageForKey(locateIpCmd, internal.Short)
	locateIpCmd.Long = internal.GetMessageForKey(locateIpCmd, internal.Long)
}
