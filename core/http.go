package core

import (
	"net/http"
)

const (
	HealthLivenessPath  = "health/liveness"
	HealthReadinessPath = "health/readiness"
	VersionPath         = "version"
	AuthorityPath       = "authority"
	AuthorityRootPath   = "/authority"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpExchange func(r *http.Request) (*http.Response, *Status)

var (
	req *http.Request
)

func init() {
	req, _ = http.NewRequest(http.MethodGet, AuthorityRootPath, nil)
}

func Authority(h HttpExchange) string {
	if h == nil {
		return ""
	}
	resp, status := h(req)
	if status.OK() {
		return resp.Header.Get(XURLAuthority)
	}
	return ""
}

/*
func VersionContent(s string) string {
	return fmt.Sprintf("{ \"version\": \"%v\" }", s)
}

func HealthContent(s string) string {
	return fmt.Sprintf("{ \"status\": \"%v\" }", s)
}


*/
