package host

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	uri2 "github.com/advanced-go/stdlib/uri"
	"sync"
)

// Proxy - key value pairs of a URI -> HttpHandler
type Proxy struct {
	m *sync.Map
}

// NewProxy - create a new Proxy
func NewProxy() *Proxy {
	p := new(Proxy)
	p.m = new(sync.Map)
	return p
}

// Register - add an HttpHandler to the proxy
func (p *Proxy) Register(uri string, handler core.HttpExchange) error {
	if len(uri) == 0 {
		return errors.New("error: proxy.Register() path is empty")
	}
	nid, _, ok := uri2.UprootUrn(uri)
	if !ok {
		return errors.New(fmt.Sprintf("error: proxy.Register() path is invalid: [%v]", uri))
	}
	if handler == nil {
		return errors.New(fmt.Sprintf("error: proxy.Register() HTTP handler is nil: [%v]", uri))
	}
	_, ok1 := p.m.Load(nid)
	if ok1 {
		return errors.New(fmt.Sprintf("error: proxy.Register() HTTP handler already exists: [%v]", uri))
	}
	p.m.Store(nid, handler)
	return nil
}

// Lookup - get an HttpHandler from the proxy, using a URI as the key
func (p *Proxy) Lookup(uri string) core.HttpExchange {
	nid, _, ok := uri2.UprootUrn(uri)
	if !ok {
		return nil //, errors.New(fmt.Sprintf("error: proxy.Lookup() URI is invalid: [%v]", uri))
	}
	return p.LookupByNID(nid)
}

// LookupByNID - get an HttpHandler from the proxy, using an NID as a key
func (p *Proxy) LookupByNID(nid string) core.HttpExchange {
	v, ok := p.m.Load(nid)
	if !ok {
		return nil //, errors.New(fmt.Sprintf("error: proxyLookupByNID() HTTP handler does not exist: [%v]", nid))
	}
	if handler, ok1 := v.(core.HttpExchange); ok1 {
		return handler //, StatusOK()
	}
	return nil //, NewStatus(StatusInvalidContent)
}
