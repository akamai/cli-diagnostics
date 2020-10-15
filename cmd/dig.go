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
	"strings"

	"github.com/spf13/cobra"
)

var typeDigFlag string

// digCmd represents the dig command
var digCmd = &cobra.Command{
	Use:   "dig <Domain name/Hostname> <Ghost location/edge server IP address> [--type Query type]",
	Short: digShortDescription,
	Args:  cobra.ExactArgs(2),
	Long:  digLongDescription,

	Run: func(cmd *cobra.Command, args []string) {
		queryType := "A"
		validQueryTypes := []string{"A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA"}
		if typeDigFlag != "A" {
			valErr := true
			for _, qType := range validQueryTypes {
				if qType == strings.ToUpper(typeDigFlag) {
					queryType = qType
					valErr = false
					break
				}
			}
			if valErr {
				printWarning("Invalid input for Query Type")
				fmt.Println("Valid Query Types:", strings.Join(validQueryTypes, ", "))
				os.Exit(1)
			}
		}
		if _, err := url.Parse(args[0]); err != nil {
			printWarning("Hostname or domain name is invalid, e.g., foo.example.com")
			os.Exit(1)
		}
		id, addr := checkEdgeServerIPorLocation(args[1])
		var url string
		switch id {
		case 0:
			url = fmt.Sprintf("/diagnostic-tools/v2/ip-addresses/%s/dig-info?hostName=%s&queryType=%s", addr, args[0], queryType)
		case 1:
			url = fmt.Sprintf("/diagnostic-tools/v2/ghost-locations/%s/dig-info?hostName=%s&queryType=%s", addr, args[0], queryType)
		case 2:
			fmt.Printf("%s", args[1])
			printWarning(" is not a valid IP address or Ghost Location")
			os.Exit(1)
		}
		resp, byt := doHTTPRequest("GET", url, nil)
		if resp.StatusCode == 200 {
			var respStruct Wrapper
			var respStructJson DigInfoJson

			err := json.Unmarshal(*byt, &respStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				respStructJson.HostName = args[0]
				respStructJson.IpAddressOrLocationId = args[1]
				respStructJson.QueryType = typeDigFlag
				respStructJson.ReportedTime = getReportedTime()
				respStructJson.DigInfo = respStruct.DigInfo
				resJson, _ := json.MarshalIndent(respStructJson, "", "  ")
				resJson = getDecodedResponse(resJson)
				fmt.Println(string(resJson))
				return
			}

			fmt.Println(respStruct.DigInfo.Result)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	digCmd.Flags().StringVarP(&typeDigFlag, "type", "t", "A", typeDigFlagDescription)
	rootCmd.AddCommand(digCmd)
}
