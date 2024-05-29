package controller

import (
	"errors"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

type Controller struct {
	Name   string
	Router *Router
}

func NewController(routeName string, primary, secondary *Resource) *Controller {
	c := new(Controller)
	c.Name = routeName
	c.Router = NewRouter(primary, secondary)
	return c
}

func (c *Controller) RouteName() string {
	return c.Name
}

func (c *Controller) Do(do core.HttpExchange, req *http.Request) (resp *http.Response, status *core.Status) {
	if req == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument : request is nil"))
	}
	authority := ""
	traffic := access.EgressTraffic
	rsc := c.Router.RouteTo()
	if rsc.Handler != nil {
		traffic = access.InternalTraffic
		do = rsc.Handler
		authority = core.Authority(do)
	} else {
		if do == nil {
			return &http.Response{StatusCode: http.StatusBadRequest}, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument : core.HttpExchange is nil"))
		}
	}
	inDuration, outDuration := durations(rsc, req)
	duration := time.Duration(0)
	flags := ""
	newURL := rsc.BuildURL(req.URL)
	req.URL = newURL
	if req.URL != nil {
		req.Host = req.URL.Host
	}
	start := time.Now().UTC()

	// if no timeout or an existing deadline and existing deadline is <= timeout, then use the existing request
	if outDuration == 0 || (inDuration > 0 && inDuration <= outDuration) {
		duration = inDuration * -1
		resp, status = do(req)
	} else {
		duration = outDuration
		// Internal call
		if rsc.Handler != nil {
			//ctx, cancel := context.WithTimeout(req.Context(), outDuration)
			//defer cancel()
			//r2 := req.Clone(ctx)
			resp, status = doInternal(outDuration, do, req)
		} else {
			resp, status = doEgress(outDuration, do, req)
		}
	}
	elapsed := time.Since(start)
	if resp != nil {
		c.Router.UpdateStats(resp.StatusCode, rsc)
		if resp.StatusCode == http.StatusGatewayTimeout {
			flags = access.TimeoutFlag
		}
	} else {
		resp = &http.Response{StatusCode: status.HttpCode()}
	}
	if !rsc.SilentAccess {
		access.Log(traffic, start, elapsed, req, resp, authority, c.RouteName(), rsc.Name, access.Milliseconds(duration), flags)
	}
	return
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
