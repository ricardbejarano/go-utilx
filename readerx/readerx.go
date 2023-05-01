package readerx

import (
	"bytes"
	"io"
	"strings"
)

func ToString(reader io.Reader) string {
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, reader); err != nil {
		return ""
	}

	return buf.String()
}

func ToBytes(reader io.Reader) []byte {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, reader); err != nil {
		return []byte{}
	}

	return buf.Bytes()
}
