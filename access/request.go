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

// Request - non HTTP request attributes
type Request struct {
	Url    string
	Header http.Header
	Method string
}

// Controller - controller attributes
type Controller struct {
	Timeout   time.Duration
	RateLimit float64
	RateBurst int
	Code      string
}
