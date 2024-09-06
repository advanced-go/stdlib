package messaging

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
)

type controlAgent struct {
	running    bool
	agentId    string
	ch         chan *Message
	handler    Handler
	shutdownFn func()
}

// NewControlAgent - create an agent that only listens on a control channel, and has a default AgentRun func
func NewControlAgent(uri string, handler Handler) (Agent, error) {
	if handler == nil {
		return nil, errors.New("error: control agent message handler is nil")
	}
	return newControlAgent(uri, handler), nil
	//return NewAgentWithChannels(uri, nil, nil, controlAgentRun, ctrlHandler)
}

func newControlAgent(uri string, handler Handler) *controlAgent {
	c := new(controlAgent)
	c.ch = make(chan *Message, ChannelSize)
	c.handler = handler
	return c
}

// Uri - identity
func (a *controlAgent) Uri() string { return a.agentId }

// String - identity
func (a *controlAgent) String() string { return a.Uri() }

// Message - message an agent
func (a *controlAgent) Message(msg *Message) {
	if msg == nil {
		return
	}
	switch msg.Channel() {
	case ChannelControl:
		if a.ch != nil {
			a.ch <- msg
		}
	default:
	}
}

// Run - run the agent
func (a *controlAgent) Run() {
	if a.running {
		return
	}
	a.running = true
	go controlAgentRun(a)
}

// Shutdown - shutdown the agent
func (a *controlAgent) Shutdown() {
	if !a.running {
		return
	}
	a.running = false
	if a.shutdownFn != nil {
		a.shutdownFn()
	}
	a.Message(NewControlMessage(a.agentId, a.agentId, ShutdownEvent))
}

// Add - add a shutdown function
func (a *controlAgent) Add(f func()) {
	a.shutdownFn = AddShutdown(a.shutdownFn, f)
}

func (a *controlAgent) shutdown() {
	close(a.ch)
}

// controlAgentRun - a simple run function that only handles control messages, and dispatches via a message handler
func controlAgentRun(c *controlAgent) {
	if c == nil || c.handler == nil {
		return
	}
	// ctrlHandler Handler
	//if h, ok := state.(Handler); ok {
	//	ctrlHandler = h
	//} else {
	//	return
	//}
	for {
		select {
		case msg, open := <-c.ch:
			if !open {
				return
			}
			switch msg.Event() {
			case ShutdownEvent:
				c.handler(NewMessageWithStatus(ChannelControl, "", "", msg.Event(), core.StatusOK()))
				c.shutdown()
				return
			default:
				c.handler(msg)
			}
		default:
		}
	}
}
