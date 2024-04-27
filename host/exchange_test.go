package host

import (
	"bytes"
	"context"
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

func testAuthExchangeOK(_ *http.Request) (*http.Response, *core.Status) {
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprint(w, "OK")
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("200 OK")))}, core.StatusOK()
}

func testAuthExchangeFail(_ *http.Request) (*http.Response, *core.Status) {
	//w.WriteHeader(http.StatusUnauthorized)
	//fmt.Fprint(w, "Missing authorization header")
	return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(bytes.NewReader([]byte("Missing authorization header")))}, core.NewStatus(http.StatusUnauthorized)
}

func testDo(r *http.Request) (*http.Response, *core.Status) {
	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
	}
	if resp == nil {
		resp = &http.Response{StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(bytes.NewReader([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]")))}
		return resp, core.NewStatus(http.StatusGatewayTimeout)
	} else {
		resp.Body = io.NopCloser(bytes.NewReader([]byte("200 OK")))
		return resp, core.NewStatus(resp.StatusCode)
	}
}

func ExampleHttpHandler2() {
	pattern := "github/advanced-go/example-domain/activity"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/activity:entry", nil)

	RegisterExchange(pattern, appHttpExchange)

	rec := httptest.NewRecorder()
	HttpHandler2(rec, r)

	fmt.Printf("test: HttpHandler() -> %v\n", rec.Result().StatusCode)

	//Output:
	//test: HttpHandler() -> 418

}

func ExampleHttpHandler_Host_OK() {
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

func ExampleHttpHandler_Host_Timeout() {
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

func ExampleHttpHandler_Auth_Authorized() {
	pattern := "github/advanced-go/example-domain/host-auth-ok"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/host-auth-ok:entry", nil)

	SetAuthExchange(testAuthExchangeOK, nil)
	SetHostTimeout2(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler2(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func ExampleHttpHandler_Auth_Unauthorized() {
	pattern := "github/advanced-go/example-domain/host-auth-unauthorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/host-auth-unauthorized:entry", nil)

	SetAuthExchange(testAuthExchangeFail, nil)
	SetHostTimeout2(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler2(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:401] [content:Missing authorization header]

}
