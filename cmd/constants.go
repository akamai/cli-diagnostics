package cmd

//command usage and example
const (
	locateIpUse     = "locate-ip IP_ADDRESS..."
	locateIpExample = ` $ akamai diagnostics locate-ip 123.123.123.123 8.8.8.8`

	verifyIpUse     = "verify-ip IP_ADDRESS..."
	verifyIpExample = ` $ akamai diagnostics verify-ip 123.123.123.123 8.8.8.8`

	verifyLocateIpUse     = "verify-locate-ip IP_ADDRESS"
	verifyLocateIpExample = ` $ akamai diagnostics verify-locate-ip 123.123.123.123`

	digUse     = "dig --hostname HOSTNAME [-l CLIENT_LOCATION | -e EDGE_SERVER_IP]  [-q QUERY_TYPE] [--gtm]"
	digExample = ` $ akamai diagnostics dig --hostname www.example.com --client-location bangalore-india --query-type NS --gtm
 $ akamai diagnostics dig --hostname  www.easybrazilinvesting.com  --edge-server-ip 123.123.123.123 --query-type A  
 $ akamai diagnostics dig --hostname  www.easybrazilinvesting.com`

	userDiagnosticsUse = "user-diagnostics"

	userDiagnosticsCreateUse     = "create {--url URL | --ipa-hostname IPA_HOSTNAME} [--notes NOTE]"
	userDiagnosticsCreateExample = ` $ akamai diagnostics user-diagnostics create --url https://www.akamai.com
 $ akamai diagnostics user-diagnostics create --url https://www.akamai.com --notes "Tokyo olympics"
 $ akamai diagnostics user-diagnostics create --ipa-hostname https://www.akamai.com --notes "Tokyo olympics"`

	userDiagnosticsListUse     = "list [--url URL|IPA_HOSTNAME] [--user USER] [--active]"
	userDiagnosticsListExample = ` $ akamai diagnostics user-diagnostics list
 $ akamai diagnostics user-diagnostics list --url https://www.akamai.com --user johnDoe --active
 $ akamai diagnostics user-diagnostics list --url https://www.akamai.com --active`

	userDiagnosticsGetUse     = "get LINK_ID [--mtr] [--dig] [--curl]"
	userDiagnosticsGetExample = ` $ akamai diagnostics user-diagnostics get ab123c
 $ akamai diagnostics user-diagnostics get ab123c --mtr --curl
 $ akamai diagnostics user-diagnostics get ab123c --dig --mtr`

	translateUrlUse     = "translate-url URL"
	translateUrlExample = ` $ akamai diagnostics translate-url http://www.example.com`

	translateErrorUse     = "translate-error-string ERROR_STRING [--chase-origin-logs]"
	translateErrorExample = ` $ akamai diagnostics translate-error-string 9.6f64d440.1318965461.2f2b078 --chase-origin-logs`

	urlHealthCheckUse     = "url-health-check URL [--client-location LOCATION] [--edge-server-ip EDGE_SERVER_IP] [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP] [--request-header REQUEST_HEADER...] [-q QUERY_TYPE] [--run-from-site-shield-map] [--logs] [--network-connectivity]"
	urlHealthCheckExample = ` $ akamai diagnostics  url-health-check http://www.example.com --run-from-site-shield-map
 $ akamai diagnostics url-health-check http://www.example.com --client-location bangalore-india --edge-server-ip 123.123.123.123 --port 80 --packet-type TCP --ip-version IPV4 --logs --network-connectivity  --request-header "X-Location: NGDT"`

	edgeLocationsUse     = "edge-locations [--search REGION]"
	edgeLocationsExample = ` $ akamai diagnostics edge-locations
 $ akamai diagnostics edge-locations --search india`

	ipaHostnamesUse     = "ipa-hostnames"
	ipaHostnamesExample = `$ akamai diagnostics ipa-hostnames`

	gtmHostnamesUse     = "gtm-hostnames [--test-target-ip GTM_HOSTNAME]"
	gtmHostnamesExample = ` $ akamai diagnostics gtm-hostnames
 $ akamai diagnostics gtm-hostnames --test-target-ip www-origin.20000puzzles.akadns.net`

	curlUse     = "curl URL [--client-location CLIENT_LOCATION | --edge-server-ip EDGE_SERVER_IP] [--ip-version IPv4|IPv6] [--request-header REQUEST_HEADER...] [--run-from-site-shield-map]"
	curlExample = ` $ akamai diagnostics curl http://www.example.com --client-location bangalore-india --ip-version IPv4 --request-header "accept:text/html"
 $ akamai diagnostics curl http://www.example.com
 $ akamai diagnostics curl http://www.example.com -l bangalore-india -i IPv4 -H "accept:text/html"
 $ akamai diagnostics curl http://www.example.com -e 56.73.66.33 --run-from-site-shield-map`

	mtrUse     = "mtr --source SOURCE_IP|SOURCE_LOCATION --destination DESTINATION_IP|HOSTNAME [--gtm-hostname GTM_HOSTNAME] [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP] [--site-shield-hostname HOSTNAME]"
	mtrExample = ` $ akamai diagnostics mtr --source bangalore-india --destination www.example.com --ip-version IPv4 --port 443 --packet-type icmp
 $ akamai diagnostics mtr --source bangalore-india --destination 121.121.121.121
 $ akamai diagnostics mtr --source 123.123.123.123 --destination 121.121.121.121
 $ akamai diagnostics mtr --source 1.1.1.1 --destination 2.2.2.2 --gtm-hostname example.com
 $ akamai diagnostics mtr --source 123.123.123.123 --destination 121.121.121.121 --site-shield-hostname www.example1.com`

	grepUse     = "grep EDGE_IP START_TIME END_TIME {--hostname HOSTNAME ... | --cp-code CP_CODE ...} [--client-ip CLIENT_IP ...] [--user-agent USER_AGENT ...] [--http-status-code HTTP_STATUS_CODE ... | --error-status-codes] [--arl ARL ...] [-r] [-f]"
	grepExample = ` $ akamai diagnostics grep 123.123.123.123  "2021-01-01T01:00:00.000Z" "2021-01-01T01:30:00.000Z" --hostname "www.akamai.com" --client-ip "123.123.123.123" --http-status-code "400, 401" -rf
 $ akamai diagnostics grep 123.123.123.123  "2021-01-01T01:00:00.000Z" "2021-01-01T01:30:00.000Z" --cp-code 12345 --client-ip "123.123.123.123" --http-status-code "400, 401" -rf`

	estatsUse     = "estats {--url URL | --cp-code CP_CODE} [--logs] [--enhanced-tls | --standard-tls] [--edge-errors] [--origin-errors]"
	estatsExample = ` $ akamai diagnostics estats --url https://www.example.com
 $ akamai diagnostics estats --cp-code 12345 --enhanced-tls --edge-errors
 $ akamai diagnostics estats --url https://www.example.com --logs
 $ akamai diagnostics estats --cp-code 12345 --logs --standard-tls`

	connectivityProblemsUse     = "connectivity-problem URL [--client-location LOCATION] [--edge-server-ip EDGE_SERVER_IP] [--client-ip CLIENT_IP] [--request-header REQUEST_HEADER...] [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP] [--run-from-site-shield-map]"
	connectivityProblemsExample = ` $ akamai diagnostics connectivity-problem http://www.example.com --client-location bangalore-india
 $ akamai diagnostics connectivity-problem http://www.example.com --run-from-site-shield-map
 $ akamai diagnostics connectivity-problem http://www.example.com --client-location bangalore-india --edge-server-ip 123.123.123.123 --client-ip 123.123.123.123 --request-header accept:text/html --port 80 --packet-type TCP --ip-version IPV4`

	contentProblemsUse     = "content-problem URL [--client-location LOCATION] [--edge-server-ip EDGE_IP] [--request-header REQUEST_HEADER...] [--ip-version IP_VERSION] [--run-from-site-shield-map]"
	contentProblemsExample = ` $ akamai diagnostics content-problem http://www.example.com
 $ akamai diagnostics content-problem http://www.example.com --client-location bangalore-india --run-from-site-shield-map
 $ akamai diagnostics content-problem http://www.example.com --client-location bangalore-india --edge-server-ip 123.123.123.123 --request-header accept:text/html --ip-version IPV4`
)
