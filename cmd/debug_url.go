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
	"net"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var headersFlag []string
var edgeIPflag string

// debugUrlCmd represents the debugUrl command
var debugUrlCmd = &cobra.Command{
	Use:     debugUrlUse,
	Aliases: []string{"debugUrl", "debugurl"},
	Args:    cobra.ExactArgs(1),
	Short:   debugUrlShortDescription,
	Long:    debugUrlLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		Url, _ := url.Parse("/diagnostic-tools/v2/url-debug")
		parameters := url.Values{}
		if !checkAbsoluteURL(args[0]) {
			printWarning("URL is invalid, e.g., http://www.example.com")
			os.Exit(1)
		}
		parameters.Add("url", args[0])
		for _, hv := range headersFlag {
			parameters.Add("header", hv)
		}
		if edgeIPflag != "" {
			if ip := net.ParseIP(edgeIPflag); ip == nil {
				printWarning("Edge IP address is invalid")
				os.Exit(1)
			}
			parameters.Add("edgeIp", edgeIPflag)
		}
		Url.RawQuery = parameters.Encode()
		resp, byt := doHTTPRequest("GET", Url.String(), nil)
		if resp.StatusCode == 200 {
			var responseStruct Wrapper
			var responseStructJson DebugUrlJson

			err := json.Unmarshal(*byt, &responseStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				responseStructJson.Url = args[0]
				responseStructJson.EdgeIP = edgeIPflag
				responseStructJson.Headers = headersFlag
				responseStructJson.ReportedTime = getReportedTime()
				responseStructJson.DebugUrl = responseStruct.URLDebug
				resJson, _ := json.MarshalIndent(responseStructJson, "", "  ")
				resJson = getDecodedResponse(resJson)
				fmt.Println(string(resJson))
				return
			}

			printDebugUrlResults(responseStruct.URLDebug)
		} else {
			printResponseError(byt)
		}

	},
}

func init() {
	debugUrlCmd.Flags().StringSliceVar(&headersFlag, "header", []string{}, headerFlagDescription)
	debugUrlCmd.Flags().StringVar(&edgeIPflag, "edge-ip", "", edgeIpFlagDescription)
	rootCmd.AddCommand(debugUrlCmd)
}
