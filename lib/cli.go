package licrep

import (
	"strings"

	cli "github.com/jawher/mow.cli"
	util "github.com/kopoli/go-util"
)

func Cli(opts util.Options, argsin []string) error {
	progName := opts.Get("program-name", "licrep")

	app := cli.App(progName, "Generate an embedded license report to your program")

	app.Spec = "[OPTIONS]"

	optDir := app.StringOpt("d directory", ".", "Directory of the package the check.")
	optOutput := app.StringOpt("o out", "", "The generated licenses file. Stdout if empty.")
	optPackage := app.StringOpt("p package", "", "The package to set for the generated file. Autodetected.")
	optPrefix := app.StringOpt("prefix", "", "Prefix given to the functions in the generated file.")
	optIgnore := app.StringsOpt("i ignore", []string{"main"}, "Ignore given package names from the list.")

	app.Version("v version", util.VersionString(opts))

	app.Action = func() {
		opts.Set("directory", *optDir)
		opts.Set("output", *optOutput)
		opts.Set("package", *optPackage)
		opts.Set("prefix", *optPrefix)
		opts.Set("ignore-packages", strings.Join(*optIgnore, " "))
	}

	return app.Run(argsin)
}
