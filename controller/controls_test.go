package controller

import (
	"fmt"
)

func ExampleRegisterController() {
	err := RegisterController(nil)
	fmt.Printf("test: RegisterController(nil) -> [err:%v]\n", err)

	ctrl := NewController("test-route", nil, nil)
	ctrl.Router = nil
	err = RegisterController(ctrl)
	fmt.Printf("test: RegisterController(ctrl) -> [err:%v]\n", err)

	ctrl = NewController("test-route", nil, nil)
	err = RegisterController(ctrl)
	fmt.Printf("test: RegisterController(ctrl) -> [err:%v]\n", err)

	ctrl = NewController("test-route", NewPrimaryResource("", "", 0, "", nil), nil)
	err = RegisterController(ctrl)
	fmt.Printf("test: RegisterController(ctrl) -> [err:%v]\n", err)

	ctrl = NewController("test-route", NewPrimaryResource("localhost:8080", "", 0, "", nil), nil)
	err = RegisterController(ctrl)
	fmt.Printf("test: RegisterController(ctrl) -> [err:%v]\n", err)

	//err = p.registerWithAuthority(path, ctrl)
	//fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)
	//path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	//err = p.register(ctrl)
	//fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	//Output:
	//test: RegisterController(nil) -> [err:invalid argument: Controller is nil]
	//test: RegisterController(ctrl) -> [err:invalid argument: Controller router is nil]
	//test: RegisterController(ctrl) -> [err:invalid argument: Controller router primary resource is nil]
	//test: RegisterController(ctrl) -> [err:invalid argument: Controller router primary resource host is empty]
	//test: RegisterController(ctrl) -> [err:<nil>]

}

func _ExampleControlsLookup() {
	ctrl := NewController("test-route", nil, nil)
	p := NewControls()
	path := "http://localhost:8080/github/advanced-go/example-domain/activity"

	_, status := p.lookup("")
	fmt.Printf("test: Lookup(\"\") -> [status:%v]\n", status)

	_, status = p.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [status:%v]\n", path, status)

	err := p.registerWithAuthority(ctrl)
	fmt.Printf("test: Register(%v) -> [err:%v]\n", path, err)

	handler, status1 := p.lookup(path)
	fmt.Printf("test: Lookup(%v) -> [status:%v] [handler:%v]\n", path, status1, handler != nil)

	path = "http://localhost:8080/github/advanced-go/example-domain/activity"
	err = p.registerWithAuthority(ctrl)
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
