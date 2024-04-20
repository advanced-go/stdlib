package sio

import (
	"errors"
	"fmt"
	"github.com/advaced-go/stdlib/core"
	"io"
	"net/http"
	"strings"
)

const (
	AcceptEncoding      = "Accept-Encoding"
	ContentEncoding     = "Content-Encoding"
	GzipEncoding        = "gzip"
	BrotliEncoding      = "br"
	DeflateEncoding     = "deflate"
	CompressEncoding    = "compress"
	NoneEncoding        = "none"
	AcceptEncodingValue = "gzip, deflate, br"

	encodingErrorFmt  = "error: content encoding not supported [%v]"
	encodingReaderLoc = PkgPath + ":EncodingReader"
)

type EncodingReader interface {
	io.ReadCloser
}

type EncodingWriter interface {
	io.WriteCloser
	ContentEncoding() string
}

func NewEncodingReader(r io.Reader, h http.Header) (EncodingReader, *core.Status) {
	encoding := contentEncoding(h)
	switch encoding {
	case GzipEncoding:
		return NewGzipReader(r)
	case BrotliEncoding, DeflateEncoding, CompressEncoding:
		return nil, core.NewStatusError(core.StatusContentEncodingError, errors.New(fmt.Sprintf(encodingErrorFmt, encoding)))
	default:
		return NewIdentityReader(r), core.StatusOK()
	}
}

func NewEncodingWriter(w io.Writer, h http.Header) (EncodingWriter, *core.Status) {
	encoding := acceptEncoding(h)
	if strings.Contains(encoding, GzipEncoding) {
		return NewGzipWriter(w), core.StatusOK()
	}
	return NewIdentityWriter(w), core.StatusOK()
}

func contentEncoding(h http.Header) string {
	if h == nil {
		return NoneEncoding
	}
	enc := h.Get(ContentEncoding)
	if len(enc) > 0 {
		return strings.ToLower(enc)
	}
	return NoneEncoding
}

func acceptEncoding(h http.Header) string {
	if h == nil {
		return NoneEncoding
	}
	enc := h.Get(AcceptEncoding)
	if len(enc) > 0 {
		return strings.ToLower(enc)
	}
	return NoneEncoding
}

type identityReader struct {
	reader io.Reader
}

// NewIdentityReader - The default (identity) encoding; the use of no transformation whatsoever
func NewIdentityReader(r io.Reader) EncodingReader {
	ir := new(identityReader)
	ir.reader = r
	return ir
}

func (i *identityReader) Read(p []byte) (n int, err error) {
	return i.reader.Read(p)
}

func (i *identityReader) Close() error {
	return nil
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
