package json

import (
	"bytes"
	"encoding/json"
	"github.com/advanced-go/stdlib/core"
	"io"
)

// NewReader - create an io.Reader from a type
func NewReader(v any) (io.Reader, *core.Status) {
	if v == nil {
		return bytes.NewReader([]byte("")), core.StatusOK()
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, core.NewStatusError(core.StatusJsonEncodeError, err)
	}
	return bytes.NewReader(buf), core.StatusOK()
}
