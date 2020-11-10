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

// translateErrorCmd represents the translateError command
var translateErrorCmd = &cobra.Command{
	Use:     translateErrorUse,
	Aliases: []string{"translateErrorString", "translate-error", "tes", "translateError", "translateerror"},
	Args:    cobra.ExactArgs(1),
	Short:   translateErrorShortDescription,
	Long:    translateErrorLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if !validErrorString(args[0]) {
			printWarning("Invalid Error String")
			os.Exit(1)
		}
		resp, byt := doHTTPRequest("GET", fmt.Sprintf("/diagnostic-tools/v2/errors/%s/translated-error?"+clientTypeKey+"=%s", args[0], clientTypeValue), nil)
		if resp.StatusCode == 200 {
			var responseStruct Wrapper
			var responseStructJson TranslatedErrorJosn
			err := json.Unmarshal(*byt, &responseStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				responseStructJson.ErrorCode = args[0]
				responseStructJson.ReportedTime = getReportedTime()
				responseStructJson.TranslatedError = responseStruct.TranslatedError
				resJson, _ := json.MarshalIndent(responseStructJson, "", "  ")
				resJson = getDecodedResponse(resJson)
				fmt.Println(string(resJson))
				return
			}

			printTranslatedError(responseStruct.TranslatedError)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	rootCmd.AddCommand(translateErrorCmd)
}
