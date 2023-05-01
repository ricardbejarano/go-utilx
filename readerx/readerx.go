package readerx

import (
	"bytes"
	"io"
	"strings"
)

func ToString(r io.Reader) string {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, r); err != nil {
		return ""
	}

	return buf.String()
}

func ToBytes(r io.Reader) []byte {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		return []byte{}
	}

	return buf.Bytes()
}
