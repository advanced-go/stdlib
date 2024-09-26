package httpx

import (
	"encoding/json"
	"errors"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
)

const (
	eofError = "EOF"
)

func Content[T any](body io.Reader) (t T, status *core.Status) {
	if body == nil {
		return t, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: body is nil"))
	}
	status = core.StatusOK()
	switch p := any(&t).(type) {
	case *[]byte:
		var buf []byte
		buf, status = io2.ReadAll(body, nil)
		if !status.OK() {
			return
		}
		if len(buf) == 0 {
			return t, core.StatusNotFound()
		}
		*p = buf
	case *string:
		var buf []byte
		buf, status = io2.ReadAll(body, nil)
		if !status.OK() {
			return
		}
		if len(buf) == 0 {
			return t, core.StatusNotFound()
		}
		*p = string(buf)
	default:
		err := json.NewDecoder(body).Decode(p)
		if err != nil {
			// If the error is "EOF", then the body was empty. If the error is "unexpected EOF", then the body has content
			// but the EOF was reached when more JSON content was expected.
			if err.Error() == eofError {
				status = core.StatusNotFound()
			} else {
				status = core.NewStatusError(core.StatusJsonDecodeError, err)
			}
		}
	}
	return
}
