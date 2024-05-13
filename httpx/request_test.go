package httpx

import (
	"fmt"
	"net/http"
)

func ExampleValidateRequest() {
	_, _, status := ValidateRequest(nil, "")
	fmt.Printf("test: ValidateRequest(nil,\"\") -> [status:%v]\n", status)

	path := "test"
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, _, status = ValidateRequest(req, path)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	_, _, status = ValidateRequest(req, path)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v]\n", req.URL.Path, path, status)

	path = "github/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/github/advanced-go/http2", nil)
	_, _, status = ValidateRequest(req, path)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [%v]\n", req.URL.Path, path, status, status.Err)

	//Output:
	//test: ValidateRequest(nil,"") -> [status:Invalid Argument [error: request is nil]]
	//test: ValidateRequest("","test") -> [status:Bad Request [error: invalid URI, path is not valid: ""]]
	//test: ValidateRequest("/search","github/advanced-go/http2") -> [status:Bad Request [error: invalid URI, authority does not match: "/search" "github/advanced-go/http2"]]
	//test: ValidateRequest("/github/advanced-go/http2","github/advanced-go/http2") -> [status:OK] [<nil>]

}

func ExampleValidateRequest_Version() {
	auth := "github/advanced-go/httpx"
	rsc := ":search?q=golang"
	uri := "https://www.google.com/" + auth + rsc
	rsc2 := ":v1/search?q=golang"
	uri2 := "https://www.google.com/" + auth + rsc2

	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	ver, path, status := ValidateRequest(req, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", uri, auth, status, ver, path)

	req, _ = http.NewRequest(http.MethodGet, uri2, nil)
	ver, path, status = ValidateRequest(req, auth)
	fmt.Printf("test: ValidateRequest(\"%v\",\"%v\") -> [status:%v] [ver:%v] [path:%v]\n", uri2, auth, status, ver, path)

	//Output:
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:search?q=golang","github/advanced-go/httpx") -> [status:OK] [ver:] [path:search]
	//test: ValidateRequest("https://www.google.com/github/advanced-go/httpx:v1/search?q=golang","github/advanced-go/httpx") -> [status:OK] [ver:v1] [path:search]

}
