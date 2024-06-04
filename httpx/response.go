package httpx

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
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

func NewResponseWithBody[E core.ErrorHandler](h http.Header, statusCode int, content any) *http.Response {
	var e E

	resp := &http.Response{StatusCode: statusCode, Header: h}
	switch ptr := (content).(type) {
	case []byte:
		resp.Body = io.NopCloser(bytes.NewReader(ptr))
	case string:
		resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr)))
	case error:
		resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr.Error())))
	default:
		if h != nil && h.Get(ContentType) == ContentTypeJson {
			var status *core.Status

			resp.Body, resp.ContentLength, status = json2.NewReadCloser(content)
			if !status.OK() {
				e.Handle(status, "")
				resp.StatusCode = status.HttpCode()
			}
		} else {
			//return 0, core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
		}
	}
	return resp
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
	return NewResponseWithStatus(core.StatusNotFound(), core.StatusNotFound().String())
}

func NewJsonResponse(content any, h http.Header) (*http.Response, *core.Status) {
	if content == nil {
		return &http.Response{StatusCode: http.StatusOK, Header: h}, core.StatusOK()
	}
	rc, length, status := json2.NewReadCloser(content)
	if !status.OK() {
		return NewResponseWithStatus(status, status.Err)
	}
	if h == nil {
		h = make(http.Header)
	}
	h.Add(ContentType, ContentTypeJson)
	return &http.Response{StatusCode: status.HttpCode(), Status: status.String(), ContentLength: length, Header: h, Body: rc}, core.StatusOK()
}
