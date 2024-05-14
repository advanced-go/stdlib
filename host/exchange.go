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

// PathHandler - struct of path and associated Exchange handler
type PathHandler struct {
	Path    string
	Handler core.HttpExchange
}

var (
	exchangeProxy = NewProxy()
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

// RegisterExchange - add a path and Http handler to the proxy
// TO DO : panic on duplicate handler and pattern combination
func RegisterExchange(path string, handler core.HttpExchange) error {
	if len(path) == 0 {
		return errors.New("error: path is empty")
	}
	if handler == nil {
		return errors.New(fmt.Sprintf("error: handler for path [%v] is nil", path))
	}
	h := handler
	if authExchange != nil {
		h = NewConditionalIntermediary(authExchange, handler, okFunc)
	}
	return exchangeProxy.Register(path, h)
}

// RegisterAuthority - add an authority the proxy
func RegisterAuthority(authority []PathHandler) error {
	if len(authority) == 0 {
		return errors.New("error: authority configuration list is empty")
	}
	for _, config := range authority {
		err := RegisterExchange(config.Path, config.Handler)
		if err != nil {
			return err
		}
	}
	return nil
}

// HttpHandler - process an HTTP request
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r == nil || w == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	p := uri.Uproot(r.URL.Path)
	if !p.Valid {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler := exchangeProxy.LookupByNID(p.Authority)
	if handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hostExchange[core.Log](w, r, hostDuration, handler)
}

func shutdownHost(msg *messaging.Message) error {
	//TO DO: authentication and implementation
	return nil
}
