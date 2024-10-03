package httpxtest

import "fmt"

func ExampleFileInfo() {
	fi := FileInfo{
		Dir:  "/test",
		Req:  "test-req.txt",
		Resp: "",
	}

	fmt.Printf("test: FileInfo() -> [req:%v] [resp:%v]\n", fi.RequestPath(), fi.ResponsePath())

	//Output:
	//test: FileInfo() -> [req:/test/test-req.txt] [resp:/test/test-req-resp.txt]

}

func ExampleFileInfo_Error() {
	fi := FileInfo{
		Dir:  "/test",
		Req:  "test-req",
		Resp: "",
	}

	fmt.Printf("test: FileInfo() -> [req:%v] [resp:%v]\n", fi.RequestPath(), fi.ResponsePath())

	//Output:
	//test: FileInfo() -> [req:error: request file name does not have a . extension : test-req] [resp:/test/test-req]

}
