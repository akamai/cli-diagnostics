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

//root
const (
	rootShortDescription = "The Diagnostic Tools allows you to diagnose many common problems Akamai customers experience when delivering content to their end users."
	rootLongDescription  = `The Diagnostic Tools allows you to diagnose many common problems Akamai customers experience when delivering content to their end users.

	Run the process in background, 
	linux : redirect the output to a textfile and use ampersand &
		   command > textfile &
	windows : use the START command with /B and redirect the output to a textfile
		   START /B command > textfile`
	edgercPathFlagDescription    = "Location of edgegrid credentials file"
	edgercSectionFlagDescription = "Name of the section in credentials file"
	forceColorFlagDescription    = "Force color to non tty output"
	jsonFlagDescription          = "Get json output"
)

//ghost_location
const (
	ghostLocationShortDescription = "Lists active Akamai edge server locations from which you can run diagnostic tools."
	ghostLocationLongDescription  = `Lists active Akamai edge server locations from which you can run diagnostic tools.`
	searchFlagDescription         = "flag to accept a search string to filter the list."

	ghostLocation = "Ghost Locations"
)

//verify_ip
const (
	verifyIpShortDescription = "Confirms if a certain IP address is that of an Akamai edge server."
	verifyIpLongDescription  = `Confirms if a certain IP address is that of an Akamai edge server.`

	ipAddress      = "IP address "
	isCdnIpSuccess = "is an Akamai IP"
	isCdnIpFailure = "is not an Akamai IP"
)

//locate_ip
const (
	locateIpShortDescription = "Provides the geographic and network location of any IP address."
	locateIpLongDescription  = `Provides the geographic and network location of any IP address.`

	geographicLocation = "Geographic Location"
	clientIp           = "Client IP"
	countryCode        = "Country Code"
	regionCode         = "Region Code"
	city               = "City"
	dma                = "DMA"
	msa                = "MSA"
	pmsa               = "PMSA"
	areaCode           = "Area Code"
	latitude           = "Latitude"
	longitude          = "Longitude"
	country            = "County"
	continent          = "Continent"
	fisp               = "FIPS"
	timeZone           = "Time Zone"
	zipCode            = "Zip Code"
	proxy              = "Proxy"

	networkLocation = "Network Location"
	network         = "Network"
	networkType     = "Network Type"
	asNum           = "As Num"
	throughput      = "Throughput"
)

//translate_url
const (
	translateUrlShortDescription = "Provides basic information about a specified URL."
	translateUrlLongDescription  = `Provides basic information about a specified URL, such as typecode,
	origin server, CP code, serial number, and TTL for a URL/ARL.`

	translateUrl = "Translate url"
	typeCode     = "Type Code"
	originServer = "Origin Server"
	cpCode       = "CP Code"
	serialNumber = "Serial Number"
	ttl          = "TTL"
)

//dig
const (
	digShortDescription = "Run DIG command from edge server."
	digLongDescription  = `Uses the DIG command to provide Domain Name Server (DNS) details for the location of the edge server 
	and hostname or domain name, enabling you to diagnose issues with domain name resolution.`
	typeDigFlagDescription = "The type of DNS record, either A, AAAA, CNAME, MX, NS, PTR, or SOA."
)

//mtr
const (
	mtrShortDescription = "Run MTR from edge server."
	mtrLongDescription  = `Uses the MTR command to provide information about the route, number of hops, and time
	that Internet traffic packets take between the edge server and a remote host or destination. 
	The results can show you where network delays are being introduced in the path.`
	resolveHostMtrFlagDescription = "Whether to use DNS to resolve hostnames. When disabled, output features only IP addresses."

	networkConnectivity = "Network Connectivity Test from"
	to                  = "to"
)

//curl
const (
	curlShortDescription = "Uses the CURL command to provide the raw html for a specified URL."
	curlLongDescription  = `Uses the CURL command to provide the raw html for a specified URL.
	Making an HTTP request from an edge server lets you gather information about the HTTP response.`
	userAgentFlagDescription = "Pass an additional User Agent"

	responseHeader = "Response Headers"
	responseBody   = "Response Body"
)

//translate_error
const (
	translateErrorShortDescription = "Fetch a summary and log information	for the error that occurred in the original request."
	translateErrorLongDescription  = `Uses the error string from the reference number to fetch a summary and log information
	for the error that occurred in the original request.`

	summary                = "SUMMARY"
	urlTranslateError      = "URL"
	httpResponseCode       = "HTTP Response Code"
	dateAndTime            = "Date and Time"
	epocTime               = "Epoch Time"
	clientIpTranslateError = "Client IP"
	connectingIp           = "Connecting IP"
	originHostName         = "Origin Hostname"
	originIp               = "Origin IP"
	userAgent              = "User Agent"
	clientRequest          = "Client Request"
	reasonForFailure       = "Reason for Failure"
	wafDetails             = "WAF Details"

	errorLogs = "ERROR LOGS"
)

//grep
const (
	grepShortDescription          = "Uses the GREP command to retrieve and parse logs from an edge server IP address, within the last 48 hours."
	grepLongDescription           = `Uses the GREP command to retrieve and parse logs from an edge server IP address, within the last 48 hours.`
	endDateFlagDescription        = "End Date of log search, specified in <yyyy:mm:dd>"
	endTimeFlagDescription        = "End Time of log search, specified in <hh:mm:ss> (UTC)"
	durationFlagDescription       = "Duration before End Time in minutes"
	maxLinesFlagDescription       = "Maximum log lines to display"
	clientRequestFlagDescription  = "Search logs of incoming client requests to the edge server"
	forwardRequestFlagDescription = "Search logs of forwarded requests from the edge server"
	findInFlagDescription         = "Where to search, specified as <field>:<value>"
)

//estats
const (
	estatsShortDescription = "Provides an understanding of the errors happening in the delivery of websites."
	estatsLongDescription  = `Provides an understanding of the errors happening in the delivery of websites
	based on real-time data of traffic of a particular CP code in terms of traffic
	from clients to edge servers and from edge servers to origin.`

	summaryEstats = "SUMMARY"

	edgeStatistics                 = "Edge Statistics"
	percentageFailurEdgeStatistics = "(Percent failure: %.1f%c)"
	edgeStatisticsDescription      = "Edge Status Code Distribution"
	statusCodeEdgeStatistics       = "Status Code"
	percentageHitEdgeStatistics    = "%Hits"

	originalStatistics                 = "Origin Statistics"
	percentageFailurOriginalStatistics = "(Percent failure: %.1f%c)"
	originalStatisticsDescription      = "Origin Status Code Distribution"
	statusCodeOriginalStatistics       = "Status Code"
	percentageHitsOriginalStatistics   = "%Hits"

	edgeErrors             = "EDGE ERRORS"
	edgeErrorsDescription  = "View last 10 edge errors."
	edgeIpEdgeErrors       = "Edge IP"
	regionEdgeErrors       = "Region"
	httpStatusEdgeErrors   = "HTTP Status"
	hitsEdgeErrors         = "Hits"
	objectStatusEdgeErrors = "Object Status"
	errorCodeEdgeErrors    = "ErrorCode"

	originErrors             = "ORIGIN ERRORS"
	originErrorsDescription  = "View last 10 origin errors"
	edgeIpOriginErrors       = "Edge IP"
	regionOriginErrors       = "Region"
	httpStatusOriginErrors   = "HTTP Status"
	hitsOriginErrors         = "Hits"
	objectStatusOriginErrors = "Object Status"
	errorCodeOriginErrors    = "ErrorCode"
)

//debug_url
const (
	debugUrlShortDescription = "Provides DNS Information, HTTP Response, Response Header, and Logs for a URL/ARL."
	debugUrlLongDescription  = `Provides DNS Information, HTTP Response, Response Header, and Logs for a URL/ARL.`
	headerFlagDescription    = "Any additional headers to add to the request,specified as <header>:<value>"
	edgeIpFlagDescription    = "The edge server IP address to test the URL against, otherwise a random server by default."

	dnsInformation         = "DNS Information"
	httpResponse           = "HTTP Response"
	responseHeaderDebugUrl = "Response Header"
	logs                   = "Logs"
)

//user_diagnostics
const (
	userDiagnosticsShortDescription = "Use this tool to create a sharable link and send it to the end users to collect diagnostic data. You can view, filter, and export the results."
	userDiagnosticsLongDescription  = `Use this tool to create a sharable link and send it to the end users to collect diagnostic data. You can view, filter, and export the results.`
)

//user_diagnostics_list
const (
	userDiagnosticsListShortDescription = "List all the End User Diagnostic Data."
	userDiagnosticsListLongDescription  = `List all the End User Diagnostic Data.`

	userDiagnosticsListNote   = "NOTE: Each link is active for 7 days and has a limit of 50 submissions."
	linkId                    = "Link ID"
	statusUserDiagnosticsLink = "Status"
	generatedOn               = "Generated On(UTC)"
	groupName                 = "Group Name"
	hostNameOrUrl             = "Hostname/URL"
	caseId                    = "Case ID"
	userSharableLink          = "User Sharable Link"
	records                   = "Records"
)

//user_diagnostics_create_group
const (
	userDiagnosticsCreateGroupShortDescription = "Create group to get user sharable link"
	userDiagnosticsCreateGroupLongDescription  = `Create group to get user sharable link`

	linkText1 = "Here is your link!"
	linkText2 = "Copy and send the link below to the end users who are experiencing the content delivery problem.\nEach link is active for 7 days and has a limit of 50 submissions."
	linkText3 = "link: "
)

//user_diagnostics_get
const (
	userDiagnosticsGetShortDescription = "get end user diagnostics result by Id"
	userDiagnosticsGetLongDescription  = `get end user diagnostics result by Id`

	viewUserDiagnosticsData           = "View User Diagnostic Data"
	generatedOnUserDiagnosticGet      = "Generated on"
	groupNameUserDiagnosticGet        = "Group Name"
	hostNameOrUrlUserDiagnosticGet    = "Hostname/URL"
	userSharableLinkUserDiagnosticGet = "User Sharable Link"
	linkStatus                        = "Link Status"

	uid                        = "UID"
	timestamp                  = "TimeStamp(UTC)"
	clientIpPreferred          = "Client IP Preferred"
	clientDnsIpv4              = "Client DNS IPv4"
	clientDnsIpv6              = "Client DNS IPv6"
	userAgentUserDiagnosticGet = "User Agent"
	cookie                     = "Cookie"
	protocol                   = "Protocol"
	connectedCipher            = "Connected Cipher"
	clientIpv4                 = "Client IPv4"
	clientIpv6                 = "Client IPv6"

	edgeIps = "Edge IPs"

	curlDescription = " - Request content from edge server, run : "
	digDescription  = " - Get domain details from an edge server, run : "

	edgeIpMessage = "Edge IP is not shown either because of a system error or the client DNS has a ECS (EDNS Client Subnet)"
)
