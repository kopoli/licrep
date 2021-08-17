package main

//go:generate licrep -o licenses.go --prefix "Licrep"

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/kopoli/appkit"
	licrep "github.com/kopoli/licrep/lib"
)

var (
	version     = "Undefined"
	timestamp   = "Undefined"
	buildGOOS   = "Undefined"
	buildGOARCH = "Undefined"
	progVersion = "" + version
)

func fault(err error, message string) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s: %s\n",
			message, err)
		os.Exit(1)
	}
}

func main() {
	opts := appkit.NewOptions()

	opts.Set("program-name", os.Args[0])
	opts.Set("program-args", strings.Join(os.Args, " "))
	opts.Set("program-version", progVersion)
	opts.Set("program-timestamp", timestamp)
	opts.Set("program-buildgoos", buildGOOS)
	opts.Set("program-buildgoarch", buildGOARCH)

	err := licrep.Cli(opts, os.Args)
	fault(err, "command line parsing failed")

	dir, err := filepath.Abs(opts.Get("directory", "."))
	fault(err, "Given directory not valid")

	// Show the licenses of this program
	if opts.IsSet("show-self-licenses") {
		var licenses map[string]LicrepLicense
		licenses, err = LicrepGetLicenses()
		fault(err, "Internal error: License decoding failed")

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
	fault(err, "Getting package licenses failed")

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
		fault(err, "Generating embedded licenses failed")
	}

	os.Exit(0)
}
