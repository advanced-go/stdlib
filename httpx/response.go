package httpx

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
)

const (
	infoFmt    = "{\n \"authority\": \"%v\",\n \"version\": \"%v\",\n \"name\": \"%v\"\n  }"
	versionFmt = "{\n \"version\": \"%v\"\n}"
)

var (
	healthOK     = []byte("{\n \"status\": \"up\"\n}")
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
	h := make(http.Header)
	h.Add(ContentType, ContentTypeText)
	return &http.Response{StatusCode: status.HttpCode(), ContentLength: int64(len(content)), Header: h, Body: io.NopCloser(bytes.NewReader([]byte(content)))}
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
