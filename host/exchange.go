package host

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"time"
)

var (
	exchangeProxy = core.NewExchangeProxy()
	hostDuration  time.Duration
	authExchange  core.HttpExchange
	okFunc        = func(code int) bool { return code == http.StatusOK }
)

func SetHostTimeout(d time.Duration) {
	hostDuration = d
}

func SetAuthExchange(h core.HttpExchange, ok func(int) bool) {
	if h != nil {
		authExchange = h
		if ok != nil {
			okFunc = ok
		}
	}
}

// RegisterExchange - add an authority and Http Exchange handler to the proxy
func RegisterExchange(authority string, handler core.HttpExchange) error {
	h := handler
	if authExchange != nil {
		h = NewConditionalIntermediary(authExchange, handler, okFunc)
	}
	return exchangeProxy.Register(authority, h)
}

// HttpHandler - process an HTTP request
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil || w == nil || r.URL == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p := uri.Uproot(r.URL.Path)
	if !p.Valid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	handler := exchangeProxy.Lookup(p.Authority)
	if handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hostExchange[core.Log](w, r, hostDuration, handler)
}
