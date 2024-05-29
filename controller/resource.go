package controller

import (
	"github.com/advanced-go/stdlib/core"
	uri2 "github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
	"time"
)

const (
	localhost = "localhost"
)

type Resource struct {
	SilentAccess bool   `json:"silent"`
	Name         string `json:"name"`
	Host         string `json:"host"`
	Authority    string `json:"authority"`
	LivenessPath string `json:"liveness"`
	Duration     time.Duration
	Handler      core.HttpExchange
}

func newResource(silent bool, name, host, authority string, duration time.Duration, livenessPath string, handler core.HttpExchange) *Resource {
	r := new(Resource)
	r.SilentAccess = silent
	r.Name = name
	r.Host = host
	r.Authority = authority
	r.LivenessPath = livenessPath
	r.Duration = duration
	if handler != nil {
		r.Handler = handler
	}
	return r
}

func NewPrimaryResource(silent bool, host, authority string, duration time.Duration, livenessPath string, handler core.HttpExchange) *Resource {
	return newResource(silent, PrimaryName, host, authority, duration, livenessPath, handler)
}

func NewSecondaryResource(silent bool, host, authority string, duration time.Duration, livenessPath string, handler core.HttpExchange) *Resource {
	return newResource(silent, SecondaryName, host, authority, duration, livenessPath, handler)
}

func (r *Resource) IsPrimary() bool {
	return r != nil && r.Name == PrimaryName
}

func (r *Resource) BuildURL(uri *url.URL) *url.URL {
	return uri2.BuildURL(r.Host, uri)
}

func (r *Resource) timeout(req *http.Request) time.Duration {
	duration := r.Duration
	if r.Duration < 0 {
		duration = 0
	}
	if req == nil || req.Context() == nil {
		return duration
	}
	ct, ok := req.Context().Deadline()
	if !ok {
		return duration
	}
	until := time.Until(ct)
	if until <= duration || duration == 0 {
		return until * -1
	}
	return duration
}

/*
	if uri == nil {
		return uri
	}
	scheme := "https"
	host := r.Host
	if host == "" {
		host = localhost
	}
	if strings.Contains(host, localhost) {
		scheme = "http"
	}
	var newUri = scheme + "://" + host
	if r.Authority == "" {
		if len(uri.Path) > 0 {
			newUri += uri.Path
		}
		if len(uri.RawQuery) > 0 {
			newUri += "?" + uri.RawQuery
		}
	} else {
		newUri += "/" + r.Authority
		if len(uri.Path) > 0 {
			newUri += ":" + uri.Path[1:]
		}
		if len(uri.RawQuery) > 0 {
			newUri += "?" + uri.RawQuery
		}
		/*
			uri2, err := url.Parse(r.Authority)
			if err != nil {
				return uri
			}
			newUri = uri2.Scheme + "://"
			if len(uri2.Host) > 0 {
				newUri += uri2.Host
			} else {
				newUri += uri.Host
			}
			if len(uri2.Path) > 0 {
				newUri += uri2.Path
			} else {
				newUri += uri.Path
			}
			if len(uri2.RawQuery) > 0 {
				newUri += "?" + uri2.RawQuery
			} else {
				if len(uri.RawQuery) > 0 {
					newUri += "?" + uri.RawQuery
				}
			}

}
u, err1 := url.Parse(newUri)
if err1 != nil {
return uri
}
return u
*/
