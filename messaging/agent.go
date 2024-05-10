package messaging

import (
	"errors"
)

const (
	ChannelSize = 16
)

// OnShutdown - add functions to be run on shutdown
type OnShutdown interface {
	Add(func())
}

// Agent - intelligent agent
// TODO : Track agent assignment as part of the URI or separate identifier??
//
//	Track agent NID or class/type?
type Agent interface {
	Uri() string
	Message(m *Message)
	Run()
	Shutdown()
}

// AgentFunc - agent function
type AgentFunc func(uri string, ctrl, data <-chan *Message, state any)

type agent struct {
	uri      string
	state    any
	ctrl     chan *Message
	data     chan *Message
	run      AgentFunc
	shutdown func()
}

func NewAgent(uri string, run AgentFunc, state any) (Agent, error) {
	return newAgent(uri, make(chan *Message, ChannelSize), make(chan *Message, ChannelSize), run, state)
}

func NewAgentWithChannel(uri string, ctrl, data chan *Message, run AgentFunc, state any) (Agent, error) {
	return newAgent(uri, ctrl, data, run, state)
}

func newAgent(uri string, ctrl, data chan *Message, run AgentFunc, state any) (Agent, error) {
	if len(uri) == 0 {
		return nil, errors.New("error: agent URI is empty")
	}
	if ctrl == nil {
		return nil, errors.New("error: agent control channel is nil")
	}
	if run == nil {
		return nil, errors.New("error: agent AgentFunc is nil")
	}
	a := new(agent)
	a.uri = uri
	a.state = state
	a.ctrl = ctrl
	a.data = data
	a.run = run
	return a, nil
}

// Uri - identity
func (a *agent) Uri() string {
	return a.uri
}

// String - identity
func (a *agent) String() string {
	return a.uri
}

// Message - message an agent
func (a *agent) Message(msg *Message) {
	if msg == nil {
		return
	}
	/*
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered in agent.Shutdown() : %v\n", r)
			}
		}()

	*/
	if msg.Channel() == ChannelControl && a.ctrl != nil {
		a.ctrl <- msg
	} else {
		if msg.Channel() == ChannelData && a.data != nil {
			a.data <- msg
		}
	}
}

// Run - run the agent
func (a *agent) Run() {
	go a.run(a.uri, a.ctrl, a.data, a.state)
}

// Shutdown - shutdown the agent
func (a *agent) Shutdown() {
	if a.shutdown != nil {
		a.shutdown()
	}
	a.Message(NewControlMessage("", "", ShutdownEvent))
}

// Add - add a shutdown function
func (a *agent) Add(f func()) {
	if f == nil {
		return
	}
	if a.shutdown == nil {
		a.shutdown = f
	} else {
		a.shutdown = func() {
			a.shutdown()
			f()
		}
	}
}
