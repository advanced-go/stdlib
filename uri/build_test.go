package uri

import (
	"fmt"
	"net/http"
)

func ExampleBuildURL_Host() {
	uri := "/search?q=golang"
	host := ""
	authority := ""

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	url := BuildURL(host, authority, req.URL)
	fmt.Printf("test: BuildURL(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, host, authority, url)

	host = "localhost:8080"
	req, _ = http.NewRequest(http.MethodGet, uri, nil)
	url = BuildURL(host, authority, req.URL)
	fmt.Printf("test: BuildURL(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, host, authority, url)

	uri = "/update"
	host = "www.google.com"
	req, _ = http.NewRequest(http.MethodGet, uri, nil)
	url = BuildURL(host, authority, req.URL)
	fmt.Printf("test: BuildURL\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, host, authority, url)

	//Output:
	//test: BuildURL("/search?q=golang") [host:] [auth:] [url:http://localhost/search?q=golang]
	//test: BuildURL("/search?q=golang") [host:localhost:8080] [auth:] [url:http://localhost:8080/search?q=golang]
	//test: BuildURL"/update") [host:www.google.com] [auth:] [url:https://www.google.com/update]

}

func ExampleBuildURL_Authority() {
	uri := "/google?q=golang"
	host := ""
	authority := "github/advanced-go/search"

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	url := BuildURL(host, authority, req.URL)
	fmt.Printf("test: BuildUri(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, host, authority, url)

	host = "www.google.com"
	req, _ = http.NewRequest(http.MethodGet, uri, nil)
	url = BuildURL(host, authority, req.URL)
	fmt.Printf("test: BuildUri(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, host, authority, url)

	/*
		// localhost
		rsc = NewPrimaryResource("localhost:8080", "", 0, "/health/liveness", httpCall)
		req, _ = http.NewRequest(http.MethodGet, uri, nil)
		url = rsc.BuildUri(req.URL)
		fmt.Printf("test: BuildUri(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, rsc.Host, rsc.Authority, url)

		// non-localhost
		uri = "/update"
		rsc = NewPrimaryResource("www.google.com", "", 0, "/health/liveness", httpCall)
		req, _ = http.NewRequest(http.MethodGet, uri, nil)
		url = rsc.BuildUri(req.URL)
		fmt.Printf("test: BuildUri(\"%v\") [host:%v] [auth:%v] [url:%v]\n", uri, rsc.Host, rsc.Authority, url

	*/

	//Output:
	//test: BuildUri("/google?q=golang") [host:] [auth:github/advanced-go/search] [url:http://localhost/github/advanced-go/search:google?q=golang]
	//test: BuildUri("/google?q=golang") [host:www.google.com] [auth:github/advanced-go/search] [url:https://www.google.com/github/advanced-go/search:google?q=golang]

}
