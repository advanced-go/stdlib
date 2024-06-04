package access

import (
	"context"
	"net/http"
	"time"
)

type Request interface {
	Url() string
	Header() http.Header
	Method() string
	RouteName() string
	Duration() time.Duration
	ContextTimeout(ctx context.Context) (context.Context, context.CancelFunc)
}

type request struct {
	url       string
	header    http.Header
	method    string
	duration  time.Duration
	routeName string
}

func NewRequest(method, url string, h http.Header, routeName string, duration time.Duration) Request {
	r := new(request)
	r.url = url
	r.method = method
	r.header = h
	r.routeName = routeName
	r.duration = duration
	return r
}

func (r *request) Url() string {
	return r.url
}

func (r *request) Method() string {
	return r.method
}

func (r *request) RouteName() string {
	return r.routeName
}

func (r *request) Duration() time.Duration {
	return r.duration
}

func (r *request) Header() http.Header {
	return r.header
}

func (r *request) ContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	} else {
		if _, ok := ctx.Deadline(); ok {
			return ctx, cancel
		}
	}
	if r.duration == 0 {
		return ctx, cancel
	}
	return context.WithTimeout(ctx, r.duration)
}

func cancel() {}
