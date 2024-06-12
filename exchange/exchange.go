package exchange

import (
	"errors"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

type Proxy struct {
	do core.HttpExchange
}

func NewProxy(do core.HttpExchange) *Proxy {
	p := new(Proxy)
	p.do = do
	return p
}

func (p *Proxy) Exchange(r *http.Request) (*http.Response, *core.Status) {
	if r == nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument : request is nil"))
	}
	ctrl, status := Lookup(r)
	if !status.OK() {
		return p.do(r)
	}
	var resp *http.Response
	var req *http.Request

	localDo := p.do
	traffic := access.EgressTraffic
	rsc := ctrl.Primary
	if rsc.Handler != nil {
		localDo = rsc.Handler
		traffic = access.InternalTraffic
	}
	inDuration, outDuration := durations(rsc, req)
	duration := time.Duration(0)
	reasonCode := ""
	newURL := rsc.BuildURL(req.URL)
	req.URL = newURL
	if req.URL != nil {
		req.Host = req.URL.Host
	}
	start := time.Now().UTC()
	from := r.Header.Get(core.XFrom)

	// if no timeout or an existing deadline and existing deadline is <= timeout, then use the existing request
	if outDuration == 0 || (inDuration > 0 && inDuration <= outDuration) {
		duration = inDuration * -1
		resp, status = localDo(req)
	} else {
		duration = outDuration
		if rsc.Handler != nil {
			resp, status = doInternal(outDuration, localDo, req)
		} else {
			resp, status = doEgress(outDuration, localDo, req)
		}
	}
	if resp != nil {
		if resp.StatusCode == http.StatusGatewayTimeout {
			reasonCode = access.TimeoutCode
		}
	} else {
		resp = &http.Response{StatusCode: status.HttpCode()}
	}
	access.Log(traffic, start, time.Since(start), req, resp, from, ctrl.RouteName, rsc.Name, duration, 0, 0, reasonCode)
	return resp, core.StatusOK()
}

func durations(rsc *Resource, req *http.Request) (in time.Duration, out time.Duration) {
	deadline, ok := req.Context().Deadline()
	if ok {
		in = time.Until(deadline) // * -1
	}
	if rsc.Duration > 0 {
		out = rsc.Duration
	}
	return
}
