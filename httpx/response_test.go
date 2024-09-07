package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"net/http"
	"net/url"
)

const (
	testResonse = "file://[cwd]/httpxtest/resource/test-response.txt"
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
	resp, _ := NewResponse[core.Log](status.HttpCode(), nil, status.Err)
	buf, _ := io2.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	status = core.NewStatusError(http.StatusGatewayTimeout, errors.New("Deadline Exceeded"))
	resp, _ = NewResponse[core.Log](status.HttpCode(), nil, status.Err)
	buf, _ = io2.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:504] [content:]
	//test: NewResponse() -> [status-code:504] [content:Deadline Exceeded]

}

func ExampleNewResponse() {
	resp, _ := NewResponse[core.Log](http.StatusOK, nil, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v]\n", resp.StatusCode)

	resp, _ = NewResponse[core.Log](core.StatusOK().HttpCode(), nil, "version 1.2.35")
	buf, _ := io2.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:200]
	//test: NewResponse() -> [status-code:200] [content:version 1.2.35]

}

func ExampleNewVersionResponse() {
	resp := NewVersionResponse("7.8.9")
	buf, _ := io2.ReadAll(resp.Body, nil)
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
	buf, _ := io2.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewHealthResponseOK() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewHealthResponseOK() -> [status-code:200] [content:{
	// "status": "up"
	//}]

}

func ExampleNewNotFoundResponseWithStatus() {
	resp := NewNotFoundResponse()
	buf, _ := io2.ReadAll(resp.Body, nil)
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
	resp, _ := NewResponse[core.Log](http.StatusOK, h, testCore)
	fmt.Printf("test: ResponseBody() -> [status-code:%v] [header:%v] [content-length:%v]\n", resp.StatusCode, resp.Header, resp.ContentLength)

	//Output:
	//test: ResponseBody() -> [status-code:200] [header:map[Content-Type:[application/json]]] [content-length:272]

}

func readAll(body io.ReadCloser) ([]byte, *core.Status) {
	if body == nil {
		return nil, core.StatusOK()
	}
	defer body.Close()
	buf, err := io.ReadAll(body)
	if err != nil {
		return nil, core.NewStatusError(core.StatusIOError, err)
	}
	return buf, core.StatusOK()
}

func Example_NewResponseFromUri() {
	s := testResonse
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [status:%v] [statusCode:%v]\n", s, status0, resp.StatusCode)

	buf, status := readAll(resp.Body)
	fmt.Printf("test: readAll() -> [status:%v] [content-length:%v]\n", status, len(buf)) //string(buf))

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/resource/test-response.txt) -> [status:OK] [statusCode:200]
	//test: readAll() -> [status:OK] [content-length:56]

}

func Example_NewResponseFromUri_URL_Nil() {
	resp, status0 := NewResponseFromUri(nil)
	fmt.Printf("test: NewResponseFromUri(nil) -> [error:[%v]] [statusCode:%v]\n", status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(nil) -> [error:[error: URL is nil]] [statusCode:500]

}

func _Example_NewResponseFromUri_Invalid_Scheme() {
	s := "https://www.google.com/search?q=golang"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%vl) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(https://www.google.com/search?q=golangl) -> [error:[error: Invalid URL scheme : https]] [statusCode:500]

}

func Example_NewResponseFromUri_HTTP_Error() {
	s := "file://[cwd]/httpxtest/resource/message.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/resource/message.txt) -> [error:[malformed HTTP status code "text"]] [statusCode:500]

}

func Example_NewResponseFromUri_NotFound() {
	s := "file://[cwd]/httpxtest/resource/not-found.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/resource/not-found.txt) -> [error:[open C:\Users\markb\GitHub\stdlib\httpx\httpxtest\resource\not-found.txt: The system cannot find the file specified.]] [statusCode:404]

}

func Example_NewResponseFromUri_EOF_Error() {
	s := "file://[cwd]/httpxtest/resource/http-503-error.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/httpxtest/resource/http-503-error.txt) -> [error:[unexpected EOF]] [statusCode:500]

}

func ExampleNewError() {
	status := core.StatusOK()
	//var resp *http.Response

	err := NewError(nil, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, nil, err)

	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", core.StatusOK(), nil, err)

	status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: invalid content"))
	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [%v]\n", status, nil, err)

	resp, _ := NewResponse[core.Output](http.StatusTeapot, nil, nil)
	err = NewError(nil, resp)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, resp != nil, err)

	resp, _ = NewResponse[core.Output](http.StatusTeapot, nil, "error: response content")
	err = NewError(nil, resp)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [%v]\n", nil, resp != nil, err)

	//Output:
	//test: NewError() -> [status:<nil>] [resp:<nil>] [err:]
	//test: NewError() -> [status:OK] [resp:<nil>] [err:]
	//test: NewError() -> [status:Invalid Content [error: invalid content]] [resp:<nil>] [error: invalid content]
	//test: NewError() -> [status:<nil>] [resp:true] [err:]
	//test: NewError() -> [status:<nil>] [resp:true] [error: response content]

}
