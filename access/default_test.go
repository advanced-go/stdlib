package access

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
	"time"
)

func ExampleFormatter() {
	//EnableTestLogger()
	start := time.Now().UTC()
	SetOrigin(core.Origin{Region: "us", Zone: "west", SubZone: "dc1", Host: "search-app", InstanceId: "123456789"})

	req, err := http.NewRequest("GET", "https://www.google.com/search?q=test", nil)
	req.Header.Add(XRequestId, "123-456")
	req.Header.Add(XRelatesTo, "your-id")
	fmt.Printf("test: NewRequest() -> [err:%v] [req:%v]\n", err, req != nil)
	resp := http.Response{StatusCode: http.StatusOK}
	resp.Header = make(http.Header)
	resp.Header.Add(ContentEncoding, "gzip")
	time.Sleep(time.Millisecond * 500)
	logTest(EgressTraffic, start, time.Since(start), req, &resp, "", "google-search", "secondary", -1, "")

	fmt.Printf("test: LogURI() -> %v\n", "success")

	//Output:
	//test: NewRequest() -> [err:<nil>] [req:true]
	//test: LogURI() -> success

}

func ExampleFormatter_Urn() {
	start := time.Now().UTC()
	values := make(url.Values)
	values.Add("region", "*")
	values.Add("zone", "texas")
	//u := uri.BuildURL()

	req, err := http.NewRequest("select", "https://github.com/advanced-go/example-domain/activity:v1/entry?"+uri.BuildQuery(values), nil)
	req.Header.Add(XRequestId, "123-456")
	req.Header.Add(XRelatesTo, "fmtlog testing")
	req.Header.Add(core.XAuthority, "github/advanced-go/auth-from")
	fmt.Printf("test: NewRequest() -> [err:%v] [req:%v]\n", err, req != nil)
	resp := http.Response{StatusCode: http.StatusOK}
	logTest(InternalTraffic, start, time.Since(start), req, &resp, "", "route", "primary", -1, "")
	fmt.Printf("test: LogURN() -> %v\n", "success")

	//Output:
	//test: NewRequest() -> [err:<nil>] [req:true]
	//test: LogURN() -> success

}

func logTest(traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, authority, routeName, routeTo string, threshold int, thresholdFlags string) {
	Log(traffic, start, duration, req, resp, authority, routeName, routeTo, threshold, thresholdFlags)
}
