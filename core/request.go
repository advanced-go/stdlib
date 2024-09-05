package core

import (
	"github.com/google/uuid"
	"net/http"
)

const (
	XRequestId = "X-Request-Id"
	XRelatesTo = "X-Relates-To"
	XAuthority = "X-Authority"
	XVersion   = "X-Version"
	XURLPath   = "x-url-path"
	XTest      = "X-Test"
	XFrom      = "X-From"
	XTo        = "X-To"
	XRoute     = "X-Route"
)

// AddRequestId - add a request to an http.Request or an http.Header
func AddRequestId(t any) http.Header {
	if t == nil {
		h := make(http.Header)
		return addRequestId(h)
	}
	if req, ok := t.(*http.Request); ok {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
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
