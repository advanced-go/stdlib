package core

import (
	"fmt"
	"net/http"
)

func appHttpExchange(r *http.Request) (*http.Response, *Status) {
	status := NewStatus(http.StatusTeapot)
	return &http.Response{StatusCode: status.Code}, status
}

func ExampleProxy_Register() {
	proxy := NewExchangeProxy()
	path := "http://localhost:8080/github/advanced-go/example-domain/activity"

	err := proxy.Register("", nil)
	fmt.Printf("test: Register(\"\") -> [err:%v]\n", err)

	err = proxy.Register(path, nil)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	err = proxy.Register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	err = proxy.Register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	err = proxy.Register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	//Output:
	//test: Register("") -> [err:invalid argument: authority is empty]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:invalid argument: HTTP Exchange is nil for authority : [http://localhost:8080/github/advanced-go/example-domain/activity]]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:invalid argument: HTTP Exchange already exists for authority : [http://localhost:8080/github/advanced-go/example-domain/activity]]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:invalid argument: HTTP Exchange already exists for authority : [http://localhost:8080/github/advanced-go/example-domain/activity]]

}

func ExampleProxy_Lookup() {
	proxy := NewExchangeProxy()
	path := "http://localhost:8080/github.com/advanced-go/example-domain/activity"

	p := proxy.Lookup("")
	fmt.Printf("test: Lookup(\"\") -> [proxy:%v]\n", p != nil)

	p = proxy.Lookup(path)
	fmt.Printf("test: Lookup(%v) -> [proxy:%v]\n", path, p != nil)

	err := proxy.Register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	handler := proxy.Lookup(path)
	fmt.Printf("test: Lookup(%v) -> [handler:%v]\n", path, handler != nil)

	path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	err = proxy.Register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)
	handler = proxy.Lookup(path)
	fmt.Printf("test: Lookup(%v) -> [handler:%v]\n", path, handler != nil)

	//Output:
	//test: Lookup("") -> [proxy:false]
	//test: Lookup(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [proxy:false]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Lookup(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [handler:true]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Lookup(http://localhost:8080/github/advanced-go/example-domain/activity) -> [handler:true]

}

func ExampleProxy_LookupByRequest() {
	proxy := NewExchangeProxy()
	req, _ := http.NewRequest("", "https://www.google.com/search?q=golang", nil)
	req2, _ := http.NewRequest("", "http://localhost:8081/github/advanced-go/search:google?q=golang", nil)
	host := "www.google.com"
	authority := "github/advanced-go/search"

	proxy.Register(host, appHttpExchange)
	p := proxy.LookupByRequest(req)
	fmt.Printf("test: Lookup(\"%v\") -> [proxy:%v]\n", req.URL.String(), p != nil)

	proxy.Register(authority, appHttpExchange)
	p = proxy.LookupByRequest(req2)
	fmt.Printf("test: Lookup(\"%v\") -> [proxy:%v]\n", req2.URL.String(), p != nil)

	//Output:
	//test: Lookup("https://www.google.com/search?q=golang") -> [proxy:true]
	//test: Lookup("http://localhost:8081/github/advanced-go/search:google?q=golang") -> [proxy:true]

}
