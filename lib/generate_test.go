package licrep

import (
	"bytes"
	"testing"
)

func Test_chunker_Write(t *testing.T) {
	tests := []struct {
		name       string
		linePos    int
		lineLength int
		input      string

		output string
	}{
		{"Empty data", 0, 0, "", ""},
		{"One line", 0, 4, "abc", "abc"},
		{"Exactly one line", 0, 3, "abc", "abc\n"},
		{"Exactly two lines", 0, 2, "abcd", "ab\ncd\n"},
		{"Shorter first line", 1, 2, "abcd", "a\nbc\nd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			w := &chunker{
				output:     buf,
				linePos:    tt.linePos,
				lineLength: tt.lineLength,
			}
			got, err := w.Write([]byte(tt.input))
			if (err != nil) != false {
				t.Errorf("chunker.Write() error = %v, wantErr %v", err, false)
				return
			}
			if got != len(tt.output) || buf.String() != tt.output {
				t.Errorf("Unexpected output:\nlength: got: %d, expected %d\ndata:\ngot: [%v]\nexpected: [%v]",
					got, len(tt.output), buf.String(), tt.output,
				)
				return
			}
		})
	}
}
