package httpx

import (
	"bytes"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
)

func NewErrorResponse(status *core.Status) *http.Response {
	if status == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}
	}
	return &http.Response{StatusCode: status.HttpCode(), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}
}
