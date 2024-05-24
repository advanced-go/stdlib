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

// BuildURL - build a URL, creating the scheme and host, based on the given URL which should only be a path
func BuildURL(host string, uri *url.URL) *url.URL {
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

	if len(uri.Path) > 0 {
		newUri += uri.Path
	}
	if len(uri.RawQuery) > 0 {
		newUri += "?" + uri.RawQuery
	}
	/*
		if authority == "" {
			if len(uri.Path) > 0 {
				newUri += uri.Path
			}
			if len(uri.RawQuery) > 0 {
				newUri += "?" + uri.RawQuery
			}
		} else {
			//parsed := Uproot(uri.Path)
			//newUri += "/" + authority
			//if len(parsed.Path) > 0 {
			//	newUri += ":" + parsed.Path //uri.Path[1:]
			//}
			if len(uri.Path) > 0 {
				newUri += uri.Path
			}
			if len(uri.RawQuery) > 0 {
				newUri += "?" + uri.RawQuery
			}
		}

	*/
	u, err1 := url.Parse(newUri)
	if err1 != nil {
		return uri
	}
	return u
}
