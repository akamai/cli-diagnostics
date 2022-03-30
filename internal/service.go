package internal

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Service struct {
	api          ApiClient
	cmd          *cobra.Command
	outputAsJson bool
}

func NewService(api ApiClient, cmd *cobra.Command, outputAsJson bool) *Service {
	return &Service{api, cmd, outputAsJson}
}

func (svc Service) LocateIp(locateIpsRequest VerifyLocateIpsRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.LocateIp(locateIpsRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var locateIpsResponse VerifyLocateIpsResponse
	unmarshalErr := json.Unmarshal(*byt, &locateIpsResponse)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintLocateIpsResponse(locateIpsResponse, GetMessageForKey(svc.cmd, FailedIps))

}

func (svc Service) VerifyIp(verifyIpsRequest VerifyLocateIpsRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.VerifyIp(verifyIpsRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var verifyIpsResponse VerifyLocateIpsResponse
	unmarshalErr := json.Unmarshal(*byt, &verifyIpsResponse)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintVerifyIpsResponse(verifyIpsResponse, GetMessageForKey(svc.cmd, FailedIps))

}

func (svc Service) VerifyLocateIp(verifyLocateIpRequest VerifyLocateIpRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.VerifyLocateIp(verifyLocateIpRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var verifyLocateIpResponse VerifyLocateIpResponse
	unmarshalErr := json.Unmarshal(*byt, &verifyLocateIpResponse)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintVerifyIpResponse(verifyLocateIpResponse)
	PrintLocateIpResponse(verifyLocateIpResponse.Result.GeoLocation, verifyLocateIpResponse.Request.IpAddress)

}

func (svc Service) Dig(digRequest DigRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.Dig(digRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var digResponse DigResponse
	if json.Unmarshal(*byt, &digResponse) != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintDigResponse(digResponse)
}

func (svc Service) UserDiagnosticsCreate(userDiagnosticsDataRequest UserDiagnosticsDataRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.UserDiagnosticsCreate(userDiagnosticsDataRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var userDiagnosticsDataGroupDetails UserDiagnosticsDataGroupDetails
	unmarshalErr := json.Unmarshal(*byt, &userDiagnosticsDataGroupDetails)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintUserDiagnosticsDataGroupDetailsAfterCreate(userDiagnosticsDataGroupDetails)
}

func (svc Service) UserDiagnosticsList(url, user string, active bool) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.UserDiagnosticsList(url, user, active)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var userDiagnosticsDataGroupDetailsList ListResponse
	unmarshalErr := json.Unmarshal(*byt, &userDiagnosticsDataGroupDetailsList)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintUserDiagnosticsDataGroupDetailsTable(userDiagnosticsDataGroupDetailsList.Groups)
}

func (svc Service) UserDiagnosticsGet(linkId string, mtr, dig, curl bool) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.UserDiagnosticsGet(linkId, mtr, dig, curl)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var userDiagnosticsDataGroupDetails UserDiagnosticsDataGroupDetails
	unmarshalErr := json.Unmarshal(*byt, &userDiagnosticsDataGroupDetails)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintUserDiagnosticsDataGroupDetails(userDiagnosticsDataGroupDetails)
}

func (svc Service) TranslateUrl(arlRequest ArlRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.TranslateUrl(arlRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var arlContainer ArlContainer
	unmarshalErr := json.Unmarshal(*byt, &arlContainer)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}
	StopSpinner(spinner, true)
	PrintTranslateUrlResponse(arlContainer)
}

func (svc Service) TranslateError(errorTranslatorRequest ErrorTranslatorRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.TranslateErrorPost(errorTranslatorRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}
	var errorTranslatorResponse ErrorTranslatorResponse
	unmarshalErr := json.Unmarshal(*byt, &errorTranslatorResponse)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	var excecutionStatus string

	for {

		log.Debug("Waiting for ", errorTranslatorResponse.RetryAfter, " seconds\n")
		time.Sleep(time.Duration(errorTranslatorResponse.RetryAfter) * time.Second)

		log.Debug("Checking Execution status..")
		byt, err := svc.api.TranslateErrorGet(errorTranslatorResponse.Link)
		if err != nil {
			StopSpinner(spinner, false)
			Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
		}

		unmarshalErr := json.Unmarshal(*byt, &errorTranslatorResponse)

		log.Debug("status recieved: ", errorTranslatorResponse.ExecutionStatus)
		excecutionStatus = errorTranslatorResponse.ExecutionStatus

		if unmarshalErr != nil {
			StopSpinner(spinner, false)
			log.Error(err)
			Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
		}

		if excecutionStatus != "IN_PROGRESS" {
			if excecutionStatus == "FAILURE" {
				StopSpinner(spinner, false)
			} else {
				StopSpinner(spinner, true)
			}
			printJsonOutput(byt)
			return
		}
	}
}

func (svc Service) EdgeLocations(searchText string) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.EdgeLocations()

	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	var edgeLocationContainer EdgeLocationContainer
	unmarshalErr := json.Unmarshal(*byt, &edgeLocationContainer)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)

	if svc.outputAsJson {
		printEdgeLocationsJsonOutput(edgeLocationContainer, searchText)
		return
	}

	printEdgeLocations(edgeLocationContainer, searchText)

}

func (svc Service) IpaHostnames() {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.IpaHostnames()

	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var ipaHostnameResponse IpaHostnameResponse
	unmarshalErr := json.Unmarshal(*byt, &ipaHostnameResponse)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)

	PrintIpaHostnamesTable(ipaHostnameResponse)

}

func (svc Service) UrlHealthCheck(urlHealthCheckRequest UrlHealthCheckRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.UrlHealthCheckPost(urlHealthCheckRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}
	var urlHealthCheckResponse UrlHealthCheckResponse
	unmarshalErr := json.Unmarshal(*byt, &urlHealthCheckResponse)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	var excecutionStatus string

	for i := 0; i < 6; i++ {

		log.Debug("Waiting for ", urlHealthCheckResponse.RetryAfter, " seconds\n")
		time.Sleep(time.Duration(urlHealthCheckResponse.RetryAfter) * time.Second)

		log.Debug("Checking Execution status..")
		byt, err := svc.api.UrlHealthCheckGet(urlHealthCheckResponse.Link)
		if err != nil {
			StopSpinner(spinner, false)
			Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
		}

		unmarshalErr := json.Unmarshal(*byt, &urlHealthCheckResponse)

		log.Debug("status recieved: ", urlHealthCheckResponse.ExecutionStatus)
		excecutionStatus = urlHealthCheckResponse.ExecutionStatus

		if unmarshalErr != nil {
			StopSpinner(spinner, false)
			log.Error(unmarshalErr)
			Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
		}

		if excecutionStatus != "IN_PROGRESS" {
			if excecutionStatus == "FAILURE" {
				StopSpinner(spinner, false)
			} else {
				StopSpinner(spinner, true)
			}
			printJsonOutput(byt)
			return
		}
	}
	StopSpinner(spinner, false)
	Abort(GetGlobalErrorMessage("500"), 500-HttpCodeExitCodeDiff)
}

func (svc Service) GtmHostnames() {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.GtmHostnames()

	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var gtmPropertyContainer GtmPropertyContainer

	unmarshalErr := json.Unmarshal(*byt, &gtmPropertyContainer)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)

	PrintGtmHostnamesTable(gtmPropertyContainer)

}

func (svc Service) GtmTestTargetIp(hostname string) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, "spinnerMessageTestTargetIp"))

	byt, err := svc.api.GtmHostnames()

	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	var gtmPropertyContainer GtmPropertyContainer

	unmarshalErr := json.Unmarshal(*byt, &gtmPropertyContainer)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	log.Debug("Got GtmHostnames response")

	for _, gtmProperty := range gtmPropertyContainer.GtmProperties {
		if gtmProperty.Property+"."+gtmProperty.Domain == hostname {
			byt, err := svc.api.GtmTestTargetIp(gtmProperty.Property, gtmProperty.Domain)

			if err != nil {
				StopSpinner(spinner, false)
				Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
			}

			if svc.outputAsJson {
				StopSpinner(spinner, true)
				printJsonOutput(byt)
				return
			}

			var gtmPropertyIpsContainer GtmPropertyIpsContainer

			unmarshalErr := json.Unmarshal(*byt, &gtmPropertyIpsContainer)
			if unmarshalErr != nil {
				StopSpinner(spinner, false)
				log.Error(err)
				Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
			}

			StopSpinner(spinner, true)

			PrintGtmTestTargetIpTable(gtmPropertyIpsContainer)
			return
		}
	}

	StopSpinner(spinner, true)
	PrintWarning("\n" + hostname + " is not a valid gtm hostname \n")

}

func (svc Service) Curl(curlRequest CurlRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.Curl(curlRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var curlResponse CurlResponse
	unmarshalErr := json.Unmarshal(*byt, &curlResponse)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintCurlResponse(curlResponse)
}

func (svc Service) Estats(estatsRequest EstatsRequest, logsEstatsFlag bool) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.Estats(estatsRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	var estatsResult EstatsResult
	unmarshalErr := json.Unmarshal(*byt, &estatsResult)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	var estatsResultWrapper EstatsResultWrapper
	estatsResultWrapper.EstatsResult = estatsResult

	var grepResponses []EstatsGrepResponse

	if logsEstatsFlag {
		if len(estatsResult.Result.TopEdgeIpsWithError) > 0 {
			var edgeIpInfo = estatsResult.Result.TopEdgeIpsWithError[0]
			byt1, err1 := svc.api.GrepGet(edgeIpInfo.EdgeLogsLink)

			if err1 != nil {
				StopSpinner(spinner, false)
				Abort(err1.errorMessage, err1.responseCode-HttpCodeExitCodeDiff)
			}
			var estatsGrepResponse1 EstatsGrepResponse
			unmarshalErr := json.Unmarshal(*byt1, &estatsGrepResponse1)

			if unmarshalErr != nil {
				StopSpinner(spinner, false)
				log.Error(unmarshalErr)
				Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
			}
			grepResponses = append(grepResponses, estatsGrepResponse1)
		}
		if len(estatsResult.Result.TopEdgeIpsWithErrorFromOrigin) > 0 {
			var edgeIpInfo = estatsResult.Result.TopEdgeIpsWithErrorFromOrigin[0]
			byt1, err1 := svc.api.GrepGet(edgeIpInfo.EdgeLogsLink)

			if err1 != nil {
				StopSpinner(spinner, false)
				Abort(err1.errorMessage, err1.responseCode-HttpCodeExitCodeDiff)
			}
			var estatsGrepResponse1 EstatsGrepResponse
			unmarshalErr := json.Unmarshal(*byt1, &estatsGrepResponse1)

			if unmarshalErr != nil {
				StopSpinner(spinner, false)
				log.Error(unmarshalErr)
				Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
			}
			grepResponses = append(grepResponses, estatsGrepResponse1)
		}
		estatsResultWrapper.EstatsLogLines = grepResponses
	}

	StopSpinner(spinner, true)
	newFsConfigBytes, _ := json.Marshal(estatsResultWrapper)
	printJsonOutput(&newFsConfigBytes)

}

func (svc Service) Mtr(mtrRequest MtrRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.Mtr(mtrRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}

	if svc.outputAsJson {
		StopSpinner(spinner, true)
		printJsonOutput(byt)
		return
	}

	var mtrResponse MtrResponse
	unmarshalErr := json.Unmarshal(*byt, &mtrResponse)
	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	StopSpinner(spinner, true)
	PrintMtrResponse(mtrRequest, mtrResponse)
}

func (svc Service) Grep(grepRequest GrepRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.GrepPost(grepRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}
	var grepResponse GrepResponse
	unmarshalErr := json.Unmarshal(*byt, &grepResponse)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	var excecutionStatus string

	for i := 0; i < 40; i++ {

		log.Debug("Waiting for ", grepResponse.RetryAfter, " seconds\n")
		time.Sleep(time.Duration(grepResponse.RetryAfter) * time.Second)

		log.Debug("Checking Execution status..")
		byt, err := svc.api.GrepGet(grepResponse.Link)
		if err != nil {
			StopSpinner(spinner, false)
			Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
		}

		unmarshalErr := json.Unmarshal(*byt, &grepResponse)

		log.Debug("status recieved: ", grepResponse.ExecutionStatus)
		excecutionStatus = grepResponse.ExecutionStatus

		if unmarshalErr != nil {
			StopSpinner(spinner, false)
			log.Error(err)
			Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
		}

		if excecutionStatus != "IN_PROGRESS" {
			if excecutionStatus == "FAILURE" {
				StopSpinner(spinner, false)
			} else {
				StopSpinner(spinner, true)
			}
			printJsonOutput(byt)
			return
		}
	}

	StopSpinner(spinner, false)
	Abort(GetGlobalErrorMessage("500"), 500-HttpCodeExitCodeDiff)
}

func (svc Service) ConnectivityProblems(connectivityProblemsRequest ConnectivityProblemsRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.ConnectivityProblemsPost(connectivityProblemsRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}
	var connectivityProblemsResponse ConnectivityProblemsResponse
	unmarshalErr := json.Unmarshal(*byt, &connectivityProblemsResponse)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(err)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	var excecutionStatus string

	for i := 0; i < 6; i++ {

		log.Debug("Waiting for ", connectivityProblemsResponse.RetryAfter, " seconds\n")
		time.Sleep(time.Duration(connectivityProblemsResponse.RetryAfter) * time.Second)

		log.Debug("Checking Execution status...")
		byt, err := svc.api.ConnectivityProblemsGet(connectivityProblemsResponse.Link)
		if err != nil {
			StopSpinner(spinner, false)
			Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
		}

		unmarshalErr := json.Unmarshal(*byt, &connectivityProblemsResponse)

		log.Debug("status recieved: ", connectivityProblemsResponse.ExecutionStatus)
		excecutionStatus = connectivityProblemsResponse.ExecutionStatus

		if unmarshalErr != nil {
			StopSpinner(spinner, false)
			log.Error(err)
			Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
		}

		if excecutionStatus != "IN_PROGRESS" {
			if excecutionStatus == "FAILURE" {
				StopSpinner(spinner, false)
			} else {
				StopSpinner(spinner, true)
			}
			printJsonOutput(byt)
			return
		}
	}

	StopSpinner(spinner, false)
	Abort(GetGlobalErrorMessage("500"), 500-HttpCodeExitCodeDiff)
}

func (svc Service) ContentProblems(contentProblemsRequest ContentProblemsRequest) {

	spinner := StartSpinner(GetMessageForKey(svc.cmd, SpinnerMessage))

	byt, err := svc.api.ContentProblemsPost(contentProblemsRequest)
	if err != nil {
		StopSpinner(spinner, false)
		Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
	}
	var contentProblemsResponse ContentProblemsResponse
	unmarshalErr := json.Unmarshal(*byt, &contentProblemsResponse)

	if unmarshalErr != nil {
		StopSpinner(spinner, false)
		log.Error(unmarshalErr)
		Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
	}

	var excecutionStatus string

	for i := 0; i < 6; i++ {

		log.Debug("Waiting for ", contentProblemsResponse.RetryAfter, " seconds\n")
		time.Sleep(time.Duration(contentProblemsResponse.RetryAfter) * time.Second)

		log.Debug("Checking Execution status..")
		byt, err := svc.api.ContentProblemsGet(contentProblemsResponse.Link)
		if err != nil {
			StopSpinner(spinner, false)
			Abort(err.errorMessage, err.responseCode-HttpCodeExitCodeDiff)
		}

		unmarshalErr := json.Unmarshal(*byt, &contentProblemsResponse)

		log.Debug("status recieved: ", contentProblemsResponse.ExecutionStatus)
		excecutionStatus = contentProblemsResponse.ExecutionStatus

		if unmarshalErr != nil {
			StopSpinner(spinner, false)
			log.Error(unmarshalErr)
			Abort(GetGlobalErrorMessage(ResponseParsingError), ParsingErrExitCode)
		}

		if excecutionStatus != "IN_PROGRESS" {
			if excecutionStatus == "FAILURE" {
				StopSpinner(spinner, false)
			} else {
				StopSpinner(spinner, true)
			}
			printJsonOutput(byt)
			return
		}
	}
	StopSpinner(spinner, false)
	Abort(GetGlobalErrorMessage("500"), 500-HttpCodeExitCodeDiff)
}
