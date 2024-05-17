package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func ExampleValidateRequestURL_Invalid() {
	_, _, status := ValidateRequestURL(nil, "")
	fmt.Printf("test: ValidateRequestURL(nil,\"\") -> [status:%v]\n", status)

	path := "test"
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, _, status = ValidateRequestURL(req, path)
	fmt.Printf("test: ValidateRequestURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, _, status = ValidateRequestURL(req, "")
	fmt.Printf("test: ValidateRequestURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	_, _, status = ValidateRequestURL(req, path)
	fmt.Printf("test: ValidateRequestURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/github/advanced-go/http2", nil)
	_, _, status = ValidateRequestURL(req, path)
	fmt.Printf("test: ValidateRequestURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	//Output:
	//test: ValidateRequestURL(nil,"") -> [status:Invalid Argument [error: request is nil]]
	//test: ValidateRequestURL("","test") -> [status:Bad Request [error: invalid input, URI is empty]]
	//test: ValidateRequestURL("","test") -> [status:Invalid Argument [error: authority is empty]]
	//test: ValidateRequestURL("/search","github/advanced-go/http2") -> [status:Bad Request [error: invalid URI, authority does not match: "/search" "github/advanced-go/http2"]]
	//test: ValidateRequestURL("/github/advanced-go/http2","github/advanced-go/http2") -> [status:Bad Request [error: invalid URI, path only contains an authority: "/github/advanced-go/http2"]]

}

func ExampleValidateRequest() {
	auth := "github/advanced-go/httpx"
	rsc := ":search?q=golang"
	uri := "https://www.google.com/" + auth + rsc
	rsc2 := ":v1/search?q=golang"
	uri2 := "https://www.google.com/" + auth + rsc2

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	ver, path, status := ValidateRequestURL(req, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", uri, auth, status, ver, path)

	req, _ = http.NewRequest(http.MethodGet, uri2, nil)
	ver, path, status = ValidateRequestURL(req, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", uri2, auth, status, ver, path)

	//Output:
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:search?q=golang","github/advanced-go/httpx") -> [status:OK] [ver:] [path:search]
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:v1/search?q=golang","github/advanced-go/httpx") -> [status:OK] [ver:v1] [path:search]

}
func ExampleValidateRequest_Info() {
	auth := "github/advanced-go/stdlib"
	req, _ := http.NewRequest(core.MethodInfo, core.InfoPath, nil)
	ver, path, status := ValidateRequestURL(req, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", core.InfoPath, auth, status, ver, path)

	//Output:
	//test: ValidateRequest("info","github/advanced-go/stdlib") -> [status:OK] [ver:] [path:info]

}
