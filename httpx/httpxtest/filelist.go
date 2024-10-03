package httpxtest

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"strings"
)

const (
	respExt = "-resp."
)

// FileList - note, req file name needs to have an extension
type FileList struct {
	Dir, Req, Resp string
}

func (f FileList) RequestPath() string {
	if !strings.Contains(f.Req, ".") {
		return fmt.Sprintf("error: request file name does not have a . extension : %v", f.Req)
	}
	return f.Dir + "/" + f.Req
}

func (f FileList) ResponsePath() string {
	if f.Resp != "" {
		return f.Dir + "/" + f.Resp
	}
	s := strings.Replace(f.Req, ".", respExt, 1)
	return f.Dir + "/" + s
}

func (f FileList) NewUrl(req *http.Request) string {
	scheme := "https"
	host := req.Host
	if strings.Contains(host, "localhost") {
		scheme = "http"
	}
	return scheme + "://" + host + req.URL.String()
}

func (f FileList) NewRequest(req *http.Request) (*http.Request, *core.Status) {
	r, err := http.NewRequest(req.Method, f.NewUrl(req), req.Body)
	if err != nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, err)
	}
	r.Header = req.Header
	return r, core.StatusOK()
}
