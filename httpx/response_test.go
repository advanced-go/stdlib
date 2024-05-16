package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

func ExampleNewErrorResponse() {
	status := core.NewStatus(http.StatusGatewayTimeout)
	resp := NewErrorResponse(status)
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewErrorResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	status = core.NewStatusError(http.StatusGatewayTimeout, errors.New("Deadline Exceeded"))
	resp = NewErrorResponse(status)
	buf, _ = io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewErrorResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewErrorResponse() -> [status-code:504] [content:]
	//test: NewErrorResponse() -> [status-code:504] [content:Deadline Exceeded]

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
