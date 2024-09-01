package core

import (
	"context"
)

const (
	ContextRequestKey  = "request"
	ContextResponseKey = "response"
)

type urlContextKey struct{}
type urlMapContextKey struct{}

var (
	urlKey    = urlContextKey{}
	urlMapKey = urlMapContextKey{}
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

// NewUrlMapContext - creates a new Context with an url map
func NewUrlMapContext(ctx context.Context, url map[string]string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		i := ctx.Value(urlMapKey)
		if i != nil {
			return ctx
		}
	}
	return context.WithValue(ctx, urlMapKey, url)
}

// UrlMapFromContext - return a url map from a context
func UrlMapFromContext(ctx context.Context) map[string]string {
	if ctx == nil {
		return nil
	}
	v := ctx.Value(urlMapKey)
	if v != nil {
		if url, ok := v.(map[string]string); ok {
			return url
		}
	}
	return nil
}
