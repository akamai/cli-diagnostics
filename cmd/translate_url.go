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
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// translateUrlCmd represents the translateUrl command
var translateUrlCmd = &cobra.Command{
	Use:     "translate-url <URL>",
	Aliases: []string{"translateUrl,translateurl"},
	Args:    cobra.ExactArgs(1),
	Short:   translateUrlShortDescription,
	Long:    translateUrlLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if !checkAbsoluteURL(args[0]) {
			printWarning("URL is invalid, e.g., http://www.example.com")
			os.Exit(1)
		}
		Url, _ := url.Parse("/diagnostic-tools/v2/translated-url")
		parameters := url.Values{}
		parameters.Add("url", args[0])
		Url.RawQuery = parameters.Encode()
		resp, byt := doHTTPRequest("GET", Url.String(), nil)

		if resp.StatusCode == 200 {
			var respStruct Wrapper
			var respStructJson TranslateURLJson

			err := json.Unmarshal(*byt, &respStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				respStructJson.Url = args[0]
				respStructJson.ReportedTime = getReportedTime()
				respStructJson.TranlatedURL = respStruct.TranlatedURL
				resJson, _ := json.MarshalIndent(respStructJson, "", "  ")
				fmt.Println(string(resJson))
				return
			}

			colorPrintf("yellow", translateUrl)
			fmt.Println()
			printLabelAndValue(typeCode, respStruct.TranlatedURL.TypeCode)
			printLabelAndValue(originServer, respStruct.TranlatedURL.OriginServer)
			printLabelAndValue(cpCode, respStruct.TranlatedURL.CpCode)
			printLabelAndValue(serialNumber, respStruct.TranlatedURL.SerialNumber)
			printLabelAndValue(ttl, respStruct.TranlatedURL.TTL)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	rootCmd.AddCommand(translateUrlCmd)

}
