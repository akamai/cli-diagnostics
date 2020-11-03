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

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"html"
	"net/url"
)

func colorPrintf(col string, value interface{}) {

	switch col {
	case "red":
		c := color.New(color.Bold).Add(color.FgHiRed)
		c.Printf("%v ", value)
	case "blue":
		c := color.New(color.Bold).Add(color.FgBlue)
		c.Printf("%v ", value)
	case "green":
		c := color.New(color.Bold).Add(color.FgGreen)
		c.Printf("%v ", value)
	case "yellow":
		c := color.New(color.Bold).Add(color.FgHiYellow)
		c.Printf("%v ", value)
	case "cyan":
		c := color.New(color.Bold).Add(color.FgHiCyan)
		c.Printf("%v ", value)
	}

}

func colorPrintln(col string, value interface{}) {
	switch col {
	case "red":
		c := color.New(color.Bold).Add(color.FgHiRed)
		c.Println(value)
	case "blue":
		c := color.New(color.Bold).Add(color.FgBlue)
		c.Println(value)
	case "green":
		c := color.New(color.Bold).Add(color.FgGreen)
		c.Println(value)
	case "yellow":
		c := color.New(color.Bold).Add(color.FgHiYellow)
		c.Println(value)
	case "cyan":
		c := color.New(color.Bold).Add(color.FgHiCyan)
		c.Println(value)
	}

}

func printWarning(msg interface{}) {

	c := color.New(color.Bold).Add(color.FgHiRed)
	c.Println(msg)
}

func printLabelAndValue(label string, value interface{}) {
	c := color.New(color.Bold).Add(color.FgHiCyan)
	c.Printf(label + ": ")
	fmt.Printf("%v\n", value)
}

func printGeoLocation(gl *GeoLocation) {
	colorPrintf("yellow", geographicLocation)
	fmt.Println()
	printLabelAndValue(clientIp, gl.ClientIP)
	printLabelAndValue(countryCode, gl.CountryCode)
	printLabelAndValue(regionCode, gl.RegionCode)
	printLabelAndValue(city, gl.City)
	if gl.DMA != nil {
		printLabelAndValue(dma, *gl.DMA)
	} else {
		printLabelAndValue(dma, "-")
	}
	if gl.MSA != nil {
		printLabelAndValue(msa, *gl.MSA)
	} else {
		printLabelAndValue(msa, "-")
	}
	if gl.PMSA != nil {
		printLabelAndValue(pmsa, *gl.PMSA)
	} else {
		printLabelAndValue(pmsa, "-")
	}

	if gl.AreaCode != "" {
		printLabelAndValue(areaCode, gl.AreaCode)
	} else {
		printLabelAndValue(areaCode, "-")
	}
	printLabelAndValue(latitude, gl.Latitude)
	printLabelAndValue(longitude, gl.Longitude)
	if gl.County != "" {
		printLabelAndValue(country, gl.County)
	} else {
		printLabelAndValue(country, "-")
	}

	printLabelAndValue(continent, gl.Continent)
	if gl.FIPS != "" {
		printLabelAndValue(fisp, gl.FIPS)
	} else {
		printLabelAndValue(fisp, "-")
	}
	printLabelAndValue(timeZone, gl.TimeZone)

	if gl.ZipCode != "" {
		printLabelAndValue(zipCode, gl.ZipCode)
	} else {
		printLabelAndValue(zipCode, "-")
	}

	if gl.Proxy != "" {
		printLabelAndValue(proxy, gl.Proxy)
	} else {
		printLabelAndValue(proxy, "-")
	}
	colorPrintf("yellow", "\n"+networkLocation)
	fmt.Println()
	if gl.Network != "" {
		printLabelAndValue(network, gl.Network)
	} else {
		printLabelAndValue(network, "-")
	}
	if gl.NetworkType != "" {
		printLabelAndValue(networkType, gl.NetworkType)
	} else {
		printLabelAndValue(networkType, "-")
	}
	if gl.AsNum != "" {
		printLabelAndValue(asNum, gl.AsNum)
	} else {
		printLabelAndValue(asNum, "-")
	}
	printLabelAndValue(throughput, gl.Throughput)

}

func printGhostLocations(obj *GhostLocationsList, searchString string) {
	colorPrintln("yellow", ghostLocation)
	count := 0
	for _, loc := range obj.Locations {
		if strings.Contains(strings.ToLower(loc["id"]), searchString) {
			colorPrintf("cyan", loc["value"]+":")
			fmt.Println(loc["id"])
			count++
		}
	}
	if count == 0 {
		printWarning("No match found")
	}
}

func printTranslatedError(obj *TranslatedError) {
	colorPrintln("yellow", summary)
	fmt.Println()
	printLabelAndValue(urlTranslateError, obj.Url)
	printLabelAndValue(httpResponseCode, obj.HttpResponseCode)
	printLabelAndValue(dateAndTime, obj.Timestamp)
	printLabelAndValue(epocTime, obj.EpochTime)
	printLabelAndValue(clientIpTranslateError, obj.ClientIP)
	printLabelAndValue(connectingIp, obj.ConnectingIP)
	printLabelAndValue(originHostName, obj.OriginHostname)
	printLabelAndValue(originIp, obj.OriginIP)
	printLabelAndValue(userAgent, obj.UserAgent)
	printLabelAndValue(clientRequest, obj.RequestMethod)
	printLabelAndValue(reasonForFailure, obj.ReasonForFailure)
	printLabelAndValue(wafDetails, obj.WafDetails)

	fmt.Println()
	colorPrintln("yellow", errorLogs)
	fmt.Println("----")
	for _, log := range obj.Logs {
		fmt.Println(log.Description)
		fmt.Println()
		for k, v := range log.Fields {
			str := fmt.Sprintf("%v", v)
			printLabelAndValue(k, html.UnescapeString(str))
		}
		fmt.Println("----")
	}

}

func printCurlResults(obj *CurlResults) {
	colorPrintln("yellow", responseHeader)
	for key, val := range obj.ResponseHeaders {

		printLabelAndValue(key, val)

	}
	fmt.Println()
	colorPrintln("yellow", responseBody)
	fmt.Println(obj.ResponseBody)

}

func printDebugUrlResults(obj *DebugUrl) {
	colorPrintln("yellow", dnsInformation)
	for _, str := range obj.DNSinformation {
		fmt.Println(str)
	}
	fmt.Println("\n----")

	colorPrintln("yellow", httpResponse)
	for _, mapObj := range obj.HTTPResponse {
		printLabelAndValue(mapObj["name"], mapObj["value"])
	}
	fmt.Println("\n----")

	colorPrintln("yellow", responseHeader)
	for _, str := range obj.ResponseHeaders {
		fmt.Println(str)
	}
	fmt.Println("\n----")

	colorPrintln("yellow", logs)
	for _, log := range obj.Logs {
		fmt.Println(log.Description)
		for k, v := range log.Fields {
			str := fmt.Sprintf("%v", v)
			printLabelAndValue(k, html.UnescapeString(str))
		}
		fmt.Println()
	}
	fmt.Println("\n----")

}

func printErrorStats(obj *Estats) {
	fmt.Println()
	colorPrintln("yellow", summaryEstats+"\n")
	colorPrintf("yellow", edgeStatistics)
	fmt.Printf(percentageFailurEdgeStatistics+"\n", obj.EdgeFailurePercentage, '%')
	color.Blue(edgeStatisticsDescription + "\n")
	var rows [][]interface{}
	for _, x := range obj.EdgeStatusCodeDistribution {
		var row []interface{}
		row = append(row, x.HTTPStatus)
		row = append(row, x.Percentage)
		rows = append(rows, row)
	}
	printTables([]interface{}{statusCodeEdgeStatistics, percentageHitEdgeStatistics}, rows)

	fmt.Println()
	colorPrintf("yellow", originalStatistics)
	fmt.Printf(percentageFailurOriginalStatistics+"\n", obj.OriginFailurePercentage, '%')
	color.Blue(originalStatisticsDescription + "\n")
	rows = rows[:0] //clear the rows
	for _, x := range obj.OriginStatusCodeDistribution {
		var row []interface{}
		row = append(row, x.HTTPStatus)
		row = append(row, x.Percentage)
		rows = append(rows, row)
	}
	printTables([]interface{}{statusCodeOriginalStatistics, percentageHitsOriginalStatistics}, rows)
	fmt.Println("\n-----")

	colorPrintln("yellow", edgeErrors)
	fmt.Println(edgeErrorsDescription)
	rows = rows[:0]
	if obj.TopEdgeIPsWithError != nil {
		for _, x := range obj.TopEdgeIPsWithError {
			row := []interface{}{x.EdgeIP, x.Region, x.HTTPstatus, x.Hits, x.ObjStatus, x.ErrorCode}
			rows = append(rows, row)
		}
		printTables([]interface{}{edgeIpEdgeErrors, regionEdgeErrors, httpStatusEdgeErrors, hitsEdgeErrors, objectStatusEdgeErrors, errorCodeEdgeErrors}, rows)
		/*
		   t := table.NewWriter()
		   t.SetOutputMirror(os.Stdout)
		   t.AppendHeader([]interface{}{"Edge IP", "Region", "HTTP Status", "Hits", "Object Status", "ErrorCode", "Ghost Log"})
		   for _, row := range rows {
		       t.AppendRow(row)
		   }
		   t.AppendSeparator()
		   t.SetColumnConfigs([]table.ColumnConfig{{Number: 7, WidthMax: 64}})

		   t.SetStyle(table.StyleLight)
		   t.Render()*/
	} else {
		fmt.Println("No error logs found.")
	}
	fmt.Println("\n-----")

	colorPrintln("yellow", originErrors)
	fmt.Println(originErrorsDescription)
	rows = rows[:0]
	if obj.TopEdgeIPsWithErrorFromOrigin != nil {
		for _, x := range obj.TopEdgeIPsWithErrorFromOrigin {
			row := []interface{}{x.EdgeIP, x.Region, x.HTTPstatus, x.Hits, x.ObjStatus, x.ErrorCode}
			rows = append(rows, row)
		}
		printTables([]interface{}{edgeIpOriginErrors, regionOriginErrors, httpStatusOriginErrors, hitsOriginErrors, objectStatusOriginErrors, errorCodeOriginErrors}, rows)

	} else {
		fmt.Println("No error logs found.")
	}
	fmt.Println("\n-----")

}

func printGenericErrorMsg() {
	printWarning("An error occurred. Please try again later, or contact Support if the problem persists.")
}

func printTables(header []interface{}, rows [][]interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(header)
	for _, row := range rows {
		t.AppendRow(row)
	}
	t.AppendSeparator()
	t.SetStyle(table.StyleLight)
	t.Render()
}

func printLogLines(obj *LogLines) {
	colorPrintln("yellow", obj.Headers)
	for _, log := range obj.Logs {
		fmt.Println(log)
		fmt.Println("----")
	}

}

func printResponseError(byt *[]byte) {
	//fmt.Println(string(*byt))
	var errorStruct ResponseError
	err := json.Unmarshal(*byt, &errorStruct)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if jsonString {
		resJson, _ := json.MarshalIndent(errorStruct, "", "  ")
		fmt.Println(string(resJson))
		return
	}
	printWarning(errorStruct.Detail)
	for _, err := range errorStruct.Errors {
		printWarning(err["error"])
		//fmt.Println(err["error"])
	}
	os.Exit(1)
}

func printUserDiagnostics(obj *UserDiagnosticData) {
	fmt.Println()
	colorPrintln("yellow", viewUserDiagnosticsData)
	fmt.Println()
	weekday, month, day, year, _ := isoToDate(obj.CreatedDate)
	printLabelAndValue(generatedOnUserDiagnosticGet, fmt.Sprintf("%s, %s %d, %d (UTC)", weekday, month, day, year))
	printLabelAndValue(groupNameUserDiagnosticGet, obj.GroupName)
	printLabelAndValue(hostNameOrUrlUserDiagnosticGet, obj.URL)
	printLabelAndValue(userSharableLink, obj.DiagnosticLink)
	printLabelAndValue(linkStatus, obj.Status)
	fmt.Println("\n----")
	printDiagnosticRecords(obj.DiagnosticRecords, obj.URL)
}

func printDiagnosticRecords(records []DiagnosticRecord, hostNameOrUrl string) {

	var hostName string
	var hostNameWithProtocol string
	if strings.HasPrefix(hostNameOrUrl, "http") {
		hostNameWithProtocol = hostNameOrUrl
		u, _ := url.Parse(hostNameOrUrl)
		hostName = u.Host
	} else {
		hostNameWithProtocol = "http://" + hostNameOrUrl
		hostName = hostNameOrUrl
	}

	for _, record := range records {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader([]interface{}{uid, timestamp, clientIpPreferred, clientDnsIpv4, clientDnsIpv6, userAgentUserDiagnosticGet, cookie, protocol, connectedCipher, clientIpv4, clientIpv6})
		weekday, month, day, year, clock := isoToDate(string(record.CreatedDate))
		var row []interface{}
		row = append(row, record.UniqueId)
		row = append(row, fmt.Sprintf("%s, %s %d, %d %s", weekday, month, day, year, clock))
		row = append(row, IPinfoRecordString(record.ClientIPv4))
		row = append(row, IPinfoRecordString(record.ClientDnsIpv4))
		row = append(row, IPinfoRecordString(record.ClientDnsIpv6))
		row = append(row, record.UserAgent)
		if record.Cookie {
			row = append(row, "Yes")
		} else {
			row = append(row, "No")
		}
		row = append(row, record.Protocol)
		row = append(row, record.Cipher)
		row = append(row, IPinfoRecordString(record.ClientIPv4))
		row = append(row, IPinfoRecordString(record.ClientIPv6))
		t.AppendRow(row)
		t.AppendSeparator()
		width := 14
		t.SetColumnConfigs([]table.ColumnConfig{{Number: 1, WidthMax: 6, WidthMaxEnforcer: func(col string, maxLen int) string { return col }},
			{Number: 2, WidthMax: width},
			{Number: 3, WidthMax: width},
			{Number: 4, WidthMax: width + 1},
			{Number: 5, WidthMax: width + 1},
			{Number: 6, WidthMax: width},
			{Number: 7, WidthMax: width},
			{Number: 8, WidthMax: width},
			{Number: 9, WidthMax: width},
			{Number: 10, WidthMax: width},
			{Number: 11, WidthMax: width},
		})
		t.SetStyle(table.StyleLight)
		t.SetAllowedRowLength(170)
		t.Style().Format.Header = text.FormatDefault
		t.Render()
		if record.EdgeIPs != nil {
			fmt.Println("\n" + edgeIps)
			for _, ip := range record.EdgeIPs {
				fmt.Println()
				colorPrintln("cyan", IPinfoRecordString(&ip))
				fmt.Printf(curlDescription)

				colorPrintln("yellow", fmt.Sprint("akamai diagnostics curl ", hostNameWithProtocol, " ", ip.IP))

				fmt.Printf(digDescription)
				colorPrintln("yellow", fmt.Sprint("akamai diagnostics dig ", hostName, " ", ip.IP))

			}

		} else {
			msg := edgeIpMessage
			printWarning(msg)
		}
		fmt.Println("\n----")
	}
}

func IPinfoRecordString(obj *IPinfoRecord) string {
	if obj == nil {
		return ""
	}
	return fmt.Sprintf("%s (%s,%s %s)(%s)", obj.IP, obj.Location["city"], obj.Location["state"], obj.Location["country"], obj.Location["asNum"])
}

func printListUserDiagnosticData(objs []EndUserDiagnosticLink) {
	fmt.Println("\n" + userDiagnosticsListNote)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader([]interface{}{linkId, statusUserDiagnosticsLink, generatedOn, groupName, hostNameOrUrl, caseId, userSharableLink, records})
	for _, obj := range objs {
		weekday, month, day, year, _ := isoToDate(fmt.Sprintf("%v", obj.CreatedDate))
		t.AppendRow([]interface{}{obj.DiagnosticLinkID, obj.Status, fmt.Sprintf("%s, %s %d, %d", weekday, month, day, year), obj.GroupName, obj.URL, strings.Join(obj.CaseIds, ","), obj.URL, obj.RecordCount})
		t.AppendSeparator()
	}
	width := 20
	t.SetColumnConfigs([]table.ColumnConfig{{Number: 8, WidthMax: width},
		{Number: 3, WidthMax: width},
		{Number: 4, WidthMax: width},
		{Number: 5, WidthMax: width},
		{Number: 6, WidthMax: width},
		{Number: 7, WidthMax: width},
	})
	t.SetStyle(table.StyleLight)
	t.SetAllowedRowLength(170)
	t.Style().Format.Header = text.FormatUpper
	t.Render()

}
