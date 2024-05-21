package json

import (
	"bytes"
	"encoding/json"
	"github.com/advanced-go/stdlib/core"
	"io"
)

// NewReadCloser - create an io.ReadCloser from a type
func NewReadCloser(v any) (io.ReadCloser, *core.Status) {
	if v == nil {
		return io.NopCloser(bytes.NewReader([]byte(""))), core.StatusOK()
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, core.NewStatusError(core.StatusJsonEncodeError, err)
	}
	return io.NopCloser(bytes.NewReader(buf)), core.StatusOK()
}
