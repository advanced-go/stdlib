package messagingtest

import (
	"fmt"
	"github.com/advanced-go/stdlib/messaging"
)

func ExampleNewAgent() {
	a := NewAgent("urn:any", nil)
	if _, ok := any(a).(messaging.Agent); ok {
		fmt.Printf("test: opsAgent() -> ok\n")
	} else {
		fmt.Printf("test: opsAgent() -> fail\n")
	}

	//Output:
	//test: opsAgent() -> ok

}

func ExampleOld() {
	fmt.Printf("test")

	//Output:
	//test

}
