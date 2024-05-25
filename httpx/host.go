package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

type ResourceMapFunc func(req *http.Request) string

type Host struct {
	Exchanges   map[string]core.HttpExchange
	Identity    *http.Response
	ResourceMap ResourceMapFunc
}

func NewHost(authority string, mapFn ResourceMapFunc, exchanges ...core.HttpExchange) *Host {
	a := new(Host)
	a.Identity = NewAuthorityResponse(authority)
	a.Exchanges = make(map[string]core.HttpExchange)
	a.ResourceMap = mapFn
	if a.ResourceMap == nil {
		a.ResourceMap = func(req *http.Request) string { return "" }
	}
	for _, ex := range exchanges {
		name := core.Authority(ex)
		a.Exchanges[name] = ex
	}
	return a
}

func (a *Host) Do(req *http.Request) *http.Response {
	if req == nil {
		return NewResponse(core.StatusBadRequest(), errors.New("bad request: http.Request is nil"))
	}
	if req.Method == http.MethodGet && req.URL.Path == core.AuthorityRootPath {
		return a.Identity
	}
	name := a.ResourceMap(req)
	if name == "" {
		return NewResponse(core.StatusBadRequest(), errors.New(fmt.Sprintf("invalid resource map, resource name is empty for: [%v]", req.URL)))
	}
	if ex, ok := a.Exchanges[name]; ok {
		resp, _ := ex(req)
		return resp
	}
	return NewResponse(core.StatusBadRequest(), errors.New(fmt.Sprintf("invalid resource map, HttpExchange not found for: [%v]", req.URL)))
}
