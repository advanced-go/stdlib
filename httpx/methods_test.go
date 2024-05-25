package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
)

var (
	testOrigins2 = []core.Origin{
		{Region: "region1", Zone: "Zone1", Host: "www.host1.com"},
		{Region: "region1", Zone: "Zone2", Host: "www.host2.com"},
		{Region: "region2", Zone: "Zone1", Host: "www.google.com"},
	}
)

func finalize(resp *http.Response) {
	if resp.Header == nil {
		resp.Header = make(http.Header)
	}
	resp.Header.Add(core.XAuthority, "github/advanced-go/stdlib")
	if resp.Request != nil {
		resp.Header.Add("x-method", resp.Request.Method)
	}
}

func originMatch2(item *core.Origin, req *http.Request) bool {
	filter := core.NewOrigin(req.URL.Query())
	//if entry, ok := item.(*core.Origin); ok {
	if core.OriginMatch(*item, filter) {
		return true
	}
	return false
}

func originPatch2(patch *Patch, list *[]core.Origin) *http.Response {
	for _, op := range patch.Updates {
		switch op.Op {
		case OpReplace:
			if op.Path == core.HostKey {
				if s, ok1 := op.Value.(string); ok1 {
					(*list)[0].Host = s
				}
			}
		default:
		}
	}
	return NewResponse(core.StatusOK(), nil)
}

type PostContent struct {
	Item string
}

func originPost(post *PostContent, list *[]core.Origin) *http.Response {
	(*list)[0].Host = "www.search.yahoo.com"
	return NewResponse(core.StatusOK(), nil)
}

func ExampleGetT() {
	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081/github/advanced-go/documents:resiliency?zone=zOne1", nil)
	resp := GetT[core.Origin](req, testOrigins2, originMatch2, finalize)

	if resp.Body != nil {
		items, status := json2.New[[]core.Origin](resp.Body, resp.Header)
		fmt.Printf("test: GetT() -> [status:%v] [status-code:%v] [header:%v] [%v]\n", status, resp.StatusCode, resp.Header, items)
	} else {
		fmt.Printf("test: GetT() -> [status-code:%v]\n", resp.StatusCode)
	}

	//Output:
	//test: GetT() -> [status:OK] [status-code:200] [header:map[Content-Type:[application/json] X-Authority:[github/advanced-go/stdlib] X-Method:[GET]]] [[{region1 Zone1  www.host1.com } {region2 Zone1  www.google.com }]]

}

func ExamplePutT() {
	reader, _, status := json2.NewReadCloser(testOrigins2)
	if !status.OK() {
		fmt.Printf("test: PutT() -> [read-closer-status:%v]\n", status)
	} else {
		var list []core.Origin
		req, _ := http.NewRequest(http.MethodPut, "https://localhost:8081/github/advanced-go/documents:resiliency", reader)
		resp := PutT[core.Origin](req, &list, finalize)
		fmt.Printf("test: PutT() -> [status-code:%v] [header:%v] [%v]\n", resp.StatusCode, resp.Header, list)
	}

	//Output:
	//test: PutT() -> [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[PUT]]] [[{region1 Zone1  www.host1.com } {region1 Zone2  www.host2.com } {region2 Zone1  www.google.com }]]

}

func ExampleDeleteT() {
	var local []core.Origin
	local = append(local, testOrigins2...)

	req, _ := http.NewRequest(http.MethodDelete, "https://localhost:8081/github/advanced-go/documents:resiliency?host=www.host2.com", nil)
	resp := DeleteT[core.Origin](req, &local, originMatch2, finalize)
	fmt.Printf("test: DeleteT-host() -> [status-code:%v] [header:%v] [%v]\n", resp.StatusCode, resp.Header, local)

	//Output:
	//test: DeleteT-host() -> [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[DELETE]]] [[{region1 Zone1  www.host1.com } {region2 Zone1  www.google.com }]]

}

func ExamplePatchT() {
	var local []core.Origin
	local = append(local, testOrigins2...)

	p := Patch{Updates: []Operation{
		{Op: OpReplace, Path: core.HostKey, Value: "www.search.yahoo.com"},
	}}
	buf, _ := json.Marshal(p)
	req, _ := http.NewRequest(http.MethodPatch, "https://localhost:8081/github/advanced-go/documents:resiliency", io.NopCloser(bytes.NewReader(buf)))
	resp := PatchT[Patch, core.Origin](req, &local, originPatch2, finalize)
	fmt.Printf("test: PatchT-host() -> [status-code:%v] [header:%v] [%v]\n", resp.StatusCode, resp.Header, local)

	//Output:
	//test: PatchT-host() -> [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[PATCH]]] [[{region1 Zone1  www.search.yahoo.com } {region1 Zone2  www.host2.com } {region2 Zone1  www.google.com }]]

}

func ExamplePostT() {
	var local []core.Origin
	local = append(local, testOrigins2...)

	p := PostContent{Item: "test"}
	buf, _ := json.Marshal(p)
	req, _ := http.NewRequest(http.MethodPost, "https://localhost:8081/github/advanced-go/documents:resiliency", io.NopCloser(bytes.NewReader(buf)))
	resp := PostT[PostContent, core.Origin](req, &local, originPost, finalize)
	fmt.Printf("test: PostT-host() -> [status-code:%v] [header:%v] [%v]\n", resp.StatusCode, resp.Header, local)

	//Output:
	//test: PostT-host() -> [status-code:200] [header:map[X-Authority:[github/advanced-go/stdlib] X-Method:[POST]]] [[{region1 Zone1  www.search.yahoo.com } {region1 Zone2  www.host2.com } {region2 Zone1  www.google.com }]]

}
