package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var httpStatusCode []string
var errorStatusCodeFlag bool
var clientRequestFlag bool
var forwardRequestFlag bool

var grepRequest internal.GrepRequest

var grepCmd = &cobra.Command{
	Use:     grepUse,
	Example: grepExample,
	Args:    cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)
		validator := internal.NewValidator(cmd, jsonData)

		validator.ValidateGrepFields(args, &grepRequest, errorStatusCodeFlag, clientRequestFlag, forwardRequestFlag, httpStatusCode)
		svc.Grep(grepRequest)

	},
}

func init() {

	rootCmd.AddCommand(grepCmd)
	grepCmd.Short = internal.GetMessageForKey(grepCmd, internal.Short)
	grepCmd.Long = internal.GetMessageForKey(grepCmd, internal.Long)

	grepCmd.Flags().StringSliceVarP(&grepRequest.Hostnames, "hostname", "d", nil, internal.GetMessageForKey(grepCmd, "hostname"))
	grepCmd.Flags().IntSliceVarP(&grepRequest.CpCodes, "cp-code", "c", nil, internal.GetMessageForKey(grepCmd, "cpCode"))
	grepCmd.Flags().StringSliceVar(&grepRequest.ClientIps, "client-ip", nil, internal.GetMessageForKey(grepCmd, "clientIp"))
	grepCmd.Flags().StringSliceVar(&grepRequest.UserAgents, "user-agent", nil, internal.GetMessageForKey(grepCmd, "userAgent"))
	grepCmd.Flags().StringSliceVar(&httpStatusCode, "http-status-code", nil, internal.GetMessageForKey(grepCmd, "httpStatusCode"))
	grepCmd.Flags().BoolVar(&errorStatusCodeFlag, "error-status-codes", false, internal.GetMessageForKey(grepCmd, "errorStatusCodes"))
	grepCmd.Flags().StringSliceVarP(&grepRequest.Arls, "arl", "a", nil, internal.GetMessageForKey(grepCmd, "arl"))
	grepCmd.Flags().BoolVarP(&clientRequestFlag, "r", "r", true, internal.GetMessageForKey(grepCmd, "r"))
	grepCmd.Flags().BoolVarP(&forwardRequestFlag, "f", "f", false, internal.GetMessageForKey(grepCmd, "f"))

}
