package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
)

func finalize2(resp *http.Response) {
	if resp.Header == nil {
		resp.Header = make(http.Header)
	}
	resp.Header.Add(core.XAuthority, "github/advanced-go/stdlib")
	if resp.Request != nil {
		resp.Header.Add("x-method", resp.Request.Method)
	}
}

var (
	content3 = NewListContent[core.Origin, Patch, postContent2](match2, patch2, post2)
	rsc3     = NewResource2[core.Origin, Patch, postContent2]("rsc-origin", content3, finalize2)
)

func getList3() []core.Origin {
	var list []core.Origin

	if l, ok := any(content3).(*ListContent[core.Origin, Patch, postContent2]); ok {
		list = l.List
	}
	return list
}

func ExampleResource2_Put() {
	reader, _, status := json2.NewReadCloser(origin2)
	if !status.OK() {
		fmt.Printf("test: Put() -> [read-closer-status:%v]\n", status)
	} else {
		req, _ := http.NewRequest(http.MethodPut, "https://localhost:8081/github/advanced-go/documents:resiliency", reader)
		resp, status1 := rsc3.Do(req)
		fmt.Printf("test: Put() -> [status:%v] [status-code:%v] [header:%v] [%v]\n", status1, resp.StatusCode, resp.Header, getList3())
	}

	//Output:
	//test: Put() -> [status:OK] [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[PUT]]] [[{region1 Zone1  www.host1.com } {region1 Zone2  www.host2.com } {region2 Zone1  www.google.com }]]

}

func ExampleResource2_Get() {
	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081/github/advanced-go/documents:resiliency?zone=zone1", nil)
	resp, status := rsc3.Do(req)
	if !status.OK() {
		fmt.Printf("test: Do() -> [status:%v]\n", status)
	} else {
		items, status1 := json2.New[[]core.Origin](resp.Body, nil)
		fmt.Printf("test: Get() -> [status:%v] [status-code:%v] [header:%v] [%v]\n", status1, resp.StatusCode, resp.Header, items)
	}

	//Output:
	//test: Get() -> [status:OK] [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[GET]]] [[{region1 Zone1  www.host1.com } {region2 Zone1  www.google.com }]]

}

func ExampleResource2_Delete() {
	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081/github/advanced-go/documents:resiliency?zone=zone1", nil)
	resp, status := rsc3.Do(req)
	if !status.OK() {
		fmt.Printf("test: Do() -> [status:%v]\n", status)
	} else {
		items, status1 := json2.New[[]core.Origin](resp.Body, nil)
		fmt.Printf("test: Get() -> [status:%v] [status-code:%v] [header:%v] [%v]\n", status1, resp.StatusCode, resp.Header, items)
	}

	//Output:
	//test: Get() -> [status:OK] [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[GET]]] [[{region1 Zone1  www.host1.com } {region2 Zone1  www.google.com }]]

}
