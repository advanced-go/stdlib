package host

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"time"
)

var (
	exchangeProxy = NewProxy()
	duration      time.Duration
	authExchange  core.HttpExchange
	okFunc        = func(code int) bool { return code == http.StatusOK }
)

func SetHostTimeout(d time.Duration) {
	duration = d
}

func SetAuthExchange(h core.HttpExchange, ok func(int) bool) {
	if h != nil {
		authExchange = h
		if ok != nil {
			okFunc = ok
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
		h = NewConditionalIntermediary(authExchange, handler, okFunc)
	}
	if duration > 0 {
		h = NewHostTimeoutIntermediary(duration, h)
	}
	err := exchangeProxy.Register(path, h)
	return err
}

// HttpExchange - process an HTTP exchange
func HttpExchange(w http.ResponseWriter, r *http.Request) {
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
	handler(w, r)
}

func shutdownHost(msg *messaging.Message) error {
	//TO DO: authentication and implementation
	return nil
}
