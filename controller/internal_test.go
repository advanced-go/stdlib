package controller

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
	"time"
)

func httpCall(r *http.Request) (resp *http.Response, status *core.Status) {
	cnt := 0
	var err0 error
	var err2 error
	var err1 error
	//var buf []byte

	resp, err0 = http.DefaultClient.Do(r)
	if err0 != nil {
		resp = new(http.Response)
		if r.Context().Err() == context.DeadlineExceeded {
			//w.WriteHeader(http.StatusGatewayTimeout)
			status = core.NewStatus(http.StatusGatewayTimeout)
		} else {
			//w.WriteHeader(http.StatusInternalServerError)
			status = core.NewStatus(http.StatusInternalServerError)
		}
		resp.StatusCode = status.Code
	} else {
		_, err1 = io.ReadAll(resp.Body)
		if err1 != nil {
			if err1 == context.DeadlineExceeded {
				//w.WriteHeader(http.StatusGatewayTimeout)
				status = core.NewStatus(http.StatusGatewayTimeout)
			} else {
				//w.WriteHeader(http.StatusInternalServerError)
				status = core.NewStatus(http.StatusInternalServerError)
			}
		} else {
			//w.WriteHeader(http.StatusOK)
			status = core.StatusOK()
			//cnt, err2 = w.Write(buf)
			//  resp.

		}
	}
	fmt.Printf("test: httpCall() -> [content:%v] [do-err:%v] [read-err:%v] [write-err:%v]\n", cnt, err0, err1, err2)
	return
}

func ExampleDoInternal() {
	uri := "https://www.google.com/search?q=golang"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	_, resp, status := doInternal(0, httpCall, req)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: DoInternal_0s() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	_, resp, status = doInternal(time.Second*5, httpCall, req)
	buf = nil
	buf, _ = io.ReadAll(resp.Body)
	fmt.Printf("test: DoInternal_5s() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	_, resp, status = doInternal(time.Millisecond*5, httpCall, req)
	buf = nil
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: DoInternal_5ms() -> [status-code:%v] [status:%v] [buff:%v]\n", resp.StatusCode, status, len(buf) > 0)

	//Output:
	//test: httpCall() -> [content:true] [do-err:<nil>] [read-err:<nil>] [write-err:<nil>]
	//test: DoInternal_0s() -> [status-code:200] [status:OK] [buff:true]
	//test: httpCall() -> [content:true] [do-err:<nil>] [read-err:<nil>] [write-err:<nil>]
	//test: DoInternal_5s() -> [status-code:200] [status:OK] [buff:true]
	//test: httpCall() -> [content:false] [do-err:Get "https://www.google.com/search?q=golang": context deadline exceeded] [read-err:<nil>] [write-err:<nil>]
	//test: DoInternal_5ms() -> [status-code:504] [status:Timeout] [buff:false]

}
