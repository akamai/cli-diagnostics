// Copyright 2020. Akamai Technologies, Inc

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	edgegrid "github.com/akamai/AkamaiOPEN-edgegrid-golang"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var edgercPath string
var edgercSection string
var forceColorFlag bool
var jsonString bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "akamai-diagnostics",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if forceColorFlag {
			color.NoColor = false
		}
		if viper.GetString("diagnostics.edgerc_path") == "" && edgercPath == "" {
			printWarning(`Value required for "edgerc" and "section" flags`)
			os.Exit(1)
		}
		if viper.GetString("diagnostics.section") == "" && edgercSection == "" {
			printWarning(`Value required for "section" flag`)
			os.Exit(1)
		}
		if edgercPath != "" {
			viper.Set("diagnostics.edgerc_path", edgercPath)
		}
		if edgercSection != "" {
			viper.Set("diagnostics.section", edgercSection)
		}

		viper.WriteConfig()
		var err error
		config, err = edgegrid.InitEdgeRc(viper.GetString("diagnostics.edgerc_path"), viper.GetString("diagnostics.section"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Suggestion: Check edgerc path and section name")
			os.Exit(1)
		}

	},
	Short: rootShortDescription,
	Long:  rootLongDescription,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVar(&edgercPath, "edgerc", home+"/.edgerc", edgercPathFlagDescription)
	rootCmd.PersistentFlags().StringVar(&edgercSection, "section", "diagnostics", edgercSectionFlagDescription)
	rootCmd.PersistentFlags().BoolVar(&forceColorFlag, "force-color", false, forceColorFlagDescription)
	rootCmd.PersistentFlags().BoolVarP(&jsonString, "json", "", false, jsonFlagDescription)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	akaHome, isPresent := os.LookupEnv("AKAMAI_CLI_HOME")
	if !isPresent {
		viper.AddConfigPath(home)
		viper.SetConfigType("ini")
		if viper.ReadInConfig() != nil {
			viper.WriteConfigAs(home + "/akamaiConfig/config.ini")
		}
	} else {
		viper.SetConfigFile(akaHome + "/config.ini")
		if err = viper.ReadInConfig(); err != nil {
			viper.Set("diagnostics.edgerc_path", "")
			viper.Set("diagnostics.section", "")
			viper.WriteConfigAs(akaHome + "/config.ini")
		}
	}
	viper.AutomaticEnv()

}
