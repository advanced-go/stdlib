package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

type ResourceMapFunc func(req *http.Request) string

type Authority struct {
	Exchanges   map[string]core.HttpExchangeable
	Identity    *http.Response
	ResourceMap ResourceMapFunc
}

func NewAuthority(authority string, mapFn ResourceMapFunc, exchanges ...core.HttpExchangeable) *Authority {
	a := new(Authority)
	a.Identity = NewAuthorityResponse(authority)
	a.Exchanges = make(map[string]core.HttpExchangeable)
	a.ResourceMap = mapFn
	for _, ex := range exchanges {
		//a.Exchanges[ex]
	}
	return a
}

func (a *Authority) Do(req *http.Request) *http.Response {

	if req.Method == http.MethodGet && req.URL.Path == core.AuthorityRootPath {
		return a.Identity
	}

	return nil
}
