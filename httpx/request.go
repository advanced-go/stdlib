package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"github.com/google/uuid"
	"net/http"
)

// AddRequestId - add a request to an http.Request or an http.Header
func AddRequestId(t any) http.Header {
	if t == nil {
		h := make(http.Header)
		return addRequestId(h)
	}
	if req, ok := t.(*http.Request); ok {
		req.Header = addRequestId(req.Header)
		return req.Header
	}
	if h, ok := t.(http.Header); ok {
		return addRequestId(h)
	}
	return make(http.Header)
}

func addRequestId(h http.Header) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	id := h.Get(XRequestId)
	if len(id) == 0 {
		uid, _ := uuid.NewUUID()
		id = uid.String()
		h.Set(XRequestId, id)
	}
	return h
}

// RequestId - return a request id from any type and will create a new one if not found
func RequestId(t any) string {
	if t == nil {
		//s, _ := uuid.NewUUID()
		return "" // s.String()
	}
	requestId := ""
	switch ptr := t.(type) {
	case string:
		requestId = ptr
	case *http.Request:
		requestId = ptr.Header.Get(XRequestId)
	case http.Header:
		requestId = ptr.Get(XRequestId)
	}
	//if len(requestId) == 0 {
	//	s, _ := uuid.NewUUID()
	//	requestId = s.String()
	//}
	return requestId
}

// ValidateRequest - validate the request given an embedded URN path
func ValidateRequest(req *http.Request, path string) (string, *core.Status) {
	if req == nil {
		return "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: Request is nil"))
	}
	reqNid, reqPath, ok := uri.UprootUrn(req.URL.Path)
	if !ok {
		return "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, path is not valid: \"%v\"", req.URL.Path)))
	}
	if reqNid != path {
		return "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, NID does not match: \"%v\" \"%v\"", req.URL.Path, path)))
	}
	return reqPath, core.StatusOK()
}
