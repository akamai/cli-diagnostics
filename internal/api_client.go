package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type ApiClient struct {
	client EdgeGridHttpClient
}

func NewApiClient(client EdgeGridHttpClient) *ApiClient {
	return &ApiClient{client}
}

// Adds query parameters to url.
// QueryMap entries with empty values are ignored
func addQueryParams(url *url.URL, queryMap map[string]string) {
	queryParams := url.Query()

	log.Debug("Adding query parameters to url %s", url)
	for k, v := range queryMap {
		log.Tracef("Processing query parameter - [%s]:[%s]", k, v)

		if v != "" {
			queryParams.Set(k, v)
		}
	}

	url.RawQuery = queryParams.Encode()
	log.Tracef("Url with query parameters: %s", url.String())
}

func (api ApiClient) LocateIp(locateIpsRequest VerifyLocateIpsRequest) (*[]byte, *CliError) {

	locateIpsRequestBytes, err := json.Marshal(locateIpsRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/locate-ip"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &locateIpsRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting geolocation details.")
}

func (api ApiClient) VerifyIp(verifyIpsRequest VerifyLocateIpsRequest) (*[]byte, *CliError) {

	verifyIpsRequestBytes, err := json.Marshal(verifyIpsRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/verify-edge-ip"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &verifyIpsRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in verifying IP address.")
}

// Uncomment when verify-locate-ip is OPEN

func (api ApiClient) VerifyLocateIp(verifyLocateIpRequest VerifyLocateIpRequest) (*[]byte, *CliError) {

	locateIpRequestBytes, err := json.Marshal(verifyLocateIpRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}
	path := "/edge-diagnostics/v1/verify-locate-ip"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &locateIpRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in verifying and getting the geolocation details.")
}

func (api ApiClient) Dig(digRequest DigRequest) (*[]byte, *CliError) {

	digRequestBytes, err := json.Marshal(digRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}
	path := "/edge-diagnostics/v1/dig"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &digRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response\n", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error getting dig response.")
}

func (api ApiClient) UserDiagnosticsCreate(userDiagnosticsDataRequest UserDiagnosticsDataRequest) (*[]byte, *CliError) {
	userDiagnosticsDataRequestBytes, err := json.Marshal(userDiagnosticsDataRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/user-diagnostic-data/groups"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &userDiagnosticsDataRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusCreated {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in generating diagnostic link")

}

func (api ApiClient) Curl(curlRequest CurlRequest) (*[]byte, *CliError) {

	curlRequestBytes, err := json.Marshal(curlRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	// get response
	path := "/edge-diagnostics/v1/curl"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &curlRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting curl results.")

}

func (api ApiClient) Estats(estatsRequest EstatsRequest) (*[]byte, *CliError) {

	estatsRequestBytes, err := json.Marshal(estatsRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	// get response
	path := "/edge-diagnostics/v1/estats"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &estatsRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting estats results.")

}

func (api ApiClient) UserDiagnosticsList(urlFilter, user string, active bool) (*[]byte, *CliError) {

	path := "/edge-diagnostics/v1/user-diagnostic-data/groups"
	parsedUrl, _ := url.Parse(path)

	// add optional query parameters
	queryMap := map[string]string{
		"url":        urlFilter,
		"user":       user,
		"activeOnly": strconv.FormatBool(active),
	}
	addQueryParams(parsedUrl, queryMap)

	// get response
	resp, byt := api.client.request(Get, parsedUrl.String(), nil, nil)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting list of diagnostic links.")
}

func (api ApiClient) UserDiagnosticsGet(linkId string, mtr, dig, curl bool) (*[]byte, *CliError) {

	path := fmt.Sprintf("/edge-diagnostics/v1/user-diagnostic-data/groups/%s/records", linkId)
	parsedUrl, _ := url.Parse(path)

	// add optional query parameters
	queryMap := map[string]string{
		"includeMtr":  strconv.FormatBool(mtr),
		"includeDig":  strconv.FormatBool(dig),
		"includeCurl": strconv.FormatBool(curl),
	}
	addQueryParams(parsedUrl, queryMap)

	// get response
	resp, byt := api.client.request(Get, parsedUrl.String(), nil, nil)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting diagnostic link")
}

func (api ApiClient) TranslateUrl(arlRequest ArlRequest) (*[]byte, *CliError) {

	arlRequestBytes, err := json.Marshal(arlRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/translated-url"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &arlRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in translating url")
}

func (api ApiClient) TranslateErrorPost(errorTranslatorRequest ErrorTranslatorRequest) (*[]byte, *CliError) {

	errorTranslatorRequestBytes, err := json.Marshal(errorTranslatorRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/error-translator"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &errorTranslatorRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusAccepted {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting logs for error string.")
}

func (api ApiClient) TranslateErrorGet(link string) (*[]byte, *CliError) {

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, link, nil, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting logs for error string.")
}

func (api ApiClient) UrlHealthCheckPost(urlHealthCheckRequest UrlHealthCheckRequest) (*[]byte, *CliError) {

	urlHealthCheckRequestBytes, err := json.Marshal(urlHealthCheckRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/url-health-check"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &urlHealthCheckRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusAccepted {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting url health check.")
}

func (api ApiClient) UrlHealthCheckGet(link string) (*[]byte, *CliError) {

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, link+"?includeContentResponseBody=true", nil, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting url health check.")
}

func (api ApiClient) EdgeLocations() (*[]byte, *CliError) {
	path := "/edge-diagnostics/v1/edge-locations"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, path, nil, requestHeaders)

	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting edge locations")
}

func (api ApiClient) IpaHostnames() (*[]byte, *CliError) {
	path := "/edge-diagnostics/v1/ipa/hostnames"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, path, nil, requestHeaders)

	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting ipa hostnames")
}

func (api ApiClient) GtmHostnames() (*[]byte, *CliError) {
	path := "/edge-diagnostics/v1/gtm/gtm-properties"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, path, nil, requestHeaders)

	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting gtm hostnames")
}

func (api ApiClient) GtmTestTargetIp(property string, domain string) (*[]byte, *CliError) {
	path := "/edge-diagnostics/v1/gtm/" + property + "/" + domain + "/gtm-property-ips"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, path, nil, requestHeaders)

	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting gtm test target ip")
}

func (api ApiClient) Mtr(mtrRequest MtrRequest) (*[]byte, *CliError) {

	mtrRequestBytes, err := json.Marshal(mtrRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/mtr"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &mtrRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting mtr results.")
}

func (api ApiClient) GrepPost(grepRequest GrepRequest) (*[]byte, *CliError) {

	grepRequestBytes, err := json.Marshal(grepRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/grep"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &grepRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusAccepted {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting grep logs.")
}

func (api ApiClient) GrepGet(link string) (*[]byte, *CliError) {

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, link, nil, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in getting grep logs.")
}

func (api ApiClient) ConnectivityProblemsPost(connectivityProblemsRequest ConnectivityProblemsRequest) (*[]byte, *CliError) {

	connectivityProblemsRequestBytes, err := json.Marshal(connectivityProblemsRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/connectivity-problems"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &connectivityProblemsRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusAccepted {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in running connectivity problems.")
}

func (api ApiClient) ConnectivityProblemsGet(link string) (*[]byte, *CliError) {

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, link+"?includeContentResponseBody=true", nil, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in running connectivity problems.")
}

func (api ApiClient) ContentProblemsPost(contentProblemsRequest ContentProblemsRequest) (*[]byte, *CliError) {

	contentProblemsRequestBytes, err := json.Marshal(contentProblemsRequest)
	if err != nil {
		log.Error(err)
		Abort(GetGlobalErrorMessage(RequestParsingError), ParsingErrExitCode)
	}

	path := "/edge-diagnostics/v1/content-problems"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &contentProblemsRequestBytes, requestHeaders)
	if resp.StatusCode == http.StatusAccepted {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in running content problems.")
}

func (api ApiClient) ContentProblemsGet(link string) (*[]byte, *CliError) {

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Get, link+"?includeContentResponseBody=true", nil, requestHeaders)
	if resp.StatusCode == http.StatusOK {
		return byt, nil
	}
	log.Debug("api error response", string(*byt))
	return nil, CliErrorFromPulsarProblemObject(*byt, resp.StatusCode, "Error in running content problems.")
}
