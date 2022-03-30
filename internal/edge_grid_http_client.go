package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	log "github.com/sirupsen/logrus"
)

type EdgeGridHttpClient struct {
	config           edgegrid.Config
	accountSwitchKey string
}

func NewEdgeGridHttpClient(filepath, section, accountSwitchKey string) *EdgeGridHttpClient {
	//checks for ENV variables, if not present looks for .edgerc file
	config, err := edgegrid.Init(filepath, section)
	if err != nil {
		Abort(GetGlobalErrorMessage("initEdgeRc"), CliErrExitCode)
	}
	return &EdgeGridHttpClient{config, accountSwitchKey}
}

func (h EdgeGridHttpClient) request(method string, path string, payload *[]byte, headers http.Header) (*http.Response, *[]byte) {
	var (
		err    error
		req    *http.Request
		client = http.Client{}
	)

	var protocol = "https://"
	if strings.Contains(h.config.Host, "http") {
		protocol = "" // For mocking API calls locally
	}

	parsedPath, _ := url.Parse(path)
	if h.accountSwitchKey != "" {
		log.Debugf("Account switch key present :: %s. Adding to URL.", h.accountSwitchKey)
		query := parsedPath.Query()
		query.Set("accountSwitchKey", h.accountSwitchKey)
		parsedPath.RawQuery = query.Encode()
	}

	log.Debugf("Sending request:: %s %s, Headers: %v, Body: %s\n", method, parsedPath, headers, payload)

	if payload != nil {
		req, err = http.NewRequest(method, protocol+h.config.Host+parsedPath.String(), bytes.NewBuffer(*payload))
		if err != nil {
			Abort(err.Error(), ParsingErrExitCode)
		}
	} else {
		req, err = http.NewRequest(method, protocol+h.config.Host+parsedPath.String(), nil)
		if err != nil {
			Abort(err.Error(), ParsingErrExitCode)
		}
	}

	if headers != nil {
		req.Header = headers
	} else {
		req.Header = make(http.Header)
	}
	req.Header.Add(X_ED_CLIENT_TYPE, "CLI")

	req = edgegrid.AddRequestHeader(h.config, req)

	resp, er := client.Do(req)
	if er != nil {
		Abort(err.Error(), CliErrExitCode)
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Abort(err.Error(), ParsingErrExitCode)
	}
	log.Debugf("Received response:: Status: %d\n", resp.StatusCode)
	log.Tracef("Response body: %s\n", byt)

	return resp, &byt
}
