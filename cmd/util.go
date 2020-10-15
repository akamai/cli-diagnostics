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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	edgegrid "github.com/akamai/AkamaiOPEN-edgegrid-golang"
)

func doHTTPRequest(method string, url string, payload *[]byte) (*http.Response, *[]byte) {

	var err error

	client := http.Client{}

	var req *http.Request

	if payload != nil {
		req, err = http.NewRequest(method, "https://"+config.Host+url, bytes.NewBuffer(*payload))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		req, err = http.NewRequest(method, "https://"+config.Host+url, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	//colorPrintln("yellow", req.URL.Query())

	req = edgegrid.AddRequestHeader(config, req)

	resp, er := client.Do(req)
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp, &byt
}

//return 0 if ip-address, 1 if ghost location, 2 for error
func checkEdgeServerIPorLocation(addr string) (int, string) {
	ip := net.ParseIP(addr)
	if ip != nil {
		return 0, addr
	}

	resp, byt := doHTTPRequest("GET", "/diagnostic-tools/v2/ghost-locations/available", nil)
	if resp.StatusCode != 200 {
		printGenericErrorMsg()
		os.Exit(1)
	}

	var obj GhostLocationsList
	err := json.Unmarshal(*byt, &obj)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, loc := range obj.Locations {
		if strings.ToLower(loc["id"]) == strings.ToLower(addr) {
			return 1, loc["id"]
		}
	}

	return 2, ""
}

func validErrorString(str string) bool {
	return true
}

func checkAbsoluteURL(str string) bool {
	urlCheck, err := url.Parse(str)
	if err != nil || !urlCheck.IsAbs() {
		return false
	}
	return true
}

func isoToDate(isoDate string) (string, string, int, int, string) {
	date, _ := time.Parse(time.RFC3339, isoDate)
	h, m, s := date.UTC().Clock()
	clock := fmt.Sprintf("%d:%d:%d", h, m, s)
	return date.UTC().Weekday().String(), date.Month().String(), date.Day(), date.Year(), clock
}

func getReportedTime() string {
	loc, _ := time.LoadLocation("UTC")
	reportedTime := time.Now().In(loc)
	return reportedTime.Format(time.RFC3339)
}

func getDecodedResponse(responseJson []byte) []byte {
	responseJson = bytes.Replace(responseJson, []byte(`\u003c`), []byte("<"), -1)
	responseJson = bytes.Replace(responseJson, []byte(`\u003e`), []byte(">"), -1)
	responseJson = bytes.Replace(responseJson, []byte(`\u0026`), []byte("&"), -1)
	return responseJson
}
