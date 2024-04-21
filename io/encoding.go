package io

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	AcceptEncoding      = "Accept-Encoding"
	AcceptEncodingValue = "gzip, deflate, br"
	ContentEncoding     = "Content-Encoding"

	GzipEncoding     = "gzip"
	BrotliEncoding   = "br"
	DeflateEncoding  = "deflate"
	CompressEncoding = "compress"
	NoneEncoding     = "none"

	ApplicationGzip    = "application/x-gzip"
	ApplicationBrotli  = "application/x-br"
	ApplicationDeflate = "application/x-deflate"

	encodingErrorFmt = "error: content encoding not supported [%v]"
)

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

func newStatusContentEncodingError(ct string) *core.Status {
	return core.NewStatusError(core.StatusContentEncodingError, errors.New(fmt.Sprintf(encodingErrorFmt, ct)))

}

// Decode - decode a []byte
func Decode(buf []byte, h http.Header) ([]byte, *core.Status) {
	if len(buf) == 0 {
		return buf, core.StatusOK()
	}
	ct := NoneEncoding
	if h == nil {
		ct = http.DetectContentType(buf)
	} else {
		ct = contentEncoding(h)
	}
	switch ct {
	case ApplicationGzip, GzipEncoding:
		zr, status := NewGzipReader(bytes.NewReader(buf))
		if !status.OK() {
			return nil, status
		}
		buf2, err1 := io.ReadAll(zr)
		err2 := zr.Close()
		if err1 != nil {
			return nil, core.NewStatusError(core.StatusIOError, err1)
		}
		if err2 != nil {
			//return nil, core.NewStatusError(core.StatusIOError,err1)
		}
		return buf2, core.StatusOK()
	case ApplicationBrotli, BrotliEncoding:
		return buf, newStatusContentEncodingError(ct)
	case ApplicationDeflate, DeflateEncoding:
		return buf, newStatusContentEncodingError(ct)
	default:
		return buf, core.StatusOK()
	}
}

func ZipFile(uri string) *core.Status {
	if len(uri) == 0 {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("error: file path is empty"))
	}
	path := FileName(uri)
	content, err0 := os.ReadFile(path)
	if err0 != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err0)
		return core.NewStatusError(core.StatusIOError, err0)
	}
	// write content
	buff := new(bytes.Buffer)
	zw := NewGzipWriter(buff)
	cnt, err := zw.Write(content)
	err1 := zw.Close()
	if err != nil {
		return core.NewStatusError(core.StatusIOError, err)
	}
	if cnt == 0 || err1 != nil {
		fmt.Printf("error: count %v err %v", cnt, err1)
	}
	i := strings.LastIndex(path, ".")
	path2 := ""
	if i > 0 {
		path2 = path[:i]
		path2 += ".gz"
	} else {
		path2 = path + ".gz"
	}
	err = os.WriteFile(path2, buff.Bytes(), 667)
	if err != nil {
		return core.NewStatusError(core.StatusIOError, err)
	}
	return core.StatusOK()
}
