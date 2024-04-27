package host

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"time"
)

var (
	exchangeProxy = NewProxy2()
	hostDuration  time.Duration
	authExchange  core.HttpExchange
	okFunc2       = func(code int) bool { return code == http.StatusOK }
)

func SetHostTimeout2(d time.Duration) {
	hostDuration = d
}

func SetAuthExchange(h core.HttpExchange, ok func(int) bool) {
	if h != nil {
		authExchange = h
		if ok != nil {
			okFunc2 = ok
		}
	}
}

// RegisterExchange - add a path and Http handler to the proxy
// TO DO : panic on duplicate handler and pattern combination
func RegisterExchange(path string, handler core.HttpExchange) error {
	if len(path) == 0 {
		return errors.New("error: path is empty")
	}
	if handler == nil {
		return errors.New(fmt.Sprintf("error: handler for path %v is nil", path))
	}
	h := handler
	if authExchange != nil {
		h = NewConditionalIntermediary2(authExchange, handler, okFunc)
	}
	return exchangeProxy.Register(path, h)

}

// HttpHandler2 - process an HTTP request
func HttpHandler2(w http.ResponseWriter, r *http.Request) {
	if r == nil || w == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	nid, _, ok := uri.UprootUrn(r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler := exchangeProxy.LookupByNID(nid)
	if handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//resp, status := handler(r)
	//httpx.WriteResponse[core.Log](w, resp.Header, status.HttpCode(), resp.Body)
	hostExchange[core.Log](w, r, hostDuration, handler)
}
