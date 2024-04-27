package host

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func appHttpExchange(r *http.Request) (*http.Response, *core.Status) {
	status := core.NewStatus(http.StatusTeapot)
	return &http.Response{StatusCode: status.Code}, status
}

/*
func testAuthExchangeOK(r *http.Request) (*http.Response, *core.Status) {
	//fmt.Fprint(w, "OK")
	status := core.StatusOK()
	return &http.Response{StatusCode: status.Code, Body: io.NopCloser(bytes.NewReader([]byte("OK")))}, status
}

func testAuthExchangeFail(r *http.Request) (*http.Response, *core.Status) {
	//fmt.Fprint(w, "Missing authorization header")
	status := core.NewStatus(http.StatusUnauthorized)
	return &http.Response{StatusCode: status.Code, Body: io.NopCloser(bytes.NewReader([]byte("Missing authorization header")))}, status
}


*/

func testDo(r *http.Request) (*http.Response, *core.Status) {
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp == nil {
		resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]")))}
		return resp, core.NewStatus(http.StatusGatewayTimeout)
	} else {
		resp.Body = io.NopCloser(bytes.NewReader([]byte("200 OK")))
		return resp, core.NewStatus(resp.StatusCode)
	}
}

func Example_TestHandler2() {
	pattern := "github/advanced-go/example-domain/activity"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/activity:entry", nil)

	RegisterExchange(pattern, appHttpExchange)

	rec := httptest.NewRecorder()
	HttpHandler2(rec, r)

	fmt.Printf("test: HttpHandler() -> %v\n", rec.Result().StatusCode)

	//Output:
	//test: HttpHandler() -> 418

}

func Example_Host_TestExchange_OK2() {
	pattern := "github/advanced-go/example-domain/slo"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/slo:entry", nil)

	SetHostTimeout2(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler2(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func _Example_Host_TestExchange_Timeout2() {
	pattern := "github/advanced-go/example-domain/timeseries"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/timeseries:entry", nil)

	SetHostTimeout2(time.Millisecond)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler2(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}
