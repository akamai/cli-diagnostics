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

	"github.com/spf13/cobra"
)

// createGroupCmd represents the createGroup command
var createGroupCmd = &cobra.Command{
	Use:     userDiagnosticsCreateGroupUse,
	Aliases: []string{"creategroup", "createGroup"},
	Args:    cobra.ExactArgs(2),
	Short:   userDiagnosticsCreateGroupShortDescription,
	Long:    userDiagnosticsCreateGroupLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		var requestStruct CreateGroup
		requestStruct.GroupName = args[0]
		requestStruct.URL = args[1]

		encPayload, err := json.Marshal(requestStruct)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resp, byt := doHTTPRequest("POST", "/diagnostic-tools/v2/end-user-links?"+clientTypeKey+"="+clientTypeValue, &encPayload)

		if resp.StatusCode == 201 {
			var responseStruct map[string]string
			var responseStructJson DiagnosticLinkResponse

			err := json.Unmarshal(*byt, &responseStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				responseStructJson.GroupName = args[0]
				responseStructJson.Url = args[1]
				responseStructJson.ReportedTime = getReportedTime()
				responseStructJson.DiagnosticLink = responseStruct["diagnosticLink"]
				resJson, _ := json.MarshalIndent(responseStructJson, "", "  ")
				fmt.Println(string(resJson))
				return
			}
			colorPrintln("green", "\n"+linkText1)
			fmt.Println(linkText2)
			fmt.Printf("\n" + linkText3)
			colorPrintln("cyan", responseStruct["diagnosticLink"])

		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	userDiagnosticsCmd.AddCommand(createGroupCmd)
}
