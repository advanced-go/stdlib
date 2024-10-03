package httpxtest

import "fmt"

func ExampleFileList() {
	fi := FileList{
		Dir:  "/test",
		Req:  "test-req.txt",
		Resp: "",
	}

	fmt.Printf("test: FileList() -> [req:%v] [resp:%v]\n", fi.RequestPath(), fi.ResponsePath())

	//Output:
	//test: FileList() -> [req:/test/test-req.txt] [resp:/test/test-req-resp.txt]

}

func ExampleFileList_Error() {
	fi := FileList{
		Dir:  "/test",
		Req:  "test-req",
		Resp: "",
	}

	fmt.Printf("test: FileList() -> [req:%v] [resp:%v]\n", fi.RequestPath(), fi.ResponsePath())

	//Output:
	//test: FileList() -> [req:error: request file name does not have a . extension : test-req] [resp:/test/test-req]

}
