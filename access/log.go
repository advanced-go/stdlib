package access

import (
	"github.com/advanced-go/stdlib/core"
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
	ContentEncoding = "Content-Encoding"
	LocationHeader  = "Location"

	Primary             = "primary"
	Secondary           = "secondary"
	ControllerTimeout   = "TO" // Controller struct code
	ControllerRateLimit = "RL" // Controller struct code
	RoutingFailover     = "FO" // Routing struct code
	RoutingRedirect     = "RD" // Routing struct code
)

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

// RequestConstraints - Request constraints
type RequestConstraints interface {
	*http.Request | Request
}

// ResponseConstraints - Response constraints
type ResponseConstraints interface {
	*http.Response | *core.Status | int
}

// Log - access logging.
// Header.Get(XRequestId)),
// Header.Get(XRelatesTo)),
// Header.Get(LocationHeader)
func Log[T RequestConstraints, U ResponseConstraints](traffic string, start time.Time, duration time.Duration, req T, resp U, routing Routing, controller Controller) {
	if logger == nil || disabled {
		return
	}
	logger(origin, traffic, start, duration, req, resp, routing, controller)
}
