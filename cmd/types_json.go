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

type GeoLocationJson struct {
	IpAddress    string       `json:"ipAddress"`
	ReportedTime string       `json:"reportedTime"`
	GeoLocation  *GeoLocation `json:"geoLocation"`
}

type VerifyIpJson struct {
	IpAddress    string `json:"ipAddress"`
	ReportedTime string `json:"reportedTime"`
	IsCdnIp      bool   `json:"isCdnIp"`
}

type TranslateURLJson struct {
	Url          string         `json:"url"`
	ReportedTime string         `json:"reportedTime"`
	TranlatedURL *TranslatedURL `json:"translatedUrl"`
}

type LogLinesJson struct {
	EdgeServerIp   string    `json:"edgeServerIp"`
	EndDate        string    `json:"endDate"`
	EndTime        string    `json:"endTime"`
	findIn         []string  `json:"findIn"`
	Duration       int       `json:"duration"`
	MaxLines       int       `json:"maxLines"`
	ClientRequest  bool      `json:"clientRequest"`
	ForwardRequest bool      `json:"forwardRequest"`
	ReportedTime   string    `json:"reportedTime"`
	LogLines       *LogLines `json:"logLines"`
}

type EstatsJson struct {
	UrlorCpCode  string  `json:"urlOrCpCode"`
	ReportedTime string  `json:"reportedTime"`
	Estats       *Estats `json:"eStats"`
}

type DebugUrlJson struct {
	Url          string    `json:"url"`
	EdgeIP       string    `json:"edgeIP"`
	Headers      []string  `json:"headers"`
	ReportedTime string    `json:"reportedTime"`
	DebugUrl     *DebugUrl `json:"urlDebug"`
}

type TranslatedErrorJosn struct {
	ErrorCode       string           `json:"errorCode"`
	ReportedTime    string           `json:"reportedTime"`
	TranslatedError *TranslatedError `json:"translatedError"`
}

type CurlResultsJson struct {
	Url                   string       `json:"url"`
	IpAddressOrLocationId string       `json:"ipAddressOrLocationId"`
	UserAgent             string       `json:"userAgent"`
	ReportedTime          string       `json:"reportedTime"`
	Curl                  *CurlResults `json:"curlResults"`
}

type MtrDataJson struct {
	DestinationDomain     string   `json:"destinationDomain"`
	IpAddressOrLocationId string   `json:"isAddressOrLocationId"`
	ResolveDns            bool     `json:"resolveDns"`
	ReportedTime          string   `json:"reportedTime"`
	Mtr                   *MtrData `json:"mtr"`
}

type DigInfoJson struct {
	HostName              string   `json:"hostName"`
	IpAddressOrLocationId string   `json:"ipAdderssOrLocationId"`
	QueryType             string   `json:"queryType"`
	ReportedTime          string   `json:"reportedTime"`
	DigInfo               *DigInfo `json:"digInfo"`
}

type EndUserDiagnosticLinkJson struct {
	ReportedTime           string                  `json:"reportedTime"`
	EndUserDiagnosticLinks []EndUserDiagnosticLink `json:"endUserDiagnosticLinks"`
}

type DiagnosticLinkResponse struct {
	GroupName      string `json:"groupName"`
	Url            string `json:"url"`
	ReportedTime   string `json:"reportedTime"`
	DiagnosticLink string `json:"diagnosticLink"`
}

type UserDiagnosticDataJson struct {
	LinkId             string             `json:"linkId"`
	ReportedTime       string             `json:"reportedTime"`
	UserDiagnosticData UserDiagnosticData `json:"endUserDiagnosticData"`
}
