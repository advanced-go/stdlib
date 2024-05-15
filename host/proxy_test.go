package host

import (
	"fmt"
)

func ExampleRegister() {
	proxy := NewProxy()
	path := "http://localhost:8080/github.com/advanced-go/example-domain/activity"

	err := proxy.register("", nil)
	fmt.Printf("test: Register(\"\") -> [%v]\n", err)

	err = proxy.register(path, nil)
	fmt.Printf("test: Register(%v) -> [%v]\n", path, err)

	err = proxy.register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [%v]\n", path, err)

	err = proxy.register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [%v]\n", path, err)

	path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	err = proxy.register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [%v]\n", path, err)

	//Output:
	//test: Register("") -> [error: authority is empty]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [error: HTTP Exchange is nil for authority : [http://localhost:8080/github.com/advanced-go/example-domain/activity]]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [<nil>]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [error: HTTP Exchange already exists for authority : [http://localhost:8080/github.com/advanced-go/example-domain/activity]]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [<nil>]

}

func ExampleLookup() {
	proxy := NewProxy()
	path := "http://localhost:8080/github.com/advanced-go/example-domain/activity"

	p := proxy.lookup("")
	fmt.Printf("test: Lookup(\"\") -> [proxy:%v]\n", p != nil)

	p = proxy.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [proxy:%v]\n", path, p != nil)

	err := proxy.register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	handler := proxy.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [handler:%v]\n", path, handler != nil)

	path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	err = proxy.register(path, appHttpExchange)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)
	handler = proxy.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [handler:%v]\n", path, handler != nil)

	//Output:
	//test: Lookup("") -> [proxy:false]
	//test: Lookup(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [proxy:false]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Lookup(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [handler:true]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Lookup(http://localhost:8080/github/advanced-go/example-domain/activity) -> [handler:true]

}
