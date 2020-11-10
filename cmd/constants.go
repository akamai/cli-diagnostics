// Copyright 2020. Akamai Technologies, Inc

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//  http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

//root
const (
	rootShortDescription = "The Diagnostic Tools CLI allows you to diagnose server, ESI, DNS, and network issues Akamai customers experience when delivering content to their end users."
	rootLongDescription  = `The Diagnostic Tools CLI allows you to diagnose server, ESI, DNS, and network issues Akamai customers experience when delivering content to their end users.

    Run the process in background, 
    linux : redirect the output to a textfile and use ampersand &
           command > textfile &
    windows : use the START command with /B and redirect the output to a textfile
           START /B command > textfile`

	edgercPathFlagDescription    = "Location of the edgegrid credentials file."
	edgercSectionFlagDescription = "Section name in the credentials file."
	forceColorFlagDescription    = "Force color to non-tty output."
	jsonFlagDescription          = "Get JSON output."

	rootUse = "diagnostics"
)

//ghost_location
const (
	ghostLocationShortDescription = "Lists active Akamai edge server locations from which you can run diagnostic tools."
	ghostLocationLongDescription  = `Lists active Akamai edge server locations from which you can run diagnostic tools.`

	searchFlagDescription = "The location to filter the list."

	ghostLocationUse = "ghost-locations --search location"

	ghostLocation = "Ghost Locations"
)

//verify_ip
const (
	verifyIpShortDescription = "Verifies whether the specified IP address is part of the Akamai edge network."
	verifyIpLongDescription  = `Verifies whether the specified IP address is part of the Akamai edge network.`

	verifyIpUse = "verify-ip IP_address"

	ipAddress      = "IP address "
	isCdnIpSuccess = "is an Akamai IP"
	isCdnIpFailure = "is not an Akamai IP"
)

//locate_ip
const (
	locateIpShortDescription = "Provides the geographic and network location of an IP address within the Akamai network."
	locateIpLongDescription  = `Provides the geographic and network location of an IP address within the Akamai network.`

	locateIpUse = "locate-ip IP_address"

	geographicLocation = "Geographic Location"
	clientIp           = "Client IP"
	countryCode        = "Country code"
	regionCode         = "Region code"
	city               = "City"
	dma                = "DMA"
	msa                = "MSA"
	pmsa               = "PMSA"
	areaCode           = "Area code"
	latitude           = "Latitude"
	longitude          = "Longitude"
	country            = "County"
	continent          = "Continent"
	fisp               = "FIPS"
	timeZone           = "Time zone"
	zipCode            = "Zip code"
	proxy              = "Proxy"

	networkLocation = "Network Location"
	network         = "Network"
	networkType     = "Network type"
	asNum           = "ASN"
	throughput      = "Throughput"
)

//translate_url
const (
	translateUrlShortDescription = "Provides high-level information about an Akamai-optimized URL (ARL), such as its time to live, origin server, and associated CP code."
	translateUrlLongDescription  = `Provides high-level information about an Akamai-optimized URL (ARL), such as its time to live, origin server, and associated CP code.`

	translateUrlUse = "translate-url URL"

	translateUrl = "Translate URL"
	typeCode     = "Type code"
	originServer = "Origin server"
	cpCode       = "CP code"
	serialNumber = "Serial number"
	ttl          = "TTL"
)

//dig
const (
	digShortDescription = "Runs DIG on a hostname or a domain name to return DNS details for the location of an Akamai edge server and the hostname or the domain name. You can use it to diagnose issues with the DNS resolution."
	digLongDescription  = `Runs the DIG command on a hostname or a domain name to return DNS details for the location of an Akamai edge server and the hostname or the domain name. You can use it to diagnose issues with the DNS resolution.`

	typeDigFlagDescription = "The type of the DNS record; possible values are: A, AAAA, CNAME, MX, NS, PTR, or SOA."

	digUse = "dig hostname source_server_location/edge_server_IP --type query_type"
)

//mtr
const (
	mtrShortDescription = "Runs MTR to check connectivity between an Akamai edge server and a remote host or destination. You can use it to diagnose network delays issues."
	mtrLongDescription  = `Runs the MTR command to provide information about the route, number of hops, and time that Internet traffic packets take between the Akamai edge server and a remote host or destination. You can use it to diagnose network delays issues.`

	resolveHostMtrFlagDescription = "Whether to use DNS to resolve hostnames. When disabled, the output features only IP addresses."

	mtrUse = "mtr domain_name/destination_IP source_server_location/edge_server_IP --resolve-hostname"

	networkConnectivity = "Network Connectivity Test from"
	to                  = "to"
)

//curl
const (
	curlShortDescription = "Runs CURL to provide a raw HTML for a URL within the Akamai network. You can use it to gather information about the HTTP response."
	curlLongDescription  = `Runs the CURL command to provide a raw HTML for a URL within the Akamai network. You can use it to gather information about the HTTP response.`

	userAgentFlagDescription = "The user agent; possible values are: android, firefox, iphone, mobile, chrome, msie, msie9, msie10, safari, safari/5, safari/6, webkit, webkit/5, webkit/6."

	curlUse = "curl URL source_server_location/edge_server_IP --user-agent additional-user-agent"

	responseHeader = "Response Headers"
	responseBody   = "Response Body"
)

//translate_error
const (
	translateErrorShortDescription = "Provides information about an error string from the reference number produced by Akamai edge servers when a request to retrieve content fails."
	translateErrorLongDescription  = `Provides a summary and logs for the error that occurred in the original request using the error string from the reference number.`

	translateErrorUse = "translate-error-string error_string"

	summary                = "Summary"
	urlTranslateError      = "URL"
	httpResponseCode       = "HTTP response code"
	dateAndTime            = "Date and time"
	epocTime               = "Epoch time"
	clientIpTranslateError = "Client IP"
	connectingIp           = "Connecting IP"
	originHostName         = "Origin hostname"
	originIp               = "Origin IP"
	userAgent              = "User agent"
	clientRequest          = "Client request"
	reasonForFailure       = "Reason for failure"
	wafDetails             = "WAF details"

	errorLogs = "Error Logs"
)

//grep
const (
	grepShortDescription = "Runs GREP to retrieve and parse logs for an IP address within the Akamai network using flags to filter the data. Data is available for 48 hours after the traffic occurs."
	grepLongDescription  = `Runs the GREP command to retrieve and parse logs for an IP address within the Akamai network using flags to filter the data. Logs provide low-level details on how each request was handled, which you can use to troubleshoot caching and performance issues and to ensure the correct set of Akamai features was applied to the traffic. Data is available for 48 hours after the traffic occurs.`

	grepUse = "grep edge_server_IP --end-date date --end-time time --duration duration --find-in Header:Value --max-lines maximum_log_lines_to_display -r | -f | -rf"

	endDateFlagDescription        = "The end date of log search, in the <yyyy:mm:dd> format."
	endTimeFlagDescription        = "The end time of log search, in the <hh:mm:ss> (UTC) format."
	durationFlagDescription       = "The number of minutes before the `end-date` and `end-time` for which to retrieve logs."
	maxLinesFlagDescription       = "The maximum log lines to display."
	clientRequestFlagDescription  = "Search logs of incoming client requests to the Akamai edge server."
	forwardRequestFlagDescription = "Search logs of forwarded requests from the Akamai edge server."
	findInFlagDescription         = "Where to search, specified as <field>:<value>. Possible `field` values are: `host-header`, `user-agent`, `http-status-code`, `arl`, `cp-code`, and `client-ip`."
)

//estats
const (
	estatsShortDescription = "Provides error statistics on a CP code’s traffic from clients to Akamai edge servers and from Akamai edge servers to origin."
	estatsLongDescription  = `Provides statistics on errors happening in the delivery of websites based on real-time data of CP code's traffic from clients to Akamai edge servers and from Akamai edge servers to origin.`

	estatsUse = "estats URL/CP_code"

	summaryEstats = "Summary"

	edgeStatistics                 = "Edge Statistics"
	percentageFailurEdgeStatistics = "(Percent failure: %.1f%c)"
	edgeStatisticsDescription      = "Edge status code distribution"
	statusCodeEdgeStatistics       = "Status code"
	percentageHitEdgeStatistics    = "% Hits"

	originalStatistics                 = "Origin Statistics"
	percentageFailurOriginalStatistics = "(Percent failure: %.1f%c)"
	originalStatisticsDescription      = "Origin status code distribution"
	statusCodeOriginalStatistics       = "Status code"
	percentageHitsOriginalStatistics   = "% Hits"

	edgeErrors             = "Edge Errors"
	edgeErrorsDescription  = "View last 10 edge errors"
	edgeIpEdgeErrors       = "Edge IP"
	regionEdgeErrors       = "Region"
	httpStatusEdgeErrors   = "HTTP status"
	hitsEdgeErrors         = "Hits"
	objectStatusEdgeErrors = "Object status"
	errorCodeEdgeErrors    = "Error code"

	originErrors             = "Origin Errors"
	originErrorsDescription  = "View last 10 origin errors"
	edgeIpOriginErrors       = "Edge IP"
	regionOriginErrors       = "Region"
	httpStatusOriginErrors   = "HTTP status"
	hitsOriginErrors         = "Hits"
	objectStatusOriginErrors = "Object status"
	errorCodeOriginErrors    = "Error code"
)

//debug_url
const (
	debugUrlShortDescription = "Provides DNS information, HTTP response, response headers, and logs for a URL on Akamai edge servers."
	debugUrlLongDescription  = `Provides DNS information, HTTP response, response headers, and logs for a URL on Akamai edge servers.`

	headerFlagDescription = "Any additional headers to add to the request, in the <header>:<value> format."
	edgeIpFlagDescription = "The Akamai edge server IP address to test the URL against, otherwise a random server by default."

	debugUrlUse = "debug-url URL --edge-ip edge_server_IP --header request_header"

	dnsInformation         = "DNS information"
	httpResponse           = "HTTP response"
	responseHeaderDebugUrl = "Response header"
	logs                   = "Logs"
)

//user_diagnostics
const (
	userDiagnosticsShortDescription = "Use this tool to create a sharable link and send it to the end users to collect diagnostic data. You can view, filter, and export the results."
	userDiagnosticsLongDescription  = `Use this tool to create a sharable link and send it to the end users to collect diagnostic data. You can view, filter, and export the results.`

	userDiagnosticsUse = "user-diagnostics"
)

//user_diagnostics_list
const (
	userDiagnosticsListShortDescription = "Lists all groups created to gather diagnostic data of end users of hostnames experiencing issues together with the generated links and number of collected data."
	userDiagnosticsListLongDescription  = `Lists all groups created to gather diagnostic data of end users of hostnames experiencing issues together with the generated links and number of collected data.`

	userDiagnosticsListUse = "list"

	userDiagnosticsListNote   = "Note: Each link is active for 7 days and has a limit of 50 submissions."
	linkId                    = "Link ID"
	statusUserDiagnosticsLink = "Status"
	generatedOn               = "Generated on (UTC)"
	groupName                 = "Group name"
	hostNameOrUrl             = "Hostname/URL"
	caseId                    = "Case ID"
	userSharableLink          = "User sharable link"
	records                   = "Records"
)

//user_diagnostics_create_group
const (
	userDiagnosticsCreateGroupShortDescription = "Creates a group for a hostname you want to gather diagnostic data for. It also generates a diagnostic link that you can send to end users of the group’s hostname or URL. When end users click the link, the tool gathers necessary diagnostic data to submit."
	userDiagnosticsCreateGroupLongDescription  = `Creates a group for a hostname you want to gather diagnostic data for. It also generates a diagnostic link that you can send to end users of the group’s hostname or URL. When end users click the link, the tool gathers necessary diagnostic data to submit.`

	userDiagnosticsCreateGroupUse = "create-group group_name hostname"

	linkText1 = "Here is your link!"
	linkText2 = "Copy and send the link below to the end users who are experiencing the content delivery problem.\nEach link is active for 7 days and has a limit of 50 submissions."
	linkText3 = "Link: "
)

//user_diagnostics_get
const (
	userDiagnosticsGetShortDescription = "Lists end users' diagnostic data submitted using a diagnostic link."
	userDiagnosticsGetLongDescription  = `Lists end users' diagnostic data submitted using a diagnostic link.`

	userDiagnosticsGetUse = "get link_id"

	viewUserDiagnosticsData           = "View User Diagnostic Data"
	generatedOnUserDiagnosticGet      = "Generated on"
	groupNameUserDiagnosticGet        = "Group name"
	hostNameOrUrlUserDiagnosticGet    = "Hostname/URL"
	userSharableLinkUserDiagnosticGet = "User sharable link"
	linkStatus                        = "Link status"

	uid                        = "UID"
	timestamp                  = "Timestamp (UTC)"
	clientIpPreferred          = "Client IP preferred"
	clientDnsIpv4              = "Client DNS IPv4"
	clientDnsIpv6              = "Client DNS IPv6"
	userAgentUserDiagnosticGet = "User agent"
	cookie                     = "Cookie"
	protocol                   = "Protocol"
	connectedCipher            = "Connected cipher"
	clientIpv4                 = "Client IPv4"
	clientIpv6                 = "Client IPv6"

	edgeIps = "Edge IPs"

	curlDescription = " - Request content from Akamai edge server, run: "
	digDescription  = " - Get domain details from an Akamai edge server, run: "

	edgeIpMessage = "Edge IP is not shown either because of a system error or the client DNS has an ECS (EDNS Client Subnet)."
)

//generic constants
const (
	clientTypeKey   = "clientType"
	clientTypeValue = "cli"
)
