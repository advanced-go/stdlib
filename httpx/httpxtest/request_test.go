package httpxtest

import (
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/io"
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

func Example_ReadRequest_GET() {
	s := "file://[cwd]/resource/get-request.txt"
	req, err := NewRequest(ParseRaw(s))
	fmt.Printf("test: NewRequest(%v) -> [status:%v] [ctx:%v] [content-location:%v]\n", s, err, req.Context(), req.Header.Get("Content-Location"))

	req, err = NewRequest(s)
	fmt.Printf("test: NewRequest(%v) -> [status:%v] [ctx:%v] [content-location:%v]\n", s, err, req.Context(), req.Header.Get("Content-Location"))

	//Output:
	//test: NewRequest(file://[cwd]/resource/get-request.txt) -> [status:OK] [ctx:context.Background.WithValue(type core.urlExchangeOverrideContextKey, val <not Stringer>)] [content-location:github/advanced-go/example-domain/activity/EntryV1]
	//test: NewRequest(file://[cwd]/resource/get-request.txt) -> [status:OK] [ctx:context.Background.WithValue(type core.urlExchangeOverrideContextKey, val <not Stringer>)] [content-location:github/advanced-go/example-domain/activity/EntryV1]

}

func Example_ReadRequest_Baseline() {
	s := "file://[cwd]/resource/baseline-request.txt"
	req, err := NewRequest(ParseRaw(s))

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
	req, status := NewRequest(ParseRaw(s))

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
