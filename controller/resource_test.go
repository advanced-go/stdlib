package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func ExampleResource_BuildURL() {
	silent := false
	uri := "/search?q=golang"

	// No host, default to localhost
	rsc := NewPrimaryResource(silent, "", "", 0, "", nil)
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	url := rsc.BuildURL(req.URL)
	fmt.Printf("test: BuildURL(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, rsc.Host, rsc.Authority, url)

	// localhost
	rsc = NewPrimaryResource(silent, "localhost:8080", "", 0, "", nil)
	req, _ = http.NewRequest(http.MethodGet, uri, nil)
	url = rsc.BuildURL(req.URL)
	fmt.Printf("test: BuildURL(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, rsc.Host, rsc.Authority, url)

	// non-localhost
	uri = "/update"
	rsc = NewPrimaryResource(silent, "www.google.com", "", 0, "", nil)
	req, _ = http.NewRequest(http.MethodGet, uri, nil)
	url = rsc.BuildURL(req.URL)
	fmt.Printf("test: BuildURL(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, rsc.Host, rsc.Authority, url)

	// authority
	uri = "/update"
	rsc = NewPrimaryResource(silent, "www.google.com", "github/advanced-go/search", 0, "", nil)
	req, _ = http.NewRequest(http.MethodGet, uri, nil)
	url = rsc.BuildURL(req.URL)
	fmt.Printf("test: BuildURL(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, rsc.Host, rsc.Authority, url)

	//Output:
	//test: BuildURL("/search?q=golang") [host:] [auth:] [url:http://localhost/search?q=golang]
	//test: BuildURL("/search?q=golang") [host:localhost:8080] [auth:] [url:http://localhost:8080/search?q=golang]
	//test: BuildURL("/update") [host:www.google.com] [auth:] [url:https://www.google.com/update]
	//test: BuildURL("/update") [host:www.google.com] [auth:github/advanced-go/search] [url:https://www.google.com/github/advanced-go/search:update]

}

func ExampleTimeout() {
	silent := false
	var dIn time.Duration = -1
	uri := "/search?q=golang"
	rsc := NewPrimaryResource(silent, "localhost:8080", "", dIn, "/health/liveness", httpCall)

	dOut := rsc.timeout(nil)
	fmt.Printf("test: timeout(nil) -> [timeout:%v]\n", dOut)

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	dOut = rsc.timeout(req)
	fmt.Printf("test: timeout(req) -> [duration:%v] [timeout:%v]\n", dIn, dOut)

	dIn = time.Millisecond * 100
	rsc = NewPrimaryResource(silent, "localhost:8080", "", dIn, "/health/liveness", httpCall)
	dOut = rsc.timeout(req)
	fmt.Printf("test: timeout(req) -> [duration:%v] [timeout:%v]\n", dIn, dOut)

	//Output:
	//test: timeout(nil) -> [timeout:0s]
	//test: timeout(req) -> [duration:-1ns] [timeout:0s]
	//test: timeout(req) -> [duration:100ms] [timeout:100ms]

}

func ExampleTimeout_Deadline() {
	silent := false
	dIn := time.Millisecond * 200
	deadline := time.Millisecond * 100
	uri := "/search?q=golang"
	rsc := NewPrimaryResource(silent, "localhost:8080", "", dIn, "/health/liveness", httpCall)

	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	dOut := rsc.timeout(req)
	fmt.Printf("test: timeout(req) -> [duration:%v] [deadline:%v] [timeout:%v]\n", dIn, deadline, dOut)

	dIn = time.Millisecond * 100
	rsc = NewPrimaryResource(silent, "localhost:8080", "", dIn, "/health/liveness", httpCall)
	dOut = rsc.timeout(req)
	fmt.Printf("test: timeout(req) -> [duration:%v] [deadline:%v] [timeout:%v]\n", dIn, deadline, dOut)

	dIn = time.Millisecond * 50
	rsc = NewPrimaryResource(silent, "localhost:8080", "", dIn, "/health/liveness", httpCall)
	dOut = rsc.timeout(req)
	fmt.Printf("test: timeout(req) -> [duration:%v] [deadline:%v] [timeout:%v]\n", dIn, deadline, dOut)

	//Output:
	//test: timeout(req) -> [duration:200ms] [deadline:100ms] [timeout:-100ms]
	//test: timeout(req) -> [duration:100ms] [deadline:100ms] [timeout:-100ms]
	//test: timeout(req) -> [duration:50ms] [deadline:100ms] [timeout:50ms]

}
