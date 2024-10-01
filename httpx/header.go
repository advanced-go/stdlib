package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"strings"
)

const (
	ContentTypeJson          = "application/json"
	ContentType              = "Content-Type"
	ContentEncoding          = "Content-Encoding"
	AcceptEncoding           = "Accept-Encoding"
	AcceptEncodingValue      = "gzip, deflate, br"
	ContentLength            = "Content-Length"
	ContentEncodingGzip      = "gzip"
	ContentTypeTextHtml      = "text/html"
	ContentTypeText          = "text/plain charset=utf-8"
	ContentLocation          = "Content-Location"
	ContentLocationExchange  = "X-Content-Location-Exchange"
	ContentLocationResolver  = "X-Content-Location-Resolver"
	ContentLocationSeparator = "->"
)

func forwardDefaults(dest http.Header, src http.Header) http.Header {
	if dest == nil {
		dest = make(http.Header)
	}
	if src == nil {
		return dest
	}
	// TODO : add other default headers
	dest.Set(core.XRequestId, src.Get(core.XRequestId))
	dest.Set(core.XRelatesTo, src.Get(core.XRelatesTo))
	dest.Set(core.XFrom, src.Get(core.XFrom))
	// Verify
	//dest.Set(core.XTest, src.Get(core.XTest))
	return dest
}

// Forward - forward headers
func Forward(dest http.Header, src http.Header, names ...string) http.Header {
	dest = forwardDefaults(dest, src)
	if src == nil {
		return dest
	}
	for _, name := range names {
		dest.Set(name, src.Get(name))
	}
	return dest
}

// GetContentType - return the content type header value
func GetContentType(headers any) string {
	if pairs, ok := headers.([]Attr); ok {
		for _, pair := range pairs {
			if pair.Key == ContentType {
				return pair.Value
			}
		}
		return ""
	}
	if h, ok := headers.(http.Header); ok {
		return h.Get(ContentType)
	}
	return ""
}

// SetHeaders - set the headers into an HTTP response writer
func SetHeaders(w http.ResponseWriter, headers any) {
	if headers == nil {
		return
	}
	if pairs, ok := headers.([]Attr); ok {
		for _, pair := range pairs {
			w.Header().Set(strings.ToLower(pair.Key), pair.Value)
		}
		return
	}
	if h, ok := headers.(http.Header); ok {
		for k, v := range h {
			if len(v) > 0 {
				w.Header().Set(strings.ToLower(k), v[0])
			}
		}
	}
}

func AddHeader(h http.Header, name, value string) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	h.Add(name, value)
	return h
}

func SetHeader(h http.Header, name, value string) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	h.Set(name, value)
	return h
}
