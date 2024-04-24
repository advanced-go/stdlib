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
	handlerProxy = NewProxy()
	duration     time.Duration
	authHandler  core.HttpHandler
	okFunc       = func(code int) bool { return code == http.StatusOK }
)

func SetHostTimeout(d time.Duration) {
	duration = d
}

func SetAuthHandler(h core.HttpHandler, ok func(int) bool) {
	if h != nil {
		authHandler = h
		if ok != nil {
			okFunc = ok
		}
	}
}

// RegisterHandler - add a path and Http handler to the proxy
// TO DO : panic on duplicate handler and pattern combination
func RegisterHandler(path string, handler core.HttpHandler) error {
	if len(path) == 0 {
		return errors.New("error: path is empty")
	}
	if handler == nil {
		return errors.New(fmt.Sprintf("error: handler for path %v is nil", path))
	}
	h := handler
	if authHandler != nil {
		h = NewConditionalIntermediary(authHandler, handler, okFunc)
	}
	if duration > 0 {
		h = NewHostTimeoutIntermediary(duration, h)
	}
	err := handlerProxy.Register(path, h)
	return err
}

// HttpHandler - process an HTTP request
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil || w == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	nid, _, ok := uri.UprootUrn(r.URL.Path)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler := handlerProxy.LookupByNID(nid)
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
