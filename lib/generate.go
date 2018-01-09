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
	"strings"
	"text/template"

	util "github.com/kopoli/go-util"
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
		return
	}

	_, err = w.Write([]byte(input))
	err2 := w.Close()
	err3 := bw.Close()
	if err != nil {
		return
	}
	if err2 != nil {
		err = err2
	}
	if err3 != nil {
		err = err3
	}
	if err != nil {
		return
	}

	out = data.String()
	return

}

// FilterPackages filters out the ignored packages from the list
func FilterPackages(opts util.Options, pkgs []Package) (ret []Package) {
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
			if ignores[j] == pkgs[i].Name || ignores[j] == pkgs[i].ImportPath {
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
func determinePackage(opts util.Options, pkgs []Package) (string, error) {
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
func GenerateEmbeddedLicenses(opts util.Options, pkgs []Package) (err error) {
	licenseData := &bytes.Buffer{}
	var str string
	for i := range pkgs {
		str = ""
		if pkgs[i].LicenseString != "" {
			str, err = encodeData(pkgs[i].LicenseString)
			if err != nil {
				return
			}
		}
		licenseData.WriteString(fmt.Sprintf(`
		"%s": EncodedLicense{
			Name: "%s",
			Text: `+"`\n%s`"+`,
		},`, pkgs[i].ImportPath, pkgs[i].License, str))
	}

	pkg, err := determinePackage(opts, pkgs)
	if err != nil {
		return
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
			return
		}
		out = fp
		defer func() {
			err = fp.Close()
		}()
	}

	tmpl, err := template.New("licrep").Parse(licenseTemplate)
	if err != nil {
		return
	}

	err = tmpl.Execute(out, argmap)
	return
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

type {{.prefix}}License struct {
	// Name of the license
	Name string

	// The text of the license
	Text string
}

func {{.prefix}}GetLicenses() (map[string]{{.prefix}}License, error) {
	type EncodedLicense struct {
		Name string
		Text string
	}
	data := map[string]EncodedLicense{
{{.data}}
	}

	decode := func(input string) (string, error) {
		data := &bytes.Buffer{}
		br := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))

		r, err := gzip.NewReader(br)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(data, r)
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
		text, err := decode(data[k].Text)
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
