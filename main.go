package main

//go:generate licrep -o licenses.go --prefix "Licrep"

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	util "github.com/kopoli/go-util"
	licrep "github.com/kopoli/licrep/lib"
)

var (
	// Version represents the program version
	Version = "0.2.1"
)

func checkFault(err error, arg ...interface{}) {
	if err != nil {
		util.E.Print(err, arg...)
		os.Exit(1)
	}
}

func main() {
	opts := util.NewOptions()

	opts.Set("program-name", "licrep")
	opts.Set("program-version", Version)
	opts.Set("program-args", strings.Join(os.Args, " "))

	err := licrep.Cli(opts, os.Args)
	checkFault(err, "command line parsing failed")

	dir, err := filepath.Abs(opts.Get("directory", "."))
	checkFault(err, "Given directory not valid")

	// Show the licenses of this program
	if opts.IsSet("show-self-licenses") {
		var licenses map[string]LicrepLicense
		licenses, err = LicrepGetLicenses()
		checkFault(err, "Internal error: License decoding failed")

		var names []string
		for i := range licenses {
			names = append(names, i)
		}
		sort.Strings(names)
		for _, i := range names {
			fmt.Printf("* %s:\n\n%s\n\n", i, licenses[i].Text)
		}
		os.Exit(0)
	}

	pkg, err := licrep.GetPackages(".", dir)
	checkFault(err, "Getting package licenses failed")

	pkg = licrep.FilterPackages(opts, pkg)

	if opts.IsSet("show-license") {
		show := opts.Get("show-license", "")
		wr := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
		for i := range pkg {
			switch show {
			case "summary":
				fmt.Fprintf(wr, "%s\t%s\n", pkg[i].ImportPath, pkg[i].License)
			case "full":
				fmt.Printf("* %s  %s\n%s\n", pkg[i].ImportPath,
					pkg[i].License, pkg[i].LicenseString)
			}
		}
		_ = wr.Flush()
	} else {
		err = licrep.GenerateEmbeddedLicenses(opts, pkg)
		checkFault(err, "Generating embedded licenses failed")
	}

	os.Exit(0)
}
