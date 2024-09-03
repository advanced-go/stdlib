package uri

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const (
	proxyKey = "proxy-key"
)

/*
type HostEntry struct {
	Key   string `json:"key"`
	Host  string `json:"host"`
	Proxy bool   `json:"proxy"`
}

*/

type Resolver struct {
	defaultHost string
	//m           *sync.Map
	//proxy       HostEntry
	//entryFn     func(host string, m *sync.Map) (HostEntry, bool)
}

func NewResolver(defaultHost string) *Resolver {
	r := new(Resolver)
	r.defaultHost = defaultHost
	return r
}

/*
func NewResolver3(entries []HostEntry) *Resolver {
	r := new(Resolver)
	r.m = new(sync.Map)
	for _, e := range entries {
		r.m.Store(e.Key, e)
		if e.Key == proxyKey {
			r.proxy = e
			r.proxy.Proxy = false
		}
	}
	r.entryFn = func(host string, m *sync.Map) (HostEntry, bool) {
		e, ok := load(host, m)
		if !ok {
			return e, ok
		}
		if !e.Proxy {
			return e, ok
		}
		if r.proxy.Host != "" {
			return r.proxy, true
		}
		return HostEntry{}, false
	}
	return r
}
*/

/*
func (r *Resolver) Override(entries []HostEntry) *Resolver {
	r2 := NewResolver(entries)
	r2.entryFn = func(host string, m *sync.Map) (HostEntry, bool) {
		e, ok := load(host, m)
		if !ok {
			return r.entryFn(host, r.m)
		}
		if !e.Proxy {
			return e, ok
		}
		if r2.proxy.Host != "" {
			return r2.proxy, true
		}
		if r.proxy.Host != "" {
			return r.proxy, true
		}
		return HostEntry{}, false
	}
	return r2
}
func (r *Resolver) Host(host string) string {
	if host == "" {
		return "error: host is empty"
	}
	e, ok := r.entryFn(host, r.m)
	if ok {
		return e.Host
	}
	return host
}

*/

func (r *Resolver) Url(host, path string, query any, h http.Header) string {
	path1 := BuildPath(path, query)
	if h != nil {
		p2 := h.Get(path1)
		if p2 != "" {
			return p2
		}
	}
	if host != "" {
		return Cat(host, path1)
	}
	return Cat(r.defaultHost, path1)
}

func (r *Resolver) UrlWithAuthority(host, authority, version, resource string, query any, h http.Header) string {
	path := BuildPathWithAuthority(authority, version, resource, query)
	if h != nil {
		p2 := h.Get(path)
		if p2 != "" {
			return p2
		}
	}
	if host != "" {
		return Cat(host, path)
	}
	return Cat(r.defaultHost, path)
}

func Cat(host, path string) string {
	origin := BuildHostWithScheme(host)
	if path[0] == '/' {
		return origin + path
	}
	return origin + "/" + path
}

func BuildPath(path string, query any) string {
	return BuildPathWithAuthority("", "", path, query)
}

func BuildPathWithAuthority(authority, version, resource string, query any) string {
	path := strings.Builder{}
	if authority != "" {
		path.WriteString(authority)
		path.WriteString(":")
		path.WriteString(formatVersion2(version))
	}
	path.WriteString(resource)
	path.WriteString(formatQuery(query))
	return path.String()
}

func BuildHostWithScheme(host string) string {
	if host == "" {
		return ""
	}
	origin := strings.Builder{}
	scheme := HttpsScheme
	if strings.Contains(host, Localhost) || strings.Contains(host, Internalhost) {
		scheme = HttpScheme
	}
	origin.WriteString(scheme)
	origin.WriteString("://")
	origin.WriteString(host)
	return origin.String()
}

func formatQuery(query any) string {
	if query == nil {
		return ""
	}
	if v, ok := query.(url.Values); ok {
		encoded := v.Encode()
		if encoded != "" {
			encoded, _ = url.QueryUnescape(encoded)
			return "?" + encoded
		}
		return ""
	}
	if s, ok := query.(string); ok {
		return "?" + s
	}
	return fmt.Sprintf("error: query type is invalid %v", reflect.TypeOf(query))
}

func formatVersion2(version string) string {
	if version == "" {
		return ""
	}
	return version + "/"
}

/*
func entry(host string, m *sync.Map) (HostEntry, bool) {
	e, ok := load(host, m)
	if !ok {
		return e, ok
	}
	return e, true
}


*/

/*
func load(host string, m *sync.Map) (HostEntry, bool) {
	if m == nil {
		return HostEntry{}, false
	}
	value, ok := m.Load(host)
	if !ok {
		return HostEntry{}, false
	}
	if e, ok1 := value.(HostEntry); ok1 {
		return e, true
	}
	return HostEntry{}, false
}


*/
/*
newUrl := strings.Builder{}
if host != "" {
scheme := HttpsScheme
if strings.Contains(host, Localhost) {
scheme = HttpScheme
}
newUrl.WriteString(scheme)
newUrl.WriteString("://")
newUrl.WriteString(host)
}
newUrl.WriteString(fmt.Sprintf(path, formatVersion(version)))
newUrl.WriteString(formatValues(values))
return newUrl.String()
newUrl := strings.Builder{}
	if host != "" {
		scheme := httpsScheme
		if strings.Contains(host, localHost) {
			scheme = httpScheme
		}
		newUrl.WriteString(scheme)
		newUrl.WriteString("://")
		newUrl.WriteString(host)
	}
	newUrl.WriteString(authority)
	newUrl.WriteString(":")
	newUrl.WriteString(path)
	newUrl.WriteString(formatValues(values))
	return newUrl.String()
*/
