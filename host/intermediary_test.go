package host

import (
	"bytes"
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/access"
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

func ExampleConditionalIntermediary_Nil() {
	ic := NewConditionalIntermediary(nil, nil, nil)
	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q-golang", nil)

	_, status := ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-auth -> [status:%v]\n", status)

	ic = NewConditionalIntermediary(authTestExchange, nil, nil)
	_, status = ic(r)
	fmt.Printf("test: ConditionalIntermediary()-nil-service -> [status:%v]\n", status)

	//Output:
	//test: ConditionalIntermediary()-nil-auth -> [status:Bad Request [error: Conditional Intermediary HttpExchange 1 is nil]]
	//test: ConditionalIntermediary()-nil-service -> [status:Bad Request [error: Conditional Intermediary HttpExchange 2 is nil]]

}

func ExampleConditionalIntermediary_AuthExchange() {
	ic := NewConditionalIntermediary(authTestExchange, serviceTestExchange, nil)
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

func ExampleAccessLogIntermediary() {
	ic := NewAccessLogIntermediary(access.InternalTraffic, testDo)

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

func proxyDo(r *http.Request) (*http.Response, *core.Status) {
	//req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, err := http.DefaultClient.Do(r)
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

func ExampleProxyIntermediary() {
	host := "www.search.yahoo.com"
	proxy := NewProxyIntermediary(host, proxyDo)

	r, _ := http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, status := proxy(r)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Printf("test: ProxyIntermediary()-OK -> [status:%v] [content:%v]\n", status, len(buf) > 0)

	/*
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5)
		defer cancel()
		r, _ = http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/search?q-golang", nil)
		resp, status = ic(r)
		buf = nil
		if resp.Body != nil {
			buf, _ = io.ReadAll(resp.Body)
		}
		fmt.Printf("test: AccessLogIntermediary()-Gateway-Timeout -> [status:%v] [content:%v]\n", status, string(buf))


	*/

	//Output:
	//test: ProxyIntermediary()-OK -> [status:OK] [content:true]

}
