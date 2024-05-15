package controller

import (
	"fmt"
)

func ExampleControlsRegister() {
	ctrl := NewController("test-route", nil, nil)
	p := NewControls()
	path := "http://localhost:8080/github.com/advanced-go/example-domain/activity"

	err := p.registerWithAuthority("", nil)
	fmt.Printf("test: Register(\"\") -> [err:%v]\n", err)

	err = p.registerWithAuthority(path, nil)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	err = p.registerWithAuthority(path, ctrl)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	//err = p.registerWithAuthority(path, ctrl)
	//fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	err = p.registerWithAuthority(path, ctrl)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	//Output:
	//test: Register("") -> [err:invalid argument: authority is empty]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [err:invalid argument: Controller is nil for authority: [http://localhost:8080/github.com/advanced-go/example-domain/activity]]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:<nil>]

}

func ExampleControlsLookup() {
	ctrl := NewController("test-route", nil, nil)
	p := NewControls()
	path := "http://localhost:8080/github.com/advanced-go/example-domain/activity"

	_, status := p.lookup("")
	fmt.Printf("test: Lookup(\"\") -> [status:%v]\n", status)

	_, status = p.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [status:%v]\n", path, status)

	err := p.registerWithAuthority(path, ctrl)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	handler, status1 := p.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [status:%v] [handler:%v]\n", path, status1, handler != nil)

	path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	err = p.registerWithAuthority(path, ctrl)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)
	handler, status1 = p.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [status:%v] [handler:%v]\n", path, status1, handler != nil)

	//Output:
	//test: Lookup("") -> [status:Invalid Argument [error: invalid input, URI is empty]]
	//test: Lookup(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [status:Invalid Argument [invalid argument: Controller does not exist: [github.com/advanced-go/example-domain/activity]]]
	//test: Register(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Lookup(http://localhost:8080/github.com/advanced-go/example-domain/activity) -> [status:OK] [handler:true]
	//test: Register(http://localhost:8080/github/advanced-go/example-domain/activity) -> [err:<nil>]
	//test: Lookup(http://localhost:8080/github/advanced-go/example-domain/activity) -> [status:OK] [handler:true]

}
