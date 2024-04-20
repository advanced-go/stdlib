package httpx

import (
	"context"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
)

func Exchange(ctx context.Context, method, url string, h http.Header, body io.Reader) (*http.Response, *core.Status) {
	req, _ := http.NewRequestWithContext(ctx, method, url, body)
	if h != nil {
		req.Header = h
	}
	ctrl, status1 := controller.Lookup(url)
	if status1.OK() {
		return ctrl.Do(Do, req)
	}
	return Do(req)
}
