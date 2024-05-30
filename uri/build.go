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

// BuildURL - build an url with the components provided, escaping the query
// TODO : escaping on path ?? url.PathEscape
func BuildURL(host, version, path string, query any) string {
	newUrl := strings.Builder{}
	scheme := HttpsScheme
	if host == "" {
		host = Localhost
	}
	if strings.Contains(host, Localhost) {
		scheme = HttpScheme
	}
	newUrl.WriteString(scheme)
	newUrl.WriteString("://")
	newUrl.WriteString(host)
	if len(path) > 0 {
		if path[:1] != "/" {
			path += "/"
		}
	}
	if version != "" {
		newUrl.WriteString("/")
		newUrl.WriteString(version)
	}
	newUrl.WriteString(path)
	q := BuildQuery(query)
	if q != "" {
		newUrl.WriteString("?")
		newUrl.WriteString(q)
	}
	return newUrl.String()
}

// BuildQuery - build a query string with escaping
func BuildQuery(query any) string {
	if query == nil {
		return ""
	}
	if s, ok := query.(string); ok {
		return url.QueryEscape(s)
	}
	if v, ok := query.(url.Values); ok {
		return v.Encode()
	}
	return ""
}

// TransformURL - build a new URL by transforming an existing URL
func TransformURL(host string, uri *url.URL) *url.URL {
	if uri == nil {
		return uri
	}
	if host == "" {
		host = uri.Host
	}
	newURL := BuildURL(host, "", uri.Path, uri.RawQuery)
	u, err1 := url.Parse(newURL)
	if err1 != nil {
		return uri
	}
	return u
}

/*
	scheme := HttpsScheme
	if host == "" {
		if uri.Host != "" {
			host = uri.Host
		} else {
			host = Localhost
		}
	}
	if strings.Contains(host, Localhost) {
		scheme = HttpScheme
	}
	var newUri = scheme + "://" + host

	if len(uri.Path) > 0 {
		if uri.Path[:1] != "/" {
			newUri += "/"
		}
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
