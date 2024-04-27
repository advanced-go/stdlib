package host

import (
	"bytes"
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"io"
	"net/http"
	"time"
)

func serviceTestExchange(_ *http.Request) (*http.Response, *core.Status) {
	//w.WriteHeader(http.StatusOK)
	//fmt.Fprint(w, "Service OK")
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("Service OK")))}, core.StatusOK()
}

func authTestExchange(r *http.Request) (*http.Response, *core.Status) {
	if r != nil {
		tokenString := r.Header.Get(Authorization)
		if tokenString == "" {
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprint(w, "Missing authorization header")
			return &http.Response{StatusCode: http.StatusUnauthorized, Body: io.NopCloser(bytes.NewReader([]byte("Missing authorization header")))}, core.NewStatus(http.StatusUnauthorized)
		}
	}
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte("Authorized")))}, core.StatusOK()
}

func ExampleConditionalIntermediary2_Nil() {
	ic := NewConditionalIntermediary2(nil, nil, nil)
	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)

	_, status := ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-auth -> [status:%v]\n", status)

	ic = NewConditionalIntermediary2(authTestExchange, nil, nil)
	_, status = ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-service -> [status:%v]\n", status)

	//Output:
	//test: ConditionalIntermediary()-nil-auth -> [status:Bad Request [error: Conditional Intermediary HttpExchange 1 is nil]]
	//test: ConditionalIntermediary()-nil-service -> [status:Bad Request [error: Conditional Intermediary HttpExchange 2 is nil]]

}

func ExampleConditionalIntermediary2_AuthExchange() {
	ic := NewConditionalIntermediary2(authTestExchange, serviceTestExchange, nil)
	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)

	resp, status := ic(r)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: ConditionalIntermediary()-auth-failure -> [status:%v] [content:%v]\n", status, string(buf))

	r.Header.Add(Authorization, "token")
	resp, status = ic(r)
	buf, _ = io.ReadAll(resp.Body)
	fmt.Printf("test: ConditionalIntermediary()-auth-success -> [status:%v] [content:%v]\n", status, string(buf))

	//Output:
	//test: ConditionalIntermediary()-auth-failure -> [status:Unauthorized] [content:Missing authorization header]
	//test: ConditionalIntermediary()-auth-success -> [status:OK] [content:Service OK]

}

func ExampleAccessLogIntermediary2() {
	ic := NewAccessLogIntermediary2("test-route", testDo)

	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)
	resp, status := ic(r)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: AccessLogIntermediary()-OK -> [status:%v] [content:%v]\n", status, len(buf) > 0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
	defer cancel()
	r, _ = http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/search?q-golang", nil)
	resp, status = ic(r)
	buf = nil
	if resp.Body != nil {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: AccessLogIntermediary()-Gateway-Timeout -> [status:%v] [content:%v]\n", status, string(buf))

	//Output:
	//test: AccessLogIntermediary()-OK -> [status:OK] [content:true]
	//test: AccessLogIntermediary()-Gateway-Timeout -> [status:Timeout] [content:Timeout [Get "https://www.google.com/search?q=golang": context deadline exceeded]]

}
