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

// listCmd represents the list command
var userDiagnosticsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   userDiagnosticsListShortDescription,
	Long:    userDiagnosticsListLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		resp, byt := doHTTPRequest("GET", "/diagnostic-tools/v2/end-user-links", nil)
		if resp.StatusCode == 200 {
			var responseStruct Wrapper
			var responseStructJson EndUserDiagnosticLinkJson

			err := json.Unmarshal(*byt, &responseStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				responseStructJson.ReportedTime = getReportedTime()
				responseStructJson.EndUserDiagnosticLinks = responseStruct.EndUserDiagnosticLinks
				resJson, _ := json.MarshalIndent(responseStructJson, "", "  ")
				fmt.Println(string(resJson))
				return
			}

			printListUserDiagnosticData(responseStruct.EndUserDiagnosticLinks)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	userDiagnosticsCmd.AddCommand(userDiagnosticsListCmd)
}
