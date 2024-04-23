package httpx

import (
	"fmt"
	"net/http"
)

func ExampleValidateRequest() {
	_, status := ValidateRequest(nil, "")
	fmt.Printf("test: ValidateRequest(nil,\"\") -> [status:%v]\n", status)

	path := "test"
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	_, status = ValidateRequest(req, path)
	fmt.Printf("test: ValidateRequest(req,%v) -> [status:%v]\n", path, status)

	path = "github.com/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/search?q=golang", nil)
	_, status = ValidateRequest(req, path)
	fmt.Printf("test: ValidateRequest(req,%v) -> [status:%v]\n", path, status)

	path = "github.com/advanced-go/http2"
	req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/github.com/advanced-go/http2", nil)
	_, status = ValidateRequest(req, path)
	fmt.Printf("test: ValidateRequest(req,%v) -> [status:%v] [%v]\n", path, status, status.Err)

	//path = "github.com/advanced-go/http2"
	//req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/github.com/advanced-go/http2:entry", nil)
	//_, status = ValidateRequest(req, path)
	//fmt.Printf("test: ValidateRequest(req,%v) -> [status:%v] [%v]\n", path, status, status.FirstError())

	//path = "github.com/advanced-go/http2"
	//req, _ = http.NewRequest(http.MethodGet, "https://www.google.com/github.com/advanced-go/http2:entry", nil)
	//_, status = ValidateRequest(req, path)
	//fmt.Printf("test: ValidateRequest(req,%v) -> [status:%v] [%v]\n", path, status, status.FirstError())

	//Output:
	//test: ValidateRequest(nil,"") -> [status:Invalid Argument [error: Request is nil]]
	//test: ValidateRequest(req,test) -> [status:Bad Request [error: invalid URI, path is not valid: ""]]
	//test: ValidateRequest(req,github.com/advanced-go/http2) -> [status:Bad Request [error: invalid URI, NID does not match: "/search" "github.com/advanced-go/http2"]]
	//test: ValidateRequest(req,github.com/advanced-go/http2) -> [status:OK] [<nil>]

}
