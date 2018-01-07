package licrep

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"

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
