package host

import "fmt"

func ExamplePing() {
	uri := "test"
	status := Ping(uri)
	fmt.Printf("test: Ping(%v) -> [status:%v]\n", uri, status)

	//Output:
	//test: Ping(test) -> [status:Internal Error [error: exchange.Send() failed as the message To is empty or invalid : [test]]]
	
}
