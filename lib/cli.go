package licrep

import (
	"strings"

	cli "github.com/jawher/mow.cli"
	"github.com/kopoli/appkit"
)

// Cli parses the command line interface
func Cli(opts appkit.Options, argsin []string) error {
	progName := opts.Get("program-name", "licrep")

	app := cli.App(progName, "Generate an embedded license report to your program.")

	app.Spec = "[OPTIONS]"

	optDir := app.StringOpt("d directory", ".", "Directory of the package the check.")
	optOutput := app.StringOpt("o out", "", "The generated licenses file. Stdout if empty.")
	optPackage := app.StringOpt("p package", "", "The package to set for the generated file. Autodetected.")
	optPrefix := app.StringOpt("prefix", "", "Prefix given to the functions in the generated file.")
	optIgnore := app.StringsOpt("i ignore", []string{"main"}, "Ignore packages that match given regexps from the list.")

	optShowSummary := app.BoolOpt("s show-summary", false, "Show a license summary (no generating).")
	optShowLicenses := app.BoolOpt("L show-licenses", false, "Show the contents of all licenses (no generating).")

	optShowLicrepLicenses := app.BoolOpt("licrep-licenses", false, "Show licenses of licrep.")

	app.Version("v version", appkit.VersionString(opts))

	app.Action = func() {
		opts.Set("directory", *optDir)
		opts.Set("output", *optOutput)
		opts.Set("package", *optPackage)
		opts.Set("prefix", *optPrefix)
		opts.Set("ignore-packages", strings.Join(*optIgnore, " "))

		if *optShowSummary {
			opts.Set("show-license", "summary")
		}
		if *optShowLicenses {
			opts.Set("show-license", "full")
		}
		if *optShowLicrepLicenses {
			opts.Set("show-self-licenses", "t")
		}
	}

	return app.Run(argsin)
}
