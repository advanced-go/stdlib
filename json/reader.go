package json

import (
	"bytes"
	"github.com/advanced-go/stdlib/core"
	"io"
)

// NewReadCloser - create an io.ReadCloser from a type
func NewReadCloser(v any) (io.ReadCloser, int64, *core.Status) {
	if v == nil {
		return io.NopCloser(bytes.NewReader([]byte(""))), 0, core.StatusOK()
	}
	buf, status := Marshal(v)
	if !status.OK() {
		return nil, 0, status
	}
	return io.NopCloser(bytes.NewReader(buf)), int64(len(buf)), core.StatusOK()
}
