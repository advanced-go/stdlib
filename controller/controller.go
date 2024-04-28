package controller

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

type Controller struct {
	RouteName string
	Router    *Router
}

func NewController(routeName string, primary, secondary *Resource) *Controller {
	c := new(Controller)
	c.RouteName = routeName
	c.Router = NewRouter(primary, secondary)
	return c
}

func (c *Controller) Do(do core.HttpExchange, req *http.Request) (resp *http.Response, status *core.Status) {
	if req == nil || do == nil {
		return &http.Response{StatusCode: http.StatusBadRequest}, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument : request is nil"))
	}
	traffic := access.EgressTraffic
	rsc := c.Router.RouteTo()
	if rsc.handler != nil {
		traffic = access.InternalTraffic
		do = rsc.handler
	}
	inDuration, outDuration := durations(rsc, req)
	duration := time.Duration(0)
	flags := ""
	newURL := rsc.BuildUri(req.URL)
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
		ctx, cancel := context.WithTimeout(req.Context(), outDuration)
		defer cancel()
		r2 := req.Clone(ctx)
		resp, status = do(r2)
		//req = r2
		/*
			if rsc.internal {
				req, resp, status = doInternal(duration, rsc.handler, req)
			} else {
				traffic = access.EgressTraffic
				if duration <= 0 {
					resp, status = do(req)
				} else {
					resp, status = doEgress(duration, do, req)
				}
			}
		*/
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
	access.Log(traffic, start, elapsed, req, resp, c.RouteName, rsc.Name, access.Milliseconds(duration), flags)
	return
}

func durations(rsc *Resource, req *http.Request) (in time.Duration, out time.Duration) {
	deadline, ok := req.Context().Deadline()
	if ok {
		in = time.Until(deadline) // * -1
	}
	if rsc.duration > 0 {
		out = rsc.duration
	}
	return
}
