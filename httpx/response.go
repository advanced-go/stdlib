package httpx

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
)

const (
	versionFmt = "{\n \"version\": \"%v\"\n}"
)

var (
	healthOK     = []byte("{\n \"status\": \"up\"\n}")
	healthLength = int64(len(healthOK))
)

/*
func NewErrorResponse(status *core.Status) *http.Response {
	if status == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}
	}
	if status.Err == nil {
		return &http.Response{StatusCode: status.HttpCode()}
	}
	return NewResponse(status, status.Err.Error())
}

func NewErrorResponseWithStatus(status *core.Status) (*http.Response, *core.Status) {
	resp := NewErrorResponse(status)
	if status == nil {
		status = core.NewStatus(http.StatusBadRequest)
	}
	return resp, status
}


*/

func NewResponse(status *core.Status, content any) *http.Response {
	if status == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}
	}
	if content == nil {
		return &http.Response{StatusCode: status.HttpCode(), Status: status.String()}
	}
	h := make(http.Header)
	h.Add(ContentType, ContentTypeText)
	if s, ok := content.(string); ok {
		return &http.Response{StatusCode: status.HttpCode(), Status: status.String(), Header: h, ContentLength: int64(len(s)), Body: io.NopCloser(bytes.NewReader([]byte(s)))}
	}
	if err, ok := content.(error); ok {
		return &http.Response{StatusCode: status.HttpCode(), Status: status.String(), Header: h, ContentLength: int64(len(err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(err.Error())))}
	}
	return &http.Response{StatusCode: http.StatusBadRequest, Header: h, Body: io.NopCloser(bytes.NewReader([]byte(fmt.Sprintf("invalid content : %v", core.NewInvalidBodyTypeError(content)))))}
}

func NewResponseWithStatus(status *core.Status, content any) (*http.Response, *core.Status) {
	if status == nil {
		status = core.NewStatus(http.StatusBadRequest)
	}
	return NewResponse(status, content), status
}

func NewVersionResponse(version string) *http.Response {
	h := make(http.Header)
	h.Add(ContentType, ContentTypeJson)
	content := fmt.Sprintf(versionFmt, version)
	return &http.Response{StatusCode: http.StatusOK, Header: h, Body: io.NopCloser(bytes.NewReader([]byte(content)))}
}

func NewAuthorityResponse(authority string) *http.Response {
	h := make(http.Header)
	h.Add(core.XAuthority, authority)
	//h.Add(ContentType, ContentTypeJson)
	return &http.Response{StatusCode: http.StatusOK, Header: h}
}

func NewHealthResponseOK() *http.Response {
	return &http.Response{StatusCode: http.StatusOK, ContentLength: healthLength, Body: io.NopCloser(bytes.NewReader(healthOK))}
}

func NewNotFoundResponseWithStatus() (*http.Response, *core.Status) {
	return NewResponseWithStatus(core.NewStatus(http.StatusNotFound), "Not Found")
}
