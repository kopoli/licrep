package licrep

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"

	util "github.com/kopoli/go-util"
)

type chunker struct {
	output     io.Writer
	pos        int
	lineLength int
}

// Write writes the given data to w.output and every w.lineLength bytes writes
// a newline character.
func (w *chunker) Write(p []byte) (int, error) {
	if w.lineLength == 0 {
		w.lineLength = 40
	}

	lines := 0

	for idx := 0; idx < len(p); {
		startpos := w.pos
		writeLen := w.lineLength - w.pos
		if writeLen+idx > len(p) {
			writeLen = len(p) - idx
		}

		w.pos = (w.pos + writeLen) % w.lineLength

		fmt.Println("Writing from idx", idx, "to", idx+writeLen, "full length", len(p), "wpos", w.pos, startpos)
		writeLen, err := w.output.Write(p[idx : idx+writeLen])
		if err != nil {
			return 0, err
		}

		if w.pos == 0 {
			fmt.Println("Writing a newline! startpos", startpos)
			_, err = w.output.Write([]byte("\n"))
			if err != nil {
				return 0, err
			}
			lines += 1
		}

		idx += writeLen
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
	var str string
	for i := range pkgs {
		if pkgs[i].LicenseString != "" {
			str, err = encodeData(pkgs[i].LicenseString)
			if err != nil {
				return
			}

			fmt.Printf("PKG: %s\nData:\n%s\nOriglen: %d\nPackedLen: %d\n",
				pkgs[i].Name, str, len(pkgs[i].LicenseString), len(str),
			)
			// fmt.Println("PKG:", pkgs[i].Name, "Packaged data:", str)
		}
	}

	str, err = encodeData("jotain sinnepäin tai sitten ei asfas g as gasgasg sag asg asgasga gagas")
	if err != nil {
		return
	}

	fmt.Println("Pakattu data:", str)
	return
}
