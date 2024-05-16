package controller

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
	"time"
)

type Resource struct {
	Name         string `json:"name"`
	Authority    string `json:"authority"`
	LivenessPath string `json:"liveness"`
	Duration     time.Duration
	Handler      core.HttpExchange
}

func newResource(name, authority string, duration time.Duration, livenessPath string, handler core.HttpExchange) *Resource {
	r := new(Resource)
	r.Name = name
	r.Authority = authority
	r.LivenessPath = livenessPath
	r.Duration = duration
	if handler != nil {
		r.Handler = handler
	}
	return r
}

func NewPrimaryResource(authority string, duration time.Duration, livenessPath string, handler core.HttpExchange) *Resource {
	return newResource(PrimaryName, authority, duration, livenessPath, handler)
}

func NewSecondaryResource(authority string, duration time.Duration, livenessPath string, handler core.HttpExchange) *Resource {
	return newResource(SecondaryName, authority, duration, livenessPath, handler)
}

func (r *Resource) IsPrimary() bool {
	return r != nil && r.Name == PrimaryName
}

func (r *Resource) BuildUri(uri *url.URL) *url.URL {
	if uri == nil || len(r.Authority) == 0 {
		return uri
	}
	uri2, err := url.Parse(r.Authority)
	if err != nil {
		return uri
	}
	var newUri = uri2.Scheme + "://"
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
	u, err1 := url.Parse(newUri)
	if err1 != nil {
		return uri
	}
	return u
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
