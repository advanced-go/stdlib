package coretest

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

// { "timestamp":"2024-09-07T15:07:01.171Z", "code":500, "status":"Internal Error", "request-id":"d59ab33d-6d2a-11ef-9cdc-00a55441ed8b",
// "errors" : [ "Get "http://localhost:8082/storage/address": dial tcp [::1]:8082: connectex: No connection could be made because the target machine actively refused it." ],
//  "trace" : [ "https://github.com/advanced-go/customer/tree/main/address1#get[...]","https://github.com/advanced-go/stdlib/tree/main/httpx#Do" ] }

func Example_FormatTraceTest() {

	s := formatTrace(core.TraceName, []string{"https://github.com/advanced-go/customer/tree/main/address1#get[...]",
		"https://github.com/advanced-go/stdlib/tree/main/httpx#Do",
	})

	fmt.Printf("test: formatTraceTest() -> %v\n", s)
	//Output:
	//fail
}

func Example_FormatErrorsTest() {
	err := errors.New("Get \"http://localhost:8082/storage/address\": dial tcp [::1]:8082: connectex: No connection could be made because the target machine actively refused it.")

	s := formatErrors(core.ErrorsName, []error{err})

	fmt.Printf("test: formatErrorsTest() -> %v\n", s)
	//Output:
	//fail
}

func Example_DefaultFormatterTest() {
	err := errors.New("Get \"http://localhost:8082/storage/address\": dial tcp [::1]:8082: connectex: No connection could be made because the target machine actively refused it.")
	errs := []error{err}

	trace := []string{"https://github.com/advanced-go/customer/tree/main/address1#get[...]",
		"https://github.com/advanced-go/stdlib/tree/main/httpx#Do",
	}
	s := defaultFormatter(time.Now().UTC(), "got", http.StatusTeapot, "I'm a teapot", "1234-5678", errs, trace)
	fmt.Printf("test: defaultFormatterTest() -> \n%v\n", s)

	//Output:
	//fail
}
