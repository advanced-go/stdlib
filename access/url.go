package access

import (
	"net/http"
	"net/url"
	"strings"
)

// CreateURLComponents - create the URL, host, path, and query
// TODO : Need to url.PathUnescape(string) and url.QueryUnescape(string)
func CreateURLComponents(req *http.Request) (uri string, host string, path string, query string) {
	// Set scheme
	scheme := req.URL.Scheme
	if scheme == "" {
		scheme = "http"
	}
	// Set host
	host = req.Host
	if len(host) == 0 {
		host = req.URL.Host
	}
	// Set path
	urlPath, _ := url.PathUnescape(req.URL.Path)
	path = urlPath
	i := strings.Index(path, ":")
	if i >= 0 {
		path = path[i+1:]
	}

	// Set query
	if req.URL.RawQuery != "" {
		query, _ = url.QueryUnescape(req.URL.RawQuery)
	}
	if query != "" {
		uri = scheme + "://" + host + urlPath + "?" + query
	} else {
		uri = scheme + "://" + host + urlPath + query
	}

	/*
		url = req.URL.String()
		if len(host) == 0 {
			//url = "urn:" + url
		} else {
			if len(req.URL.Scheme) == 0 {
				url = "http://" + host + req.URL.Path
			}
		}
		path = req.URL.Path
		i := strings.Index(path, ":")
		if i >= 0 {
			path = path[i+1:]
		}

	*/
	return
}
