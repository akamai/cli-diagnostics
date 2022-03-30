package internal

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Validator struct {
	cmd      *cobra.Command
	jsonData []byte
}

func NewValidator(cmd *cobra.Command, jsonData []byte) *Validator {
	return &Validator{cmd, jsonData}
}

func isValidIp(str string) bool {
	ip := net.ParseIP(strings.TrimSpace(str))
	return ip != nil
}

func (validator Validator) validateIpArg(ipAddress string, field string) {

	if !isValidIp(ipAddress) {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForArg(validator.cmd, Invalid, field), CmdErrExitCode)
	}
}

func (validator Validator) validateIpFlag(ipAddress, field string, required bool) {

	if ipAddress == "" {
		if required {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, field), CmdErrExitCode)
		}
		return
	}
	if !isValidIp(ipAddress) {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, field), CmdErrExitCode)
	}
}

func isAbsoluteURL(str string) bool {
	urlCheck, err := url.Parse(str)
	if err != nil || !urlCheck.IsAbs() {
		return false
	}
	return true
}

func (validator Validator) validateUrlArg(url string, field string) {

	if !isAbsoluteURL(url) {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForArg(validator.cmd, Invalid, field), CmdErrExitCode)
	}
}

func (validator Validator) validateUrlFlag(url, field string, required bool) {

	if url == "" {
		if required {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, field), CmdErrExitCode)
		}
		return
	}
	if !isAbsoluteURL(url) {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, field), CmdErrExitCode)
	}
}

func (validator Validator) ValidatePortFlag(port string, portInt *int, field string, required bool) {

	if port == "" {
		if required {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, field), CmdErrExitCode)
		}
		return
	}

	if port != "80" && port != "443" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, field), CmdErrExitCode)
	}

	*portInt, _ = strconv.Atoi(port)

}

func (validator Validator) ValidatePacketTypeFlag(packetType *string, field string, required bool) {

	if *packetType == "" {
		if required {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, field), CmdErrExitCode)
		}
		return
	}

	*packetType = strings.ToUpper(*packetType)

	if *packetType != TCP && *packetType != ICMP {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, field), CmdErrExitCode)
	}

}

func (validator Validator) ValidateIpVersionFlag(ipVersion *string, field string, required bool) {

	if *ipVersion == "" {
		if required {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, field), CmdErrExitCode)
		}
		return
	}

	*ipVersion = strings.ToUpper(*ipVersion)

	if *ipVersion != IPV4 && *ipVersion != IPV6 {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, field), CmdErrExitCode)
	}

}

func (validator Validator) ValidateVerifyLocateIpFields(args []string, verifyLocateIpRequest *VerifyLocateIpRequest) {

	if validator.jsonData != nil {
		if len(args) > 0 {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*verifyLocateIpRequest = VerifyLocateIpRequest{}
		ByteArrayToStruct(validator.jsonData, &verifyLocateIpRequest)
		return
	}

	if len(args) != 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}

	validator.validateIpArg(args[0], "ipAddress")
	verifyLocateIpRequest.IpAddress = args[0]
}

func (validator Validator) ValidateVerifyIpOrLocateIpFields(args []string, verifyLocateIpsRequest *VerifyLocateIpsRequest) {

	if validator.jsonData != nil {
		if len(args) > 0 {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*verifyLocateIpsRequest = VerifyLocateIpsRequest{}
		ByteArrayToStruct(validator.jsonData, &verifyLocateIpsRequest)
		return
	}

	if len(args) < 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}

	for _, arg := range args {
		validator.validateIpArg(arg, "ipAddress")
		verifyLocateIpsRequest.IpAddresses = append(verifyLocateIpsRequest.IpAddresses, arg)
	}
}

func (validator Validator) ValidateUserDiagnosticsCreateFields(args []string, userDiagnosticsDataRequest *UserDiagnosticsDataRequest) {

	if validator.jsonData != nil {
		if userDiagnosticsDataRequest.Url != "" || userDiagnosticsDataRequest.IpaHostname != "" || userDiagnosticsDataRequest.Note != "" {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*userDiagnosticsDataRequest = UserDiagnosticsDataRequest{}
		ByteArrayToStruct(validator.jsonData, userDiagnosticsDataRequest)
		return
	}

	if len(args) != 0 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 0, len(args)), CmdErrExitCode)
	}

	if userDiagnosticsDataRequest.Url == "" && userDiagnosticsDataRequest.IpaHostname == "" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "any"), CmdErrExitCode)
	} else if userDiagnosticsDataRequest.Url != "" && userDiagnosticsDataRequest.IpaHostname != "" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "exclusive"), CmdErrExitCode)
	}

	validator.validateUrlFlag(userDiagnosticsDataRequest.Url, "url", false)

}

func (validator Validator) ValidateCurlFields(args []string, curlRequest *CurlRequest) {

	if validator.jsonData != nil {
		if len(args) > 0 || curlRequest.EdgeLocationId != "" || curlRequest.EdgeIp != "" || curlRequest.IpVersion != IPV4 || len(curlRequest.RequestHeaders) > 0 {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*curlRequest = CurlRequest{}
		ByteArrayToStruct(validator.jsonData, curlRequest)
		return
	}

	if len(args) != 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}

	validator.validateUrlArg(args[0], "url")
	curlRequest.Url = args[0]

	if curlRequest.EdgeIp != "" && curlRequest.EdgeLocationId != "" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "exclusive"), CmdErrExitCode)
	}

	validator.validateIpFlag(curlRequest.EdgeIp, "edgeServerIp", false)
	validator.ValidateIpVersionFlag(&curlRequest.IpVersion, "ipVersion", false)

}

func (validator Validator) ValidateDigFields(args []string, digRequest *DigRequest) {

	if validator.jsonData != nil {
		if digRequest.Hostname != "" || digRequest.ClientLocation != "" || digRequest.EdgeServerIp != "" || digRequest.QueryType != "A" || digRequest.IsGtmHostName {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*digRequest = DigRequest{}
		ByteArrayToStruct(validator.jsonData, digRequest)
		return
	}

	if len(args) != 0 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 0, len(args)), CmdErrExitCode)
	}

	if digRequest.Hostname == "" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "hostname"), CmdErrExitCode)
	}
	if digRequest.ClientLocation != "" && digRequest.EdgeServerIp != "" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "exclusive"), CmdErrExitCode)
	}

	validator.validateIpFlag(digRequest.EdgeServerIp, "edgeServerIp", false)

}

func (validator Validator) ValidateEstatsFields(args []string, estatsRequest *EstatsRequest, logs, enhancedTls, standardTls, edgeErrors, originErrors bool) {

	if validator.jsonData != nil {
		if estatsRequest.Url != "" || estatsRequest.CpCode != 0 || enhancedTls || standardTls || edgeErrors || originErrors {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*estatsRequest = EstatsRequest{}
		ByteArrayToStruct(validator.jsonData, estatsRequest)
		return
	}

	if len(args) != 0 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 0, len(args)), CmdErrExitCode)
	}

	if estatsRequest.Url == "" && estatsRequest.CpCode == 0 {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "any"), CmdErrExitCode)
	} else if estatsRequest.Url != "" && estatsRequest.CpCode != 0 {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "exclusive"), CmdErrExitCode)
	}

	validator.validateUrlFlag(estatsRequest.Url, "url", false)

	if enhancedTls && standardTls {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "exclusiveDelivery"), CmdErrExitCode)
	}

	if enhancedTls {
		estatsRequest.Delivery = "ENHANCED_TLS"
	} else if standardTls {
		estatsRequest.Delivery = "STANDARD_TLS"
	}

	if edgeErrors && !originErrors {
		estatsRequest.ErrorType = "EDGE_ERRORS"
	} else if !edgeErrors && originErrors {
		estatsRequest.ErrorType = "ORIGIN_ERRORS"
	}

}
func (validator Validator) ValidateMtrFields(args []string, portStr string, mtrRequest *MtrRequest) {

	if validator.jsonData != nil {
		if mtrRequest.Source != "" || mtrRequest.Destination != "" || mtrRequest.GtmHostname != "" || mtrRequest.PacketType != TCP || portStr != "" || mtrRequest.IPVersion != "" {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*mtrRequest = MtrRequest{}
		ByteArrayToStruct(validator.jsonData, mtrRequest)
		return
	}

	if len(args) != 0 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 0, len(args)), CmdErrExitCode)
	}

	if mtrRequest.Source == "" && mtrRequest.SiteShieldHostname == "" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "source"), CmdErrExitCode)
	}

	if mtrRequest.Destination == "" {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "destination"), CmdErrExitCode)
	}

	//GTM scenario
	if mtrRequest.GtmHostname != "" {
		if !isValidIp(mtrRequest.Source) {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "gtmScenario"), CmdErrExitCode)
		}
		if portStr != "" {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Redundant, "port"), CmdErrExitCode)
		}
		if mtrRequest.IPVersion != "" {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Redundant, "ipVersion"), CmdErrExitCode)
		}
		return
	}

	if mtrRequest.Source != "" {
		if isValidIp(mtrRequest.Source) {
			mtrRequest.SourceType = sourceTypeEdgeIp
		} else {
			mtrRequest.SourceType = sourceTypeLocation
		}
	}

	// destination is IP, --port and --ip-version are redundant
	if isValidIp(mtrRequest.Destination) {
		if portStr != "" {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Redundant, "port"), CmdErrExitCode)
		}
		if mtrRequest.IPVersion != "" {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Redundant, "ipVersion"), CmdErrExitCode)
		}
		mtrRequest.DestinationType = destinationTypeIp
	} else {
		if mtrRequest.SiteShieldHostname != "" {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "destination"), CmdErrExitCode)
		}
		//set default values for port and ipversion if destination is hostname
		if portStr == "" {
			mtrRequest.Port = DefaultPort
		} else {
			validator.ValidatePortFlag(portStr, &mtrRequest.Port, "port", false)
		}
		if mtrRequest.IPVersion == "" {
			mtrRequest.IPVersion = IPV4
		} else {
			validator.ValidateIpVersionFlag(&mtrRequest.IPVersion, "ipVersion", false)
		}
		mtrRequest.DestinationType = destinationTypeHost
	}
	validator.ValidatePacketTypeFlag(&mtrRequest.PacketType, "packetType", true)
}

func (validator Validator) ValidateTranslateUrlFields(args []string, translateUrlRequest *ArlRequest) {

	if validator.jsonData != nil {
		if len(args) > 0 {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		ByteArrayToStruct(validator.jsonData, translateUrlRequest)
		return
	}

	if len(args) != 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}

	validator.validateUrlArg(args[0], "url")
	translateUrlRequest.Url = args[0]

}

func (validator Validator) ValidateUrlHealthCheckFields(args []string, portStr string, urlHealthCheckRequest *UrlHealthCheckRequest, logs, networkConnectivity bool) {

	if validator.jsonData != nil {
		if len(args) > 0 || urlHealthCheckRequest.EdgeIp != "" || urlHealthCheckRequest.PacketType != "" || urlHealthCheckRequest.IpVersion != "" || urlHealthCheckRequest.QueryType != "" ||
			len(urlHealthCheckRequest.RequestHeaders) > 0 || portStr != "" || logs || networkConnectivity {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*urlHealthCheckRequest = UrlHealthCheckRequest{}
		ByteArrayToStruct(validator.jsonData, urlHealthCheckRequest)
		return
	}

	if len(args) != 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}

	validator.validateUrlArg(args[0], "url")
	validator.validateIpFlag(urlHealthCheckRequest.EdgeIp, "edgeServerIp", false)
	validator.ValidatePortFlag(portStr, &urlHealthCheckRequest.Port, "port", false)
	validator.ValidateIpVersionFlag(&urlHealthCheckRequest.IpVersion, "ipVersion", false)
	validator.ValidatePacketTypeFlag(&urlHealthCheckRequest.PacketType, "packetType", false)

	urlHealthCheckRequest.Url = args[0]

	if logs {
		urlHealthCheckRequest.ViewsAllowed = append(urlHealthCheckRequest.ViewsAllowed, "LOGS")
	}
	if networkConnectivity {
		urlHealthCheckRequest.ViewsAllowed = append(urlHealthCheckRequest.ViewsAllowed, "CONNECTIVITY")
	}

}

func (validator Validator) ValidateConnectivityProblemsFields(args []string, portStr string, connectivityProblemsRequest *ConnectivityProblemsRequest) {

	if validator.jsonData != nil {
		if len(args) > 0 || len(connectivityProblemsRequest.RequestHeaders) > 0 || connectivityProblemsRequest.SpoofEdgeIp != "" || portStr != "" ||
			connectivityProblemsRequest.ClientIp != "" || connectivityProblemsRequest.PacketType != "" || connectivityProblemsRequest.IpVersion != "" {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*connectivityProblemsRequest = ConnectivityProblemsRequest{}
		ByteArrayToStruct(validator.jsonData, &connectivityProblemsRequest)
		return
	}

	if len(args) != 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}

	validator.validateUrlArg(args[0], "url")

	connectivityProblemsRequest.Url = args[0]

	validator.validateIpFlag(connectivityProblemsRequest.SpoofEdgeIp, "edgeServerIp", false)
	validator.validateIpFlag(connectivityProblemsRequest.ClientIp, "clientIp", false)
	validator.ValidatePortFlag(portStr, &connectivityProblemsRequest.Port, "port", false)
	validator.ValidatePacketTypeFlag(&connectivityProblemsRequest.PacketType, "packetType", false)
	validator.ValidateIpVersionFlag(&connectivityProblemsRequest.IpVersion, "ipVersion", false)

}

func (validator Validator) ValidateGrepFields(args []string, grepRequest *GrepRequest, errorStatusCodeFlag, clientRequest, forwardRequest bool, httpStatusCodes []string) {

	if validator.jsonData != nil {
		if len(args) > 0 || len(grepRequest.CpCodes) > 0 || len(grepRequest.Hostnames) > 0 || len(grepRequest.ClientIps) > 0 || len(httpStatusCodes) > 0 || len(grepRequest.Arls) > 0 ||
			len(grepRequest.UserAgents) > 0 || errorStatusCodeFlag || forwardRequest || !clientRequest {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*grepRequest = GrepRequest{}
		ByteArrayToStruct(validator.jsonData, grepRequest)
		return
	}

	if len(args) != 3 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 3, len(args)), CmdErrExitCode)
	}

	if len(grepRequest.CpCodes) == 0 && len(grepRequest.Hostnames) == 0 {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "any"), CmdErrExitCode)
	} else if len(grepRequest.CpCodes) != 0 && len(grepRequest.Hostnames) != 0 {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "exclusive"), CmdErrExitCode)
	}

	if len(httpStatusCodes) != 0 && errorStatusCodeFlag {
		AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "exclusiveHttpStatusCode"), CmdErrExitCode)
	}

	validator.validateIpArg(args[0], "edgeServerIp")
	grepRequest.EdgeIp = args[0]
	grepRequest.Start = args[1]
	grepRequest.End = args[2]

	for _, clientIp := range grepRequest.ClientIps {
		if !isValidIp(clientIp) {
			AbortWithUsageAndMessage(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "clientIp"), CmdErrExitCode)
		}
	}

	if clientRequest && forwardRequest {
		grepRequest.LogType = "BOTH"
	} else if forwardRequest {
		grepRequest.LogType = "F"
	} else {
		grepRequest.LogType = "R"
	}

	if len(httpStatusCodes) > 0 {
		grepRequest.HttpStatusCodes = &HttpStatusCode{Comparison: "EQUALS", Value: httpStatusCodes}
	} else if errorStatusCodeFlag {
		grepRequest.HttpStatusCodes = &HttpStatusCode{Comparison: "NOT_EQUALS", Value: []string{"200"}}
	}

}

func (validator Validator) ValidateTranslateErrorFields(args []string, errorTranslatorRequest *ErrorTranslatorRequest) {

	if validator.jsonData != nil {
		if len(args) > 0 {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		ByteArrayToStruct(validator.jsonData, errorTranslatorRequest)
		return
	}

	if len(args) != 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}
	errorTranslatorRequest.ErrorCode = args[0]
}

func (validator Validator) ValidateContentProblemsFields(args []string, contentProblemsRequest *ContentProblemsRequest) {

	if validator.jsonData != nil {
		if len(args) > 0 || len(contentProblemsRequest.RequestHeaders) > 0 || contentProblemsRequest.EdgeIp != "" || contentProblemsRequest.IpVersion != "" {
			AbortWithUsageAndMessage(validator.cmd, GetGlobalErrorMessage(FieldsNotRequired), CmdErrExitCode)
		}
		*contentProblemsRequest = ContentProblemsRequest{}
		ByteArrayToStruct(validator.jsonData, &contentProblemsRequest)
		return
	}

	if len(args) != 1 {
		AbortWithUsageAndMessage(validator.cmd, fmt.Sprintf(GetGlobalErrorMessage(MissingArgs), 1, len(args)), CmdErrExitCode)
	}

	validator.validateUrlArg(args[0], "url")

	contentProblemsRequest.Url = args[0]

	validator.validateIpFlag(contentProblemsRequest.EdgeIp, "edgeServerIp", false)
	validator.ValidateIpVersionFlag(&contentProblemsRequest.IpVersion, "ipVersion", false)

}
