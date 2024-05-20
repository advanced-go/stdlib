package uri

import (
	"net/url"
	"strings"
)

const (
	HttpScheme  = "http"
	HttpsScheme = "https"
	Localhost   = "localhost"
)

func BuildURL(host, authority string, uri *url.URL) *url.URL {
	if uri == nil {
		return uri
	}
	scheme := HttpsScheme
	if host == "" {
		host = Localhost
	}
	if strings.Contains(host, Localhost) {
		scheme = HttpScheme
	}
	var newUri = scheme + "://" + host
	if authority == "" {
		if len(uri.Path) > 0 {
			newUri += uri.Path
		}
		if len(uri.RawQuery) > 0 {
			newUri += "?" + uri.RawQuery
		}
	} else {
		newUri += "/" + authority
		if len(uri.Path) > 0 {
			newUri += ":" + uri.Path[1:]
		}
		if len(uri.RawQuery) > 0 {
			newUri += "?" + uri.RawQuery
		}
	}
	u, err1 := url.Parse(newUri)
	if err1 != nil {
		return uri
	}
	return u
}