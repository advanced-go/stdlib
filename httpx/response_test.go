package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/io"
	"net/http"
)

func ExampleNewErrorResponse() {
	status := core.NewStatusError(http.StatusGatewayTimeout, errors.New("Deadline Exceeded"))
	resp := NewErrorResponse(status)
	buf, _ := io.ReadAll(resp.Body, nil)
	fmt.Printf("test: NewErrorResponse() -> [status-code:%v] [content:%v]\n", resp.StatusCode, string(buf))

	//Output:
	//test: NewErrorResponse() -> [status-code:504] [content:Deadline Exceeded]
	
}
