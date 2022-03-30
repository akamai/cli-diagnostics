package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var verifyLocateIpRequest internal.VerifyLocateIpRequest

var verifyLocateIpCmd = &cobra.Command{
	Use:     verifyLocateIpUse,
	Example: verifyLocateIpExample,
	Aliases: []string{"vli"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateVerifyLocateIpFields(args, &verifyLocateIpRequest)
		svc.VerifyLocateIp(verifyLocateIpRequest)

	},
}

func init() {

	rootCmd.AddCommand(verifyLocateIpCmd)
	verifyLocateIpCmd.Flags().SortFlags = false

	verifyLocateIpCmd.Short = internal.GetMessageForKey(verifyLocateIpCmd, internal.Short)
	verifyLocateIpCmd.Long = internal.GetMessageForKey(verifyLocateIpCmd, internal.Long)

}
