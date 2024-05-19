package controller

import (
	"fmt"
	"time"
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

func ExampleControlsLookup() {
	p := NewControls()
	//path := "http://localhost:8080/github/advanced-go/search:google?q=golang"
	auth := "github/advanced-go/search"
	ctrl := NewController("test-route", NewPrimaryResource("", auth, time.Second*2, "", httpCall), nil)

	_, status := p.lookup("")
	fmt.Printf("test: Lookup(\"\") -> [status:%v]\n", status)

	_, status = p.lookup(auth)
	fmt.Printf("test: Lookup(%v) -> [status:%v]\n", auth, status)

	err := p.registerWithAuthority(ctrl)
	fmt.Printf("test: Register() -> [err:%v]\n", err)

	handler, status1 := p.lookup(auth)
	fmt.Printf("test: Lookup(%v) -> [status:%v] [handler:%v]\n", auth, status1, handler != nil)

	host := "www.google.com"
	ctrl = NewController("test-route", NewPrimaryResource(host, "", time.Second*2, "", httpCall), nil)

	err = p.register(ctrl)
	fmt.Printf("test: Register() -> [err:%v]\n", err)
	handler, status1 = p.lookup(host)
	fmt.Printf("test: Lookup(%v) -> [status:%v] [handler:%v]\n", host, status1, handler != nil)

	//Output:
	//test: Lookup("") -> [status:Invalid Argument [invalid argument: authority is empty]]
	//test: Lookup(github/advanced-go/search) -> [status:Invalid Argument [invalid argument: Controller does not exist: [github/advanced-go/search]]]
	//test: Register() -> [err:<nil>]
	//test: Lookup(github/advanced-go/search) -> [status:OK] [handler:true]
	//test: Register() -> [err:<nil>]
	//test: Lookup(www.google.com) -> [status:OK] [handler:true]

}
