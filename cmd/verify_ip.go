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
	"os"

	"github.com/spf13/cobra"
)

//verifyIpCmd represents the verifyIp command
var verifyIpCmd = &cobra.Command{
	Use:     verifyIpUse,
	Aliases: []string{"verifyIp", "verifyip"},
	Args:    cobra.ExactArgs(1),
	Short:   verifyIpShortDescription,
	Long:    verifyIpLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if ip := net.ParseIP(args[0]); ip == nil {
			printWarning("IP address is invalid")
			os.Exit(1)
		}

		resp, byt := doHTTPRequest("GET", fmt.Sprintf("/diagnostic-tools/v2/ip-addresses/%s/is-cdn-ip?"+clientTypeKey+"=%s", args[0], clientTypeValue), nil)
		if resp.StatusCode == 200 {
			var respStruct map[string]bool
			var respStructJson VerifyIpJson

			err := json.Unmarshal(*byt, &respStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				respStructJson.IpAddress = args[0]
				respStructJson.ReportedTime = getReportedTime()
				respStructJson.IsCdnIp = respStruct["isCdnIp"]
				resJson, _ := json.MarshalIndent(respStructJson, "", "  ")
				fmt.Println(string(resJson))
				return
			}

			fmt.Printf(ipAddress)
			colorPrintf("yellow", args[0])
			if respStruct["isCdnIp"] {
				colorPrintln("green", isCdnIpSuccess)
			} else {
				colorPrintln("red", isCdnIpFailure)
			}
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyIpCmd)
}
