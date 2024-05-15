package httpx

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

// Get - process an HTTP Get request
func Get(ctx context.Context, uri string, h http.Header) (resp *http.Response, status *core.Status) {
	if len(uri) == 0 {
		return serverErrorResponse(), core.NewStatusError(http.StatusBadRequest, errors.New("error: URI is empty"))
	}
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return serverErrorResponse(), core.NewStatusError(http.StatusBadRequest, err)
	}
	if h != nil {
		req.Header = h
	}
	// exchange.Do() will always return a non nil *http.Response
	resp, status = Do(req)
	if !status.OK() {
		status.AddLocation()
	}
	return
}

// GetExchange - process an HTTP Get using an exchange if available
func GetExchange(ctx context.Context, uri string, h http.Header) (resp *http.Response, status *core.Status) {
	if len(uri) == 0 {
		return serverErrorResponse(), core.NewStatusError(http.StatusBadRequest, errors.New("error: URI is empty"))
	}
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return serverErrorResponse(), core.NewStatusError(http.StatusBadRequest, err)
	}
	if h != nil {
		req.Header = h
	}
	// exchange.Do() will always return a non nil *http.Response
	resp, status = DoExchange(req)
	if !status.OK() {
		status.AddLocation()
	}
	return
}
