package core

import (
	"context"
)

type urlContextKey struct{}

var (
	urlKey = urlContextKey{}
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
