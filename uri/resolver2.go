package uri

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type HostKey string

const (
	intermediaryText = "intermediary"
	intermediaryKey  = HostKey(intermediaryText)
)

type HostEntry struct {
	Key          string `json:"key"`
	Host         string `json:"host"`
	Intermediary bool   `json:"intermediary"`
}

type Resolver struct {
	m map[HostKey]HostEntry
}

func NewResolver() *Resolver {
	r := new(Resolver)
	r.m = make(map[HostKey]HostEntry)
	return r
}

func (r *Resolver) Resolve(host any, authority, resourcePath string, values url.Values, h http.Header) string {
	path := BuildPath(authority, resourcePath, values)
	if h != nil {
		p2 := h.Get(path)
		if p2 != "" {
			return p2
		}
	}
	if host == nil {
		return path
	}
	if s, ok := host.(string); ok {
		if s == "" {
			return path
		}
		if path[0] == '/' {
			return BuildOrigin(s) + path
		}
		return BuildOrigin(s) + "/" + path
	}
	if k, ok := host.(HostKey); !ok {
		return fmt.Sprintf("error: invalid key type: %v", reflect.TypeOf(host))
	} else {
		entry, ok1 := r.entry(k)
		if ok1 {
			if path[0] == '/' {
				return BuildOrigin(entry.Host) + path
			}
			return BuildOrigin(entry.Host) + "/" + path
		}
	}
	return fmt.Sprintf("error: missing host entry for key: %v", host)
}

func (r *Resolver) entry(key HostKey) (HostEntry, bool) {
	entry, ok := r.m[key]
	if !ok {
		return HostEntry{}, false
	}
	if !entry.Intermediary {
		return entry, true
	}
	entry, ok = r.m[intermediaryKey]
	if ok {
		return entry, true
	}
	return HostEntry{}, false
}

/*
	//if host == "" {
		return path
	}
	//if path[0] == '/' {
	//	return BuildOrigin(host) + path
	//}
	//return BuildOrigin(host) + "/" + path


*/

func BuildPath(authority, resourcePath string, values url.Values) string {
	path := strings.Builder{}
	if authority != "" {
		path.WriteString(authority)
		path.WriteString(":")
	}
	//path.WriteString(formatVersion(version))
	path.WriteString(resourcePath)
	path.WriteString(formatValues(values))
	return path.String()
}

func BuildRsc(version, resource string) string {
	return formatVersion(version) + resource
}

func BuildOrigin(host string) string {
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
