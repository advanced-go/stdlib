package core

import (
	"net/http"
	"time"
)

const (
	HealthLivenessPath  = "health/liveness"
	HealthReadinessPath = "health/readiness"
	VersionPath         = "version"
	AuthorityPath       = "authority"
	AuthorityRootPath   = "/authority"
)

// Routing - routing attributes
type Routing struct {
	FromAuthority string // Authority
	RouteName     string
	To            string // Primary, secondary
	Percent       int
	Code          string
}

// Controller - controller attributes
type Controller struct {
	Timeout   time.Duration
	RateLimit float64
	RateBurst int
	Code      string
}

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpExchange func(r *http.Request) (*http.Response, *Status)
type HttpAccess func(o Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller)

var (
	authorityReq *http.Request
)

func init() {
	authorityReq, _ = http.NewRequest(http.MethodGet, AuthorityRootPath, nil)
}

func Authority(h HttpExchange) string {
	if h == nil {
		return ""
	}
	resp, status := h(authorityReq)
	if status.OK() {
		return resp.Header.Get(XAuthority)
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
