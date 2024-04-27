package controller

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func doInternal(duration time.Duration, handler core.HttpExchange, req *http.Request) (r2 *http.Request, resp *http.Response, status *core.Status) {
	//w := NewResponseWriter()
	if duration > 0 {
		ctx, cancel := context.WithTimeout(req.Context(), duration)
		defer cancel()
		r2 = req.Clone(ctx)
		resp, status = handler(r2)
	} else {
		r2 = req
		resp, status = handler(req)
	}
	//resp = w.Response()
	//resp.ContentLength = 0
	return r2, resp, status //core.NewStatus(resp.StatusCode)
}
