package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

/*
func _ExampleInfoFmt() {
	info := core.ModuleInfo{
		Authority: "github/advanced/go/stdlib",
		Version:   "7.8.9",
		Name:      "library",
	}
	s := fmt.Sprintf(infoFmt, info.Authority, info.Version, info.Name)
	fmt.Printf("test: InfoFmt() -> %v\n", s)

	//Output:
	//fail

}


*/

var testCore = []core.Origin{
	{Region: "region1", Zone: "Zone1", Host: "www.host1.com"},
	{Region: "region1", Zone: "Zone2", Host: "www.host2.com"},
	{Region: "region2", Zone: "Zone1", Host: "www.google.com"},
}

func ExampleNewResponse_Error() {
	status := core.NewStatus(http.StatusGatewayTimeout)
	resp, _ := NewResponse(status.HttpCode(), nil, status.Err)
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	status = core.NewStatusError(http.StatusGatewayTimeout, errors.New("Deadline Exceeded"))
	resp, _ = NewResponse(status.HttpCode(), nil, status.Err)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:504] [content:]
	//test: NewResponse() -> [status-code:504] [content:Deadline Exceeded]

}

func ExampleNewResponse() {
	resp, _ := NewResponse(http.StatusOK, nil, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v]\n", resp.StatusCode)

	resp, _ = NewResponse(core.StatusOK().HttpCode(), nil, "version 1.2.35")
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:200]
	//test: NewResponse() -> [status-code:200] [content:version 1.2.35]

}

func ExampleNewVersionResponse() {
	resp := NewVersionResponse("7.8.9")
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewVersionResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewVersionResponse() -> [status-code:200] [content:{
	// "version": "7.8.9"
	//  }]

}

func ExampleNewAuthorityResponse() {
	resp := NewAuthorityResponse("github/advanced-go/stdlib")
	fmt.Printf("test: NewAuthorityResponse() -> [status-code:%v] [auth:%v]\n", resp.StatusCode, resp.Header.Get(core.XAuthority))

	//Output:
	//test: NewAuthorityResponse() -> [status-code:200] [auth:github/advanced-go/stdlib]

}

func ExampleNewHealthResponseOK() {
	resp := NewHealthResponseOK()
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewHealthResponseOK() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewHealthResponseOK() -> [status-code:200] [content:{
	// "status": "up"
	//}]

}

func ExampleNewNotFoundResponseWithStatus() {
	resp := NewNotFoundResponse()
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewNotFoundResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewNotFoundResponse() -> [status-code:404] [content:Not Found]

}

/*
func ExampleNewJsonResponse() {
	resp, status := NewJsonResponse(nil, nil)
	fmt.Printf("test: NewJsonResponse(nil,nil) -> [status:%v] [status-code:%v] [content-length:%v]\n", status, resp.StatusCode, resp.ContentLength)

	resp, status = NewJsonResponse(testCore, nil)
	//if status.OK() && resp.Body != nil {
	//	buff, _ = io.ReadAll(resp.Body, nil)
	//}
	fmt.Printf("test: NewJsonResponse(testCore,nil) -> [status:%v] [status-code:%v] [header:%v] [content-length:%v]\n", status, resp.StatusCode, resp.Header, resp.ContentLength)
	//fmt.Printf("test: NewJsonResponse(testCore,nil) -> [status:%v] [status-code:%v] [content:%v]\n", status, resp.StatusCode, string(buff))

	h := make(http.Header)
	h.Add(ContentLocation, "http://localhost:8081/search?q=golang")
	resp, status = NewJsonResponse(testCore, h)
	fmt.Printf("test: NewJsonResponse(testCore,nil) -> [status:%v] [status-code:%v] [header:%v] [content-length:%v]\n", status, resp.StatusCode, resp.Header, resp.ContentLength)

	//Output:
	//test: NewJsonResponse(nil,nil) -> [status:OK] [status-code:200] [content-length:0]
	//test: NewJsonResponse(testCore,nil) -> [status:OK] [status-code:200] [header:map[Content-Type:[application/json]]] [content-length:272]
	//test: NewJsonResponse(testCore,nil) -> [status:OK] [status-code:200] [header:map[Content-Location:[http://localhost:8081/search?q=golang] Content-Type:[application/json]]] [content-length:272]

}


*/

func ExampleNewResponseWithBody() {
	h := make(http.Header)
	h.Add(ContentType, ContentTypeJson)
	resp, status := NewResponse(http.StatusOK, h, testCore)
	fmt.Printf("test: ResponseBody() -> [status:%v] [status-code:%v] [header:%v] [content-length:%v]\n", status, resp.StatusCode, resp.Header, resp.ContentLength)

	//Output:
	//test: ResponseBody() -> [status:OK] [status-code:200] [header:map[Content-Type:[application/json]]] [content-length:272]

}
