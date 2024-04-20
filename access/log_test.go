package access

import (
	"fmt"
	"net/http"
	"time"
)

func _Example_LogAccess() {
	start := time.Now().UTC()
	r, _ := http.NewRequest("PUT", "/github.com/advanced-go/example-domain/activity:entry", nil)
	r.Host = "localhost:8080"
	s := DefaultFormat(&Origin{Region: "us", Zone: "zone", App: "ai-agent"}, EgressTraffic, start, time.Since(start), r, &http.Response{StatusCode: 200, Status: "OK"}, "route", "primary", -1, "")

	fmt.Printf("test: fmtlog() -> %v\n", s)

	//Output:

}
