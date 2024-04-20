package host

import (
	"fmt"
	"github.com/advanced-go/stdlib/access"
	"net/http"
	"time"
)

const (
	Authorization = "Authorization"
	TimeoutFlag   = "TO"
	XRequestId    = "X-Request-Id"
)

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request)

func NewConditionalIntermediary(c1 HttpHandlerFunc, c2 HttpHandlerFunc, ok func(int) bool) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if c2 == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: component 2 is nil")
			return
		}
		w2 := newWrapper(w)
		if c1 != nil {
			c1(w2, r)
		}
		if (ok == nil && w2.statusCode == http.StatusOK) || (ok != nil && ok(w2.statusCode)) {
			c2(w, r)
		}
	}
}

func NewAccessLogIntermediary(routeName string, c2 HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if c2 == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: component 2 is nil")
			return
		}
		w2 := newWrapper(w)
		flags := ""
		var dur time.Duration
		if ct, ok := r.Context().Deadline(); ok {
			dur = time.Until(ct) * -1
		}
		start := time.Now().UTC()
		c2(w2, r)
		if w2.statusCode == http.StatusGatewayTimeout {
			flags = TimeoutFlag
		}
		access.Log(access.InternalTraffic, start, time.Since(start), r, &http.Response{StatusCode: w2.statusCode, ContentLength: w2.written}, routeName, "", Milliseconds(dur), flags)
	}
}
