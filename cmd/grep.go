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
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var endTimeFlag string
var endDateFlag string
var maxLinesFlag, durationFlag int
var clientRequestFlag bool
var forwardRequestFlag bool
var findInFlag []string

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep <Edge server IP> <--end-date Date> <--end-time Time> <--duration Duration> [--find-in Header:Value ...] [--max-lines Maximum log lines to display] <-r | -f | -rf> ",
	Args:  cobra.ExactArgs(1),
	Short: grepShortDescription,
	Long:  grepLongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if ip := net.ParseIP(args[0]); ip == nil {
			printWarning("IP address is invalid")
			os.Exit(1)
		}
		if durationFlag > 360 {
			printWarning("Duration before end time cannot exceed 6 hours")
			os.Exit(1)
		}
		if maxLinesFlag > 1000 || maxLinesFlag < 0 {
			printWarning("Max log lines ranges from 0 to 1000")
			os.Exit(1)
		}
		if !clientRequestFlag && !forwardRequestFlag {
			printWarning("Select atleast one type of record")
			fmt.Println("-r : Search logs of client requests to the edge server")
			fmt.Println("-f : Search logs of forwarded requests from the edge server")
			os.Exit(1)

		}
		datetime, err := time.Parse(time.RFC3339, endDateFlag+"T"+endTimeFlag+"Z")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		loc, _ := time.LoadLocation("UTC")
		diff := time.Now().In(loc).Sub(datetime).Hours()
		if diff > 48 || diff < 0 {
			printWarning("End time (UTC) of log search: End time is invalid. Logs are generally available for last 48 hours")
			os.Exit(1)
		}

		parameters := url.Values{}
		parameters.Add("endTime", endDateFlag+"T"+endTimeFlag+"Z")

		logType := "r"
		Url, _ := url.Parse(fmt.Sprintf("/diagnostic-tools/v2/ip-addresses/%s/log-lines", args[0]))

		validQueries := []string{"host-header", "user-agent", "http-status-code", "arl", "cp-code", "client-ip"}
		for _, str := range findInFlag {
			kv := strings.SplitN(str, ":", 2)
			if len(kv) != 2 {
				printWarning("Invalid Syntax, specify as <field>:<value>")
				os.Exit(1)
			}
			qType := strings.Trim(kv[0], " ")
			searchFor := strings.Trim(kv[1], " ")

			matched := false
			for _, q := range validQueries {
				if strings.ToLower(q) == strings.ToLower(qType) {
					matched = true
					if "host-header" == strings.ToLower(q) {
						parameters.Add("hostHeader", searchFor)
					} else if "user-agent" == strings.ToLower(q) {
						parameters.Add("userAgent", searchFor)
					} else if "http-status-code" == strings.ToLower(q) {
						parameters.Add("httpStatusCode", searchFor)
					} else if "arl" == strings.ToLower(q) {
						parameters.Add("arl", searchFor)
					} else if "cp-code" == strings.ToLower(q) {
						parameters.Add("cpcode", searchFor)
					} else if "client-ip" == strings.ToLower(q) {
						parameters.Add("clientIp", searchFor)
					}
					break
				}
			}

			if !matched {
				printWarning("Invalid query field " + qType)
				fmt.Println("Valid Query fields:", strings.Join(validQueries, ", "))
				os.Exit(1)
			}

		}
		parameters.Add("duration", strconv.Itoa(durationFlag))
		parameters.Add("maxLogLines", strconv.Itoa(maxLinesFlag))
		if clientRequestFlag && forwardRequestFlag {
			logType = "both"
		} else if !clientRequestFlag {
			logType = "f"
		}
		parameters.Add("logType", logType)

		Url.RawQuery = parameters.Encode()

		resp, byt := doHTTPRequest("GET", Url.String(), nil)

		if resp.StatusCode == 200 {
			var responseStruct Wrapper
			var responseStructJson LogLinesJson
			err := json.Unmarshal(*byt, &responseStruct)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if jsonString {
				responseStructJson.EdgeServerIp = args[0]
				responseStructJson.Duration = durationFlag
				responseStructJson.EndDate = endDateFlag
				responseStructJson.EndTime = endTimeFlag
				responseStructJson.MaxLines = maxLinesFlag
				responseStructJson.findIn = findInFlag
				responseStructJson.ClientRequest = clientRequestFlag
				responseStructJson.ForwardRequest = forwardRequestFlag
				responseStructJson.ReportedTime = getReportedTime()
				responseStructJson.LogLines = responseStruct.LogLines
				resJson, _ := json.MarshalIndent(responseStructJson, "", "  ")
				resJson = getDecodedResponse(resJson)
				fmt.Println(string(resJson))
				return
			}

			printLogLines(responseStruct.LogLines)
		} else {
			printResponseError(byt)
		}
	},
}

func init() {
	grepCmd.Flags().StringVar(&endDateFlag, "end-date", "", endDateFlagDescription)
	grepCmd.Flags().StringVar(&endTimeFlag, "end-time", "", endTimeFlagDescription)
	grepCmd.MarkFlagRequired("end-time")
	grepCmd.Flags().IntVar(&durationFlag, "duration", 30, durationFlagDescription)
	grepCmd.MarkFlagRequired("duration")
	grepCmd.Flags().IntVar(&maxLinesFlag, "max-lines", 200, maxLinesFlagDescription)
	grepCmd.Flags().BoolVarP(&clientRequestFlag, "r", "r", true, clientRequestFlagDescription)
	grepCmd.Flags().BoolVarP(&forwardRequestFlag, "f", "f", false, forwardRequestFlagDescription)
	grepCmd.Flags().StringSliceVar(&findInFlag, "find-in", []string{}, findInFlagDescription)
	rootCmd.AddCommand(grepCmd)

}
