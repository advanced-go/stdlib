package core

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var testTS time.Time

func init() {
	testTS = time.Date(2024, 3, 1, 18, 23, 50, 205*1e6, time.UTC)

}

func Example_FormatUri() {
	s := "github/advanced-go/stdlib/core:testFunc"

	fmt.Printf("test: formatUri(%v) -> %v\n", s, formatUri(s))

	s = "gitlab/advanced-go/stdlib/core:testFunc"
	fmt.Printf("test: formatUri(%v) -> %v\n", s, formatUri(s))

	//Output:
	//test: formatUri(github/advanced-go/stdlib/core:testFunc) -> https://github.com/advanced-go/stdlib/tree/main/core#testFunc
	//test: formatUri(gitlab/advanced-go/stdlib/core:testFunc) -> gitlab/advanced-go/stdlib/core:testFunc

}

func Example_FormatUri_Test() {
	s := "http://localhost:8080/github.com/advanced-go/stdlib/core/testFunc"
	req, err := http.NewRequest("", s, nil)
	fmt.Printf("test: http.URL -> [req:%v] [url:%v] [err:%v]\n", req != nil, req.URL, err)

	s = "http://localhost:8080/github.com/advanced-go/stdlib/core:testFunc"
	req, err = http.NewRequest("", s, nil)
	fmt.Printf("test: http.URL -> [req:%v] [url:%v] [err:%v]\n", req != nil, req.URL, err)

	s = "http://localhost:8080/github.com:advanced-go/stdlib/core.testFunc"
	req, err = http.NewRequest("", s, nil)
	fmt.Printf("test: http.URL -> [req:%v] [url:%v] [err:%v]\n", req != nil, req.URL, err)

	//Output:
	//test: http.URL -> [req:true] [url:http://localhost:8080/github.com/advanced-go/stdlib/core/testFunc] [err:<nil>]
	//test: http.URL -> [req:true] [url:http://localhost:8080/github.com/advanced-go/stdlib/core:testFunc] [err:<nil>]
	//test: http.URL -> [req:true] [url:http://localhost:8080/github.com:advanced-go/stdlib/core.testFunc] [err:<nil>]

}

func Example_DefaultFormat() {
	s := NewStatusError(http.StatusNotFound, errors.New("test error message 1"))

	str := formatter(testTS, s.Code, HttpStatus(s.Code), "1234-5678", []error{s.Error()}, s.Trace())
	fmt.Printf("test: formatter() -> %v", str)

	//Output:
	//test: formatter() -> { "timestamp":"2024-03-01T18:23:50.205Z", "code":404, "status":"Not Found", "request-id":"1234-5678", "errors" : [ "test error message 1" ], "trace" : [ "https://github.com/advaced-go/stdlib/tree/main/core#Example_DefaultFormat" ] }

}

/*
func ExampleOutputHandler_Handle() {
	ctx := NewRequestIdContext(nil, "123-request-id")
	err := errors.New("test error")
	var h Output

	s := h.Handle(NewStatus(http.StatusInternalServerError), RequestId(ctx))
	fmt.Printf("test: Handle(status,id) -> [%v] [errors:%v]\n", s, s.Error() != nil)

	s = h.Handle(NewStatusError(http.StatusInternalServerError, err), RequestId(ctx))
	fmt.Printf("test: Handle(status,id) -> [%v] [handled:%v]\n", s, s.handled)

	s = NewStatusError(http.StatusInternalServerError, nil)
	fmt.Printf("test: Handle(nil,id) -> [%v] [handled:%v]\n", h.Handle(nil, RequestId(ctx)), s.handled)

	//Output:
	//test: Handle(status,id) -> [Internal Error] [errors:false]
	//{ "timestamp":"2024-03-01T18:23:50.205Z", "code":500, "status":"Internal Error", "request-id":"123-request-id", "errors" : [ "test error" ], "trace" : [ "https://github.com/advanced-go/core/tree/main/runtime#ExampleOutputHandler_Handle" ] }
	//test: Handle(status,id) -> [Internal Error [test error]] [handled:true]
	//test: Handle(nil,id) -> [OK] [handled:false]

}

func ExampleLogHandler_Handle() {
	ctx := NewRequestIdContext(nil, "")
	err := errors.New("test error")
	var h Log

	//s := h.Handle(GetOrCreateRequestId(ctx), location, nil)
	s := h.Handle(NewStatus(http.StatusOK), RequestId(ctx))
	fmt.Printf("test: Handle(ctx,location,nil) -> [%v] [errors:%v]\n", s, s.Error() != nil)

	//	s = h.Handle(GetOrCreateRequestId(ctx), location, err)
	s = h.Handle(NewStatusError(http.StatusInternalServerError, err), RequestId(ctx))
	fmt.Printf("test: Handle(ctx,location,err) -> [%v] [errors:%v]\n", s, s.Error() != nil)

	s = NewStatusError(http.StatusInternalServerError, nil)
	fmt.Printf("test: Handle(nil,s) -> [%v] [errors:%v]\n", h.Handle(nil, RequestId(ctx)), s.Error() != nil)

	s = NewStatusError(http.StatusInternalServerError, err)
	errors := s.Error() != nil
	s1 := h.Handle(s, RequestId(ctx))
	fmt.Printf("test: Handle(nil,s) -> [prev:%v] [prev-errors:%v] [curr:%v] [curr-errors:%v]\n", s, errors, s1, s1.Error() != nil)

	//Output:
	//test: Handle(ctx,location,nil) -> [OK] [errors:false]
	//test: Handle(ctx,location,err) -> [Internal Error [test error]] [errors:true]
	//test: Handle(nil,s) -> [OK] [errors:false]
	//test: Handle(nil,s) -> [prev:Internal Error [test error]] [prev-errors:true] [curr:Internal Error [test error]] [curr-errors:true]

}

func Example_InvalidTypeError() {
	fmt.Printf("test: NewInvalidBodyTypeError(nil) -> %v\n", NewInvalidBodyTypeError(nil))
	fmt.Printf("test: NewInvalidBodyTypeError(string) -> %v\n", NewInvalidBodyTypeError("test data"))
	fmt.Printf("test: NewInvalidBodyTypeError(int) -> %v\n", NewInvalidBodyTypeError(500))

	req, _ := http.NewRequest("patch", "https://www.google.com/search", nil)
	fmt.Printf("test: NewInvalidBodyTypeError(*http.Request) -> %v\n", NewInvalidBodyTypeError(req))

	//Output:
	//test: NewInvalidBodyTypeError(nil) -> invalid body type: <nil>
	//test: NewInvalidBodyTypeError(string) -> invalid body type: string
	//test: NewInvalidBodyTypeError(int) -> invalid body type: int
	//test: NewInvalidBodyTypeError(*http.Request) -> invalid body type: *http.Request

}

func ExampleFmtAttrs() {
	var attrs []any

	s := formatAttrs(attrs)
	fmt.Printf("test: FmtAttrs() -> [empty:%v]\n", len(s) == 0)

	attrs = append(attrs, StatusName)
	attrs = append(attrs, "Bad Request")

	attrs = append(attrs, CodeName)
	attrs = append(attrs, http.StatusBadRequest)

	attrs = append(attrs, "isError")
	attrs = append(attrs, false)

	attrs = append(attrs, "empty-string")
	attrs = append(attrs, "")

	attrs = append(attrs, TimestampName)
	attrs = append(attrs, testTS)

	s = formatAttrs(attrs)
	fmt.Printf("test: FmtAttrs-Even() -> %v\n", s)

	attrs = append(attrs, "name-only")
	s = formatAttrs(attrs)
	fmt.Printf("test: FmtAttrs-Odd() -> %v\n", s)

	//Output:
	//test: FmtAttrs() -> [empty:true]
	//test: FmtAttrs-Even() -> "status":"Bad Request", "code":400, "isError":false, "empty-string":null, "timestamp":"2024-03-01T18:23:50.205Z"
	//test: FmtAttrs-Odd() -> "status":"Bad Request", "code":400, "isError":false, "empty-string":null, "timestamp":"2024-03-01T18:23:50.205Z", "name-only":null

}


*/

/*

// ErrorHandleFn - function type for error handling
//type ErrorHandleFn func(requestId, location string, errs ...error) *Status
// NewErrorHandler - templated function providing an error handle function via a closure

func NewErrorHandler[E ErrorHandler]() ErrorHandleFn {
	var e E
	return func(requestId string, location string, errs ...error) *Status {
		return e.Handle(NewStatusError(http.StatusInternalServerError, location, errs...), requestId, "")
	}
}

func ExampleErrorHandleFn() {
	loc := PkgUri + "/ErrorHandleFn"

	fn := NewErrorHandler[LogError]()
	fn("", loc, errors.New("log - error message"))
	fmt.Printf("test: Handle[LogErrorHandler]()\n")

	Output:
	test: Handle[LogErrorHandler]()

}


*/
