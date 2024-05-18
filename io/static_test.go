package io

import (
	"fmt"
	"io"
	"net/http"
)

func ExampleBuffer_Read() {
	bytes := NewBuffer("test data")
	buf := make([]byte, 64)

	count, err := bytes.Read(buf)
	fmt.Printf("test: Buffer() -> [cnt:%v] [err:%v] [buf:%v]\n", count, err, string(buf))

	//Output:
	//fail

}

func ExampleBuffer_Reader() {
	bytes := NewBuffer("test data")

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes)}
	buf, err := io.ReadAll(resp.Body)
	fmt.Printf("test: Reader() -> [err:%v] [buf:%v]\n", err, string(buf))

	//Output:
	//fail

}
