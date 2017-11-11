package main

import (
	"fmt"
	"os"

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

	err := licrep.Cli(opts, os.Args)
	checkFault(err, "command line parsing failed")

	cwd, _ := os.Getwd()
	pkg, err := licrep.GetPackages(".", cwd)
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
	}

	os.Exit(0)
}
