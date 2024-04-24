package core

import (
	"fmt"
	"net/http"
)

func ExampleHttpHandler() {
	ok := exchange(func(w http.ResponseWriter, r *http.Request) {})
	fmt.Printf("test: HttpHandler(anonymous-function) -> [ok:%v|\n", ok)

	ok = exchange(handler2)
	fmt.Printf("test: HttpHandler(function) -> [ok:%v|\n", ok)

	ok = exchange(handler3())
	fmt.Printf("test: HttpHandler(return-function) -> [ok:%v|\n", ok)

	//Output:
	//test: HttpHandler(anonymous-function) -> [ok:true|
	//test: HttpHandler(function) -> [ok:true|
	//test: HttpHandler(return-function) -> [ok:true|

}

func exchange(fn HttpHandler) bool {
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
