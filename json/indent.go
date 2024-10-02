package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"net/http"
)

func Indent(body io.ReadCloser, h http.Header, prefix, indent string) (io.ReadCloser, *core.Status) {
	var buf bytes.Buffer

	if body == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: body is nil"))
	}
	buf2, status := io2.ReadAll(body, h)
	if !status.OK() {
		return nil, status
	}
	err := json.Indent(&buf, buf2, prefix, indent)
	if err != nil {
		return nil, core.NewStatusError(core.StatusJsonDecodeError, err)
	}
	return io.NopCloser(bytes.NewReader(buf.Bytes())), core.StatusOK()
}
