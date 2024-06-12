package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

var (
	exchangeProxy = core.NewExchangeProxy()
)

// RegisterExchange - add an authority and Http Exchange handler to the proxy
func RegisterExchange(authority string, handler core.HttpExchange) error {
	return exchangeProxy.Register(authority, handler)
}

// Exchange - process an HTTP call utilizing an Exchange
func Exchange(req *http.Request) (*http.Response, *core.Status) {
	ex := exchangeProxy.LookupByRequest(req)
	if ex != nil {
		return ex(req)
	}
	return Do(req)
}
