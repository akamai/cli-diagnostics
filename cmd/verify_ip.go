package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var verifyIpsRequest internal.VerifyLocateIpsRequest

var verifyIpCmd = &cobra.Command{
	Use:     verifyIpUse,
	Example: verifyIpExample,
	Aliases: []string{"vi"},
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateVerifyIpOrLocateIpFields(args, &verifyIpsRequest)
		svc.VerifyIp(verifyIpsRequest)

	},
}

func init() {

	rootCmd.AddCommand(verifyIpCmd)
	verifyIpCmd.Flags().SortFlags = false

	verifyIpCmd.Short = internal.GetMessageForKey(verifyIpCmd, internal.Short)
	verifyIpCmd.Long = internal.GetMessageForKey(verifyIpCmd, internal.Long)

}
