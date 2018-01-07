package licrep

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
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
			lines += 1
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
	w.Close()
	bw.Close()
	if err != nil {
		return
	}

	out = data.String()
	return

}

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

			fmt.Printf("PKG: %s\nData:\n%s\nOriglen: %d\nPackedLen: %d\n",
				pkgs[i].Name, str, len(pkgs[i].LicenseString), len(str),
			)
		}
		licenseData.WriteString(fmt.Sprintf(`
		"%s": Data{
			Name: "%s",
			Data: `+"`\n%s`"+`,
		},`, pkgs[i].Name, pkgs[i].License, str))
	}

	argmap := map[string]string{
		"programName":    opts.Get("program-name", ""),
		"programVersion": opts.Get("program-version", ""),
		"programArgs":    opts.Get("program-args", ""),
		"prefix":         opts.Get("prefix", ""),
		"package":        opts.Get("package", ""),
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
		defer fp.Close()
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
	type Data struct {
		Name string
		Data string
	}
	data := map[string]Data{
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
		text, err := decode(data[k].Data)
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
