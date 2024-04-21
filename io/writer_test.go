package io

import (
	"bytes"
	"fmt"
	"reflect"
)

func ExampleIdentityWriter() {
	s := "identity encoding"
	buf := new(bytes.Buffer)

	ew, status := NewEncodingWriter(buf, nil)
	fmt.Printf("test: NewEncodingWriter(none) -> [ew:%v] [status:%v]\n", reflect.TypeOf(ew).String(), status)

	cnt, err := ew.Write([]byte(s))
	fmt.Printf("test: Write() -> [cnt:%v] [err:%v] [content:\"%v\"]\n", cnt, err, string(buf.Bytes()))

	//Output:
	//test: NewEncodingWriter(none) -> [ew:*io.identityWriter] [status:OK]
	//test: Write() -> [cnt:17] [err:<nil>] [content:"identity encoding"]

}

/*
func ExampleEncodingWriter_Gzip() {
	content, err0 := os.ReadFile(FileName(htmlResponse))
	if err0 != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err0)
		return
	}
	buff := new(bytes.Buffer)
	h := make(http.Header)
	h.Set(AcceptEncoding, GzipEncoding)

	// write, flush and close
	zw := gzip.NewWriter(buff)
	cnt, err := zw.Write(content)
	ferr := zw.Flush()
	cerr := zw.Close()
	fmt.Printf("test: gzip.Writer() -> [cnt:%v] [err:%v] [flush-err:%v] [close_err:%v]\n", cnt, err, ferr, cerr)

	// encoding results
	buff2 := bytes.Clone(buff.Bytes())
	fmt.Printf("test: DetectContent -> [input:%v] [output:%v] [in-len:%v]\n", http.DetectContentType(content), http.DetectContentType(buff2), len(content))

	// decode the content
	r := bytes.NewReader(buff2)
	zr, rerr := gzip.NewReader(r)
	buff1, err1 := io.ReadAll(zr)
	cerr = zr.Close()
	fmt.Printf("test: gzip.Reader() -> [new-err:%v] [read-err:%v] [close-err:%v]\n", rerr, err1, cerr)

	fmt.Printf("test: DetectContent -> [input:%v] [output:%v] [out-len:%v]\n", http.DetectContentType(buff2), http.DetectContentType(buff1), len(buff1))

	//Output:
	//test: gzip.Writer() -> [input:text/plain; charset=utf-8] [output:application/x-gzip]

}


*/
