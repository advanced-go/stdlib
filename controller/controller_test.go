package controller

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
	"time"
)

/*
func testDo(req *http.Request) (*http.Response, *core.Status) {
	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		return resp, core.NewStatus(resp.StatusCode)
	}
	resp = &http.Response{StatusCode: core.StatusDeadlineExceeded}
	return resp, core.NewStatusError(core.StatusDeadlineExceeded, err)
}
*/

func testDo(r *http.Request) (*http.Response, *core.Status) {
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if resp == nil {
			resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]")))}
			return resp, core.NewStatus(core.StatusDeadlineExceeded)
		}
		resp.Body = io.NopCloser(bytes.NewReader([]byte(err.Error())))
		return resp, core.NewStatus(http.StatusInternalServerError)
	}
	resp.Body = io.NopCloser(bytes.NewReader([]byte(fmt.Sprintf("%v OK", resp.StatusCode))))
	return resp, core.NewStatus(resp.StatusCode)
}

func ExampleDo_Error() {
	ctrl := NewController("google-search", NewPrimaryResource("https://www.google.com", "/health/liveness", 0, httpCall), nil)
	uri := "/search?q=golang"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	_, status := ctrl.Do(nil, req)
	fmt.Printf("test: Do(nil,req) -> [status:%v]\n", status)

	_, status = ctrl.Do(testDo, nil)
	fmt.Printf("test: Do(testDo,nil) -> [status:%v]\n", status)

	//Output:
	//test: Do(nil,req) -> [status:Invalid Argument [invalid argument : request is nil]]
	//test: Do(testDo,nil) -> [status:Invalid Argument [invalid argument : request is nil]]

}

func _ExampleDo_Internal() {
	ctrl := NewController("google-search", NewPrimaryResource("https://www.google.com", "/health/liveness", 0, httpCall), nil)
	uri := "/search?q=golang"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, status := ctrl.Do(nil, req)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: Do_0s() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	ctrl = NewController("google-search", NewPrimaryResource("https://www.google.com", "/health/liveness", time.Millisecond*5, httpCall), nil)
	resp, status = ctrl.Do(nil, req)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: Do_5ms() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	//Output:
	//test: httpCall() -> [content:134585] [do-err:<nil>] [read-err:<nil>] [write-err:<nil>]
	//test: Do_0s() -> [status-code:200] [status:OK] [buff:true]
	//test: httpCall() -> [content:0] [do-err:Get "https://www.google.com/search?q=golang": context deadline exceeded] [read-err:<nil>] [write-err:<nil>]
	//test: Do_5ms() -> [status-code:504] [status:Timeout] [buff:true]

}

func _ExampleDo_Internal_Deadline() {
	ctrl := NewController("google-search", NewPrimaryResource("https://www.google.com", "/health/liveness", 0, httpCall), nil)
	uri := "/search?q=golang"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, status := ctrl.Do(nil, req)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: Do_0s() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	ctrl = NewController("google-search", NewPrimaryResource("https://www.google.com", "/health/liveness", time.Millisecond*5, httpCall), nil)
	resp, status = ctrl.Do(nil, req)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: Do_5ms() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	//Output:
	//test: httpCall() -> [content:134585] [do-err:<nil>] [read-err:<nil>] [write-err:<nil>]
	//test: Do_0s() -> [status-code:200] [status:OK] [buff:true]
	//test: httpCall() -> [content:0] [do-err:Get "https://www.google.com/search?q=golang": context deadline exceeded] [read-err:<nil>] [write-err:<nil>]
	//test: Do_5ms() -> [status-code:504] [status:Timeout] [buff:true]

}

func _ExampleDo_Egress() {
	ctrl := NewController("google-search", NewPrimaryResource("https://www.google.com", "/health/liveness", 0, nil), nil)
	uri := "/search?q=golang"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, status := ctrl.Do(testDo, req)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: Do_0s() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	ctrl = NewController("google-search", NewPrimaryResource("https://www.google.com", "/health/liveness", time.Millisecond*5, nil), nil)
	resp, status = ctrl.Do(testDo, req)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: Do_5ms() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	//Output:
	//test: Do_0s() -> [status-code:200] [status:OK] [buff:true]
	//test: Do_5ms() -> [status-code:504] [status:Deadline Exceeded [context deadline exceeded]] [buff:true]

}
