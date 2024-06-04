package access

import (
	"context"
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

var defaultLog = func(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routeName, routeTo string, threshold any, thresholdCode string) {
	s := formatter(o, traffic, start, duration, req, resp, routeName, routeTo, threshold, thresholdCode)
	log.Default().Printf("%v\n", s)
}

func DefaultFormat(o core.Origin, traffic string, start time.Time, duration time.Duration, req any, resp any, routeName, routeTo string, threshold any, thresholdCode string) string {
	newReq := NewRequest(req)
	newResp := NewResponse(resp)
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
		"\"th-code\":%v }",

		// Origin, traffic, timestamp, duration
		fmt2.JsonString(o.Region),
		fmt2.JsonString(o.Zone),
		fmt2.JsonString(o.SubZone),
		fmt2.JsonString(o.InstanceId),
		traffic,
		fmt2.FmtRFC3339Millis(start),
		strconv.Itoa(milliseconds(duration)),

		// Request
		fmt2.JsonString(newReq.Header.Get(XRequestId)),
		fmt2.JsonString(newReq.Header.Get(XRelatesTo)),
		fmt2.JsonString(newReq.Proto),
		fmt2.JsonString(newReq.Method),
		fmt2.JsonString(o.Host),
		fmt2.JsonString(newReq.Header.Get(core.XAuthority)),
		fmt2.JsonString(uri.UprootAuthority(newReq.URL)),
		fmt2.JsonString(url),
		fmt2.JsonString(parsed.Path),
		fmt2.JsonString(parsed.Query),

		// Response
		newResp.StatusCode,
		//fmt2.JsonString(resp.Status),
		fmt2.JsonString(Encoding(newResp)),
		fmt.Sprintf("%v", newResp.ContentLength),

		// Routing, controller threshold
		fmt2.JsonString(routeName),
		fmt2.JsonString(routeTo),
		Threshold(threshold),
		fmt2.JsonString(thresholdCode),
	)

	return s
}

// milliseconds - convert time.Duration to milliseconds
func milliseconds(duration time.Duration) int {
	return int(duration / time.Duration(1e6))
}

func NewRequest(r any) *http.Request {
	if r == nil {
		newReq, _ := http.NewRequest("", "https://somehost.com/search?q=test", nil)
		return newReq
	}
	if req, ok := r.(*http.Request); ok {
		return req
	}
	if req, ok := r.(Request); ok {
		newReq, _ := http.NewRequest(req.Method(), req.Url(), nil)
		newReq.Header = req.Header()
		return newReq
	}
	newReq, _ := http.NewRequest("", "https://somehost.com/search?q=test", nil)
	return newReq
}

func NewResponse(r any) *http.Response {
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

func Threshold(threshold any) int {
	if threshold == nil {
		return 0
	}
	if dur, ok := threshold.(time.Duration); ok {
		return milliseconds(dur)
	}
	if i, ok1 := threshold.(int); ok1 {
		return i
	}
	if f, ok2 := threshold.(float64); ok2 {
		return int(f)
	}
	if ctx, ok := threshold.(context.Context); ok {
		if deadline, ok1 := ctx.Deadline(); ok1 {
			return milliseconds(time.Until(deadline))
		}
	}
	return 0
}
