package licrep

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/kopoli/appkit"
)

// chunker is a writer which writes given data to output and emits a newline
// every lineLength bytes received.
type chunker struct {
	output     io.Writer
	linePos    int
	lineLength int
}

// Write writes the given data to w.output and every w.lineLength bytes writes
// a newline character.
func (w *chunker) Write(p []byte) (int, error) {
	if w.lineLength == 0 {
		w.lineLength = 80
	}

	lines := 0

	for idx := 0; idx < len(p); {
		writeLen := w.lineLength - w.linePos
		if writeLen+idx > len(p) {
			writeLen = len(p) - idx
		}

		w.linePos = (w.linePos + writeLen) % w.lineLength

		writeLen, err := w.output.Write(p[idx : idx+writeLen])
		if err != nil {
			return 0, err
		}

		idx += writeLen

		if w.linePos == 0 {
			_, err = w.output.Write([]byte("\n"))
			if err != nil {
				return 0, err
			}
			lines++
		}
	}

	return len(p) + lines, nil
}

func encodeData(input string) (out string, err error) {
	data := &bytes.Buffer{}
	ch := &chunker{output: data}
	bw := base64.NewEncoder(base64.StdEncoding, ch)
	w, err := gzip.NewWriterLevel(bw, gzip.BestCompression)
	if err != nil {
		return "", err
	}

	_, err = w.Write([]byte(input))
	err2 := w.Close()
	err3 := bw.Close()
	if err != nil {
		return "", err
	}
	if err2 != nil {
		err = err2
	}
	if err3 != nil {
		err = err3
	}
	if err != nil {
		return "", err
	}

	out = data.String()
	return out, nil
}

// FilterPackages filters out the ignored packages from the list
func FilterPackages(opts appkit.Options, pkgs []Package) (ret []Package) {
	// Prepare the list of ignored packages
	ignorestr := opts.Get("ignore-packages", "")
	var ignores []string
	if ignorestr != "" {
		ignores = strings.Split(ignorestr, " ")
	}

	// Check if the package has been ignored
	var ignored bool
	for i := range pkgs {
		ignored = false
		for j := range ignores {
			re, err := regexp.Compile(ignores[j])
			if err != nil {
				continue
			}
			if re.MatchString(pkgs[i].Name) || re.MatchString(pkgs[i].ImportPath) {
				ignored = true
			}
		}
		if !ignored {
			ret = append(ret, pkgs[i])
		}
	}
	return
}

// determinePackage determines the package that is written to the generated
// file
func determinePackage(opts appkit.Options) (string, error) {
	pkg := opts.Get("package", "")

	// Return the package given explicitly in the command line
	if pkg != "" {
		return pkg, nil
	}

	outfile := opts.Get("output", "")

	var dir string
	var err error
	if outfile == "" {
		// If outfile is stdout, return the package in the current working
		// directory
		dir, err = os.Getwd()
		if err != nil {
			return "", err
		}
	} else {
		// If outfile is a filename, get the package of that directory.
		dir, err = filepath.Abs(outfile)
		if err != nil {
			return "", err
		}
		dir = filepath.Dir(dir)
	}

	p, err := build.Import(".", dir, 0)
	if err != nil {
		return "", err
	}
	return p.Name, nil
}

// GenerateEmbeddedLicenses creates a compressed license representation to a
// file or os.Stdout of the given Packages. The representation is go so it can
// be built as part of a program.
func GenerateEmbeddedLicenses(opts appkit.Options, pkgs []Package) error {
	licenseData := &bytes.Buffer{}
	var str string
	var err error

	for i := range pkgs {
		str = ""
		if pkgs[i].LicenseString != "" {
			str, err = encodeData(pkgs[i].LicenseString)
			if err != nil {
				return err
			}
		}
		licenseData.WriteString(fmt.Sprintf(`
		"%s": {
			Name: "%s",
			Text: `+"`\n%s`"+`,
			length: %d,
		},`, pkgs[i].ImportPath, pkgs[i].License, str,
			len(pkgs[i].LicenseString)))
	}

	pkg, err := determinePackage(opts)
	if err != nil {
		return err
	}

	argmap := map[string]string{
		"programName":    opts.Get("program-name", ""),
		"programVersion": opts.Get("program-version", ""),
		"programArgs":    opts.Get("program-args", ""),
		"prefix":         opts.Get("prefix", ""),
		"package":        pkg,
		"data":           licenseData.String(),
	}

	var out io.Writer = os.Stdout

	outfile := opts.Get("output", "")
	if outfile != "" {
		var fp *os.File
		fp, err = os.Create(outfile)
		if err != nil {
			return err
		}
		out = fp
		defer func() {
			err = fp.Close()
		}()
	}

	tmpl, err := template.New("licrep").Parse(licenseTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(out, argmap)
	return nil
}

const (
	licenseTemplate string = `// Generated with by {{.programName}} version {{.programVersion}}
// https://github.com/kopoli/licrep
// Called with: {{.programArgs}}

package {{.package}}

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

// {{.prefix}}License is a representation of an embedded license.
type {{.prefix}}License struct {
	// Name of the license
	Name string

	// The text of the license
	Text string
}

// {{.prefix}}GetLicenses gets a map of {{.prefix}}Licenses where the keys are
// the package names.
func {{.prefix}}GetLicenses() (map[string]{{.prefix}}License, error) {
	type EncodedLicense struct {
		Name   string
		Text   string
		length int64
	}
	data := map[string]EncodedLicense{
{{.data}}
	}

	decode := func(input string, length int64) (string, error) {
		data := &bytes.Buffer{}
		br := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))

		r, err := gzip.NewReader(br)
		if err != nil {
			return "", err
		}

		_, err = io.CopyN(data, r, length)
		if err != nil {
			return "", err
		}

		// Make sure the gzip is decoded successfully
		err = r.Close()
		if err != nil {
			return "", err
		}
		return data.String(), nil
	}

	ret := make(map[string]{{.prefix}}License)

	for k := range data {
		text, err := decode(data[k].Text, data[k].length)
		if err != nil {
			return nil, err
		}
		ret[k] = {{.prefix}}License{
			Name: data[k].Name,
			Text: text,
		}
	}

	return ret, nil
}
`
)
