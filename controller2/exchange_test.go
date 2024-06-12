package controller2

import (
	"fmt"
	uri2 "github.com/advanced-go/stdlib/uri"
	"io"
	"net/http"
	"time"
)

func ExampleExchange_Error() {
	//ctrl := NewController("google-search", NewPrimaryResource("www.google.com", "", 0,  httpCall), nil)
	//RegisterController(ctrl)
	_, status := Exchange(nil)
	fmt.Printf("test: Exchange(nil) -> [status:%v]\n", status)

	//Output:
	//test: Exchange(nil) -> [status:Invalid Argument [invalid argument : request is nil]]

}

func ExampleExchange_Internal() {
	//defer DisableLogging(true)()
	//authority := "github/advanced-go/search"
	ctrl := NewController("google-search", NewPrimaryResource("www.google.com", "", 0, httpCall), nil)
	uri := "https://www.google.com/search?" + uri2.BuildQuery("q=golang")
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	RegisterController(ctrl)
	resp, status := Exchange(req)
	var buf []byte
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: Exchange_0s() -> [status-code:%v] [status:%v] [buf:%v]\n", resp.StatusCode, status, len(buf) > 0)
	ctrlMap.remove("www.google.com")
	//ctrl = NewController("yahoo-search", NewPrimaryResource("www.search.yahoo.com", "", time.Millisecond*5, httpCall), nil)
	ctrl = NewController("google-search", NewPrimaryResource("www.google.com", "", time.Millisecond*5, httpCall), nil)
	err := RegisterController(ctrl)
	if err != nil {
		fmt.Printf("test: RegisterController() -> [err:%v]\n", err)
	}
	resp, status = Exchange(req)
	if status.OK() && resp.Body != nil {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: Exchange_5ms() -> [status-code:%v] [status:%v] [buf:%v]\n", resp.StatusCode, status, len(buf) > 0)

	//Output:
	//test: httpCall() -> [content:true] [do-err:<nil>] [read-err:<nil>]
	//test: Exchange_0s() -> [status-code:200] [status:OK] [buf:false]
	//test: httpCall() -> [content:false] [do-err:Get "https://www.google.com/search?q=golang": context deadline exceeded] [read-err:<nil>]
	//test: Exchange_5ms() -> [status-code:504] [status:Timeout] [buf:false]

}
