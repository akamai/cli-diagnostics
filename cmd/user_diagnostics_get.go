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
	"strconv"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   userDiagnosticsGetUse,
	Short: userDiagnosticsGetShortDescription,
	Args:  cobra.ExactArgs(1),
	Long:  userDiagnosticsGetLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		linkId, err := strconv.Atoi(args[0])
		if err != nil {
			printWarning("Diagnostics Link Id should be numeric.")
			os.Exit(1)
		}
		resp, byt := doHTTPRequest("GET", fmt.Sprintf("/diagnostic-tools/v2/api/end-user-links/%d", linkId), nil)

		if resp.StatusCode == 200 {
			var responseStruct UserDiagnosticData
			var responseStructJson UserDiagnosticDataJson
			err := json.Unmarshal(*byt, &responseStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				responseStructJson.LinkId = args[0]
				responseStructJson.ReportedTime = getReportedTime()
				responseStructJson.UserDiagnosticData = responseStruct
				resJson, _ := json.MarshalIndent(responseStructJson, "", "  ")
				fmt.Println(string(resJson))
				return
			}

			printUserDiagnostics(&responseStruct)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	userDiagnosticsCmd.AddCommand(getCmd)
}
