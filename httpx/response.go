package httpx

import (
	"bytes"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
)

var (
	healthOK     = []byte("up")
	healthLength = int64(len(healthOK))
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

func NewInfoResponse(status *core.Status, authority, version string) *http.Response {
	if status == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}
	}
	h := make(http.Header)
	h.Add(core.XVersion, version)
	h.Add(core.XAuthority, authority)
	return &http.Response{StatusCode: status.HttpCode(), Header: h}
}

func NewHealthResponseOK() *http.Response {
	return &http.Response{StatusCode: http.StatusOK, ContentLength: healthLength, Body: io.NopCloser(bytes.NewReader(healthOK))}

}
