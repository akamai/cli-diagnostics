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

var searchString string

var ghostLocationsCmd = &cobra.Command{
	Use:     ghostLocationUse,
	Aliases: []string{"ghostLocations", "ghostlocations"},
	Short:   ghostLocationShortDescription,
	Long:    ghostLocationLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

		resp, byt := doHTTPRequest("GET", "/diagnostic-tools/v2/ghost-locations/available", nil)

		if resp.StatusCode == 200 {
			var respStruct GhostLocationsList
			var respStructFiltered GhostLocationsList

			err := json.Unmarshal(*byt, &respStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				for _, loc := range respStruct.Locations {
					if strings.Contains(strings.ToLower(loc["id"]), searchString) {
						respStructFiltered.Locations = append(respStructFiltered.Locations, loc)
					}
				}
				resJson, _ := json.MarshalIndent(respStructFiltered, "", "  ")
				fmt.Println(string(resJson))
				return
			}

			printGhostLocations(&respStruct, strings.ToLower(searchString))
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	ghostLocationsCmd.Flags().StringVar(&searchString, "search", "", searchFlagDescription)
	rootCmd.AddCommand(ghostLocationsCmd)
}
