package host

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/access"
	"github.com/advanced-go/stdlib/core"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func newIngressIntermediary(routeName string, d time.Duration, c2 core.HttpExchange, traffic string) core.HttpExchange {
	return func(r *http.Request) (*http.Response, *core.Status) {
		if c2 == nil {
			//w.WriteHeader(http.StatusInternalServerError)
			//fmt.Fprintf(w, "error: component 2 is nil")
			return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: HttpExchange is nil"))
		}
		if traffic == access.IngressTraffic {
			if r.Header.Get(XRequestId) == "" {
				uid, _ := uuid.NewUUID()
				r.Header.Add(XRequestId, uid.String())
			}
		}
		//w2 := newWrapper(w)
		return apply2(nil, r, routeName, d, c2, traffic, "")
	}
}

func apply2(w *wrapper, r *http.Request, routeName string, duration time.Duration, handler core.HttpExchange, traffic, routeTo string) (resp *http.Response, status *core.Status) {
	if handler == nil {
		return
	}
	flags := ""
	var start time.Time

	if ct, ok := r.Context().Deadline(); ok {
		duration = time.Until(ct) * -1
	}
	if duration > 0 {
		ctx, cancel := context.WithTimeout(r.Context(), duration)
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
	if traffic == "" {
		traffic = access.InternalTraffic
	}
	access.Log(traffic, start, time.Since(start), r, resp, routeName, routeTo, Milliseconds(duration), flags)
	return
}
