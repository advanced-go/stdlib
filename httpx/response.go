package httpx

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	json2 "github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
	"reflect"
	"strings"
)

const (
	versionFmt      = "{\n \"version\": \"%v\"\n}"
	authorityFmt    = "{\n \"authority\": \"%v\"\n}"
	fileExistsError = "The system cannot find the file specified"
)

var (
	healthOK     = []byte("{\n \"status\": \"up\"\n}")
	healthLength = int64(len(healthOK))
)

func NewError(status *core.Status, resp *http.Response) string {
	if status != nil && status.Err != nil {
		return status.Err.Error()
	}
	if resp != nil && resp.Body != nil {
		s, status1 := Content[string](resp.Body)
		if status1.OK() {
			return s
		}
		return status1.String()
	}
	return ""
}

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

func NewResponse(statusCode int, h http.Header, content any) (resp *http.Response, status *core.Status) {
	resp = &http.Response{StatusCode: statusCode, ContentLength: -1, Header: h, Body: io.NopCloser(bytes.NewReader([]byte{}))}
	if h == nil {
		resp.Header = make(http.Header)
	}
	if content == nil {
		return resp, core.NewStatus(statusCode)
	}
	switch ptr := (content).(type) {
	case []byte:
		if len(ptr) > 0 {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader(ptr))
		}
	case string:
		if ptr != "" {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr)))
		}
	case error:
		if ptr.Error() != "" {
			resp.ContentLength = int64(len(ptr.Error()))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr.Error())))
		}
	default:
		status = core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
		return &http.Response{StatusCode: http.StatusInternalServerError, Header: SetHeader(nil, ContentType, ContentTypeText), ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
	}
	return resp, core.NewStatus(statusCode)
}

func NewResponse1[E core.ErrorHandler](statusCode int, h http.Header, content any) (resp *http.Response, status *core.Status) {
	var e E

	resp = &http.Response{StatusCode: statusCode, Header: h, Body: io.NopCloser(bytes.NewReader([]byte{}))}
	if h == nil {
		resp.Header = make(http.Header)
	}
	if content == nil {
		return resp, core.NewStatus(statusCode)
	}
	switch ptr := (content).(type) {
	case []byte:
		if len(ptr) > 0 {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader(ptr))
		}
	case string:
		if ptr != "" {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr)))
		}
	case error:
		if ptr.Error() != "" {
			resp.ContentLength = int64(len(ptr.Error()))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr.Error())))
		}
	default:
		if h != nil && h.Get(ContentType) == ContentTypeJson {
			resp.Body, resp.ContentLength, status = json2.NewReadCloser(content)
			if !status.OK() {
				e.Handle(status)
				h = SetHeader(nil, ContentType, ContentTypeText)
				if status.Err != nil {
					return &http.Response{StatusCode: http.StatusInternalServerError, Header: h, ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
				} else {
					return &http.Response{StatusCode: http.StatusInternalServerError, Header: h}, status
				}
			}
		} else {
			status = core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
			return &http.Response{StatusCode: http.StatusInternalServerError, Header: SetHeader(nil, ContentType, ContentTypeText), ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
		}
	}
	return resp, core.NewStatus(statusCode)
}

func NewVersionResponse(version string) *http.Response {
	content := fmt.Sprintf(versionFmt, version)
	resp, _ := NewResponse(http.StatusOK, SetHeader(nil, ContentType, ContentTypeText), content)
	return resp
}

func NewAuthorityResponse(authority string) *http.Response {
	resp, _ := NewResponse(http.StatusOK, SetHeader(nil, core.XAuthority, authority), nil)
	return resp
}

func NewHealthResponseOK() *http.Response {
	resp, _ := NewResponse(http.StatusOK, SetHeader(nil, ContentType, ContentTypeText), healthOK)
	return resp
}

func NewNotFoundResponse() *http.Response {
	resp, _ := NewResponse(http.StatusNotFound, SetHeader(nil, ContentType, ContentTypeText), core.StatusNotFound().String())
	return resp
}

// NewResponseFromUri - read a Http response given a URL
func NewResponseFromUri(uri any) (*http.Response, *core.Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: internalError, Header: make(http.Header)}
	if uri == nil {
		return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	//}
	buf, status := io2.ReadFile(uri)
	if !status.OK() {
		if strings.Contains(status.Err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found", Header: make(http.Header)}, core.NewStatusError(core.StatusInvalidArgument, status.Err)
		}
		return serverErr, core.NewStatusError(core.StatusIOError, status.Err)
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, core.NewStatusError(core.StatusIOError, err2)
	}
	return resp1, core.StatusOK()

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
