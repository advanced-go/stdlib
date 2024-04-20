package access

import (
	"fmt"
	"github.com/advaced-go/stdlib/sfmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ContentEncoding = "Content-Encoding"
)

var defaultLogger = func(o *Origin, traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, routeName, routeTo string, threshold int, thresholdFlags string) {
	s := formatter(o, traffic, start, duration, req, resp, routeName, routeTo, threshold, thresholdFlags)
	log.Default().Printf("%v\n", s)
}

func DefaultFormatter(o *Origin, traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, routeName, routeTo string, threshold int, thresholdFlags string) string {
	if o == nil {
		o = &origin
	}
	req = SafeRequest(req)
	resp = SafeResponse(resp)
	url, host, path := CreateUrlHostPath(req)
	s := fmt.Sprintf("{"+
		"\"region\":%v, "+
		"\"zone\":%v, "+
		"\"sub-zone\":%v, "+
		"\"app\":%v, "+
		"\"instance-id\":%v, "+
		"\"traffic\":\"%v\", "+
		"\"start\":%v, "+
		"\"duration\":%v, "+
		"\"request-id\":%v, "+
		"\"relates-to\":%v, "+
		"\"protocol\":%v, "+
		"\"method\":%v, "+
		"\"uri\":%v, "+
		"\"host\":%v, "+
		"\"path\":%v, "+
		"\"status-code\":%v, "+
		"\"encoding\":%v, "+
		"\"bytes-written\":%v, "+
		"\"route\":%v, "+
		"\"route-to\":%v, "+
		"\"threshold\":%v, "+
		"\"threshold-flags\":%v }",
		sfmt.JsonString(o.Region),
		sfmt.JsonString(o.Zone),
		sfmt.JsonString(o.SubZone),
		sfmt.JsonString(o.App),
		sfmt.JsonString(o.InstanceId),

		traffic,
		sfmt.FmtRFC3339Millis(start),
		strconv.Itoa(Milliseconds(duration)),

		sfmt.JsonString(req.Header.Get(XRequestId)),
		sfmt.JsonString(req.Header.Get(XRelatesTo)),
		sfmt.JsonString(req.Proto),
		sfmt.JsonString(req.Method),
		sfmt.JsonString(url),
		sfmt.JsonString(host),
		sfmt.JsonString(path),

		resp.StatusCode,
		//sfmt.JsonString(resp.Status),
		sfmt.JsonString(Encoding(resp)),
		fmt.Sprintf("%v", resp.ContentLength),

		sfmt.JsonString(routeName),
		sfmt.JsonString(routeTo),
		threshold,
		sfmt.JsonString(thresholdFlags),
	)

	return s
}

// Milliseconds - convert time.Duration to milliseconds
func Milliseconds(duration time.Duration) int {
	return int(duration / time.Duration(1e6))
}

// CreateUrlHostPath - create the URL, host and path
func CreateUrlHostPath(req *http.Request) (url string, host string, path string) {
	host = req.Host
	if len(host) == 0 {
		host = req.URL.Host
	}
	url = req.URL.String()
	if len(host) == 0 {
		//url = "urn:" + url
	} else {
		if len(req.URL.Scheme) == 0 {
			url = "http://" + host + req.URL.Path
		}
	}
	path = req.URL.Path
	i := strings.Index(path, ":")
	if i >= 0 {
		path = path[i+1:]
	}
	return
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
