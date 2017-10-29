package main

import (
	"fmt"
	"os"
	"strings"

	util "github.com/kopoli/go-util"
	licrep "github.com/kopoli/licrep/lib"
)

var (
	Version = "0.1.0"
)

func printErr(err error, message string, arg ...string) {
	msg := ""
	if err != nil {
		msg = fmt.Sprintf(" (error: %s)", err)
	}
	fmt.Fprintf(os.Stderr, "Error: %s%s.%s\n", message, strings.Join(arg, " "), msg)
}

func checkFault(err error, message string, arg ...string) {
	if err != nil {
		printErr(err, message, arg...)
		os.Exit(1)
	}
}

func main() {
	opts := util.NewOptions()

	opts.Set("program-name", "licrep")
	opts.Set("program-version", Version)

	err := licrep.Cli(opts, os.Args)
	checkFault(err, "command line parsing failed")

	cwd, _ := os.Getwd()
	licrep.GetPackages(".", cwd)

	os.Exit(0)
}
