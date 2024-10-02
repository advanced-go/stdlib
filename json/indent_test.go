package json

import (
	"bytes"
	"fmt"
	io2 "github.com/advanced-go/stdlib/io"
	"io"
)

const (
	customerAddr = "file://[cwd]/jsontest/customer-address.txt"
)

func ExampleIndent() {
	buf, status := io2.ReadFile(customerAddr)
	fmt.Printf("test: io.ReadFile() -> [status:%v] %v\n", status, string(buf))

	if status.OK() {
		//fmt.Printf("test:")
		c := io.NopCloser(bytes.NewReader(buf))
		c2, status1 := Indent(c, nil, "", "  ")
		if status1.OK() {
			buf2, status := io2.ReadAll(c2, nil)
			fmt.Printf("test: Indent() -> [status:%v] %v\n", status, string(buf2))
		}
	}

	//Output:
	//test: io.ReadFile() -> [status:OK] [{"customer-id":"C001","created-ts":"0001-01-01T00:00:00Z","address-1":"1514 Cedar Ridge Road","address-2":"","city":"Vinton","state":"IA","postal-code":"52349","email":"before-email@hotmail.com"}]
	//test: Indent() -> [status:OK] [
	//  {
	//    "customer-id": "C001",
	//    "created-ts": "0001-01-01T00:00:00Z",
	//    "address-1": "1514 Cedar Ridge Road",
	//    "address-2": "",
	//    "city": "Vinton",
	//    "state": "IA",
	//    "postal-code": "52349",
	//    "email": "before-email@hotmail.com"
	//  }
	//]

}
