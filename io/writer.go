package io

import (
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
	"strings"
)

type EncodingWriter interface {
	io.WriteCloser
	ContentEncoding() string
}

func NewEncodingWriter(w io.Writer, h http.Header) (EncodingWriter, *core.Status) {
	encoding := acceptEncoding(h)
	if strings.Contains(encoding, GzipEncoding) {
		return NewGzipWriter(w), core.StatusOK()
	}
	// TODO : implement additional encoding support
	return NewIdentityWriter(w), core.StatusOK()
}

type identityWriter struct {
	writer io.Writer
}

// NewIdentityWriter - The default (identity) encoding; the use of no transformation whatsoever
func NewIdentityWriter(w io.Writer) EncodingWriter {
	iw := new(identityWriter)
	iw.writer = w
	return iw
}

func (i *identityWriter) Write(p []byte) (n int, err error) {
	return i.writer.Write(p)
}

func (i *identityWriter) ContentEncoding() string {
	return NoneEncoding
}

func (i *identityWriter) Close() error {
	return nil
}
