package host

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func appHttpHandler(w http.ResponseWriter, r *http.Request) {
	status := core.NewStatus(http.StatusTeapot)
	w.WriteHeader(status.Code)
	//w.Write([]byte(status.String()))
}

func testAuthHandlerOK(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprint(w, "OK")
}

func testAuthHandlerFail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "Missing authorization header")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp == nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		w.Write([]byte("Timeout [Get \"https://www.google.com/search?q=golang\": context deadline exceeded]"))
	} else {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Status))
	}
}

func Example_TestHandler() {
	pattern := "github/advanced-go/example-domain/activity"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/activity:entry", nil)

	RegisterExchange(pattern, appHttpHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)

	fmt.Printf("test: HttpExchange() -> %v\n", rec.Result().StatusCode)

	//Output:
	//test: HttpExchange() -> 418

}

func Example_Host_TestExchange_OK() {
	pattern := "github/advanced-go/example-domain/slo"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/slo:entry", nil)

	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:200] [content:200 OK]

}

func Example_Host_TestExchange_Timeout() {
	pattern := "github/advanced-go/example-domain/timeseries"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/timeseries:entry", nil)

	SetHostTimeout(time.Millisecond)
	RegisterExchange(pattern, testHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func Example_Auth_TestExchange_OK() {
	pattern := "github/advanced-go/example-domain/auth-ok"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/auth-ok:entry", nil)

	SetHostTimeout(0)
	SetAuthExchange(testAuthHandlerOK, nil)
	RegisterExchange(pattern, testHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:200] [content:200 OK]

}

func Example_Auth_TestExchange_Fail() {
	pattern := "github/advanced-go/example-domain/auth-fail"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/auth-fail:entry", nil)

	SetAuthExchange(testAuthHandlerFail, nil)
	RegisterExchange(pattern, testHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:401] [content:Missing authorization header]

}

func Example_Host_Auth_TestExchange_OK() {
	pattern := "github/advanced-go/example-domain/host-auth-ok"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/host-auth-ok:entry", nil)

	SetAuthExchange(testAuthHandlerOK, nil)
	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:200] [content:200 OK]

}

func Example_Host_Auth_TestExchange_Timeout() {
	pattern := "github/advanced-go/example-domain/host-auth-timeout"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/host-auth-timeout:entry", nil)

	SetAuthExchange(testAuthHandlerOK, nil)
	SetHostTimeout(time.Millisecond * 2)
	RegisterExchange(pattern, testHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:504] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}

func Example_Host_Auth_TestExchange_Unauthorized() {
	pattern := "github/advanced-go/example-domain/host-auth-unauthorized"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/host-auth-unauthorized:entry", nil)

	SetAuthExchange(testAuthHandlerFail, nil)
	SetHostTimeout(time.Second * 2)
	RegisterExchange(pattern, testHandler)

	rec := httptest.NewRecorder()
	HttpExchange(rec, r)
	buf, _ := io.ReadAll(rec.Result().Body)
	fmt.Printf("test: HttpExchange() -> [status-code:%v] [content:%v]\n", rec.Result().StatusCode, string(buf))

	//Output:
	//test: HttpExchange() -> [status-code:401] [content:Missing authorization header]

}

func _ExamplePing() {
	uri1 := "github/advanced-go/example-domain/activity"
	r, _ := http.NewRequest("", "github/advanced-go/example-domain/activity:ping", nil)
	err := messaging.HostExchange.Add(messaging.NewMailbox(uri1, nil))
	if err != nil {
		fmt.Printf("test: processPing() -> [err:%v]\n", err)
	}
	//nid, rsc, ok := UprootUrn(r.URL.Path)
	status := messaging.Ping(nil, r.URL) //ProcessPing[core.Bypass](w, nid)
	fmt.Printf("test: messaging.Ping() -> [nid:%v] [nss:%v] [ok:%v] [status:%v]\n", "", "", true, status)

	//Output:
	//test: messaging.Ping() -> [nid:] [nss:] [ok:true] [status:504] [content:ping response time out: [github/advanced-go/example-domain/activity]]

}

func ExampleHttpExchange() {
	ok := exchange(func(w http.ResponseWriter, r *http.Request) {})
	fmt.Printf("test: HttpExchange(anonymous-function) -> [ok:%v|\n", ok)

	ok = exchange(handler2)
	fmt.Printf("test: HttpExchange(function) -> [ok:%v|\n", ok)

	ok = exchange(handler3())
	fmt.Printf("test: HttpExchange(return-function) -> [ok:%v|\n", ok)

	//Output:
	//test: HttpExchange(anonymous-function) -> [ok:true|
	//test: HttpExchange(function) -> [ok:true|
	//test: HttpExchange(return-function) -> [ok:true|

}

func exchange(fn core.HttpExchange) bool {
	if fn == nil {
		return false
	}
	return true
}

func handler2(w http.ResponseWriter, r *http.Request) {
}

func handler3() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
