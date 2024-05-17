package core

import (
	"fmt"
	"net/http"
)

func authExchange(req *http.Request) (*http.Response, *Status) {
	if req.URL.Path == InfoRootPath {
		h := make(http.Header)
		h.Add(XAuthority, "github/advanced-go/stdlib")
		h.Add(XVersion, "1.2.345")
		return &http.Response{StatusCode: http.StatusOK, Header: h}, StatusOK()
	}
	return &http.Response{StatusCode: http.StatusBadRequest}, NewStatus(http.StatusBadRequest)
}

func ExampleAuthority() {
	auth, vers := Authority(authExchange)
	fmt.Printf("test: Authority() -> [auth:%v] [vers:%v]\n", auth, vers)

	//Output:
	//test: Authority() -> [auth:github/advanced-go/stdlib] [vers:1.2.345]

}
