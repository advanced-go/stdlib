package host

import (
	"context"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"time"
)

const (
	HostRoute = "host"
	EtcRoute  = "etc"
)

func hostExchange[E core.ErrorHandler](w http.ResponseWriter, r *http.Request, dur time.Duration, handler core.HttpExchange) {
	controllerCode := ""
	var start time.Time
	var resp *http.Response
	var status *core.Status

	core.AddRequestId(r)
	from := r.Header.Get(core.XFrom)
	if from == "" {
		r.Header.Set(core.XFrom, HostRoute)
	}
	if dur > 0 {
		ctx, cancel := context.WithTimeout(r.Context(), dur)
		defer cancel()
		r2 := r.Clone(ctx)
		start = time.Now().UTC()
		r2.Header.Set(core.XFrom, HostRoute)
		resp, status = handler(r2)
	} else {
		start = time.Now().UTC()
		r.Header.Set(core.XFrom, HostRoute)
		resp, status = handler(r)
	}
	if status.Code == http.StatusGatewayTimeout {
		controllerCode = access.ControllerTimeout
	}
	resp.ContentLength = httpx.WriteResponse[E](w, resp.Header, resp.StatusCode, resp.Body, r.Header)
	r.Header.Set(core.XTo, HostRoute)
	access.Log(access.IngressTraffic, start, time.Since(start), r, resp, access.Routing{From: from, Route: HostRoute, To: "", Percent: -1}, access.Controller{Timeout: dur, RateLimit: 0, RateBurst: 0, Code: controllerCode})
}
