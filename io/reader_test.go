package io

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func ExampleIdentityReader() {
	s := "identity encoding"

	br := bytes.NewReader([]byte(s))

	er, status := NewEncodingReader(br, nil)
	fmt.Printf("test: NewEncodingReader(none) -> [er:%v] [status:%v]\n", reflect.TypeOf(er).String(), status)

	buf, err := io.ReadAll(er)
	fmt.Printf("test: Read() -> [err:%v] [content:\"%v\"]\n", err, string(buf))

	//h := make(http.Header)
	//h.Add(ContentEncoding, GzipEncoding)
	//er, status = NewEncodingReader(br, h)
	//fmt.Printf("test: NewEncodingReader(gzip) -> [er:%v] [status:%v]\n", reflect.TypeOf(er).String(), status)

	//Output:
	//test: NewEncodingReader(none) -> [er:*io.identityReader] [status:OK]
	//test: Read() -> [err:<nil>] [content:"identity encoding"]

}

/*
func ExampleEncodingReader_Error() {
	s := address3Url
	buf0, err := os.ReadFile(FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}
	r := strings.NewReader(string(buf0))

	h := make(http.Header)
	h.Set(ContentEncoding, BrotliEncoding)
	reader, status := EncodingReader(r, h)
	fmt.Printf("test: EncodingReader() -> [reader:%v] [status:%v]\n", reader, status)

	h.Set(ContentEncoding, DeflateEncoding)
	reader, status = EncodingReader(r, h)
	fmt.Printf("test: EncodingReader() -> [reader:%v] [status:%v]\n", reader, status)

	h.Set(ContentEncoding, CompressEncoding)
	reader, status = EncodingReader(r, h)
	fmt.Printf("test: EncodingReader() -> [reader:%v] [status:%v]\n", reader, status)

	//Output:
	//test: EncodingReader() -> [reader:<nil>] [status:Content Decoding Failure [error: content encoding not supported [br]]]
	//test: EncodingReader() -> [reader:<nil>] [status:Content Decoding Failure [error: content encoding not supported [deflate]]]
	//test: EncodingReader() -> [reader:<nil>] [status:Content Decoding Failure [error: content encoding not supported [compress]]]

}


*/

/*
func ExampleEncodingReader_Gzip() {
	s := searchResultsGzip
	buf0, err0 := os.ReadFile(FileName(s))
	if err0 != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err0)
		return
	}
	r := bytes.NewReader(buf0)

	h := make(http.Header)
	h.Set(ContentEncoding, GzipEncoding)
	zr, _ := EncodingReader(r, h)
	buf, err := io.ReadAll(zr)
	fmt.Printf("test: io.ReadAll() -> [input:%v] [output:%v] [err:%v]\n", http.DetectContentType(buf0), http.DetectContentType(buf), err)

	//Output:
	//test: io.ReadAll() -> [input:application/x-gzip] [output:text/html; charset=utf-8] [err:<nil>]

}

*/
