package core

import (
	"context"
)

const (
	requestKey  = "request"
	responseKey = "response"
	statusKey   = "status"
)

type ExchangeMap struct {
	m map[string]string
}

func NewExchangeMap(request, response, status string) *ExchangeMap {
	e := new(ExchangeMap)
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

func (e *ExchangeMap) Request() string {
	return e.m[requestKey]
}

func (e *ExchangeMap) Response() string {
	return e.m[responseKey]
}

func (e *ExchangeMap) Status() string {
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

// NewExchangeMapContext - creates a new Context with an exchange map
func NewExchangeMapContext(ctx context.Context, ex *ExchangeMap) context.Context {
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

// ExchangeMapFromContext - return a url map from a context
func ExchangeMapFromContext(ctx context.Context) *ExchangeMap {
	if ctx == nil {
		return nil
	}
	v := ctx.Value(urlExchangeKey)
	if v != nil {
		if url, ok := v.(*ExchangeMap); ok {
			return url
		}
	}
	return nil
}
