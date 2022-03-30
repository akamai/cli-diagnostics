package main

import (
	"github.com/akamai/cli-diagnostics/cmd"
	"github.com/akamai/cli-diagnostics/internal"
)

var (
	VERSION string = "1.0.0"
)

func main() {
	internal.InitLoggingConfig()
	cmd.Execute(VERSION)
}