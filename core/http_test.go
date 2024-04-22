package core

import (
	"fmt"
	"net/http"
)

func ExampleHttpExchange() {
	ok := exchange(func(w http.ResponseWriter, r *http.Request) {})
	fmt.Printf("test: HttpExchange(anonymous-function) -> [ok:%v|\n", ok)

	ok = exchange(handler2)
	fmt.Printf("test: HttpExchange(function) -> [ok:%v|\n", ok)

	ok = exchange(handler3())
	fmt.Printf("test: HttpExchange(return-function) -> [ok:%v|\n", ok)

	//Output:
	//test: HttpExchange(anonymous-function) -> [ok:true|
	//test: HttpExchange(function) -> [ok:true|
	//test: HttpExchange(return-function) -> [ok:true|
	
}

func exchange(fn HttpExchange) bool {
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
