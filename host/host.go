package host

import (
	"context"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"time"
)

func hostExchange[E core.ErrorHandler](w http.ResponseWriter, r *http.Request, dur time.Duration, handler core.HttpExchange) {
	flags := ""
	var start time.Time
	var resp *http.Response
	var status *core.Status

	if dur > 0 {
		ctx, cancel := context.WithTimeout(r.Context(), dur)
		defer cancel()
		r2 := r.Clone(ctx)
		start = time.Now().UTC()
		resp, status = handler(r2)
	} else {
		start = time.Now().UTC()
		resp, status = handler(r)
	}
	if status.Code == http.StatusGatewayTimeout {
		flags = access.TimeoutFlag
	}
	resp.ContentLength = httpx.WriteResponse[E](w, resp.Header, resp.StatusCode, resp.Body)
	access.Log(access.IngressTraffic, start, time.Since(start), r, resp, RouteName, "", Milliseconds(dur), flags)
}
