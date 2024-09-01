package core

import (
	"context"
)

const (
	requestKey  = "request"
	responseKey = "response"
	statusKey   = "status"
)

type Exchange struct {
	m map[string]string
}

func NewExchange(request, response, status string) *Exchange {
	e := new(Exchange)
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

func (e *Exchange) Request() string {
	return e.m[requestKey]
}

func (e *Exchange) Response() string {
	return e.m[responseKey]
}

func (e *Exchange) Status() string {
	return e.m[statusKey]
}

type urlContextKey struct{}
type urlExchangeContextKey struct{}

var (
	urlKey         = urlContextKey{}
	urlExchangeKey = urlExchangeContextKey{}
)

// NewUrlContext - creates a new Context with an url
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

// NewExchangeContext - creates a new Context with an exchange
func NewExchangeContext(ctx context.Context, ex *Exchange) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		i := ctx.Value(urlExchangeKey)
		if i != nil {
			return ctx
		}
	}
	return context.WithValue(ctx, urlExchangeKey, ex)
}

// ExchangeFromContext - return an exchange
func ExchangeFromContext(ctx context.Context) *Exchange {
	if ctx == nil {
		return nil
	}
	v := ctx.Value(urlExchangeKey)
	if v != nil {
		if url, ok := v.(*Exchange); ok {
			return url
		}
	}
	return nil
}
