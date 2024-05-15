package host

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"sync"
)

// Proxy - key value pairs of a URI -> HttpHandler
type Proxy struct {
	m *sync.Map
}

// NewProxy - create a new Proxy2
func NewProxy() *Proxy {
	p := new(Proxy)
	p.m = new(sync.Map)
	return p
}

// Register - add an HttpExchange to the proxy
func (p *Proxy) register(authority string, handler core.HttpExchange) error {
	if len(authority) == 0 {
		return errors.New("error: proxy.register() authority is empty")
	}
	if handler == nil {
		return errors.New(fmt.Sprintf("error: proxy.register() HTTP Exchange is nil: [%v]", authority))
	}
	//parsed := uri2.Uproot(authority)
	//if !parsed.Valid {
	//	return errors.New(fmt.Sprintf("error: proxy.register() authority is invalid: [%v] [%v]", authority, parsed.Err))
	//}
	_, ok1 := p.m.Load(authority)
	if ok1 {
		return errors.New(fmt.Sprintf("error: proxy.register() HTTP Exchange already exists: [%v]", authority))
	}
	p.m.Store(authority, handler)
	return nil
}

// Lookup - get an HttpExchange from the proxy, using a URI as the key
/*
func (p *Proxy) Lookup(authority string) core.HttpExchange {
	parsed := uri2.Uproot(authority)
	if !parsed.Valid {
		return nil //, errors.New(fmt.Sprintf("error: proxy.Lookup() URI is invalid: [%v]", authority))
	}
	return p.LookupByAuthority(parsed.Authority)
}


*/

// Lookup - get an HttpExchange from the proxy, using an authority as a key
func (p *Proxy) Lookup(authority string) core.HttpExchange {
	v, ok := p.m.Load(authority)
	if !ok {
		return nil //, errors.New(fmt.Sprintf("error: proxyLookupByauthority() HTTP handler does not exist: [%v]", authority))
	}
	if handler, ok1 := v.(core.HttpExchange); ok1 {
		return handler //, StatusOK()
	}
	return nil //, NewStatus(StatusInvalidContent)
}
