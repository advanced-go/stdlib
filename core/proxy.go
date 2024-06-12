package core

import (
	"errors"
	"fmt"
	uri2 "github.com/advanced-go/stdlib/uri"
	"net/http"
	"sync"
)

// ExchangeProxy - key value pairs of an authority -> HttpExchange
type ExchangeProxy struct {
	m *sync.Map
}

// NewExchangeProxy - create a new Exchange Proxy
func NewExchangeProxy() *ExchangeProxy {
	p := new(ExchangeProxy)
	p.m = new(sync.Map)
	return p
}

func (p *ExchangeProxy) Register(authority string, handler HttpExchange) error {
	if len(authority) == 0 {
		return errors.New("invalid argument: authority is empty")
	}
	if handler == nil {
		return errors.New(fmt.Sprintf("invalid argument: HTTP Exchange is nil for authority : [%v]", authority))
	}
	_, ok1 := p.m.Load(authority)
	if ok1 {
		return errors.New(fmt.Sprintf("invalid argument: HTTP Exchange already exists for authority : [%v]", authority))
	}
	p.m.Store(authority, handler)
	return nil
}

// LookupByRequest - find an HttpExchange from a request
func (p *ExchangeProxy) LookupByRequest(req *http.Request) HttpExchange {
	if req == nil || req.URL == nil {
		return nil
	}
	// Try host first
	ex := p.Lookup(req.Host)
	if ex != nil {
		return ex
	}

	// Default to embedded authority
	parsed := uri2.Uproot(req.URL.Path)
	if parsed.Valid {
		ex = p.Lookup(parsed.Authority)
	}
	return ex
}

// Lookup - get an HttpExchange from the proxy, using an authority as a key
func (p *ExchangeProxy) Lookup(authority string) HttpExchange {
	v, ok := p.m.Load(authority)
	if !ok {
		return nil //, errors.New(fmt.Sprintf("error: proxyLookupByauthority() HTTP handler does not exist: [%v]", authority))
	}
	if handler, ok1 := v.(HttpExchange); ok1 {
		return handler
	}
	return nil
}
