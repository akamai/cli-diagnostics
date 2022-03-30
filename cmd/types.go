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
	edgegrid "github.com/akamai/AkamaiOPEN-edgegrid-golang"
)

var econfig edgegrid.Config

type GeoLocation struct {
	AreaCode    string  `json:"areaCode"`
	AsNum       string  `json:"asNum"`
	City        string  `json:"city"`
	ClientIP    string  `json:"clientIp"`
	Continent   string  `json:"continent"`
	CountryCode string  `json:"countryCode"`
	County      string  `json:"county"`
	DMA         *int    `json:"dma"`
	FIPS        string  `json:"fips"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	MSA         *int    `json:"msa"`
	Network     string  `json:"network"`
	NetworkType string  `json:"networkType"`
	PMSA        *int    `json:"pmsa"`
	Proxy       string  `json:"proxy"`
	RegionCode  string  `json:"regionCode"`
	Throughput  string  `json:"throughput"`
	TimeZone    string  `json:"timeZone"`
	ZipCode     string  `json:"zipCode"`
}

type GhostLocationsList struct {
	Locations []map[string]string `json:"locations"`
}

type DigInfo struct {
	HostName         string      `json:"hostName"`
	QueryType        string      `json:"queryType"`
	AnswerSection    []DnsRecord `json:"answerSection"`
	AuthoritySection []DnsRecord `json:"authoritySection"`
	Result           string      `json:"result"`
}

type CurlResults struct {
	HttpStatusCode  int               `json:"httpStatusCode"`
	ResponseHeaders map[string]string `json:"responseHeaders"`
	ResponseBody    string            `json:"responseBody"`
}

type Wrapper struct {
	DigInfo                *DigInfo                `json:"digInfo"`
	Mtr                    *MtrData                `json:"mtr"`
	GeoLocation            *GeoLocation            `json:"geoLocation"`
	TranlatedURL           *TranslatedURL          `json:"translatedUrl"`
	TranslatedError        *TranslatedError        `json:"translatedError"`
	Curl                   *CurlResults            `json:"curlResults"`
	URLDebug               *DebugUrl               `json:"urlDebug"`
	Estats                 *Estats                 `json:"eStats"`
	LogLines               *LogLines               `json:"logLines"`
	EndUserDiagnosticLinks []EndUserDiagnosticLink `json:"endUserDiagnosticLinks"`
}

type Hop struct {
	Avg    float32 `json:"avg"`
	Best   float32 `json:"best"`
	Host   string  `json:"host"`
	Last   float32 `json:"last"`
	Loss   float32 `json:"loss"`
	Number int     `json:"number"`
	Sent   int     `json:"sent"`
	StdDev float32 `json:"stDev"`
	Worst  float32 `json:"worst"`
}

type MtrData struct {
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	StartTime   string  `json:"startTime"`
	Host        string  `json:"host"`
	PacketLoss  float32 `json:"packetLoss"`
	AvgLatency  float32 `json:"avgLatency"`
	Analysis    string  `json:"analysis"`
	Hops        []Hop   `json:"hops"`
	Result      string  `json:"result"`
}

type TranslatedURL struct {
	TypeCode     string `json:"typeCode"`
	OriginServer string `json:"originServer"`
	CpCode       int    `json:"cpCode"`
	SerialNumber int    `json:"serialNumber"`
	TTL          string `json:"ttl"`
}

type Log struct {
	Description string                 `json:"description"`
	Fields      map[string]interface{} `json:"fields"`
}

type TranslatedError struct {
	Url              string `json:"url"`
	HttpResponseCode int    `json:"httpResponseCode"`
	ClientIP         string `json:"clientIp"`
	ConnectingIP     string `json:"connectingIp"`
	CpCode           string `json:"cpCode"`
	EpochTime        int    `json:"epochTime"`
	Logs             []Log  `json:"logs"`
	OriginHostname   string `json:"originHostname"`
	OriginIP         string `json:"originIp"`
	ReasonForFailure string `json:"reasonForFailure"`
	RequestMethod    string `json:"requestMethod"`
	ServerIP         string `json:"serverIp"`
	Timestamp        string `json:"timestamp"`
	UserAgent        string `json:"userAgent"`
	WafDetails       string `json:wafDetails`
}

type DebugUrl struct {
	DNSinformation  []string            `json:"dnsInformation"`
	HTTPResponse    []map[string]string `json:"httpResponse"`
	Logs            []Log               `json:"logs"`
	ResponseHeaders []string            `json:"responseHeaders"`
}

type Estats struct {
	CpCode                          int                      `json:"cpCode"`
	EdgeErrors                      int                      `json:"edgeErrors"`
	EdgeFailurePercentage           float32                  `json:"edgeFailurePercentage"`
	EdgeHits                        int                      `json:"edgeHits"`
	EdgeStatusCodeDistribution      []StatusCodeDistribution `json:"edgeStatusCodeDistribution"`
	OriginErrors                    int                      `json:"originError"`
	OriginFailurePercentage         float32                  `json:"originFailuerPercentage"`
	OriginHits                      int                      `json:"originHits"`
	OriginStatusCodeDistribution    []StatusCodeDistribution `json:"originStatusCodeDistribution"`
	TopEdgeIPsWithError             []EdgeIPinfo             `json:"topEdgeIpsWithError"`
	TopEdgeIPsWithSuccess           []EdgeIPinfo             `json:"topEdgeIpsWithSuccess"`
	TopEdgeIPsWithErrorFromOrigin   []EdgeIPinfo             `json:"topEdgeIpsWithErrorFromOrigin"`
	TopEdgeIPsWithSuccessFromOrigin []EdgeIPinfo             `json:"topEdgeIpsWithSuccessFromOrigin"`
}

type StatusCodeDistribution struct {
	Hits       int     `json:"hits"`
	HTTPStatus int     `json:"httpStatus"`
	Percentage float32 `json:"percentage"`
}

type EdgeIPinfo struct {
	EdgeIP       string `json:"edgeIp"`
	EdgeLogsLink string `json:"edgeLogsLink"`
	ErrorCode    string `json:"errorCode"`
	Hits         int    `json:"hits"`
	HTTPstatus   int    `json:"httpStatus"`
	ObjStatus    string `json:"objStatus"`
	Region       int    `json:"region"`
}

type LogLines struct {
	Headers string   `json:"headers"`
	Logs    []string `json:"logs"`
}

type ResponseError struct {
	Type   string              `json:"type"`
	Title  string              `json:"title"`
	Status int                 `json:"status"`
	Detail string              `json:"detail"`
	Errors []map[string]string `json:"errors"`
}

type IPinfoRecord struct {
	ID              string            `json:"id"`
	IP              string            `json:"ip"`
	IPtype          string            `json:"IPtype"`
	AssociatedDnsIp string            `json:"associatedDnsIp"`
	Ecs             string            `json:"ecs"`
	Location        map[string]string `json:"location"`
}

type DiagnosticRecord struct {
	EndUserDataId     string         `json:"endUserDataId"`
	Cipher            string         `json:"cipher"`
	Cookie            bool           `json:"cookie"`
	Protocol          string         `json:"protocol"`
	UserAgent         string         `json:"userAgent"`
	CreatedDate       string         `json:"createdDate"`
	UniqueId          int            `json:"uniqueId"`
	UserKey           string         `json:"userKey"`
	ClientDnsIpv6     *IPinfoRecord  `json:"clientDnsIpv6"`
	EdgeIPs           []IPinfoRecord `json:"edgeIps"`
	ClientIPv4        *IPinfoRecord  `json:"clientIpv4"`
	ClientIPv6        *IPinfoRecord  `json:"clientIpv6"`
	ClientDnsIpv4     *IPinfoRecord  `json:"clientDnsIpv4"`
	PreferredClientIP *IPinfoRecord  `json:"preferredClientIp"`
}

type UserDiagnosticData struct {
	GroupName         string             `json:"groupName"`
	CreatedDate       string             `json:"createdDate"`
	URL               string             `json:"url"`
	DiagnosticLink    string             `json:"diagnosticLink"`
	CaseIds           []string           `json:"caseIds"`
	Status            string             `json:"status"`
	DiagnosticRecords []DiagnosticRecord `json:"diagnosticRecords"`
}

type EndUserDiagnosticLink struct {
	DiagnosticLinkID string   `json:"diagnosticLinkId"`
	GroupName        string   `json:"groupName"`
	CaseIds          []string `json:"caseIds"`
	URL              string   `json:"url"`
	CreatedDate      string   `json:"createdDate"`
	DiagnosticLink   string   `json:"diagnosticLink"`
	Status           string   `json:"status"`
	RecordCount      int      `json:"recordCount"`
	DiagLinkCode     string   `json:"diagLinkCode"`
}

type CreateGroup struct {
	URL       string `json:"url"`
	GroupName string `json:"groupName"`
}

type DnsRecord struct {
	Domain          string `json:"domain"`
	Ttl             int    `json:"ttl"`
	RecordClass     string `json:"recordClass"`
	RecordType      string `json:"recordType"`
	PreferenceValue string `json:"preferenceValue"`
	Value           string `json:"value"`
}

type ConnectivityProblemsRequest struct {
	Url            string   `json:"url"`
	EdgeLocationId string   `json:"edgeLocationId"`
	SpoofEdgeIp    string   `json:"spoofEdgeIp,omitempty"`
	ClientIp       string   `json:"clientIp,omitempty"`
	RequestHeaders []string `json:"requestHeaders,omitempty"`
	IpVersion      string   `json:"ipVersion,omitempty"`
	PacketType     string   `json:"packetType,omitempty"`
	Port           int      `json:"port,omitempty"`
}

type ConnectivityProblemsResponse struct {
	ExecutionStatus string `json:"executionStatus"`
	RetryAfter      int    `json:"retryAfter"`
	Link            string `json:"link"`
}
