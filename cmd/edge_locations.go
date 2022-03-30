package cmd

import (
	"strings"

	"github.com/akamai/cli-diagnostics/internal"
	"github.com/spf13/cobra"
)

var (
	searchText string
)

var edgeLocationsCmd = &cobra.Command{
	Use:     edgeLocationsUse,
	Example: edgeLocationsExample,
	Aliases: []string{"el"},
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(globalFlags.edgeRcPath, globalFlags.edgeRcSection, globalFlags.accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, globalFlags.json)

		searchText = strings.ToLower(searchText)
		svc.EdgeLocations(searchText)
	},
}

func init() {

	rootCmd.AddCommand(edgeLocationsCmd)
	edgeLocationsCmd.Flags().SortFlags = false
	edgeLocationsCmd.Flags().StringVar(&searchText, "search", "", internal.GetMessageForKey(edgeLocationsCmd, internal.Search))

	edgeLocationsCmd.Short = internal.GetMessageForKey(edgeLocationsCmd, internal.Short)
	edgeLocationsCmd.Long = internal.GetMessageForKey(edgeLocationsCmd, internal.Long)

}
