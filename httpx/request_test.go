package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func ExampleValidateURL_Invalid() {
	_, _, status := ValidateURL(nil, "")
	fmt.Printf("test: ValidateURL(nil,\"\") -> [status:%v]\n", status)

	path := "test"
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, _, status = ValidateURL(req.URL, path)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, _, status = ValidateURL(req.URL, "")
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	_, _, status = ValidateURL(req.URL, path)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/github/advanced-go/http2", nil)
	_, _, status = ValidateURL(req.URL, path)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	//Output:
	//test: ValidateURL(nil,"") -> [status:Invalid Argument [error: request is nil]]
	//test: ValidateURL("","test") -> [status:Bad Request [error: invalid input, URI is empty]]
	//test: ValidateURL("","test") -> [status:Invalid Argument [error: authority is empty]]
	//test: ValidateURL("/search","github/advanced-go/http2") -> [status:Bad Request [error: invalid URI, authority does not match: "/search" "github/advanced-go/http2"]]
	//test: ValidateURL("/github/advanced-go/http2","github/advanced-go/http2") -> [status:Bad Request [error: invalid URI, path only contains an authority: "/github/advanced-go/http2"]]

}

func ExampleValidateRequest() {
	auth := "github/advanced-go/httpx"
	rsc := ":search?q=golang"
	uri := "https://www.google.com/" + auth + rsc
	rsc2 := ":v1/search?q=golang"
	uri2 := "https://www.google.com/" + auth + rsc2

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	ver, path, status := ValidateURL(req.URL, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", uri, auth, status, ver, path)

	req, _ = http.NewRequest(http.MethodGet, uri2, nil)
	ver, path, status = ValidateURL(req.URL, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", uri2, auth, status, ver, path)

	//Output:
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:search?q=golang","github/advanced-go/httpx") -> [status:OK] [ver:] [path:search]
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:v1/search?q=golang","github/advanced-go/httpx") -> [status:OK] [ver:v1] [path:search]

}
func ExampleValidateURL_Authority() {
	auth := "github/advanced-go/stdlib"
	req, _ := http.NewRequest(http.MethodGet, core.AuthorityRootPath, nil)
	ver, path, status := ValidateURL(req.URL, auth)
	fmt.Printf("test: ValidateURL(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", core.AuthorityPath, auth, status, ver, path)

	//Output:
	//test: ValidateURL("info","github/advanced-go/stdlib") -> [status:OK] [ver:] [path:info]

}
