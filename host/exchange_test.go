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

/*
	func ExampleRegisterAuthority() {
		authority := []PathHandler{
			{"test", nil},
		}
		err := RegisterAuthority(authority)
		fmt.Printf("test: RegisterAuthority() -> [%v]\n", err)

		authority = []PathHandler{
			{"test", testAuthExchangeOK},
			{"test2", testAuthExchangeOK},
		}
		err = RegisterAuthority(authority)
		fmt.Printf("test: RegisterAuthority() -> [%v]\n", err)

		//Output:
		//test: RegisterAuthority() -> [error: handler for path [test] is nil]
		//test: RegisterAuthority() -> [<nil>]

}
*/

func ExampleRegisterExchange() {
	a := "github/advanced-go/stdlib"

	err := RegisterExchange("", nil)
	fmt.Printf("test: RegisterExchange(_,nil) -> [%v]\n", err)

	err = RegisterExchange(a, nil)
	fmt.Printf("test: RegisterExchange(\"%v\",nil) -> [%v]\n", a, err)

	err = RegisterExchange(a, appHttpExchange)
	fmt.Printf("test: RegisterExchange(\"%v\",appHttpExchange) -> [%v]\n", a, err)

	h := exchangeProxy.lookup(a)
	fmt.Printf("test: Lookup(\"%v\") -> [ok:%v]\n", a, h != nil)

	err = RegisterExchange(a, appHttpExchange)
	fmt.Printf("test: RegisterExchange(\"%v\",appHttpExchange) -> [%v]\n", a, err)

	//Output:
	//test: RegisterExchange(_,nil) -> [error: authority is empty]
	//test: RegisterExchange("github/advanced-go/stdlib",nil) -> [error: HTTP Exchange is nil for authority : [github/advanced-go/stdlib]]
	//test: RegisterExchange("github/advanced-go/stdlib",appHttpExchange) -> [<nil>]
	//test: Lookup("github/advanced-go/stdlib") -> [ok:true]
	//test: RegisterExchange("github/advanced-go/stdlib",appHttpExchange) -> [error: HTTP Exchange already exists for authority : [github/advanced-go/stdlib]]

}

func ExampleHttpHandler() {
	pattern := "github/advanced-go/host/HttpHandler"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/HttpHandler:entry", nil)

	RegisterExchange(pattern, appHttpExchange)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)

	fmt.Printf("test: HttpHandler() -> %v\n", rec.Result().StatusCode)

	//Output:
	//test: HttpHandler() -> 418

}

func ExampleHttpHandler_Host_OK() {
	pattern := "github/advanced-go/host/ok"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/ok:entry", nil)

	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func ExampleHttpHandler_Host_Timeout() {
	pattern := "github/advanced-go/host/timeout"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/timeout:entry", nil)

	SetHostTimeout(time.Millisecond)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func ExampleHttpHandler_Auth_Authorized() {
	pattern := "github/advanced-go/host/authorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/authorized:entry", nil)

	SetAuthExchange(testAuthExchangeOK, nil)
	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func ExampleHttpHandler_Auth_Unauthorized() {
	pattern := "github/advanced-go/host/unauthorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/unauthorized:entry", nil)

	SetAuthExchange(testAuthExchangeFail, nil)
	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testDo)

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:401] [content:Missing authorization header]

}

func ExampleHttpHandler_AccessLog_Service_OK() {
	pattern := "github/advanced-go/host/access-log-service-ok"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/access-log-service-ok:entry", nil)

	SetAuthExchange(testAuthExchangeOK, nil)
	SetHostTimeout(time.Second * 4)
	RegisterExchange(pattern, NewAccessLogIntermediary("log-route-ok", testDo))

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:200] [content:200 OK]

}

func ExampleHttpHandler_AccessLog_Service_Timeout() {
	pattern := "github/advanced-go/host/access-log-service-timeout"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/access-log-service-timeout:entry", nil)

	SetAuthExchange(testAuthExchangeOK, nil)
	SetHostTimeout(time.Millisecond * 4)
	RegisterExchange(pattern, NewAccessLogIntermediary("log-route-timeout", testDo))

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func ExampleHttpHandler_AccessLog_Service_Unauthorized() {
	pattern := "github/advanced-go/host/access-log-service-unauthorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/host/access-log-service-unauthorized:entry", nil)

	SetAuthExchange(testAuthExchangeFail, nil)
	SetHostTimeout(time.Second * 4)
	RegisterExchange(pattern, NewAccessLogIntermediary("log-route-unauthorized", testDo))

	rec := httptest.NewRecorder()
	HttpHandler(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpHandler() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpHandler() -> [status-code:401] [content:Missing authorization header]

}
