package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

func printJsonOutput(byt *[]byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, *byt, "", indentString)
	if err != nil {
		PrintError("Failed to indent the json response")
		fmt.Println("\n" + string(*byt))
		return
	}
	fmt.Println("\n" + prettyJSON.String())
}

func printEdgeLocationsJsonOutput(edgeLocationContainer EdgeLocationContainer, searchText string) {
	var edgeLocationContainerFiltered EdgeLocationContainer
	if searchText != "" {
		for _, edgeLocation := range edgeLocationContainer.EdgeLocations {
			if strings.Contains(strings.ToLower(edgeLocation.Id), searchText) {
				edgeLocationContainerFiltered.EdgeLocations = append(edgeLocationContainerFiltered.EdgeLocations, edgeLocation)
			}
		}
		resJson, _ := json.MarshalIndent(edgeLocationContainerFiltered, "", indentString)
		fmt.Println(string(resJson))
	} else {
		resJson, _ := json.MarshalIndent(edgeLocationContainer, "", indentString)
		fmt.Println(string(resJson))
	}
}

func PrintLocateIpResponse(geoLocation GeoLocation, ip string) {

	PrintHeader("\n" + geographicLocation + "\n")

	printLabelAndValue(ipAddress, ifNilReturnAlternative(ip, "-"))
	printLabelAndValue(countryCode, ifNilReturnAlternative(geoLocation.CountryCode, "-"))
	printLabelAndValue(regionCode, ifNilReturnAlternative(geoLocation.RegionCode, "-"))
	printLabelAndValue(city, ifNilReturnAlternative(geoLocation.City, "-"))
	printLabelAndValue(dma, ifNilIntPtrReturnAlternative(geoLocation.DMA, "-"))
	printLabelAndValue(msa, ifNilIntPtrReturnAlternative(geoLocation.MSA, "-"))
	printLabelAndValue(pmsa, ifNilIntPtrReturnAlternative(geoLocation.PMSA, "-"))
	printLabelAndValue(areaCode, ifNilReturnAlternative(geoLocation.AreaCode, "-"))
	printLabelAndValue(latitude, ifNilFloatPtrReturnAlternative(geoLocation.Latitude, "-"))
	printLabelAndValue(longitude, ifNilFloatPtrReturnAlternative(geoLocation.Longitude, "-"))
	printLabelAndValue(county, ifNilReturnAlternative(geoLocation.County, "-"))
	printLabelAndValue(continent, ifNilReturnAlternative(geoLocation.Continent, "-"))
	printLabelAndValue(fisp, ifNilReturnAlternative(geoLocation.FIPS, "-"))
	printLabelAndValue(timeZone, ifNilReturnAlternative(geoLocation.TimeZone, "-"))
	printLabelAndValue(zipCode, ifNilReturnAlternative(geoLocation.ZipCode, "-"))
	printLabelAndValue(proxy, ifNilReturnAlternative(geoLocation.Proxy, "-"))

	PrintHeader("\n" + networkLocation + "\n")

	printLabelAndValue(network, ifNilReturnAlternative(geoLocation.Network, "-"))
	printLabelAndValue(networkType, ifNilReturnAlternative(geoLocation.NetworkType, "-"))
	printLabelAndValue(asNum, ifNilIntPtrReturnAlternative(geoLocation.AsNum, "-"))
	printLabelAndValue(throughput, ifNilReturnAlternative(geoLocation.Throughput, "-"))

}

func PrintLocateIpsResponse(verifyLocateIpsResponse VerifyLocateIpsResponse, failedIpsMessage string) {
	var failedIps []string
	for _, verifyLocateIpsData := range verifyLocateIpsResponse.Result {
		if verifyLocateIpsData.ExecutionStatus == "SUCCESS" {
			PrintLocateIpResponse(verifyLocateIpsData.GeoLocation, verifyLocateIpsData.IpAddress)
		} else {
			failedIps = append(failedIps, verifyLocateIpsData.IpAddress)
		}
	}
	if len(failedIps) > 0 {
		PrintFailedIps(failedIps, failedIpsMessage)
	}
}

func ifNilReturnAlternative(og, alternate interface{}) interface{} {
	if og != "" {
		return og
	}
	return alternate
}

func ifNilIntPtrReturnAlternative(og *int, alternate interface{}) interface{} {
	if og == nil {
		return alternate
	}
	return *og
}

func ifNilFloatPtrReturnAlternative(og *float32, alternate interface{}) interface{} {
	if og == nil {
		return alternate
	}
	return *og
}

func PrintVerifyIpResponse(verifyLocateIpResponse VerifyLocateIpResponse) {

	funcMap := template.FuncMap{
		"bold": bold,
	}
	var tmp string
	if verifyLocateIpResponse.Result.IsEdgeIp {
		tmp = isEdgeIpTemplate
	} else {
		tmp = isNotEdgeIpTemplate
	}
	printTemplate(funcMap, tmp, verifyLocateIpResponse.Request)

}

func PrintVerifyIpsResponse(verifyLocateIpsResponse VerifyLocateIpsResponse, failedIpsMessage string) {

	funcMap := template.FuncMap{
		"bold": bold,
	}
	var tmp string
	var failedIps []string
	for _, verifyLocateIpsData := range verifyLocateIpsResponse.Result {
		if verifyLocateIpsData.ExecutionStatus == "SUCCESS" {
			if verifyLocateIpsData.IsEdgeIp {
				tmp = isEdgeIpTemplate
			} else {
				tmp = isNotEdgeIpTemplate
			}
		} else {
			failedIps = append(failedIps, verifyLocateIpsData.IpAddress)
		}
		printTemplate(funcMap, tmp, verifyLocateIpsData)
	}
	if len(failedIps) > 0 {
		PrintFailedIps(failedIps, failedIpsMessage)
	}

}

func PrintFailedIps(failedIps []string, failedIpsMessage string) {
	fmt.Println("\n" + failedIpsMessage)
	for _, failedIp := range failedIps {
		fmt.Println(failedIp)
	}
}

func PrintDigResponse(digResponse DigResponse) {
	funcMap := template.FuncMap{
		"header":           header,
		"bold":             bold,
		"locationToString": EdgeIpLocation.toString,
		"dateToString":     isoToDate,
	}

	printTemplate(funcMap, digTemplate, digResponse)
	PrintSuggestedActions(digResponse.SuggestedActions)
}

func PrintUserDiagnosticsDataGroupDetailsAfterCreate(userDiagnosticsDataGroupDetails UserDiagnosticsDataGroupDetails) {

	// standard functions for templates
	funcMap := template.FuncMap{
		"bold": bold,
	}

	printTemplate(funcMap, userDiagnosticsCreateTemplate, userDiagnosticsDataGroupDetails)
}

func PrintUserDiagnosticsDataGroupDetailsTable(groupsList []UserDiagnosticsDataGroupDetails) {

	if len(groupsList) <= 0 {
		PrintWarning("No diagnostic links found")
		fmt.Println()
		return
	}

	//as createdTime is in ISO format, normal string comparison suffices
	sort.SliceStable(groupsList, func(i, j int) bool {
		return groupsList[i].CreatedTime > groupsList[j].CreatedTime
	})

	tableHeaders := []string{linkId, hostNameOrUrl, statusUserDiagnosticsLink, diagnosticLink, results, user, requestDate}
	content := make([][]string, len(groupsList))

	for i, grp := range groupsList {

		log.Debug("created time string: ", grp.CreatedTime)
		convertedTime := isoToDate(grp.CreatedTime)
		urlRow := grp.URL
		if grp.URL == "" {
			urlRow = grp.IpaHostname
		}
		grp.DiagnosticLinkStatus = CapsToTitle(grp.DiagnosticLinkStatus)
		if grp.DiagnosticLinkStatus == "Active" {
			grp.DiagnosticLinkStatus = success(grp.DiagnosticLinkStatus)
		}
		row := []string{grp.GroupID, urlRow, grp.DiagnosticLinkStatus, grp.DiagnosticLink, strconv.Itoa(grp.RecordCount), grp.CreatedBy, convertedTime}
		content[i] = row
	}

	fmt.Println("\n" + userDiagnosticsListNote + "\n")
	ShowTable(tableHeaders, content)
}

func PrintUserDiagnosticsDataGroupDetails(userDiagnosticsDataGroupDetails UserDiagnosticsDataGroupDetails) {

	// standard functions for templates
	funcMap := template.FuncMap{
		"bold":        bold,
		"header":      header,
		"italic":      italic,
		"capsToTitle": CapsToTitle,
	}

	printTemplate(funcMap, userDiagnosticsGetTemplate, userDiagnosticsDataGroupDetails)

}

func PrintTranslateUrlResponse(arlContainer ArlContainer) {

	funcMap := template.FuncMap{
		"bold":   bold,
		"header": header,
	}
	printTemplate(funcMap, translateUrlTemplate, arlContainer)
}
func printEdgeLocations(edgeLocationContainer EdgeLocationContainer, searchText string) {
	PrintHeader("\n" + edgeLocations + "\n")
	count := 0
	if searchText != "" {
		for _, edgeLocation := range edgeLocationContainer.EdgeLocations {
			if strings.Contains(strings.ToLower(edgeLocation.Id), searchText) {
				printLabelAndValue(edgeLocation.Value, edgeLocation.Id)
				count++
			}
		}
	} else {
		for _, edgeLocation := range edgeLocationContainer.EdgeLocations {
			printLabelAndValue(edgeLocation.Value, edgeLocation.Id)
			count++
		}
	}
	if count == 0 {
		PrintWarning("No edge locations found\n")
	}
}

func PrintIpaHostnamesTable(ipaHostnamesResponse IpaHostnameResponse) {
	tableHeaders := []string{ipaHostname}
	content := make([][]string, len(ipaHostnamesResponse.Hostnames))

	for i, hostname := range ipaHostnamesResponse.Hostnames {
		row := []string{hostname}
		content[i] = row
	}

	PrintHeader("\n" + ipaHostnames + "\n")
	ShowTable(tableHeaders, content)
}

func PrintGtmHostnamesTable(gtmPropertyContainer GtmPropertyContainer) {
	tableHeaders := []string{gtmHostname}
	content := make([][]string, len(gtmPropertyContainer.GtmProperties))

	for i, gtmProperties := range gtmPropertyContainer.GtmProperties {
		row := []string{gtmProperties.Hostname}
		content[i] = row
	}

	PrintHeader("\n" + gtmHostnames + "\n")
	ShowTable(tableHeaders, content)
}

func PrintGtmTestTargetIpTable(gtmPropertyIpsContainer GtmPropertyIpsContainer) {
	tableHeaders := []string{gtmHostname, testIp, target}
	content := make([][]string, 1)

	var testIpConcat string
	var targetConcat string

	for _, testIp := range gtmPropertyIpsContainer.GTMPropertyIps.TestIps {
		testIpConcat = testIpConcat + testIp + "\n"
	}

	for _, target := range gtmPropertyIpsContainer.GTMPropertyIps.Targets {
		targetConcat = targetConcat + target + "\n"
	}

	row := []string{gtmPropertyIpsContainer.GTMPropertyIps.Property + "." + gtmPropertyIpsContainer.GTMPropertyIps.Domain, testIpConcat, targetConcat}
	content[0] = row

	PrintHeader("\n" + gtmHostnames + "\n")
	ShowTable(tableHeaders, content)
}

func PrintCurlResponse(curlResponse CurlResponse) {
	funcMap := template.FuncMap{
		"header":           header,
		"bold":             bold,
		"italic":           italic,
		"capsToTitle":      CapsToTitle,
		"locationToString": EdgeIpLocation.toString,
		"splitHeaderValue": SplitCurlHeaderString,
		"isoDate":          isoToDate,
	}

	printTemplate(funcMap, curlTemplate, curlResponse)
	PrintSuggestedActions(curlResponse.SuggestedActions)
}

func PrintMtrResponse(mtrRequest MtrRequest, mtrResponse MtrResponse) {
	funcMap := template.FuncMap{
		"header":           header,
		"bold":             bold,
		"locationToString": EdgeIpLocation.toString,
		"dateToString":     isoToDate,
	}

	if net.ParseIP(mtrRequest.Destination) != nil {
		mtrResponse.DestinationInternalIP = mtrRequest.Destination
	}
	printTemplate(funcMap, mtrDetailsTemplate, mtrResponse)

	tableHeaders := strings.Split("#, Host:"+mtrResponse.Result.Host+", "+mtrTableCol, ",")
	content := make([][]string, len(mtrResponse.Result.Hops))

	for i, hop := range mtrResponse.Result.Hops {
		hostAndIP := hop.Host
		if hop.Host == "" {
			hostAndIP = "???"
		}
		if hop.IP != "" {
			hostAndIP += "\n" + hop.IP
		}
		content[i] = []string{fmt.Sprint(hop.Number), hostAndIP, fmt.Sprint(hop.PacketLoss), fmt.Sprint(hop.SentPackets), fmt.Sprint(hop.LastPacketLatency), fmt.Sprint(hop.AverageLatency), fmt.Sprint(hop.BestRtt), fmt.Sprint(hop.WorstRtt), fmt.Sprint(hop.StandardDeviation), hop.IPLocation.toStringNoBrackets()}
	}
	ShowTable(tableHeaders, content)
	PrintSuggestedActions(mtrResponse.SuggestedActions)
}

func PrintSuggestedActions(suggestedActions []string) {
	if suggestedActions != nil {
		PrintHeader("Suggested Actions\n")
		for _, suggestedAction := range suggestedActions {
			fmt.Println(suggestedAction)
		}
	}
}
