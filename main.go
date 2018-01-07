package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	util "github.com/kopoli/go-util"
	licrep "github.com/kopoli/licrep/lib"
)

var (
	Version = "0.1.0"
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

	pkg, err := licrep.GetPackages(".", dir)
	checkFault(err, "Getting package licenses failed")

	if opts.IsSet("show-license") {
		show := opts.Get("show-license", "")
		for i := range pkg {
			switch show {
			case "summary":
				fmt.Println(pkg[i].Name, "  ", pkg[i].License)
			case "full":
				fmt.Println("* ", pkg[i].Name, "  ", pkg[i].License,
					"\n", pkg[i].LicenseString, "\n")
			}
		}
	} else {
		err = licrep.GenerateEmbeddedLicenses(opts, pkg)
		checkFault(err,"Generating embedded licenses failed")
	}

	os.Exit(0)
}
