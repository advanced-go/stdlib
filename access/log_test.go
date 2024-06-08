package access

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func ExampleLogAccess() {
	var start time.Time //:= time.Now().UTC()
	r, _ := http.NewRequest("PUT", "/github/advanced-go/example-domain/activity:entry", nil)
	r.Host = "localhost:8080"
	s := DefaultFormat(core.Origin{Region: "us", Zone: "zone", Host: r.Host, InstanceId: "id-012"}, EgressTraffic, start, time.Millisecond*345, r, &http.Response{StatusCode: 200, Status: "OK"}, "from", "route", "primary", -1, 0, 0, "TO")

	fmt.Printf("test: log() -> %v\n", s)

	//Output:
	//test: log() -> {"region":"us", "zone":"zone", "sub-zone":null, "instance-id":"id-012", "traffic":"egress", "start":0001-01-01T00:00:00.000Z, "duration":345, "request-id":null, "relates-to":null, "protocol":"HTTP/1.1", "method":"PUT", "host":"localhost:8080", "from":"from", "to":"github/advanced-go/example-domain/activity", "uri":"http://localhost:8080/github/advanced-go/example-domain/activity:entry", "path":"entry", "query":null, "status-code":200, "encoding":null, "bytes":0, "route":"route", "route-to":"primary", "timeout":0, "rate-limit":0, "rate-burst":0, "rc":"TO" }

}
