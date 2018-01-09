package main

//go:generate licrep -o licenses.go

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	optLicenseSummary := flag.Bool("summary", true, "Show packages and license types")
	optLicenses := flag.Bool("licenses", false, "Show actual license texts")

	licenses, err := GetLicenses()
	errors.WithMessage(err, "Getting licenses failed")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	flag.Parse()

	if *optLicenseSummary {
		fmt.Println("Licenses of dependencies:")
		for i := range licenses {
			fmt.Printf("%s: %s\n", i, licenses[i].Name)
		}
		fmt.Println("")
	}
	if *optLicenses {
		fmt.Println("License texts of depending packages:")
		for i := range licenses {
			fmt.Printf("* %s:\n\n%s\n\n", i, licenses[i].Text)
		}
	}
}
