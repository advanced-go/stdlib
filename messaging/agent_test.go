package messaging

import (
	"fmt"
	"time"
)

type testAgent struct {
	running    bool
	agentId    string
	ctrl       chan *Message
	data       chan *Message
	handler    Handler
	shutdownFn func()
}

func NewTestAgent(uri string) Agent {
	return newTestAgent(uri, nil, nil)
}

func newTestAgent(uri string, ctrl, data chan *Message) *testAgent {
	t := new(testAgent)
	t.agentId = uri
	if ctrl == nil {
		t.ctrl = make(chan *Message, ChannelSize)
	} else {
		t.ctrl = ctrl
	}
	if data == nil {
		t.data = make(chan *Message, ChannelSize)
	} else {
		t.data = data
	}
	return t
}
func (t *testAgent) Uri() string    { return t.agentId }
func (t *testAgent) String() string { return t.Uri() }
func (t *testAgent) Message(msg *Message) {
	if msg == nil {
		return
	}
	switch msg.Channel() {
	case ChannelControl:
		if t.ctrl != nil {
			t.ctrl <- msg
		}
	case ChannelData:
		if t.data != nil {
			t.data <- msg
		}
	default:
	}
}
func (t *testAgent) Run() {
	if t.running {
		return
	}
	t.running = true
	go testAgentRun(t)
}

// Shutdown - shutdown the agent
func (t *testAgent) Shutdown() {
	if !t.running {
		return
	}
	t.running = false
	if t.shutdownFn != nil {
		t.shutdownFn()
	}
	t.Message(NewControlMessage(t.agentId, t.agentId, ShutdownEvent))
}

// Add - add a shutdown function
func (t *testAgent) Add(f func()) {
	t.shutdownFn = AddShutdown(t.shutdownFn, f)
}

func testAgentRun(t *testAgent) {
	for {
		select {
		case msg, open := <-t.ctrl:
			if !open {
				return
			}
			fmt.Printf("test: AgentRun() -> %v\n", msg)
			if msg.Event() == ShutdownEvent {
				return
			}
		default:
		}
		select {
		case msg, open := <-t.data:
			if !open {
				return
			}
			fmt.Printf("test: AgentRun() -> %v\n", msg)
		default:
		}
	}
}

/*
func printAgentRun(uri string, ctrl, data <-chan *Message, state any) {
	fmt.Printf("test: AgentRun() -> [uri:%v] [ctrl:%v] [data:%v] [state:%v]\n", uri, ctrl != nil, data != nil, state != nil)
}

func ExampleNewAgent_Error() {
	a, err := newAgent("", nil, nil, nil, nil)
	fmt.Printf("test: newAgent() -> [agent:%v] [%v]\n", a, err)

	a, err = newAgent("urn:agent7", nil, nil, nil, nil)
	fmt.Printf("test: newAgent() -> [agent:%v] [%v]\n", a, err)

	//Output:
	//test: newAgent() -> [agent:<nil>] [error: agent URI is empty]
	//test: newAgent() -> [agent:<nil>] [error: agent AgentFunc is nil]

}

func ExampleNewAgent() {
	uri := "urn:agent007"
	uri1 := "urn:agent009"

	a := newTestAgent(uri)
	a.Run()
	time.Sleep(time.Second)

	a, _ = NewAgentWithChannels(uri1, nil, nil, printAgentRun, "data")
	a.Run()
	time.Sleep(time.Second)

	//Output:
	//test: AgentRun() -> [uri:urn:agent007] [ctrl:true] [data:true] [state:false]
	//test: AgentRun() -> [uri:urn:agent009] [ctrl:true] [data:false] [state:true]

}

*/

func ExampleAgentRun() {
	uri := "urn:agent007"
	a := newTestAgent(uri, nil, nil)
	a.Run()
	a.Message(NewControlMessage(uri, "ExampleAgentRun()", StartupEvent))
	a.Message(NewDataMessage(uri, "ExampleAgentRun()", DataEvent))
	time.Sleep(time.Second)
	a.Shutdown()
	time.Sleep(time.Second)

	//Output:
	//test: AgentRun() -> [chan:CTRL] [from:ExampleAgentRun()] [to:urn:agent007] [event:startup]
	//test: AgentRun() -> [chan:DATA] [from:ExampleAgentRun()] [to:urn:agent007] [event:data]
	//test: AgentRun() -> [chan:CTRL] [from:urn:agent007] [to:urn:agent007] [event:shutdown]

}

func ExampleOnShutdown() {
	uri := "urn:agent007"

	a := newTestAgent(uri, nil, nil)
	a.running = true
	a.Shutdown()

	a1 := newTestAgent(uri, nil, nil)
	if sd, ok := any(a1).(OnShutdown); ok {
		sd.Add(func() { fmt.Printf("test: OnShutdown() -> func-1()\n") })
		sd.Add(func() { fmt.Printf("test: OnShutdown() -> func-2()\n") })
		sd.Add(func() { fmt.Printf("test: OnShutdown() -> func-3()\n") })
	}
	a1.running = true
	a1.Shutdown()

	//Output:
	//test: OnShutdown() -> func-1()
	//test: OnShutdown() -> func-2()
	//test: OnShutdown() -> func-3()

}
