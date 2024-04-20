package controller

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func doInternal(duration time.Duration, handler func(w http.ResponseWriter, r *http.Request), req *http.Request) (r2 *http.Request, resp *http.Response, status *core.Status) {
	w := NewResponseWriter()
	if duration > 0 {
		ctx, cancel := context.WithTimeout(req.Context(), duration)
		defer cancel()
		r2 = req.Clone(ctx)
		handler(w, r2)
	} else {
		r2 = req
		handler(w, req)
	}
	resp = w.Response()
	resp.ContentLength = w.Written()
	return r2, resp, core.NewStatus(resp.StatusCode)
}
