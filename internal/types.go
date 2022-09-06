package internal

//Geolocation
type GeoLocation struct {
	AreaCode    string   `json:"areaCode"`
	AsNum       *int     `json:"asNumber"`
	City        string   `json:"city"`
	ClientIP    string   `json:"clientIp"`
	Continent   string   `json:"continent"`
	CountryCode string   `json:"countryCode"`
	County      string   `json:"county"`
	DMA         *int     `json:"dma"`
	FIPS        string   `json:"fips"`
	Latitude    *float32 `json:"latitude"`
	Longitude   *float32 `json:"longitude"`
	MSA         *int     `json:"msa"`
	Network     string   `json:"network"`
	NetworkType string   `json:"networkType"`
	PMSA        *int     `json:"pmsa"`
	Proxy       string   `json:"proxy"`
	RegionCode  string   `json:"regionCode"`
	Throughput  string   `json:"throughput"`
	TimeZone    string   `json:"timeZone"`
	ZipCode     string   `json:"zipCode"`
}

// VerifyLocateIp Result
type VerifyLocateIpData struct {
	IsEdgeIp    bool        `json:"isEdgeIp"`
	GeoLocation GeoLocation `json:"geoLocation"`
}

// VerifyLocateIpRequest
type VerifyLocateIpRequest struct {
	IpAddress string `json:"ipAddress"`
}

// VerifyLocateIpResponse
type VerifyLocateIpResponse struct {
	Request         VerifyLocateIpRequest `json:"request"`
	CreatedTime     string                `json:"createdTime"`
	CreatedBy       string                `json:"createdBy"`
	ExecutionStatus string                `json:"executionStatus"`
	Result          VerifyLocateIpData    `json:"result"`
}

// VerifyLocateIpsRequest
type VerifyLocateIpsRequest struct {
	IpAddresses []string `json:"ipAddresses"`
}

// VerifyLocateIpsData
type VerifyLocateIpsData struct {
	ExecutionStatus string      `json:"executionStatus"`
	IpAddress       string      `json:"ipAddress"`
	IsEdgeIp        bool        `json:"isEdgeIp"`
	GeoLocation     GeoLocation `json:"geoLocation"`
}

// VerifyLocateIpsResponse
type VerifyLocateIpsResponse struct {
	Request         VerifyLocateIpsRequest `json:"request"`
	CreatedTime     string                 `json:"createdTime"`
	CreatedBy       string                 `json:"createdBy"`
	ExecutionStatus string                 `json:"executionStatus"`
	Result          []VerifyLocateIpsData  `json:"results"`
}

//DigInfo
type DigInfo struct {
	Result string `json:"result"`
}

//DnsRecord
type DnsRecord struct {
	Domain          string `json:"domain"`
	Hostname        string `json:"hostName"`
	Ttl             string `json:"ttl"`
	RecordClass     string `json:"recordClass"`
	RecordType      string `json:"recordType"`
	PreferenceValue int    `json:"preferenceValue"`
	Value           string `json:"value"`
}

//DigRequest
type DigRequest struct {
	Hostname       string `json:"hostname"`
	QueryType      string `json:"queryType"`
	ClientLocation string `json:"edgeLocationId,omitempty"`
	EdgeServerIp   string `json:"edgeIp,omitempty"`
	IsGtmHostName  bool   `json:"isGtmHostname"`
}

//DigResponse
type DigResponse struct {
	Request          DigRequest     `json:"request"`
	CreatedTime      string         `json:"createdTime"`
	CreatedBy        string         `json:"createdBy"`
	CompletedTime    string         `json:"completedTime"`
	ExecutionStatus  string         `json:"executionStatus"`
	EdgeIpLocation   EdgeIpLocation `json:"edgeIpLocation"`
	InternalIp       string         `json:"internalIp"`
	Result           DigInfo        `json:"result"`
	SuggestedActions []string       `json:"suggestedActions"`
	Type             string         `json:"type"`
	Title            string         `json:"title"`
	Status           int            `json:"status"`
	Detail           string         `json:"detail"`
}

// UserDiagnosticsDataRequest
type UserDiagnosticsDataRequest struct {
	Url         string `json:"url,omitempty"`
	Note        string `json:"note,omitempty"`
	IpaHostname string `json:"ipaHostname,omitempty"`
}

// UserDiagnosticsDataGroupDetails
type UserDiagnosticsDataGroupDetails struct {
	GroupID              string `json:"groupId"`
	Note                 string `json:"note"`
	URL                  string `json:"url"`
	IpaHostname          string `json:"ipaHostname"`
	CreatedTime          string `json:"createdTime"`
	DiagnosticLink       string `json:"diagnosticLink"`
	DiagnosticLinkStatus string `json:"diagnosticLinkStatus"`
	RecordCount          int    `json:"recordCount"`
	CreatedBy            string `json:"createdBy"`
}

//add all list responses, use this as a container
type ListResponse struct {
	Groups []UserDiagnosticsDataGroupDetails `json:"groups,omitempty"`
}

type ArlOutput struct {
	TypeCode         string `json:"typeCode"`
	CacheKeyHostname string `json:"cacheKeyHostname"`
	CpCode           string `json:"cpCode"`
	SerialNumber     string `json:"serialNumber"`
	Ttl              string `json:"ttl"`
	Pragma           string `json:"pragma"`
	CacheControl     string `json:"cacheControl"`
}

type ArlRequest struct {
	Url string `json:"url"`
}

type ArlContainer struct {
	Request       ArlRequest `json:"request"`
	TranslatedUrl ArlOutput  `json:"translatedUrl"`
}

// ErrorTranslatorRequest
type ErrorTranslatorRequest struct {
	ErrorCode       string `json:"errorCode"`
	ChaseOriginLogs bool   `json:"chaseOriginLogs,omitempty"`
}

// ErrorTranslatorResponse
type ErrorTranslatorResponse struct {
	ExecutionStatus string `json:"executionStatus"`
	RetryAfter      int    `json:"retryAfter"`
	Link            string `json:"link"`
}

type UrlHealthCheckRequest struct {
	Url               string   `json:"url"`
	EdgeLocationId    string   `json:"edgeLocationId,omitempty"`
	EdgeIp            string   `json:"spoofEdgeIp,omitempty"`
	Port              int      `json:"port,omitempty"`
	PacketType        string   `json:"packetType,omitempty"`
	IpVersion         string   `json:"ipVersion,omitempty"`
	RequestHeaders    []string `json:"requestHeaders,omitempty"`
	QueryType         string   `json:"queryType,omitempty"`
	ViewsAllowed      []string `json:"viewsAllowed,omitempty"`
	RunFromSiteShield bool     `json:"runFromSiteShield,omitempty"`
}

type UrlHealthCheckResponse struct {
	ExecutionStatus string `json:"executionStatus"`
	RetryAfter      int    `json:"retryAfter"`
	Link            string `json:"link"`
}

type EdgeLocation struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type EdgeLocationContainer struct {
	EdgeLocations []EdgeLocation `json:"edgeLocations"`
}

type IpaHostnameResponse struct {
	Hostnames []string `json:"hostnames"`
}

type GtmProperty struct {
	Property string `json:"property"`
	Domain   string `json:"domain"`
	Hostname string `json:"hostname"`
	DomainId int    `json:"domainId"`
}

type GtmPropertyContainer struct {
	GtmProperties []GtmProperty `json:"gtmProperties"`
}

type GTMPropertyIps struct {
	Property string   `json:"property"`
	Domain   string   `json:"domain"`
	TestIps  []string `json:"testIps"`
	Targets  []string `json:"targets"`
}

type GtmPropertyIpsContainer struct {
	GTMPropertyIps GTMPropertyIps `json:"gtmPropertyIps"`
}

// CurlRequest
type CurlRequest struct {
	Url               string   `json:"url"`
	EdgeIp            string   `json:"edgeIp,omitempty"`
	EdgeLocationId    string   `json:"edgeLocationId,omitempty"`
	IpVersion         string   `json:"ipVersion,omitempty"`
	RequestHeaders    []string `json:"requestHeaders,omitempty"`
	RunFromSiteShield bool     `json:"runFromSiteShield,omitempty"`
}

//EstatsWrapper
type EstatsResultWrapper struct {
	EstatsResult   EstatsResult         `json:"estatsResult"`
	EstatsLogLines []EstatsGrepResponse `json:"estatsLogLines,omitempty"`
}

//EstatsRequest
type EstatsRequest struct {
	Url       string `json:"url,omitempty"`
	CpCode    int    `json:"cpCode,omitempty"`
	Delivery  string `json:"delivery,omitempty"`
	ErrorType string `json:"errorType,omitempty"`
}

type EstatsData struct {
	EdgeErrors                      int                      `json:"edgeErrors"`
	EdgeHits                        int                      `json:"edgeHits"`
	EdgeFailurePercentage           float64                  `json:"edgeFailurePercentage"`
	OriginErrors                    int                      `json:"originErrors"`
	OriginHits                      int                      `json:"originHits"`
	OriginFailurePercentage         float64                  `json:"originFailurePercentage"`
	TopEdgeIpsWithError             []EdgeIpInfo             `json:"topEdgeIpsWithError"`
	TopEdgeIpsWithErrorFromOrigin   []EdgeIpInfo             `json:"topEdgeIpsWithErrorFromOrigin"`
	TopEdgeIpsWithSuccess           []EdgeIpInfo             `json:"topEdgeIpsWithSuccess"`
	TopEdgeIpsWithSuccessFromOrigin []EdgeIpInfo             `json:"topEdgeIpsWithSuccessFromOrigin"`
	EdgeStatusCodeDistribution      []StatusCodeDistribution `json:"edgeStatusCodeDistribution"`
	OriginStatusCodeDistribution    []StatusCodeDistribution `json:"originStatusCodeDistribution"`
}

type EstatsResult struct {
	Request         EstatsRequest `json:"request"`
	RequestId       int           `json:"requestId"`
	CreatedBy       string        `json:"createdBy"`
	CreatedTime     string        `json:"createdTime"`
	CompletedTime   string        `json:"completedTime"`
	ExecutionStatus string        `json:"executionStatus"`
	Result          EstatsData    `json:"result"`
}

type EdgeIpInfo struct {
	ErrorId        string         `json:"errorId,omitempty"`
	LogLinesCount  int            `json:"logLinesCount"`
	EdgeIp         string         `json:"edgeIp"`
	EdgeIpLocation EdgeIpLocation `json:"edgeIpLocation"`
	HttpStatus     int            `json:"httpStatus"`
	Hits           int            `json:"hits"`
	ObjectStatus   []ObjectStatus `json:"objectStatus"`
	ErrorCode      string         `json:"errorCode"`
	EdgeLogsLink   string         `json:"edgeLogsLink"`
}

type StatusCodeDistribution struct {
	HttpStatus int     `json:"httpStatus"`
	Hits       int     `json:"hits"`
	Percentage float64 `json:"percentage"`
}

type ObjectStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type EdgeIpLocation struct {
	City        string   `json:"city"`
	RegionCode  string   `json:"regionCode,omitempty"`
	CountryCode string   `json:"countryCode"`
	AsNum       *int     `json:"asNumber,omitempty"`
	Latitude    *float32 `json:"latitude,omitempty"`
	Longitude   *float32 `json:"longitude,omitempty"`
}

type Warning struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type Timing struct {
	DnsLookupTime     float64 `json:"dnsLookupTime,omitempty"`
	TcpConnectionTime float64 `json:"tcpConnectionTime,omitempty"`
	SslConnectionTime float64 `json:"sslConnectionTime,omitempty"`
	TimeToFirstByte   float64 `json:"timeToFirstByte,omitempty"`
	TotalTime         float64 `json:"totalTime,omitempty"`
}

type CurlOutput struct {
	HttpStatusCode     int               `json:"httpStatusCode"`
	ResponseHeaders    map[string]string `json:"responseHeaders"`
	Timing             Timing            `json:"timing,omitempty"`
	ResponseBody       string            `json:"responseBody"`
	ResponseHeaderList []string          `json:"responseHeaderList"`
	HttpVersion        string            `json:"httpVersion"`
	ReasonPhrase       string            `json:"reasonPhrase"`
	PartialSuccess     bool              `json:"partialSuccess"`
}

type CurlResponse struct {
	Request              CurlRequest    `json:"request"`
	CreatedBy            string         `json:"createdBy"`
	CreatedTime          string         `json:"createdTime"`
	CompletedTime        string         `json:"completedTime"`
	ExecutionStatus      string         `json:"executionStatus"`
	InternalIp           string         `json:"internalIp"`
	EdgeIpLocation       EdgeIpLocation `json:"edgeIpLocation"`
	SiteShieldIp         string         `json:"siteShieldIp"`
	SiteShieldIpLocation EdgeIpLocation `json:"siteShieldIpLocation"`
	CurlOutput           CurlOutput     `json:"result"`
	SuggestedActions     []string       `json:"suggestedActions"`
	Warning              Warning        `json:"warning"`
}

// MtrRequest
type MtrRequest struct {
	Source             string `json:"source"`
	Destination        string `json:"destination"`
	PacketType         string `json:"packetType"`
	SourceType         string `json:"sourceType,omitempty"`
	DestinationType    string `json:"destinationType,omitempty"`
	IPVersion          string `json:"ipVersion,omitempty"`
	Port               int    `json:"port,omitempty"`
	GtmHostname        string `json:"gtmHostname,omitempty"`
	SiteShieldHostname string `json:"siteShieldHostname,omitempty"`
}

// Mtr hop
type MtrHop struct {
	Number            int            `json:"number"`
	IP                string         `json:"ip"`
	IPLocation        EdgeIpLocation `json:"ipLocation"`
	Host              string         `json:"host"`
	PacketLoss        float64        `json:"packetLoss"`
	SentPackets       int            `json:"sentPackets"`
	LastPacketLatency float64        `json:"lastPacketLatency"`
	AverageLatency    float64        `json:"averageLatency"`
	BestRtt           float64        `json:"bestRtt"`
	WorstRtt          float64        `json:"worstRtt"`
	StandardDeviation float64        `json:"standardDeviation"`
}

// MtrResult
type MtrResult struct {
	Host string   `json"host"`
	Hops []MtrHop `json:"hops"`
}

// MtrResponse
type MtrResponse struct {
	SourceIPLocation      EdgeIpLocation `json:"sourceIpLocation"`
	DestinationIPLocation EdgeIpLocation `json:"destinationIpLocation"`
	SourceContext         string         `json:"sourceContext"`
	DestinationContext    string         `json:"destinationContext"`
	SourceInternalIP      string         `json:"sourceInternalIp"`
	DestinationInternalIP string         `json:"destinationInternalIp"`
	CreatedBy             string         `json:"createdBy"`
	CreatedTime           string         `json:"createdTime"`
	SiteShieldIp          string         `json:"siteShieldIp"`
	SiteShieldIpLocation  EdgeIpLocation `json:"siteShieldIpLocation"`
	Result                MtrResult      `json:"result"`
	SuggestedActions      []string       `json:"suggestedActions"`
}

type HttpStatusCode struct {
	Comparison string   `json:"comparison"`
	Value      []string `json:"value"`
}

type Arl struct {
	Comparison string   `json:"comparison"`
	Value      []string `json:"value"`
}

type GrepRequest struct {
	CpCodes         []int           `json:"cpCodes,omitempty"`
	Hostnames       []string        `json:"hostnames,omitempty"`
	EdgeIp          string          `json:"edgeIp"`
	LogType         string          `json:"logType"`
	Start           string          `json:"start"`
	End             string          `json:"end"`
	UserAgents      []string        `json:"userAgent,omitempty"`
	HttpStatusCodes *HttpStatusCode `json:"httpStatusCodes,omitempty"`
	Arls            *Arl            `json:"arls,omitempty"`
	ClientIps       []string        `json:"clientIps,omitempty"`
}

type GrepResponse struct {
	ExecutionStatus string `json:"executionStatus"`
	RetryAfter      int    `json:"retryAfter"`
	Link            string `json:"link"`
}

type EstatsGrepResponse struct {
	LogLinesCount    int            `json:"logLinesCount"`
	ExecutionStatus  string         `json:"executionStatus"`
	CreatedTime      string         `json:"createdTime"`
	CompletedTime    string         `json:"completedTime"`
	CreatedBy        string         `json:"createdBy"`
	LogsContainer    LogsContainer  `json:"result"`
	SuggestedActions []string       `json:"suggestedActions"`
	EdgeIpLocation   EdgeIpLocation `json:"edgeIpLocation"`
	Warning          Warning        `json:"warning"`
}

type LogsContainer struct {
	Logs   []map[string]string `json:"logs"`
	Legend Legend              `json:"legend"`
}

type Legend struct {
	LogType       map[string]string `json:"logType"`
	ObjectStatus  map[string]string `json:"objectStatus"`
	ObjectStatus2 map[string]string `json:"objectStatus2"`
}

type ConnectivityProblemsRequest struct {
	Url               string   `json:"url"`
	EdgeLocationId    string   `json:"edgeLocationId,omitempty"`
	SpoofEdgeIp       string   `json:"spoofEdgeIp,omitempty"`
	ClientIp          string   `json:"clientIp,omitempty"`
	RequestHeaders    []string `json:"requestHeaders,omitempty"`
	IpVersion         string   `json:"ipVersion,omitempty"`
	PacketType        string   `json:"packetType,omitempty"`
	Port              int      `json:"port,omitempty"`
	RunFromSiteShield bool     `json:"runFromSiteShield,omitempty"`
}

type ConnectivityProblemsResponse struct {
	ExecutionStatus string `json:"executionStatus"`
	RetryAfter      int    `json:"retryAfter"`
	Link            string `json:"link"`
}

type ContentProblemsRequest struct {
	Url               string   `json:"url"`
	EdgeLocationId    string   `json:"edgeLocationId,omitempty"`
	EdgeIp            string   `json:"spoofEdgeIp,omitempty"`
	RequestHeaders    []string `json:"requestHeaders,omitempty"`
	IpVersion         string   `json:"ipVersion,omitempty"`
	RunFromSiteShield bool     `json:"runFromSiteShield,omitempty"`
}

type ContentProblemsResponse struct {
	ExecutionStatus string `json:"executionStatus"`
	RetryAfter      int    `json:"retryAfter"`
	Link            string `json:"link"`
}
