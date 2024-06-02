package access

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	fmt2 "github.com/advanced-go/stdlib/fmt"
	"github.com/advanced-go/stdlib/uri"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ContentEncoding = "Content-Encoding"
)

var defaultLog = func(o core.Origin, traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, authority, routeName, routeTo string, threshold int, thresholdFlags string) {
	s := formatter(o, traffic, start, duration, req, resp, authority, routeName, routeTo, threshold, thresholdFlags)
	log.Default().Printf("%v\n", s)
}

func DefaultFormat(o core.Origin, traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, authority, routeName, routeTo string, threshold int, thresholdFlags string) string {
	req = SafeRequest(req)
	resp = SafeResponse(resp)
	//url, host, path, query := CreateURLComponents(req)
	url, parsed := uri.ParseURL(req.Host, req.URL)
	o.Host = Conditional(o.Host, parsed.Host)
	authority = Conditional(authority, o.Host)
	s := fmt.Sprintf("{"+
		"\"region\":%v, "+
		"\"zone\":%v, "+
		"\"sub-zone\":%v, "+
		"\"instance-id\":%v, "+
		"\"traffic\":\"%v\", "+
		"\"start\":%v, "+
		"\"duration\":%v, "+
		"\"request-id\":%v, "+
		"\"relates-to\":%v, "+
		"\"protocol\":%v, "+
		"\"method\":%v, "+
		"\"host\":%v, "+
		"\"auth-from\":%v, "+
		"\"auth-to\":%v, "+
		"\"uri\":%v, "+
		"\"path\":%v, "+
		"\"query\":%v, "+
		"\"status-code\":%v, "+
		"\"encoding\":%v, "+
		"\"bytes\":%v, "+
		"\"route\":%v, "+
		"\"route-to\":%v, "+
		"\"threshold\":%v, "+
		"\"threshold-flags\":%v }",

		// Origin, traffic, timestamp, duration
		fmt2.JsonString(o.Region),
		fmt2.JsonString(o.Zone),
		fmt2.JsonString(o.SubZone),
		fmt2.JsonString(o.InstanceId),
		traffic,
		fmt2.FmtRFC3339Millis(start),
		strconv.Itoa(Milliseconds(duration)),

		// Request
		fmt2.JsonString(req.Header.Get(XRequestId)),
		fmt2.JsonString(req.Header.Get(XRelatesTo)),
		fmt2.JsonString(req.Proto),
		fmt2.JsonString(req.Method),
		fmt2.JsonString(o.Host),
		fmt2.JsonString(req.Header.Get(core.XAuthority)),
		fmt2.JsonString(authority),
		fmt2.JsonString(url),
		fmt2.JsonString(parsed.Path),
		fmt2.JsonString(parsed.Query),

		// Response
		resp.StatusCode,
		//fmt2.JsonString(resp.Status),
		fmt2.JsonString(Encoding(resp)),
		fmt.Sprintf("%v", resp.ContentLength),

		// Routing, controller threshold
		fmt2.JsonString(routeName),
		fmt2.JsonString(routeTo),
		threshold,
		fmt2.JsonString(thresholdFlags),
	)

	return s
}

// Milliseconds - convert time.Duration to milliseconds
func Milliseconds(duration time.Duration) int {
	return int(duration / time.Duration(1e6))
}

func SafeRequest(r *http.Request) *http.Request {
	if r == nil {
		r, _ = http.NewRequest("", "https://somehost.com/search?q=test", nil)
	}
	return r
}

func SafeResponse(r *http.Response) *http.Response {
	if r == nil {
		r = new(http.Response)
	}
	return r
}

func Encoding(resp *http.Response) string {
	encoding := ""
	if resp != nil && resp.Header != nil {
		encoding = resp.Header.Get(ContentEncoding)
	}
	// normalize encoding
	if strings.Contains(strings.ToLower(encoding), "none") {
		encoding = ""
	}
	return encoding
}

func Conditional(primary, secondary string) string {
	if len(primary) == 0 {
		return secondary
	}
	return primary
}
