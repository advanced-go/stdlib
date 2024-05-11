package messaging

import (
	"fmt"
)

func emptyRun(uri string, ctrl, data <-chan *Message, state any) {
}

func ExampleRegister() {
	testDir := NewExchange()

	fmt.Printf("test: Count() -> : %v\n", testDir.Count())

	uri1 := "urn:test:one"
	a := testDir.Get(uri1)
	fmt.Printf("test: Get(%v) -> : [agent:%v]\n", uri1, a)

	a1, _ := NewAgent(uri1, emptyRun, nil)
	err := testDir.Register(a1)
	fmt.Printf("test: Register(%v) -> : [err:%v]\n", uri1, err)

	fmt.Printf("test: Count() -> : %v\n", testDir.Count())
	a1 = testDir.Get(uri1)
	fmt.Printf("test: Get(%v) -> : [agent:%v]\n", uri1, a1)

	uri2 := "urn:test:two"
	a2, _ := NewAgent(uri2, emptyRun, nil)
	err = testDir.Register(a2)
	fmt.Printf("test: Register(%v) -> : [err:%v]\n", uri2, err)
	fmt.Printf("test: Count() -> : %v\n", testDir.Count())
	a2 = testDir.Get(uri2)
	fmt.Printf("test: Get(%v) -> : [agent:%v]\n", uri2, a2)

	fmt.Printf("test: List() -> : %v\n", testDir.List())

	//Output:
	//test: Count() -> : 0
	//test: Get(urn:test:one) -> : [agent:<nil>]
	//test: Register(urn:test:one) -> : [err:<nil>]
	//test: Count() -> : 1
	//test: Get(urn:test:one) -> : [agent:urn:test:one]
	//test: Register(urn:test:two) -> : [err:<nil>]
	//test: Count() -> : 2
	//test: Get(urn:test:two) -> : [agent:urn:test:two]
	//test: List() -> : [urn:test:one urn:test:two]

}

/*
func Example_SendError() {
	uri := "urn:test"
	testDir := NewExchange()

	fmt.Printf("test: Send(%v) -> : %v\n", uri, testDir.Send(NewControlMessage(uri, "", "")))

	m := NewMailboxWithCtrl(uri, false, nil, nil)
	status := testDir.Add(m)
	fmt.Printf("test: Add(%v) -> : [status:%v]\n", uri, status)

	//Output:
	//test: Send(urn:test) -> : error: exchange.SendCtrl() failed as the message To is empty or invalid [urn:test]
	//test: Add(urn:test) -> : [status:error: exchange.Add() mailbox command channel is nil]

}

func Example_Send() {
	uri1 := "urn:test-1"
	uri2 := "urn:test-2"
	uri3 := "urn:test-3"
	c := make(chan *Message, 16)
	testDir := NewExchange()

	testDir.Add(NewMailboxWithCtrl(uri1, false, c, nil))
	testDir.Add(NewMailboxWithCtrl(uri2, false, c, nil))
	testDir.Add(NewMailboxWithCtrl(uri3, false, c, nil))

	testDir.Send(NewControlMessage(uri1, PkgPath, StartupEvent))
	testDir.Send(NewControlMessage(uri2, PkgPath, StartupEvent))
	testDir.Send(NewControlMessage(uri3, PkgPath, StartupEvent))

	time.Sleep(time.Second * 1)
	resp1 := <-c
	resp2 := <-c
	resp3 := <-c
	fmt.Printf("test: <- c -> : [%v] [%v] [%v]\n", resp1.To(), resp2.To(), resp3.To())
	close(c)

	//Output:
	//test: <- c -> : [urn:test-1] [urn:test-2] [urn:test-3]

}

func Example_ListCount() {
	testDir := NewExchange()

	testDir.Add(newDefaultMailbox("test:uri1"))
	testDir.Add(newDefaultMailbox("test:uri2"))

	fmt.Printf("test: Count() -> : %v\n", testDir.Count())

	fmt.Printf("test: List() -> : %v\n", testDir.List())

	//Output:
	//test: Count() -> : 2
	//test: List() -> : [test:uri1 test:uri2]

}

func Example_Remove() {
	uri := "urn:test/one"

	m := newDefaultMailbox(uri)
	testDir := NewExchange()

	status := testDir.Add(m)
	fmt.Printf("test: Add(%v) -> : [%v]\n", uri, status)

	status = testDir.Send(NewControlMessage(uri, "", PingEvent))
	fmt.Printf("test: Send(%v) -> : [%v]\n", uri, status)

	m.close()

	status = testDir.Send(NewControlMessage(uri, "", PingEvent))
	fmt.Printf("test: Send(%v) -> : [%v]\n", uri, status)

	//Output:
	//test: Add(urn:test/one) -> : [<nil>]
	//test: Send(urn:test/one) -> : [<nil>]
	//test: Send(urn:test/one) -> : [error: exchange.SendCtrl() failed as the message To is empty or invalid [urn:test/one]]

}


*/
