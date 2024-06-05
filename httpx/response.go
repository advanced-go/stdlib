package httpx

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
	"reflect"
)

const (
	versionFmt = "{\n \"version\": \"%v\"\n}"
)

var (
	healthOK     = []byte("{\n \"status\": \"up\"\n}")
	healthLength = int64(len(healthOK))
)

/*
func NewResponse2(status *core.Status, content any) *http.Response {
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

*/

func NewResponse[E core.ErrorHandler](statusCode int, h http.Header, content any) (resp *http.Response, status *core.Status) {
	var e E

	resp = &http.Response{StatusCode: statusCode, Header: h}
	if content == nil {
		return resp, core.NewStatus(statusCode)
	}
	switch ptr := (content).(type) {
	case []byte:
		resp.ContentLength = int64(len(ptr))
		resp.Body = io.NopCloser(bytes.NewReader(ptr))
	case string:
		resp.ContentLength = int64(len(ptr))
		resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr)))
	case error:
		if ptr.Error() != "" {
			resp.ContentLength = int64(len(ptr.Error()))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr.Error())))
		}
	default:
		if h != nil && h.Get(ContentType) == ContentTypeJson {
			resp.Body, resp.ContentLength, status = json2.NewReadCloser(content)
			if !status.OK() {
				e.Handle(status, "")
				h = make(http.Header)
				h.Add(ContentType, ContentTypeText)
				if status.Err != nil {
					return &http.Response{StatusCode: http.StatusInternalServerError, Header: h, ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
				} else {
					return &http.Response{StatusCode: http.StatusInternalServerError, Header: h}, status
				}
			}
		} else {
			status = core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
			h = make(http.Header)
			h.Add(ContentType, ContentTypeText)
			return &http.Response{StatusCode: http.StatusInternalServerError, Header: h, ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
		}
	}
	return resp, core.NewStatus(statusCode)
}

func NewVersionResponse(version string) *http.Response {
	h2 := make(http.Header)
	h2.Add(ContentType, ContentTypeJson)
	content := fmt.Sprintf(versionFmt, version)
	resp, _ := NewResponse[core.Log](http.StatusOK, h2, content)
	return resp
}

func NewAuthorityResponse(authority string) *http.Response {
	h2 := make(http.Header)
	h2.Add(core.XAuthority, authority)
	//h.Add(ContentType, ContentTypeJson)
	resp, _ := NewResponse[core.Log](http.StatusOK, h2, nil)
	return resp
}

func NewHealthResponseOK() *http.Response {
	h2 := make(http.Header)
	h2.Add(ContentType, ContentTypeText)
	resp, _ := NewResponse[core.Log](http.StatusOK, h2, healthOK)
	return resp
	///&http.Response{StatusCode: http.StatusOK, Header: h2, ContentLength: healthLength, Body: io.NopCloser(bytes.NewReader(healthOK))}
}

func NewNotFoundResponse() *http.Response {
	h2 := make(http.Header)
	h2.Add(ContentType, ContentTypeText)
	resp, _ := NewResponse[core.Log](http.StatusNotFound, h2, core.StatusNotFound().String())
	return resp
}

/*
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


*/
