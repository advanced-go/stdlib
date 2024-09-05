package httpxtest

import (
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/io"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type entryTest struct {
	Traffic    string
	Duration   time.Duration
	Region     string
	Zone       string
	SubZone    string
	Service    string
	Url        string
	Protocol   string
	Host       string
	Path       string
	Method     string
	StatusCode int32
}

// parseRaw - parse a raw Uri without error
func parseRaw(rawUri string) *url.URL {
	u, _ := url.Parse(rawUri)
	return u
}

func Example_ReadRequest_GET() {
	s := "file://[cwd]/resource/get-request.txt"
	req, err := NewRequest(parseRaw(s))
	fmt.Printf("test: NewRequest(%v) -> [status:%v] [ctx:%v] [content-location:%v]\n", s, err, req.Context(), req.Header.Get("Content-Location"))

	req, err = NewRequest(s)
	fmt.Printf("test: NewRequest(%v) -> [status:%v] [ctx:%v] [content-location:%v]\n", s, err, req.Context(), req.Header.Get("Content-Location"))

	//Output:
	//test: NewRequest(file://[cwd]/resource/get-request.txt) -> [status:OK] [ctx:context.Background] [content-location:github/advanced-go/example-domain/activity/EntryV1]
	//test: NewRequest(file://[cwd]/resource/get-request.txt) -> [status:OK] [ctx:context.Background] [content-location:github/advanced-go/example-domain/activity/EntryV1]

}

func Example_ReadRequest_Baseline() {
	s := "file://[cwd]/resource/baseline-request.txt"
	req, err := NewRequest(parseRaw(s))

	if req != nil {
	}
	// print content
	//fmt.Printf("test: ReadRequest(%v) -> [err:%v] [%v]\n", s, err, req)
	fmt.Printf("test: NewRequest(%v) -> [status:%v]\n", s, err)

	//Output:
	//test: NewRequest(file://[cwd]/resource/baseline-request.txt) -> [status:OK]

}

func Example_ReadRequest_PUT() {
	s := "file://[cwd]/resource/put-req.txt"
	req, status := NewRequest(parseRaw(s))

	if !status.OK() {
		fmt.Printf("test: NewRequest(%v) -> [status:%v]\n", s, status)
	} else {
		buf, err1 := io.ReadAll(req.Body, nil)
		if err1 != nil {
		}
		var entry []entryTest
		json.Unmarshal(buf, &entry)
		fmt.Printf("test: NewRequest(%v) -> [cnt:%v] [fields:%v]\n", s, len(entry), entry)
	}

	//Output:
	//test: NewRequest(file://[cwd]/resource/put-req.txt) -> [cnt:2] [fields:[{ingress 800µs usa west  access-log https://access-log.com/example-domain/timeseries/entry http access-log.com /example-domain/timeseries/entry GET 200} {egress 100µs usa east  access-log https://access-log.com/example-domain/timeseries/entry http access-log.com /example-domain/timeseries/entry PUT 202}]]

}

func _ExampleNewRequest_Overrides() {
	s := "file://[cwd]/resource/get-request-overrides.txt"
	req, err := NewRequest(parseRaw(s))

	if req != nil {
	}
	// print content
	//fmt.Printf("test: ReadRequest(%v) -> [err:%v] [%v]\n", s, err, req)
	fmt.Printf("test: NewRequest(%v) -> [status:%v] [header:%v][type:%v]\n", s, err, req.Header[httpx.ContentLocationResolver], reflect.TypeOf(any(req.Header[httpx.ContentLocationResolver])))

	//Output:
	//test: NewRequest(file://[cwd]/resource/baseline-request.txt) -> [status:OK]

}

func _ExampleNewRequest_Overrides_Empty() {
	s := "file://[cwd]/resource/get-request-overrides.txt"
	req, err := NewRequest(parseRaw(s))

	if req != nil {
	}
	str, ok := req.Header["Content-Location-Empty"]
	fmt.Printf("test: NewRequest(%v) -> [err:%v] [ok:%v] [str:%v]\n", s, err, ok, len(str))

	// print content
	//fmt.Printf("test: ReadRequest(%v) -> [err:%v] [%v]\n", s, err, req)
	//fmt.Printf("test: NewRequest(%v) -> [status:%v] [header:%v][type:%v]\n", s, err, req.Header["Content-Location-Empty"], reflect.TypeOf(any(req.Header["Content-Location-Test"])))

	//Output:
	//test: NewRequest(file://[cwd]/resource/baseline-request.txt) -> [status:OK]

}

func ExampleCreateExchange() {
	h := make(http.Header)

	h.Add(httpx.ContentLocationExchange, "")

	ex := createExchange(h)
	fmt.Printf("test: createExchange() -> [ex:%v]\n", ex)

	h = make(http.Header)
	h.Add(httpx.ContentLocationExchange, "request->file:///f:/resource/request.json")
	h.Add(httpx.ContentLocationExchange, "response->file:///f:/resource/response.json")
	h.Add(httpx.ContentLocationExchange, "status->file:///f:/resource/status.json")

	ex = createExchange(h)
	fmt.Printf("test: createExchange() -> [ex:%v]\n", ex)

	//Output:
	//test: createExchange() -> [ex:<nil>]
	//test: createExchange() -> [ex:&{map[request:file:///f:/resource/request.json response:file:///f:/resource/response.json status:file:///f:/resource/status.json]}]

}
