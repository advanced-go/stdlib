package access

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func ExampleNewRequest() {
	h := make(http.Header)
	h.Add(core.XAuthority, "github/advanced-go/search")
	req := NewRequest(http.MethodPatch, "https://www.google.com/search?q=golang", h, "google-search", time.Second*2)

	fmt.Printf("test: NewRequest() -> [method:%v] [url:%v] [route:%v] [dur:%v] [h:%v]\n", req.Method(), req.Url(), req.RouteName(), req.Duration(), req.Header())

	//Output:
	//test: NewRequest() -> [method:PATCH] [url:https://www.google.com/search?q=golang] [route:google-search] [dur:2s] [h:map[X-Authority:[github/advanced-go/search]]]

}
