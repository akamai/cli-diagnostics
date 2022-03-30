package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var (
	testTargetIp string
)

var gtmHostnamesCmd = &cobra.Command{
	Use:     gtmHostnamesUse,
	Example: gtmHostnamesExample,
	Aliases: []string{"gtm"},
	Args:    cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)

		if testTargetIp != "" {
			svc.GtmTestTargetIp(testTargetIp)
		} else {
			svc.GtmHostnames()
		}
	},
}

func init() {

	rootCmd.AddCommand(gtmHostnamesCmd)
	gtmHostnamesCmd.Flags().SortFlags = false
	gtmHostnamesCmd.Flags().StringVarP(&testTargetIp, "test-target-ip", "t", "", internal.GetMessageForKey(gtmHostnamesCmd, internal.TestTargetIp))

	gtmHostnamesCmd.Short = internal.GetMessageForKey(gtmHostnamesCmd, internal.Short)
	gtmHostnamesCmd.Long = internal.GetMessageForKey(gtmHostnamesCmd, internal.Long)

}
