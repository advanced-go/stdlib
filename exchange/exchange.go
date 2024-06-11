package exchange

import (
	"errors"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"time"
)

func Exchange(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument : request is nil"))
	}
	ctrl, status := Lookup(r)

	if !status.OK() {
		return httpx.Do(r)
	}
	var resp *http.Response
	var req *http.Request
	var duration time.Duration

	traffic := access.EgressTraffic
	rsc := ctrl.Primary
	if rsc.Handler != nil {
		traffic = access.InternalTraffic
	}
	reasonCode := ""
	start := time.Now().UTC()

	from := r.Header.Get(core.XFrom)

	access.Log(traffic, start, time.Since(start), req, resp, from, ctrl.RouteName, rsc.Name, duration, 0, 0, reasonCode)
	return resp, core.StatusOK()
}
