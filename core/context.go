package core

import (
	"context"
)

const (
	requestKey  = "request"
	responseKey = "response"
	statusKey   = "status"
)

type ExchangeOverride struct {
	m map[string]string
}

func NewExchangeOverride(request, response, status string) *ExchangeOverride {
	e := new(ExchangeOverride)
	e.m = make(map[string]string)
	if request != "" {
		e.m[requestKey] = request
	}
	if response != "" {
		e.m[responseKey] = response
	}
	if status != "" {
		e.m[statusKey] = status
	}
	return e
}

func (e *ExchangeOverride) Request() string {
	return e.m[requestKey]
}

func (e *ExchangeOverride) Response() string {
	return e.m[responseKey]
}

func (e *ExchangeOverride) Status() string {
	return e.m[statusKey]
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
