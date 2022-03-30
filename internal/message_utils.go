package internal

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/oleiade/reflections"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var (
	//go:embed en_US.json
	messageJsonBytes []byte

	messageJson       gjson.Result
	rootCommandName   = "akamai"
	jsonPathSeparator = "."
	flagKey           = "flag"
	argKey            = "arg"
	placeHolderRegex  = "{{(.*?)}}"
	fallbackKey       = "fallback"
)

//Get global errors for given key
func GetGlobalErrorMessage(key string) string {
	return getMessageForJsonPathOrFallback(strings.Join([]string{rootCommandName, Global, key}, jsonPathSeparator))
}

// Return message for given key under command.
func GetMessageForKey(baseCmdPath *cobra.Command, key string) string {
	jsonPath := getJsonPathForCommand(strings.Join([]string{baseCmdPath.CommandPath(), key}, " "))
	return getMessageForJsonPathOrFallback(jsonPath)
}

// Get different type of message for flag
func GetErrorMessageForFlag(cmd *cobra.Command, errorType, flagKeyInJson string) string {

	jsonPath := getJsonPathForCommand(cmd.CommandPath())
	jsonPath = strings.Join([]string{jsonPath, flagKey, errorType, flagKeyInJson}, jsonPathSeparator)

	log.Debugf("Get message for json path [%s], error type - [%s], flag key in json - [%s]", jsonPath, errorType, flagKeyInJson)
	return getMessageForJsonPathOrFallback(jsonPath)
}

// Get different type of message for positional args
func GetErrorMessageForArg(cmd *cobra.Command, errorType, argKeyInJson string) string {

	jsonPath := getJsonPathForCommand(cmd.CommandPath())
	switch errorType {
	case Missing:
		jsonPath = strings.Join([]string{jsonPath, argKey, Missing, argKeyInJson}, jsonPathSeparator)
	case Invalid:
		jsonPath = strings.Join([]string{jsonPath, argKey, Invalid, argKeyInJson}, jsonPathSeparator)
	}

	log.Debugf("Get message for json path [%s], error type - [%s], flag key in json - [%s]", jsonPath, errorType, argKeyInJson)
	return getMessageForJsonPathOrFallback(jsonPath)
}

func GetApiErrorMessagesForCommand(cmd *cobra.Command, apiError ApiError, subResource, operation string) []string {
	jsonPathForCommand := getJsonPathForCommand(cmd.CommandPath())
	parentErrorKey := getErrorJsonKeyForErrorType(apiError.Type)

	if len(apiError.Errors) != 0 {
		return GetApiSubErrorMessagesForCommand(cmd, apiError.Errors, parentErrorKey, subResource, operation)
	} else {
		errorPath := getJsonPathForCommand(strings.Join([]string{jsonPathForCommand, subResource, operation, parentErrorKey}, " "))
		errorMessage := getMessageForJsonPathOrFallback(errorPath)
		return []string{getReplacedPlaceholderMessage(apiError, errorMessage)}
	}
}

// Get All the error messages for api sub errors
func GetApiSubErrorMessagesForCommand(cmd *cobra.Command, apiSubError []ApiSubError, parentErrorKey, subResource, operation string) []string {
	jsonPathForCommand := getJsonPathForCommand(cmd.CommandPath())
	var errorMessages = make([]string, len(apiSubError))

	for i, subError := range apiSubError {
		subErrorKey := getErrorJsonKeyForErrorType(subError.Type)
		subErrorRequestField := getErrorJsonKeyForErrorType(subError.RequestField)
		log.Debug("subErrorKey", subErrorKey)
		errorPath := getJsonPathForCommand(strings.Join([]string{jsonPathForCommand, subResource, operation, parentErrorKey, subErrorKey, subErrorRequestField}, " "))

		// /*Custom logic starts here*/

		// //Pulsar object sometimes contains same error type for different objects, currently not able to figure out how to show those messages differently for different objects
		// //One other possible solution is show generic message
		// // For now this is done to only support submit test run
		// // First custom logic
		// if strings.Contains("resourceNotFound,resourceInDeletedState", subErrorKey) && checkIfMessageExist(errorPath+jsonPathSeparator+subError.RequestField+strings.Title(subErrorKey)) {

		// 	errorMessage := getMessageForJsonPathOrFallback(errorPath + jsonPathSeparator + subError.RequestField + strings.Title(subErrorKey))
		// 	// Replace placeholder values in string from json if there are any
		// 	errorMessages[i] = getReplacedPlaceholderMessage(subError, errorMessage)
		// 	continue
		// }

		errorMessage := getMessageForJsonPathOrFallback(errorPath)
		// Replace placeholder values in string from json if there are any
		errorMessages[i] = getReplacedPlaceholderMessage(subError, errorMessage)
	}
	return errorMessages
}

func getErrorJsonKeyForErrorType(errorType string) string {

	str := strings.Split(errorType, jsonPathSeparator)
	var jsonPath = make([]string, len(str))
	for i2, s := range str {
		if i2 == 0 {
			jsonPath[i2] = s
		} else {
			jsonPath[i2] = strings.Title(s)
		}
	}
	return strings.Join(jsonPath, "")
}

// Replace placeholder values in string from json if there are any
func getReplacedPlaceholderMessage(error interface{}, errorMessage string) string {

	for _, str := range GetPlaceHoldersInString(errorMessage, placeHolderRegex) {
		value, _ := reflections.GetField(error, strings.Title(str))
		errorMessage = strings.ReplaceAll(errorMessage, fmt.Sprintf("{{%s}}", str), fmt.Sprintf("%v", value))
	}
	return errorMessage
}

// Return json path for given command chain, e.g. -  `test-center test-suite view` converted to akamai.testCenter.testSuite.view
func getJsonPathForCommand(cmdString string) string {
	log.Debugf("Get json path for command [%s]", cmdString)
	givenString := strings.Fields(cmdString)
	var jsonPath = make([]string, len(givenString))

	for i, str := range givenString {
		var dashRemovedString = make([]string, len(str))
		for i2, dashedString := range strings.Split(str, "-") {
			if i2 == 0 {
				dashRemovedString[i2] = dashedString
			} else {
				dashRemovedString[i2] = strings.Title(dashedString)
			}
		}
		jsonPath[i] = strings.Join(dashRemovedString, "")
	}

	convertedString := strings.Join(jsonPath, jsonPathSeparator)
	if strings.Contains(convertedString, rootCommandName) {
		return convertedString
	}

	return strings.Join([]string{rootCommandName, convertedString}, jsonPathSeparator)
}

// standard function to get message from json for given json path
func checkIfMessageExist(jsonPath string) bool {
	message := gjson.Get(messageJson.String(), jsonPath)
	log.Debugf("Message for json path [%s] : [%s]", jsonPath, message.String())
	return message.Exists()
}

// standard function to get message from json for given json path
func getMessageForJsonPathOrFallback(jsonPath string) string {
	message := gjson.Get(messageJson.String(), jsonPath)
	if message.Exists() && message.Type == gjson.String {
		log.Debugf("Message for json path [%s] : [%s]", jsonPath, message.String())
		return message.String()
	} else {
		log.Infof("Message for json path [%s] : [%s]", jsonPath, message.String())
		log.Debugf("Message is not configured for jsonPath [%s]", jsonPath)
		return gjson.Get(messageJson.String(), rootCommandName+jsonPathSeparator+fallbackKey).String()
	}
}

// get fallback message
func GetGlobalFallBackMessage() string {
	return gjson.Get(messageJson.String(), rootCommandName+jsonPathSeparator+fallbackKey).String()
}

// Initialize message file
func init() {
	messageJson = gjson.ParseBytes(messageJsonBytes)
}
