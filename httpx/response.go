package httpx

import (
	"bytes"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
)

var (
	HealthResponseUp = NewResponse(core.StatusOK(), core.HealthContent("up"))
)

func NewErrorResponse(status *core.Status) *http.Response {
	if status == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}
	}
	if status.Err == nil {
		return &http.Response{StatusCode: status.HttpCode()}
	}
	return NewResponse(status, status.Err.Error())
}

func NewResponse(status *core.Status, content string) *http.Response {
	if status == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}
	}
	if len(content) == 0 {
		return &http.Response{StatusCode: status.HttpCode()}
	}
	return &http.Response{StatusCode: status.HttpCode(), ContentLength: int64(len(content)), Body: io.NopCloser(bytes.NewReader([]byte(content)))}
}
