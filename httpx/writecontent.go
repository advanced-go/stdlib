package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"reflect"
	"strings"
)

const (
	jsonToken = "json"
)

func writeContent(w io.Writer, content any, contentType string) (length int64, status *core.Status) {
	var err error
	var cnt int

	if content == nil {
		return 0, core.StatusOK()
	}
	switch ptr := (content).(type) {
	case []byte:
		cnt, err = w.Write(ptr)
	case string:
		cnt, err = w.Write([]byte(ptr))
	case error:
		cnt, err = w.Write([]byte(ptr.Error()))
	case io.Reader:
		var buf []byte

		buf, status = io2.ReadAll(ptr, nil)
		if !status.OK() {
			return 0, status.AddLocation()
		}
		cnt, err = w.Write(buf)
	case io.ReadCloser:
		var buf []byte

		buf, status = io2.ReadAll(ptr, nil)
		_ = ptr.Close()
		if !status.OK() {
			return 0, status.AddLocation()
		}
		cnt, err = w.Write(buf)
	default:
		if strings.Contains(contentType, jsonToken) {
			var buf []byte

			buf, err = json.Marshal(content)
			if err != nil {
				status = core.NewStatusError(core.StatusJsonEncodeError, err)
				if !status.OK() {
					return
				}
			}
			cnt, err = w.Write(buf)
		} else {
			return 0, core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
		}
	}
	if err != nil {
		return 0, core.NewStatusError(core.StatusIOError, err)
	}
	return int64(cnt), core.StatusOK()
}
