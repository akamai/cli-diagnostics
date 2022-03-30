package cmd

import (
	"github.com/akamai/cli-diagnostics/internal"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var jsonData []byte

type GlobalFlags struct {
	edgeRcPath       string
	edgeRcSection    string
	forceColor       bool
	accountSwitchKey string
	json             bool
}

var globalFlags GlobalFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "diagnostics",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if globalFlags.forceColor {
			color.NoColor = false
		}

		if runtime.GOOS == "windows" {
			color.NoColor = true
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(Version string) {
	rootCmd.Version = Version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(internal.CliErrExitCode)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true // Remove this if we choose to offer a completion command

	jsonData = internal.ReadStdin()

	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false

	rootCmd.Short = internal.GetMessageForKey(rootCmd, internal.Short)
	rootCmd.Long = internal.GetMessageForKey(rootCmd, internal.Long)

	defaultEdgercPath := os.Getenv("AKAMAI_EDGERC")
	defaultEdgercSection := os.Getenv("AKAMAI_EDGERC_SECTION")
	jsonOutput, _ := strconv.ParseBool(os.Getenv("AKAMAI_OUTPUT_JSON"))

	if defaultEdgercPath == "" {
		if home, err := homedir.Dir(); err == nil {
			defaultEdgercPath = filepath.Join(home, ".edgerc")
		}
	}

	if defaultEdgercSection == "" {
		defaultEdgercSection = "diagnostics"
	}

	rootCmd.PersistentFlags().StringVar(&globalFlags.edgeRcPath, "edgerc", defaultEdgercPath, internal.GetMessageForKey(rootCmd, "edgerc"))
	rootCmd.PersistentFlags().StringVar(&globalFlags.edgeRcSection, "section", defaultEdgercSection, internal.GetMessageForKey(rootCmd, "section"))
	rootCmd.PersistentFlags().StringVar(&globalFlags.accountSwitchKey, "account-key", "", internal.GetMessageForKey(rootCmd, "account-key"))
	rootCmd.PersistentFlags().BoolVar(&globalFlags.forceColor, "force-color", false, internal.GetMessageForKey(rootCmd, "force-color"))
	rootCmd.PersistentFlags().BoolVar(&globalFlags.json, "json", jsonOutput, internal.GetMessageForKey(rootCmd, "json"))

	rootCmd.SetUsageTemplate(internal.CustomUsageTemplate)
	cobra.AddTemplateFuncs(internal.UsageFuncMap)

}
