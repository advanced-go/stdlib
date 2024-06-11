package httpx

import (
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

// Exchange - process an HTTP call utilizing a controller if configured
func Exchange(req *http.Request) (*http.Response, *core.Status) {
	//req, _ := http.NewRequestWithContext(ctx, method, url, body)
	//if h != nil {
	//	req.Header = h
	//}
	ctrl, status := controller.Lookup(req)
	if status.OK() {
		return ctrl.Do(Do, req)
	}
	return Do(req)
}
