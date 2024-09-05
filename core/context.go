package core

import (
	"context"
)

const (
	ExchangeRequestKey  = "request"
	ExchangeResponseKey = "response"
	ExchangeStatusKey   = "status"
)

type ExchangeOverride struct {
	m map[string]string
}

func NewExchangeOverrideEmpty() *ExchangeOverride {
	e := new(ExchangeOverride)
	e.m = make(map[string]string)
	return e
}

func NewExchangeOverride(request, response, status string) *ExchangeOverride {
	e := NewExchangeOverrideEmpty()
	if request != "" {
		e.m[ExchangeRequestKey] = request
	}
	if response != "" {
		e.m[ExchangeResponseKey] = response
	}
	if status != "" {
		e.m[ExchangeStatusKey] = status
	}
	return e
}

func (e *ExchangeOverride) Request() string {
	return e.m[ExchangeRequestKey]
}

func (e *ExchangeOverride) SetRequest(s string) {
	e.m[ExchangeRequestKey] = s
}

func (e *ExchangeOverride) Response() string {
	return e.m[ExchangeResponseKey]
}

func (e *ExchangeOverride) SetResponse(s string) {
	e.m[ExchangeResponseKey] = s
}

func (e *ExchangeOverride) Status() string {
	return e.m[ExchangeStatusKey]
}

func (e *ExchangeOverride) SetStatus(s string) {
	e.m[ExchangeStatusKey] = s
}

// type urlContextKey struct{}
type urlExchangeOverrideContextKey struct{}

var (
	//	urlKey         = urlContextKey{}
	urlExchangeOverrideKey = urlExchangeOverrideContextKey{}
)

// NewUrlContext - creates a new Context with an url
/*
func NewUrlContext(ctx context.Context, url string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		i := ctx.Value(urlKey)
		if i != nil {
			return ctx
		}
	}
	return context.WithValue(ctx, urlKey, url)
}

// UrlFromContext - return the url from a context
func UrlFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v := ctx.Value(urlKey)
	if v != nil {
		if url, ok := v.(string); ok {
			return url
		}
	}
	return ""
}


*/

// NewExchangeOverrideContext - creates a new Context with an exchange override
func NewExchangeOverrideContext(ctx context.Context, ex *ExchangeOverride) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		i := ctx.Value(urlExchangeOverrideKey)
		if i != nil {
			return ctx
		}
	}
	return context.WithValue(ctx, urlExchangeOverrideKey, ex)
}

// ExchangeOverrideFromContext - return an exchange override
func ExchangeOverrideFromContext(ctx context.Context) *ExchangeOverride {
	if ctx == nil {
		return nil
	}
	v := ctx.Value(urlExchangeOverrideKey)
	if v != nil {
		if url, ok := v.(*ExchangeOverride); ok {
			return url
		}
	}
	return nil
}
