package access

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func ExampleNewRequest() {
	h := make(http.Header)
	h.Add(core.XAuthority, "github/advanced-go/search")
	req := NewRequest(http.MethodPatch, "https://www.google.com/search?q=golang", h)

	fmt.Printf("test: NewRequest() -> [method:%v] [url:%v] [h:%v]\n", req.Method(), req.Url(), req.Header())

	//Output:
	//test: NewRequest() -> [method:PATCH] [url:https://www.google.com/search?q=golang] [route:google-search] [dur:2s] [h:map[X-Authority:[github/advanced-go/search]]]

}
