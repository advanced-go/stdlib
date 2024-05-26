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

func postProcess(list *[]core.Origin, content *PostData) *http.Response {
	return &http.Response{StatusCode: http.StatusBadRequest}
}

func patchProcess(list *[]core.Origin, content *Patch) *http.Response {
	return &http.Response{StatusCode: http.StatusBadRequest}
}

func ExampleNewBasicResource() {
	a := NewBasicResource[core.Origin]("github/advanced-go/stdlib", originMatch2, finalize)

	reader, _, status := json2.NewReadCloser(testOrigins2)
	if !status.OK() {
		fmt.Printf("test: PutT() -> [read-closer-status:%v]\n", status)
	} else {
		var list []core.Origin
		req, _ := http.NewRequest(http.MethodPut, "https://localhost:8081/github/advanced-go/documents:resiliency", reader)
		resp, _ := a.Do(req)
		fmt.Printf("test: PutT() -> [status-code:%v] [header:%v] [%v]\n", resp.StatusCode, resp.Header, list)
	}

	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081/github/advanced-go/documents:resiliency?zone=zOne1", nil)
	resp, _ := a.Do(req)
	if resp.Body != nil {
		items, status1 := json2.New[[]core.Origin](resp.Body, resp.Header)
		fmt.Printf("test: GetT() -> [status:%v] [status-code:%v] [header:%v] [%v]\n", status1, resp.StatusCode, resp.Header, items)
	} else {
		fmt.Printf("test: GetT() -> [status-code:%v]\n", resp.StatusCode)
	}

	//Output:
	//test: PutT() -> [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[PUT]]] [[]]
	//test: GetT() -> [status:OK] [status-code:200] [header:map[Content-Type:[application/json] X-Authority:[github/advanced-go/stdlib] X-Method:[GET]]] [[{region1 Zone1  www.host1.com } {region2 Zone1  www.google.com }]]

}

func ExampleNewResource() {
	a := NewResource[core.Origin, Patch, PostData]("github/advanced-go/stdlib", originMatch2, finalize, patchProcess, postProcess)

	reader, _, status := json2.NewReadCloser(testOrigins2)
	if !status.OK() {
		fmt.Printf("test: PutT() -> [read-closer-status:%v]\n", status)
	} else {
		var list []core.Origin
		req, _ := http.NewRequest(http.MethodPut, "https://localhost:8081/github/advanced-go/documents:resiliency", reader)
		resp, _ := a.Do(req)
		fmt.Printf("test: PutT() -> [status-code:%v] [header:%v] [%v]\n", resp.StatusCode, resp.Header, list)
	}

	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081/github/advanced-go/documents:resiliency?zone=zOne1", nil)
	resp, _ := a.Do(req)
	if resp.Body != nil {
		items, status1 := json2.New[[]core.Origin](resp.Body, resp.Header)
		fmt.Printf("test: GetT() -> [status:%v] [status-code:%v] [header:%v] [%v]\n", status1, resp.StatusCode, resp.Header, items)
	} else {
		fmt.Printf("test: GetT() -> [status-code:%v]\n", resp.StatusCode)
	}

	//Output:
	//test: PutT() -> [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[PUT]]] [[]]
	//test: GetT() -> [status:OK] [status-code:200] [header:map[Content-Type:[application/json] X-Authority:[github/advanced-go/stdlib] X-Method:[GET]]] [[{region1 Zone1  www.host1.com } {region2 Zone1  www.google.com }]]

}
