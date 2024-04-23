package httpx

import (
	"fmt"
)

const (
	testLocation = "github/advanced-go/core/exchange:ExampleGet"
)

func ExampleGet() {
	//var e core.Output
	r, status := Get(nil, "", nil)

	//e.Handle(status, "123-456")
	fmt.Printf("test: Get(\"\") -> [resp:%v] [status:%v]\n", r.Status, status)

	//Output:
	//test: Get("") -> [resp:Internal Error] [status:Bad Request [error: URI is empty]]

}
