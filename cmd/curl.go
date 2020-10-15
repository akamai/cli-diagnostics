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
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var userAgentFlag string

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl <URL> <Ghost Location/edge server IP address> [--user-agent Additional user agent]",
	Args:  cobra.ExactArgs(2),
	Short: curlShortDescription,
	Long:  curlLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		id, addr := checkEdgeServerIPorLocation(args[1])
		var urlstr string
		switch id {
		case 0:
			urlstr = fmt.Sprintf("/diagnostic-tools/v2/ip-addresses/%s/curl-results", addr)
		case 1:
			urlstr = fmt.Sprintf("/diagnostic-tools/v2/ghost-locations/%s/curl-results", addr)
		case 2:
			fmt.Printf("%s", args[1])
			printWarning(" is not a valid IP address or Ghost Location")
			os.Exit(1)
		}

		var requestStruct CurlRequest

		if !checkAbsoluteURL(args[0]) {
			printWarning("URL is invalid, e.g., http://www.example.com")
			os.Exit(1)
		}
		requestStruct.Url = args[0]

		if userAgentFlag != "" {
			valErr := true
			availableAgents := []string{"Android", "Firefox", "iPhone", "Mobile", "Chrome", "MSIE", "MSIE 9", "MSIE 10", "Safari", "Safari/5", "Safari/6", "Webkit", "Webkit/5", "Webkit/6"}
			for _, agent := range availableAgents {
				if strings.ToLower(agent) == strings.ToLower(userAgentFlag) {
					valErr = false
					requestStruct.UserAgent = agent
					break
				}
			}
			if valErr {
				fmt.Printf(userAgentFlag)
				printWarning(" is not an available User Agent")
				fmt.Println("Available User Agents:", strings.Join(availableAgents, ", "))
				os.Exit(1)
			}
		}

		encPayload, err := json.Marshal(requestStruct)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resp, byt := doHTTPRequest("POST", urlstr, &encPayload)

		if resp.StatusCode == 200 {
			var responseStruct Wrapper
			var responseStructJson CurlResultsJson

			err := json.Unmarshal(*byt, &responseStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				responseStructJson.Url = args[0]
				responseStructJson.IpAddressOrLocationId = args[1]
				responseStructJson.UserAgent = userAgentFlag
				responseStructJson.ReportedTime = getReportedTime()
				responseStructJson.Curl = responseStruct.Curl
				resJson, _ := json.MarshalIndent(responseStructJson, "", "  ")
				fmt.Println(string(resJson))
				return
			}
			printCurlResults(responseStruct.Curl)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	curlCmd.Flags().StringVar(&userAgentFlag, "user-agent", "", userAgentFlagDescription)
	rootCmd.AddCommand(curlCmd)
}
