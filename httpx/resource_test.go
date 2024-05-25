package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
)

type PostData struct {
	Item string
}

func postContent[POST any, T any](content *POST, list *[]T) *http.Response {
	return &http.Response{StatusCode: http.StatusBadRequest}
}

func patchContent[PATCH any, T any](content *PATCH, list *[]T) *http.Response {
	return &http.Response{StatusCode: http.StatusBadRequest}
}

func ExampleNewResource() {
	a := NewResource[core.Origin, Patch, PostData]("github/advanced-go/stdlib", originMatch2, finalize, nil, nil)
	//fmt.Printf("test: NewAuthority() -> [%v]\n", a)

	reader, _, status := json2.NewReadCloser(testOrigins2)
	if !status.OK() {
		fmt.Printf("test: PutT() -> [read-closer-status:%v]\n", status)
	} else {
		var list []core.Origin
		req, _ := http.NewRequest(http.MethodPut, "https://localhost:8081/github/advanced-go/documents:resiliency", reader)
		resp := a.Do(req) //PutT[core.Origin](req, &list, finalize)
		fmt.Printf("test: PutT() -> [status-code:%v] [header:%v] [%v]\n", resp.StatusCode, resp.Header, list)
	}

	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081/github/advanced-go/documents:resiliency?zone=zOne1", nil)
	resp := a.Do(req)
	if resp.Body != nil {
		items, status := json2.New[[]core.Origin](resp.Body, resp.Header)
		fmt.Printf("test: GetT() -> [status:%v] [status-code:%v] [header:%v] [%v]\n", status, resp.StatusCode, resp.Header, items)
	} else {
		fmt.Printf("test: GetT() -> [status-code:%v]\n", resp.StatusCode)
	}

	//Output:
	//test: PutT() -> [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[PUT]]] [[]]
	//test: GetT() -> [status:OK] [status-code:200] [header:map[Content-Type:[application/json] X-Authority:[github/advanced-go/stdlib] X-Method:[GET]]] [[{region1 Zone1  www.host1.com } {region2 Zone1  www.google.com }]]

}
