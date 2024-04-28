package access

import (
	"net/http"
	"time"
)

const (
	InternalTraffic = "internal"
	EgressTraffic   = "egress"
	IngressTraffic  = "ingress"
	failsafeUri     = "https://invalid-uri.com"
	XRequestId      = "x-request-id"
	XRelatesTo      = "x-relates-to"
	TimeoutFlag     = "TO"
)

// Origin - log source location
type Origin struct {
	Region     string
	Zone       string
	SubZone    string
	App        string
	InstanceId string
}

// SetOrigin - initialize the origin
func SetOrigin(o Origin) {
	origin = o
}

// FormatFunc - formatting
type FormatFunc func(o *Origin, traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, routeName, routeTo string, threshold int, thresholdFlags string) string

// SetFormatFunc - override formatting
func SetFormatFunc(fn FormatFunc) {
	if fn != nil {
		formatter = fn
	}
}

// LogFn - log function
type LogFn func(o *Origin, traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, routeName string, routeTo string, threshold int, thresholdFlags string)

// SetLogFn - override logging
func SetLogFn(fn LogFn) {
	if fn != nil {
		logger = fn
	}
}

var (
	origin    = Origin{}
	formatter = DefaultFormat
	logger    = defaultLog
)

// Log - access logging
func Log(traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, routeName, routeTo string, threshold int, thresholdFlags string) {
	if logger == nil {
		return
	}
	logger(&origin, traffic, start, duration, req, resp, routeName, routeTo, threshold, thresholdFlags)
}

/*
// LogDeferred - deferred accessing logging
func LogDeferred(traffic string, req *http.Request, routeName, routeTo string, threshold int, thresholdFlags string, statusCode func() int) func() {
	start := time.Now().UTC()
	return func() {
		Log(traffic, start, time.Since(start), req, &http.Response{StatusCode: statusCode(), Status: ""}, routeName, routeTo, threshold, thresholdFlags)
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
