package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var ipaHostnamesCmd = &cobra.Command{
	Use:     ipaHostnamesUse,
	Example: ipaHostnamesExample,
	Aliases: []string{"ipa"},
	Args:    cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)

		svc.IpaHostnames()
	},
}

func init() {

	rootCmd.AddCommand(ipaHostnamesCmd)
	ipaHostnamesCmd.Flags().SortFlags = false

	ipaHostnamesCmd.Short = internal.GetMessageForKey(ipaHostnamesCmd, internal.Short)
	ipaHostnamesCmd.Long = internal.GetMessageForKey(ipaHostnamesCmd, internal.Long)

}
