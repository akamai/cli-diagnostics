package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/camelcase"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

func CamelToTitle(inp string) string {
	inp = strings.Title(inp)
	arr := camelcase.Split(inp)
	out := strings.Join(arr, " ")
	out = strings.Replace(out, "Url", "URL", -1)
	return out
}

func CapsToTitle(inp string) string {
	return strings.Title(strings.ToLower(inp))
}

func StartSpinner(message string) *spinner.Spinner {
	// making different variable name from import package name.
	s := spinner.New([]string{".     ", "..    ", "...   ", "....  ", "..... ", "......"}, 300*time.Millisecond)
	s.Writer = os.Stderr // Out-of-band info like progress info should go to stderr
	s.Prefix = message
	s.Start()
	return s
}

func StopSpinner(s *spinner.Spinner, success bool) {
	if success {
		s.FinalMSG = s.Prefix + "...... " + color.GreenString("[OK]")
	} else {
		s.FinalMSG = s.Prefix + "...... " + color.RedString("[FAIL]")
	}
	s.Stop()
	fmt.Fprintln(s.Writer) // Add a new line after the s
}

// Check contains and irrespective of case.
func ContainsIgnoreCase(a string, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

// Check contains and irrespective of case.
func ContainsInArray(array []string, inputString string) bool {
	var result = false
	for _, x := range array {
		if x == strings.ToLower(inputString) {
			result = true
			break
		}
	}

	return result
}

func ConvertBooleanToYesOrNo(input bool) string {
	if input {
		return "Yes"
	}
	return "No"
}

func FormatTime(inputTime string) string {
	layout := "2006-01-02T15:04:05+0000"
	myDate, err := time.Parse(layout, inputTime)
	if err != nil {
		fmt.Println(err)
	}
	// convert this date to desired format when decided.
	return myDate.Format("01/02/2006, 15:04 PM -07:00")
}

//Get all placeholders in string inside {{}}
func GetPlaceHoldersInString(errorMessage, regex string) []string {

	r := regexp.MustCompile(regex)
	matches := r.FindAllStringSubmatch(errorMessage, -1)
	var placeHolders = make([]string, len(matches))
	for i, v := range matches {
		placeHolders[i] = v[1]
	}

	return placeHolders
}

// 2021-10-14T16:14:12+0000 to Thursday, October 14, 2021
func isoToDate(isoDate string) string {
	isoDate = strings.Replace(isoDate, "+0000", "Z", 1)
	date, _ := time.Parse(time.RFC3339, isoDate)
	return date.Format("Monday, January 2, 2006")
}

// EdgeIpLocation to (city, regionCode, countryCode)(ASN asNumber)
func (loc EdgeIpLocation) toString() string {
	str := ""
	sep := ", "
	var arr []string

	if loc.City != "" {
		arr = append(arr, CapsToTitle(loc.City))
	}
	if loc.RegionCode != "" {
		arr = append(arr, loc.RegionCode)
	}
	if loc.CountryCode != "" {
		arr = append(arr, loc.CountryCode)
	}

	if len(arr) > 0 {
		str = "(" + strings.Join(arr, sep) + ")"
	}

	if loc.AsNum != nil {
		str += " (ASN " + strconv.Itoa(*loc.AsNum) + ")"
	}

	return str
}

func (loc EdgeIpLocation) toStringNoBrackets() string {
	str := ""
	sep := ", "
	var arr []string

	if loc.City != "" {
		arr = append(arr, CapsToTitle(loc.City))
	}
	if loc.RegionCode != "" {
		arr = append(arr, loc.RegionCode)
	}
	if loc.CountryCode != "" {
		arr = append(arr, loc.CountryCode)
	}

	if len(arr) > 0 {
		str = strings.Join(arr, sep)
	}

	if loc.AsNum != nil {
		str += " (ASN " + strconv.Itoa(*loc.AsNum) + ")"
	}

	return str
}

func ReadStdin() []byte {
	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		Abort("Error in stdin", CliErrExitCode)
	}
	size := fi.Size()
	if size > 0 {
		log.Debug("%v bytes available in Stdin\n", size)
		var reader = bufio.NewReader(os.Stdin)
		jsonData, err := ioutil.ReadAll(reader)
		if err != nil {
			Abort("Error in reading stdin", ParsingErrExitCode)
		}
		return jsonData
	}
	log.Debug("Stdin is empty")
	return nil
}

func ByteArrayToStruct(byt []byte, objPtr interface{}) {
	if err := json.Unmarshal(byt, objPtr); err != nil {
		log.Debug(err)
		Abort("Unable to parse the json", ParsingErrExitCode)
	}
}

func SplitCurlHeaderString(responseHeaderString string) string {
	splitArr := strings.SplitAfterN(responseHeaderString, curlHeaderSeparator, 2)
	value := ""
	if len(splitArr) == 2 {
		value = splitArr[1]
	}
	return bold(splitArr[0]) + value
}
