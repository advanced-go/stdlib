package access

import (
	"net/http"
	"time"
)

// Routing - routing attributes
type Routing struct {
	From    string // Authority
	Route   string
	To      string // Primary, secondary
	Percent int
	Code    string
}

// NilRouting - used when Routing is not applicable
func NilRouting() Routing {
	return Routing{Percent: -1}
}

// Controller - controller attributes
type Controller struct {
	Timeout   time.Duration
	RateLimit float64
	RateBurst int
	Code      string
}

// NilController - used when Controller is not applicable
func NilController() Controller {
	return Controller{RateLimit: -1, RateBurst: -1}
}

// Request - request attributes interface for non HTTP traffic
type Request interface {
	Url() string
	Header() http.Header
	Method() string
	Protocol() string
}

// RequestImpl - non HTTP request attributes
type RequestImpl struct {
	Url      string
	Header   http.Header
	Method   string
	Protocol string
}
