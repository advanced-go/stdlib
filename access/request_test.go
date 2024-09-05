package access

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func ExampleNewRequest() {
	h := make(http.Header)
	h.Add(core.XAuthority, "github/advanced-go/search")
	req := RequestImpl{Method: http.MethodPatch, Url: "https://www.google.com/search?q=golang", Header: h}

	fmt.Printf("test: NewRequest() -> [method:%v] [url:%v] [h:%v]\n", req.Method, req.Url, req.Header)

	//Output:
	//test: NewRequest() -> [method:PATCH] [url:https://www.google.com/search?q=golang] [h:map[X-Authority:[github/advanced-go/search]]]

}
