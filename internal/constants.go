package internal

// HTTP headers
const (
	ContentType = "Content-Type"
)

// HTTP header values
const (
	ApplicationJson = "application/json"
)

const (
	X_ED_CLIENT_TYPE = "X-ED-CLIENT-TYPE"
)

// constant message keys
const (
	Short                = "short"
	Long                 = "long"
	Empty                = ""
	Global               = "global"
	Missing              = "missing"
	Exclusive            = "exclusive"
	Invalid              = "invalid"
	Redundant            = "redundant"
	RequestParsingError  = "requestParsingError"
	ResponseParsingError = "responseParsingError"
	SpinnerMessage       = "spinnerMessage"
	MissingArgs          = "missingArgs"
	FieldsNotRequired    = "fieldsNotRequired"
	DefaultPort          = 80
	IPV4                 = "IPV4"
	IPV6                 = "IPV6"
	TCP                  = "TCP"
	ICMP                 = "ICMP"
	FailedIps            = "failedIps"
)

// HTTP request methods
const (
	Get  = "GET"
	Post = "POST"
)

// JSON output
const (
	indentString = "\t"
)

// Geolocation fields
const (
	geographicLocation = "GEOGRAPHIC LOCATION"
	ipAddress          = "IP address"
	countryCode        = "Country code"
	regionCode         = "Region code"
	city               = "City"
	dma                = "DMA"
	msa                = "MSA"
	pmsa               = "PMSA"
	areaCode           = "Area code"
	latitude           = "Latitude"
	longitude          = "Longitude"
	county             = "County"
	continent          = "Continent"
	fisp               = "FIPS"
	timeZone           = "Time zone"
	zipCode            = "Zip code"
	proxy              = "Proxy"

	networkLocation = "NETWORK LOCATION"
	network         = "Network"
	networkType     = "Network type"
	asNum           = "ASN"
	throughput      = "Throughput"
)

// User diagnostics data table
const (
	userDiagnosticsListNote   = "NOTE: Each link is active for 7 days and has a limit of 50 submissions."
	linkId                    = "LINK ID"
	hostNameOrUrl             = "URL"
	statusUserDiagnosticsLink = "LINK STATUS"
	diagnosticLink            = "DIAGNOSTIC LINK"
	results                   = "RESULTS"
	user                      = "user"
	requestDate               = "REQUEST DATE"
)

// Edge Locations
const (
	Search        = "search"
	edgeLocations = "EDGE SERVER LOCATIONS"
)

// Ipa Hostnames
const (
	ipaHostnames = "IP ACCELERATION HOSTNAMES"
	ipaHostname  = "IP ACCELERATION HOSTNAME"
)

// Gtm Hostnames
const (
	TestTargetIp = "testTargetIp"
	gtmHostnames = "GTM HOSTNAMES"
	gtmHostname  = "GTM HOSTNAME"
	testIp       = "TEST IP"
	target       = "TARGET"
)

// Mtr constants
const (
	sourceTypeLocation  = "LOCATION"
	sourceTypeEdgeIp    = "EDGE_IP"
	destinationTypeHost = "HOST"
	destinationTypeIp   = "IP"
	mtrTableCol         = "Loss %, Sent, Last, Avg, Best, Worst, StDev, Location"
)

const curlHeaderSeparator = ":"

const (
	CliErrExitCode     int = 1
	CmdErrExitCode     int = 2
	ParsingErrExitCode int = 3
)

const HttpCodeExitCodeDiff = 300
