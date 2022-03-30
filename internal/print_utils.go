package internal

//This class will have only akamai/global cli standard print functions
import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Start Standard Print Functions
func PrintError(message string, args ...interface{}) {
	c := color.New(color.FgRed)
	c.Fprintf(os.Stderr, message, args...)
}

func PrintWarning(message string, args ...interface{}) {
	c := color.New(color.FgCyan)
	c.Fprintf(os.Stderr, message, args...)
}

func PrintHeader(message string, args ...interface{}) {
	c := color.New(color.FgYellow).Add(color.Bold)
	c.Printf(message, args...)
}

func PrintSuccess(message string, args ...interface{}) {
	c := color.New(color.FgGreen)
	c.Printf(message, args...)
}

func success(a ...interface{}) string {
	c := color.New(color.FgGreen)
	return c.Sprint(a...)
}

func blue(a ...interface{}) string {
	c := color.New(color.FgBlue)
	return c.Sprint(a...)
}

func bold(a ...interface{}) string {
	c := color.New(color.Bold)
	return c.Sprint(a...)
}

func italic(a ...interface{}) string {
	c := color.New(color.Italic)
	return c.Sprint(a...)
}

func header(a ...interface{}) string {
	c := color.New(color.FgYellow).Add(color.Bold)
	return c.Sprint(a...)
}

func printLabelAndValue(label string, value interface{}) {
	c := color.New(color.Bold)
	c.Printf(label + ": ")
	fmt.Printf("%v\n", value)
}

func Abort(message string, errorCode int) {
	PrintError(message + "\n")
	os.Exit(errorCode)
}

func AbortWithUsageAndMessage(cmd *cobra.Command, message string, errorCode int) {
	PrintError(message + "\n\n")
	err := cmd.Usage()
	if err != nil {
		return
	}
	os.Exit(errorCode)
}
func AbortForCommand(cmd *cobra.Command, cliError *CliError) {
	AbortForCommandWithSubResource(cmd, cliError, Empty, Empty)
}

func AbortForCommandWithSubResource(cmd *cobra.Command, cliError *CliError, subResource, operation string) {

	responseCode := strconv.Itoa(cliError.responseCode)
	if len(responseCode) == 3 && strings.Contains("500,502,503,504", responseCode) {
		PrintError(GetGlobalErrorMessage(responseCode) + "\n\n")
	} else if cliError.apiError != nil {
		printErrorMessages(GetApiErrorMessagesForCommand(cmd, *cliError.apiError, subResource, operation))
		println()
	} else if len(cliError.apiSubErrors) != 0 {
		printErrorMessages(GetApiSubErrorMessagesForCommand(cmd, cliError.apiSubErrors, "", subResource, operation))
		println()
	} else {
		PrintError(cliError.errorMessage + "\n\n")
	}

	os.Exit(1)
}

func printErrorMessages(errorMessages []string) {
	for _, message := range errorMessages {
		PrintError(message + "\n")
	}
}

// ShowTable Standard function to show table in same format across Test Center CLI
func ShowTable(tableHeaders []string, tableContents [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	table.SetHeader(tableHeaders)
	table.SetAutoFormatHeaders(false)
	table.AppendBulk(tableContents)
	table.Render()
	fmt.Printf("\nTotal items: %d\n", len(tableContents))
}

func printTemplate(funcMap template.FuncMap, templateToParse string, data interface{}) {

	tmp, err := template.New("template").Funcs(funcMap).Parse(templateToParse)

	// if there is error,
	if err != nil {
		log.Error(err)
		PrintError("Output can not shown because of internal cli error.\n")
		os.Exit(1)
	} else {
		// standard output to print merged data
		err = tmp.Execute(os.Stdout, data)
		println()
		if err != nil {
			log.Error(err)
			PrintError("Output can not shown because of internal cli error.\n")
			os.Exit(1)
		}
	}
}
