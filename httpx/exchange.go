package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

var (
	exchangeProxy = core.NewExchangeProxy()
)

// registerExchange - add an authority and Http Exchange handler to the proxy
func registerExchange(authority string, handler core.HttpExchange) error {
	return exchangeProxy.Register(authority, handler)
}

// Exchange - process an HTTP call utilizing an Exchange
func Exchange(req *http.Request) (*http.Response, *core.Status) {
	ex := exchangeProxy.LookupByRequest(req)
	if ex != nil {
		return ex(req)
	}
	//ctrl, status := controller2.Lookup(req)
	//if status.OK() {
	//	return controller2.Exchange(req, Do, ctrl)
	//}
	return Do(req)
}
