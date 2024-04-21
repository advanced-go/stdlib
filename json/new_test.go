package json

import (
	"fmt"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
	"net/url"
	"os"
	"strings"
)

const (
	address1Url = "file://[cwd]/jsontest/address1.json"
	address2Url = "file://[cwd]/jsontest/address2.json"
	address3Url = "file://[cwd]/jsontest/address3.json"
	status504   = "file://[cwd]/jsontest/status-504.json"
)

type newAddress struct {
	City    string
	State   string
	ZipCode string
}

// parseRaw - parse a raw Uri without error
func parseRaw(rawUri string) *url.URL {
	u, _ := url.Parse(rawUri)
	return u
}

func ExampleNew_String_Error() {
	_, status := New[newAddress]("", nil)
	fmt.Printf("test: New(\"\") -> [status:%v]\n", status)

	s := "https://www.google.com/search"
	_, status = New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	s = "file://[cwd]/jsontest/address.txt"
	_, status = New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	//Output:
	//test: New("") -> [status:Invalid Argument [error: URI is empty]]
	//test: New(https://www.google.com/search) -> [status:Invalid Argument [error: URI is not of scheme file: https://www.google.com/search]]
	//test: New(file://[cwd]/jsontest/address.txt) -> [status:I/O Failure [open C:\Users\markb\GitHub\stdlib\json\jsontest\address.txt: The system cannot find the file specified.]]

}

/*
func ExampleNew_String_Status() {
	s := StatusOKUri
	addr, status := New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status)

	s = StatusNotFoundUri
	bytes, status0 := New[[]byte](s, nil)
	fmt.Printf("test: New(%v) -> [bytes:%v] [status:%v]\n", s, bytes, status0)

	s = StatusTimeoutUri
	addr, status = New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status)

	// status from uri, generic type is ignored
	s = status504
	bytes1, status1 := New[[]byte](s, nil)
	fmt.Printf("test: New(%v) -> [bytes:%v] [status:%v]\n", s, bytes1, status1)

	//Output:
	//test: New(urn:status:ok) -> [addr:{  }] [status:OK]
	//test: New(urn:status:notfound) -> [bytes:[]] [status:Not Found]
	//test: New(urn:status:timeout) -> [addr:{  }] [status:Timeout]
	//test: New(file://[cwd]/jsontest/status-504.json) -> [bytes:[]] [status:Timeout [error 1]]

}


*/

func ExampleNew_String_URI() {
	// bytes
	s := address1Url
	//bytes, status := New[[]byte](s)
	//fmt.Printf("test: New(%v) -> [bytes:%v] [status:%v]\n", s, len(bytes), status)

	// type
	s = address1Url
	addr, status1 := New[newAddress](s, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/jsontest/address1.json) -> [addr:{frisco texas 75034}] [status:OK]

}

func ExampleNew_URL_Error() {
	_, status := New[newAddress](nil, nil)
	fmt.Printf("test: New(\"\") -> [status:%v]\n", status)

	s := "https://www.google.com/search"
	_, status = New[newAddress](parseRaw(s), nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	s = "file://[cwd]/jsontest/address.txt"
	_, status = New[newAddress](parseRaw(s), nil)
	fmt.Printf("test: New(%v) -> [status:%v]\n", s, status)

	//Output:
	//test: New("") -> [status:Invalid Argument [error: invalid type [<nil>]]]
	//test: New(https://www.google.com/search) -> [status:Invalid Argument [error: URI is not of scheme file: https://www.google.com/search]]
	//test: New(file://[cwd]/jsontest/address.txt) -> [status:I/O Failure [open C:\Users\markb\GitHub\stdlib\json\jsontest\address.txt: The system cannot find the file specified.]]

}

/*
func ExampleNew_URL_Status() {
	s := status504
	u, _ := url.Parse(s)

	addr, status0 := New[newAddress](u, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status0)

	//Output:
	//test: New(file://[cwd]/jsontest/status-504.json) -> [addr:{  }] [status:Timeout [error 1]]

}


*/

func ExampleNew_URL() {
	s := address1Url
	u, _ := url.Parse(s)
	addr, status1 := New[newAddress](u, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/jsontest/address1.json) -> [addr:{frisco texas 75034}] [status:OK]

}

func ExampleNew_Bytes() {
	s := address2Url
	buf, err := os.ReadFile(io2.FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}

	addr, status := New[newAddress](buf, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status)

	//Output:
	//test: New(file://[cwd]/jsontest/address2.json) -> [addr:{vinton iowa 52349}] [status:<nil>]

}

func ExampleNew_Reader() {
	s := address2Url
	buf, err := os.ReadFile(io2.FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}

	r := strings.NewReader(string(buf))
	addr, status1 := New[newAddress](r, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/jsontest/address2.json) -> [addr:{vinton iowa 52349}] [status:OK]

}

func ExampleNew_ReadCloser() {
	s := address3Url
	buf, err := os.ReadFile(io2.FileName(s))
	if err != nil {
		fmt.Printf("test: os.ReadFile() -> [err:%v]\n", err)
	}

	body := io.NopCloser(strings.NewReader(string(buf)))
	addr, status1 := New[newAddress](body, nil)
	fmt.Printf("test: New(%v) -> [addr:%v] [status:%v]\n", s, addr, status1)

	//Output:
	//test: New(file://[cwd]/jsontest/address3.json) -> [addr:{forest city iowa 50436}] [status:OK]

}
