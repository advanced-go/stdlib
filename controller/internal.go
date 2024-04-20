package controller

import (
	"context"
	"github.com/advaced-go/stdlib/core"
	"github.com/advaced-go/stdlib/shttp"
	"net/http"
	"time"
)

func doInternal(duration time.Duration, handler func(w http.ResponseWriter, r *http.Request), req *http.Request) (r2 *http.Request, resp *http.Response, status *core.Status) {
	w := controller.NewResponseWriter()
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
	resp.ContentLength = w.written
	return r2, resp, core.NewStatus(resp.StatusCode)
}
