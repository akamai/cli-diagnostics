package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var logsEstatsFlag bool
var enhancedTlsEstatsFlag bool
var standardTlsEstatsFlag bool
var edgeErrorsEstatsFlag bool
var originErrorsEstatsFlag bool

var estatsRequest internal.EstatsRequest

var errorStatsCmd = &cobra.Command{
	Use:     estatsUse,
	Example: estatsExample,
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateEstatsFields(args, &estatsRequest, logsEstatsFlag, enhancedTlsEstatsFlag, standardTlsEstatsFlag, edgeErrorsEstatsFlag, originErrorsEstatsFlag)
		svc.Estats(estatsRequest, logsEstatsFlag)

	},
}

func init() {
	rootCmd.AddCommand(errorStatsCmd)

	errorStatsCmd.Flags().SortFlags = false

	errorStatsCmd.Short = internal.GetMessageForKey(errorStatsCmd, internal.Short)
	errorStatsCmd.Long = internal.GetMessageForKey(errorStatsCmd, internal.Long)

	errorStatsCmd.Flags().StringVarP(&estatsRequest.Url, "url", "u", "", internal.GetMessageForKey(errorStatsCmd, "url"))
	errorStatsCmd.Flags().IntVarP(&estatsRequest.CpCode, "cp-code", "c", 0, internal.GetMessageForKey(errorStatsCmd, "cpCode"))
	errorStatsCmd.Flags().BoolVar(&logsEstatsFlag, "logs", false, internal.GetMessageForKey(errorStatsCmd, "logs"))
	errorStatsCmd.Flags().BoolVar(&enhancedTlsEstatsFlag, "enhanced-tls", false, internal.GetMessageForKey(errorStatsCmd, "enhancedTls"))
	errorStatsCmd.Flags().BoolVar(&standardTlsEstatsFlag, "standard-tls", false, internal.GetMessageForKey(errorStatsCmd, "standardTls"))
	errorStatsCmd.Flags().BoolVarP(&edgeErrorsEstatsFlag, "edge-errors", "e", false, internal.GetMessageForKey(errorStatsCmd, "edgeErrors"))
	errorStatsCmd.Flags().BoolVarP(&originErrorsEstatsFlag, "origin-errors", "o", false, internal.GetMessageForKey(errorStatsCmd, "originErrors"))

}
