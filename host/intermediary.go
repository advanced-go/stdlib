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
	XRequestId    = "X-Request-Id"
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

func NewAccessLogIntermediary(routeName string, c2 core.HttpExchange) core.HttpExchange {
	return func(r *http.Request) (resp *http.Response, status *core.Status) {
		if c2 == nil {
			return badRequest("error: AccessLog Intermediary HttpExchange is nil")
		}
		flags := ""
		authority := ""
		var dur time.Duration
		if ct, ok := r.Context().Deadline(); ok {
			dur = time.Until(ct) * -1
		}
		start := time.Now().UTC()
		resp, status = c2(r)
		if status.Code == http.StatusGatewayTimeout {
			flags = access.TimeoutFlag
		}
		authority, _ = core.Authority(c2)
		access.Log(access.InternalTraffic, start, time.Since(start), r, resp, authority, routeName, "", access.Milliseconds(dur), flags)
		return
	}
}
