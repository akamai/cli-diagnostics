package internal

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// ApiError object for API error payloads
type ApiError struct {
	Type     string        `json:"type"`
	Title    string        `json:"title"`
	Status   int           `json:"status"`
	Instance string        `json:"instance"`
	Detail   string        `json:"detail"`
	Errors   []ApiSubError `json:"errors"`
}

// ApiSubError object represents sub-errors of an error payload or error response in 207
type ApiSubError struct {
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	RequestField  string        `json:"requestField"`
	RequestValues []interface{} `json:"requestValues"`
	Detail        string        `json:"detail"`
}

// CliError is used to transmit errors across the app
type CliError struct {
	apiError     *ApiError
	apiSubErrors []ApiSubError
	errorMessage string
	responseCode int
}

func CliErrorWithMessage(message string) *CliError {
	return &CliError{errorMessage: message}
}

func CliErrorFromPulsarProblemObject(apiErrorByte []byte, responseCode int, fallbackMessage string) *CliError {
	var cliError CliError
	cliError.responseCode = responseCode

	// set error message as fallback message by default
	cliError.errorMessage = fallbackMessage

	var apiError ApiError
	apiParsingError := json.Unmarshal(apiErrorByte, &apiError)
	if apiParsingError != nil {
		log.Debugf("Failed to parse api error response: [%s]", apiParsingError)
		if GetGlobalErrorMessage(fmt.Sprint(responseCode)) != GetGlobalFallBackMessage() {
			cliError.errorMessage = GetGlobalErrorMessage(fmt.Sprint(responseCode))
		}
	} else {
		log.Debug("Error response parsed")
		cliError.apiError = &apiError
		if apiError.Detail != "" {
			log.Debug("error message set from detail field")
			cliError.errorMessage = apiError.Detail
		} else if apiError.Errors != nil && apiError.Errors[0].Detail != "" {
			log.Debug("error message set from nested detail field")
			cliError.errorMessage = apiError.Errors[0].Detail
		}
	}

	return &cliError
}
