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

var resolveHostMtrFlag bool

// mtrCmd represents the mtr command
var mtrCmd = &cobra.Command{
	Use:   mtrUse,
	Args:  cobra.ExactArgs(2),
	Short: mtrShortDescription,
	Long:  mtrLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var destinationDomain string
		if ip := net.ParseIP(args[0]); ip != nil {
			destinationDomain = args[0]
		} else if _, err := url.Parse(args[0]); err == nil {
			destinationDomain = args[0]

		} else {
			printWarning("IP or domain name is invalid, e.g., 123.123.123.123 or example.com")
			os.Exit(1)
		}
		id, addr := checkEdgeServerIPorLocation(args[1])
		var urlstr string
		switch id {
		case 0:
			urlstr = fmt.Sprintf("/diagnostic-tools/v2/ip-addresses/%s/mtr-data", addr)
		case 1:
			urlstr = fmt.Sprintf("/diagnostic-tools/v2/ghost-locations/%s/mtr-data", addr)
		case 2:
			fmt.Printf("%s", args[1])
			printWarning(" is not a valid IP address or Ghost Location")
			os.Exit(1)
		}
		Url, _ := url.Parse(urlstr)
		parameters := url.Values{}
		parameters.Add("destinationDomain", destinationDomain)
		if resolveHostMtrFlag {
			parameters.Add("resolveDns", "true")
		} else {
			parameters.Add("resolveDns", "false")
		}
		Url.RawQuery = parameters.Encode()
		resp, byt := doHTTPRequest("GET", Url.String(), nil)

		if resp.StatusCode == 200 {
			var respStruct Wrapper
			var respStructJson MtrDataJson
			err := json.Unmarshal(*byt, &respStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				respStructJson.DestinationDomain = args[0]
				respStructJson.IpAddressOrLocationId = args[1]
				respStructJson.ResolveDns = resolveHostMtrFlag
				respStructJson.ReportedTime = getReportedTime()
				respStructJson.Mtr = respStruct.Mtr
				resJson, _ := json.MarshalIndent(respStructJson, "", "  ")
				fmt.Println(string(resJson))
				return
			}

			colorPrintf("blue", networkConnectivity)
			colorPrintf("yellow", respStruct.Mtr.Source)
			colorPrintf("blue", to)
			colorPrintln("yellow", respStruct.Mtr.Destination)
			fmt.Println(respStruct.Mtr.Result)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	mtrCmd.Flags().BoolVarP(&resolveHostMtrFlag, "resolve-hostname", "r", false, resolveHostMtrFlagDescription)
	rootCmd.AddCommand(mtrCmd)
}
