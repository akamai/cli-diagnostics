package internal

var isEdgeIpTemplate = "IP address {{.IpAddress}} is an Edge IP"
var isNotEdgeIpTemplate = "IP address {{.IpAddress}} is not an Edge IP"

var digTemplate = `
{{ header "DIG RESULTS" }}

{{ header "REQUEST"}}
{{ bold "Hostname:"}} {{.Request.Hostname}}
{{ bold "Edge server:"}} {{.InternalIp}} {{locationToString .EdgeIpLocation}}
{{ bold "Query Type:"}} {{.Request.QueryType}}
{{ bold "Request Date:"}} {{dateToString .CreatedTime}}
{{ bold "Request By:"}} {{.CreatedBy}}

{{ header "RESPONSE"}}
{{ .Result.Result}} `

var userDiagnosticsCreateTemplate = `
{{ bold "Here is your link!"}}
Copy and send the link below to the end users who are experiencing the content delivery problem.
Each link is active for 7 days and has a limit of 50 submissions."

link: {{ bold .DiagnosticLink}}`

var userDiagnosticsGetTemplate = `
{{ header "User Diagnostic Data Result"}}
{{ bold "Generated on"}}: {{  .CreatedTime}}
{{ bold "Notes" }}: {{ .Note }}
{{ bold "URL" }}: {{ .URL }}
{{ bold "Diagnostic Link" }}: {{ .DiagnosticLink }}
{{ bold "Link Status" }}: {{ capsToTitle .DiagnosticLinkStatus }}
{{ italic "View the JSON output for diagnostic details, use --json"}}
`
var translateUrlTemplate = `
{{ header "Translate url"}}

{{ bold "Type Code"}}: {{  .TranslatedUrl.TypeCode}}
{{ bold "Cache Key Hostname"}}: {{  .TranslatedUrl.CacheKeyHostname}}
{{ bold "CP Code"}}: {{  .TranslatedUrl.CpCode}}
{{ bold "Serial Number"}}: {{  .TranslatedUrl.SerialNumber}}
{{ bold "TTL"}}: {{  .TranslatedUrl.Ttl}} `

var curlTemplate = `
{{ header "CURL RESULTS"}}

{{ header "REQUEST"}}
{{ bold "URL"}}: {{ .Request.Url}}
{{ bold "Edge server:"}} {{.InternalIp}} {{locationToString .EdgeIpLocation}} {{if .SiteShieldIp}}
{{ bold "Site shield IP"}}: {{.SiteShieldIp}} {{locationToString .SiteShieldIpLocation}} {{end}}
{{ bold "IP version"}}: {{.Request.IpVersion}}
{{ bold "Request header"}}: {{.Request.RequestHeaders}}
{{ bold "Request date"}}: {{isoDate .CreatedTime}}
{{ bold "Request by"}}: {{.CreatedBy}}

{{ header "Response Headers"}}{{ range $ind, $value := .CurlOutput.ResponseHeaderList }}
{{ splitHeaderValue $value }}{{ end }}{{if .CurlOutput.Timing.DnsLookupTime}}

{{ header "Timing Data"}}
{{ bold "DNS Lookup Time"}}: {{ .CurlOutput.Timing.DnsLookupTime}}
{{ bold "TCP Connection Time"}}: {{ .CurlOutput.Timing.TcpConnectionTime}}
{{ bold "SSL Connection Time"}}: {{ .CurlOutput.Timing.SslConnectionTime}}
{{ bold "Time To First Byte"}}: {{ .CurlOutput.Timing.TimeToFirstByte}}
{{ bold "Total Time"}}: {{ .CurlOutput.Timing.TotalTime}}{{ end }}
`
var mtrDetailsTemplate = `
{{ header "MTR RESULTS" }}

{{ header "REQUEST" }}
{{ bold "Source" }}: {{ .SourceInternalIP }} {{ locationToString .SourceIPLocation }}
{{ bold "Destination" }}: {{ .DestinationInternalIP }} {{ locationToString .DestinationIPLocation }} {{if .SiteShieldIp}}
{{ bold "Site shield IP"}}: {{.SiteShieldIp}} {{locationToString .SiteShieldIpLocation}} {{end}}
{{ bold "Request Date:"}} {{dateToString .CreatedTime}}
{{ bold "Request By:"}} {{.CreatedBy}}

{{ header "Network Connectivity Test from" }} {{ header .SourceInternalIP }} {{ header "to" }} {{ header .DestinationInternalIP }}
`
