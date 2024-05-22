package json

import (
	"bytes"
	"encoding/json"
	"github.com/advanced-go/stdlib/core"
	"io"
)

// NewReadCloser - create an io.ReadCloser from a type
func NewReadCloser(v any) (io.ReadCloser, int64, *core.Status) {
	if v == nil {
		return io.NopCloser(bytes.NewReader([]byte(""))), 0, core.StatusOK()
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, 0, core.NewStatusError(core.StatusJsonEncodeError, err)
	}
	return io.NopCloser(bytes.NewReader(buf)), int64(len(buf)), core.StatusOK()
}
