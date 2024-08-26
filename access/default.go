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

const ()

var defaultLog = func(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) {
	s := formatter(o, traffic, start, duration, req, resp, routing, controller)
	log.Default().Printf("%v\n", s)
}

func DefaultFormat(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routing Routing, controller Controller) string {
	newReq := BuildRequest(req)
	newResp := BuildResponse(resp)
	url, parsed := uri.ParseURL(newReq.Host, newReq.URL)
	o.Host = Conditional(o.Host, parsed.Host)
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
		"\"location\":%v, "+
		"\"protocol\":%v, "+
		"\"method\":%v, "+
		"\"host\":%v, "+
		"\"from\":%v, "+
		"\"to\":%v, "+
		"\"uri\":%v, "+
		"\"path\":%v, "+
		"\"query\":%v, "+
		"\"status-code\":%v, "+
		"\"encoding\":%v, "+
		"\"bytes\":%v, "+
		"\"route\":%v, "+
		"\"route-to\":%v, "+
		"\"route-percent\":%v, "+
		"\"timeout\":%v, "+
		"\"rate-limit\":%v, "+
		"\"rate-burst\":%v, "+
		"\"cc\":%v, "+
		"\"rc\":%v }",

		// Origin, traffic, timestamp, duration
		fmt2.JsonString(o.Region),
		fmt2.JsonString(o.Zone),
		fmt2.JsonString(o.SubZone),
		fmt2.JsonString(o.InstanceId),
		traffic,
		fmt2.FmtRFC3339Millis(start),
		strconv.Itoa(Milliseconds(duration)),

		// Request
		fmt2.JsonString(newReq.Header.Get(XRequestId)),
		fmt2.JsonString(newReq.Header.Get(XRelatesTo)),
		fmt2.JsonString(newReq.Header.Get(LocationHeader)),
		fmt2.JsonString(newReq.Proto),
		fmt2.JsonString(newReq.Method),
		fmt2.JsonString(o.Host),
		fmt2.JsonString(routing.From),
		fmt2.JsonString(CreateTo(newReq)),
		fmt2.JsonString(url),
		fmt2.JsonString(parsed.Path),
		fmt2.JsonString(parsed.Query),

		// Response
		newResp.StatusCode,
		//fmt2.JsonString(resp.Status),
		fmt2.JsonString(Encoding(newResp)),
		fmt.Sprintf("%v", newResp.ContentLength),

		// Routing
		fmt2.JsonString(routing.Route),
		fmt2.JsonString(routing.To),
		fmt.Sprintf("%v", routing.Percent),

		// Controller thresholds
		//Threshold(threshold),
		Milliseconds(controller.Timeout),
		fmt.Sprintf("%v", controller.RateLimit),
		strconv.Itoa(controller.RateBurst),
		//fmt2.JsonString(thresholdCode),
		fmt2.JsonString(controller.Code),
		fmt2.JsonString(routing.Code),
	)

	return s
}

// Milliseconds - convert time.Duration to milliseconds
func Milliseconds(duration time.Duration) int {
	return int(duration / time.Duration(1e6))
}

func BuildRequest(r any) *http.Request {
	if r == nil {
		newReq, _ := http.NewRequest("", failsafeUri, nil)
		return newReq
	}
	if req, ok := r.(*http.Request); ok {
		return req
	}
	if req, ok := r.(*Request); ok {
		newReq, _ := http.NewRequest(req.Method, req.Url, nil)
		newReq.Header = req.Header
		return newReq
	}
	newReq, _ := http.NewRequest("", "https://somehost.com/search?q=test", nil)
	return newReq
}

func BuildResponse(r any) *http.Response {
	if r == nil {
		newResp := &http.Response{StatusCode: http.StatusOK}
		return newResp
	}
	if newResp, ok := r.(*http.Response); ok {
		return newResp
	}
	if sc, ok := r.(int); ok {
		return &http.Response{StatusCode: sc}
	}
	if status, ok := r.(*core.Status); ok {
		return &http.Response{StatusCode: status.HttpCode()}
	}
	newResp := &http.Response{StatusCode: http.StatusOK}
	return newResp
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

/*
func threshold(threshold any) int {
	if threshold == nil {
		return 0
	}
	if dur, ok := threshold.(time.Duration); ok {
		return Milliseconds(dur)
	}
	if i, ok1 := threshold.(int); ok1 {
		return i
	}
	if f, ok2 := threshold.(float64); ok2 {
		return int(f)
	}
	if ctx, ok := threshold.(context.Context); ok {
		if deadline, ok1 := ctx.Deadline(); ok1 {
			return Milliseconds(time.Until(deadline))
		}
	}
	return 0
}


*/

func CreateTo(req *http.Request) string {
	if req == nil {
		return ""
	}
	to := req.Header.Get(core.XTo)
	if to != "" {
		return to
	}
	return uri.UprootAuthority(req.URL)
}
