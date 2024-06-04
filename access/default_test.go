package access

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
	"time"
)

type testRequest struct {
	url    string
	header http.Header
	method string
}

func newTestRequest(url, method string, h http.Header) Request {
	r := new(testRequest)
	r.url = url
	r.method = method
	r.header = h
	return r
}

func (r *testRequest) Url() string {
	return r.url
}

func (r *testRequest) Method() string {
	return r.method
}

func (r *testRequest) Header() http.Header {
	return r.header
}

func ExampleDefault_Host() {
	start := time.Now().UTC()
	SetOrigin(core.Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	req, _ := http.NewRequest("GET", "https://www.google.com/search?q=test", nil)
	req.Header.Add(XRequestId, "123-456")
	req.Header.Add(XRelatesTo, "your-id")
	//fmt.Printf("test: NewRequest() -> [err:%v] [req:%v]\n", err, req != nil)
	resp := http.Response{StatusCode: http.StatusOK}
	resp.Header = make(http.Header)
	resp.Header.Add(ContentEncoding, "gzip")
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), req, &resp, "google-search", "secondary", -1, "")

	fmt.Printf("test: Default-Host() -> %v\n", "success")

	//Output:
	//test: Default-Host() -> success

}

func ExampleDefault_Authority() {
	start := time.Now().UTC()
	values := make(url.Values)
	values.Add("region", "*")
	values.Add("zone", "texas")
	//u := uri.BuildURL()

	req, _ := http.NewRequest("select", "https://github.com/advanced-go/example-domain/activity:v1/entry?"+uri.BuildQuery(values), nil)
	req.Header.Add(XRequestId, "123-456")
	req.Header.Add(XRelatesTo, "fmtlog testing")
	req.Header.Add(core.XAuthority, "github/advanced-go/auth-from")
	//fmt.Printf("test: NewRequest() -> [err:%v] [req:%v]\n", err, req != nil)
	resp := http.Response{StatusCode: http.StatusOK}
	logTest(InternalTraffic, start, time.Since(start), req, &resp, "route", "primary", -1, "")

	fmt.Printf("test: Default-Authority() -> %v\n", "success")

	//Output:
	//test: Default-Authority() -> success

}

func ExampleDefault_Access_Request_Status() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	h.Add(XRelatesTo, "test-request-interface")
	req := newTestRequest("https://www.google.com/search?q=test", http.MethodPut, h)
	start := time.Now().UTC()
	SetOrigin(core.Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := core.StatusNotFound()
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), req, resp, "google-search", "secondary", -1, "")

	fmt.Printf("test: Default-Access-Request-Status() -> %v\n", "success")

	//Output:
	//test: Default-Access-Request-Status() -> success

}

func ExampleDefault_Access_Request_Status_Code() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	h.Add(XRelatesTo, "test-request-interface")
	req := newTestRequest("https://www.google.com/search?q=test", http.MethodPut, h)
	start := time.Now().UTC()
	SetOrigin(core.Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), req, resp, "google-search", "secondary", -1, "")

	fmt.Printf("test: Default-Access-Request-Status-Code() -> %v\n", "success")

	//Output:
	//test: Default-Access-Request-Status-Code() -> success

}

func ExampleDefault_Threshold_Duration() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	h.Add(XRelatesTo, "test-request-interface")
	req := newTestRequest("https://www.google.com/search?q=test", http.MethodPut, h)
	start := time.Now().UTC()
	SetOrigin(core.Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), req, resp, "google-search", "secondary", time.Second*4, "")

	fmt.Printf("test: Default-Threshold-Duration() -> %v\n", "success")

	//Output:
	//test: Default-Threshold-Duration() -> success

}

func ExampleDefault_Threshold_Int() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	h.Add(XRelatesTo, "test-request-interface")
	req := newTestRequest("https://www.google.com/search?q=test", http.MethodPut, h)
	start := time.Now().UTC()
	SetOrigin(core.Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), req, resp, "google-search", "secondary", 345, "")

	fmt.Printf("test: Default-Threshold-Int() -> %v\n", "success")

	//Output:
	//test: Default-Threshold-Int() -> success

}

func ExampleDefault_Threshold_Deadline() {
	h := make(http.Header)
	h.Add(XRequestId, "987-654")
	h.Add(XRelatesTo, "test-request-interface")
	req := newTestRequest("https://www.google.com/search?q=test", http.MethodPut, h)
	start := time.Now().UTC()
	SetOrigin(core.Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	resp := http.StatusGatewayTimeout
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), req, resp, "google-search", "secondary", ctx, "")

	fmt.Printf("test: Default-Threshold-Int() -> %v\n", "success")

	//Output:
	//test: Default-Threshold-Int() -> success

}

func logTest(traffic string, start time.Time, duration time.Duration, req any, resp any, routeName, routeTo string, threshold any, thresholdFlags string) {
	Log(traffic, start, duration, req, resp, routeName, routeTo, threshold, thresholdFlags)
}
