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

func ExampleNewResponse_Error() {
	status := core.NewStatus(http.StatusGatewayTimeout)
	resp := NewResponse(status, status.Err)
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	status = core.NewStatusError(http.StatusGatewayTimeout, errors.New("Deadline Exceeded"))
	resp = NewResponse(status, status.Err)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:504] [content:]
	//test: NewResponse() -> [status-code:504] [content:Deadline Exceeded]

}

func ExampleNewResponse() {
	resp := NewResponse(nil, "")
	fmt.Printf("test: NewResponse() -> [status-code:%v]\n", resp.StatusCode)

	resp = NewResponse(core.StatusOK(), "")
	fmt.Printf("test: NewResponse() -> [status-code:%v]\n", resp.StatusCode)

	resp = NewResponse(core.StatusOK(), "version 1.2.35")
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewResponse() -> [status-code:400]
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
	resp, status := NewNotFoundResponseWithStatus()
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewNotFoundResponse() -> [status-code:%v] [status:%v] [content:%v]\n", resp.StatusCode, status, string(buf))

	//Output:
	//test: NewNotFoundResponse() -> [status-code:404] [status:Not Found] [content:Not Found]

}
