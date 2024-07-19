package access

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	InternalTraffic = "internal"
	EgressTraffic   = "egress"
	IngressTraffic  = "ingress"
	failsafeUri     = "https://invalid-uri.com"
	XRequestId      = "x-request-id"
	XRelatesTo      = "x-relates-to"
	TimeoutCode     = "TO"
	RateLimitCode   = "RL"
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

// SetOrigin - initialize the origin
func SetOrigin(o core.Origin) {
	origin = o
}

// FormatFunc - formatting
type FormatFunc func(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) string

// SetFormatFunc - override formatting
func SetFormatFunc(fn FormatFunc) {
	if fn != nil {
		formatter = fn
	}
}

// LogFn - log function
type LogFn func(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller)

// SetLogFn - override logging
func SetLogFn(fn LogFn) {
	if fn != nil {
		logger = fn
	}
}

func DisableLogging(v bool) {
	disabled = v
}

var (
	origin    = core.Origin{}
	formatter = DefaultFormat
	logger    = defaultLog
	disabled  = false
)

// Log - access logging
func Log(traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) {
	if logger == nil || disabled {
		return
	}
	logger(origin, traffic, start, duration, req, resp, routing, controller)
}

func LogEgress(start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) {
	Log(EgressTraffic, start, duration, req, resp, routing, controller)
}

/*
// LogDeferred - deferred accessing logging
func LogDeferred(traffic string, req *http.Request, routeName, routeTo string, threshold int, thresholdCode string, statusCode func() int) func() {
	start := time.Now().UTC()
	return func() {
		Log(traffic, start, time.Since(start), req, &http.Response{StatusCode: statusCode(), Status: ""}, routeName, routeTo, threshold, thresholdCode)
	}
}

// NewRequest - create a new request
func NewRequest(h http.Header, method, uri string) *http.Request {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		req, err = http.NewRequest(method, failsafeUri, nil)
	}
	req.Header = h
	return req
}


*/
