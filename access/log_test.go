package access

import (
	"fmt"
	"net/http"
	"time"
)

func Example_LogAccess() {
	var start time.Time //:= time.Now().UTC()
	r, _ := http.NewRequest("PUT", "/github.com/advanced-go/example-domain/activity:entry", nil)
	r.Host = "localhost:8080"
	s := DefaultFormat(&Origin{Region: "us", Zone: "zone", App: "ai-agent"}, EgressTraffic, start, 0, r, &http.Response{StatusCode: 200, Status: "OK"}, "route", "primary", -1, "")

	fmt.Printf("test: fmtlog() -> %v\n", s)

	//Output:
	//test: fmtlog() -> {"region":"us", "zone":"zone", "sub-zone":null, "app":"ai-agent", "instance-id":null, "traffic":"egress", "start":0001-01-01T00:00:00.000Z, "duration":0, "request-id":null, "relates-to":null, "protocol":"HTTP/1.1", "method":"PUT", "uri":"http://localhost:8080/github.com/advanced-go/example-domain/activity:entry", "host":"localhost:8080", "path":"entry", "status-code":200, "encoding":null, "bytes":0, "route":"route", "route-to":"primary", "threshold":-1, "threshold-flags":null }

}
