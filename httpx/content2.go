package httpx

import (
	"encoding/json"
	"errors"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
)

func Content[T any](body io.Reader) (t T, status *core.Status) {
	if body == nil {
		return t, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: body is nil"))
	}
	var buf []byte
	buf, status = io2.ReadAll(body, nil)
	if !status.OK() {
		return
	}
	if len(buf) == 0 {
		return t, core.StatusNotFound()
	}
	switch p := any(&t).(type) {
	case *[]byte:
		*p = buf
	case *string:
		*p = string(buf)
	default:
		err := json.NewDecoder(body).Decode(p)
		if err != nil {
			status = core.NewStatusError(core.StatusJsonDecodeError, err)
		}
	}
	return
}
