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

func NewHost(authority string, mapFn ResourceMapFunc, exchanges ...core.HttpExchange) (*Host, error) {
	if authority == "" {
		return nil, errors.New("error: authority is empty")
	}
	if mapFn == nil {
		return nil, errors.New("resource map function is nil")
	}
	a := new(Host)
	a.Identity = NewAuthorityResponse(authority)
	a.Exchanges = make(map[string]core.HttpExchange)
	a.ResourceMap = mapFn
	for _, ex := range exchanges {
		err := a.AddExchange(ex)
		if err != nil {
			return a, err
		}
	}
	return a, nil
}

func (a *Host) AddExchange(ex core.HttpExchange) error {
	name := core.Authority(ex)
	if name == "" {
		return errors.New(fmt.Sprintf("error: invalid resource map, resource name is empty"))
	}
	if _, ok := a.Exchanges[name]; ok {
		return errors.New(fmt.Sprintf("error: invalid resource name, Exchange already exists for: %v", name))
	}
	a.Exchanges[name] = ex
	return nil
}

func (a *Host) Exchange(name string) (core.HttpExchange, error) {
	if name == "" {
		return nil, errors.New(fmt.Sprintf("invalid arguement: resource name is empty"))
	}
	if ex, ok := a.Exchanges[name]; ok {
		return ex, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid argument: Exchange not found for resource name: %v", name))
}

func (a *Host) Do(req *http.Request) (*http.Response, *core.Status) {
	if req == nil {
		return NewResponse(core.StatusBadRequest(), errors.New("bad request: http.Request is nil")), core.StatusBadRequest()
	}
	if req.Method == http.MethodGet && req.URL.Path == core.AuthorityRootPath {
		return a.Identity, core.StatusOK()
	}
	ex, err := a.Exchange(a.ResourceMap(req))
	if ex == nil {
		return NewResponse(core.StatusBadRequest(), errors.New(fmt.Sprintf("invalid resource map: %v, HttpExchange not found for: [%v]", err, req.URL))), core.StatusBadRequest()
	}
	return ex(req)
}

/*
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

*/
