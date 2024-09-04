package host

import (
	"errors"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

const (
	Authorization = "Authorization"
)

func badRequest(msg string) (*http.Response, *core.Status) {
	return &http.Response{StatusCode: http.StatusBadRequest}, core.NewStatusError(http.StatusBadRequest, errors.New(msg))
}

func NewConditionalIntermediary(c1 core.HttpExchange, c2 core.HttpExchange, ok func(int) bool) core.HttpExchange {
	return func(r *http.Request) (resp *http.Response, status *core.Status) {
		if c1 == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 1 is nil")
		}
		if c2 == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 2 is nil")
		}
		resp, status = c1(r)
		if resp == nil {
			return badRequest("error: Conditional Intermediary HttpExchange 1 response is nil")
		}
		if (ok == nil && resp.StatusCode == http.StatusOK) || (ok != nil && ok(resp.StatusCode)) {
			resp, status = c2(r)
		}
		return
	}
}

func NewAccessLogIntermediary(c2 core.HttpExchange) core.HttpExchange {
	return func(r *http.Request) (resp *http.Response, status *core.Status) {
		if c2 == nil {
			return badRequest("error: AccessLog Intermediary HttpExchange is nil")
		}
		reasonCode := ""
		from := r.Header.Get(core.XFrom)

		var dur time.Duration
		if ct, ok := r.Context().Deadline(); ok {
			dur = time.Until(ct) * -1
		}
		start := time.Now().UTC()
		resp, status = c2(r)
		if status.Code == http.StatusGatewayTimeout {
			reasonCode = access.ControllerTimeout
		}
		access.Log(access.InternalTraffic, start, time.Since(start), r, resp, access.Routing{From: from, Route: resp.Header.Get(core.XRoute), To: "", Percent: -1}, access.Controller{Timeout: dur, RateLimit: 0, RateBurst: 0, Code: reasonCode})
		return
	}
}
