package messaging

import (
	"errors"
)

const (
	ChannelSize = 16
)

// Agent - Intelligent Agent
// TODO : Track agent assignment as part of the URI or separate identifier??
//
//	Track agent NID or class/type?
type Agent interface {
	Uri() string
	Message(m *Message)
	Run()
	Shutdown()
}

type RunFunc func(agent any)

type agent struct {
	uri  string
	ctrl chan *Message
	data chan *Message
	run  RunFunc
}

func NewAgent(uri string, run RunFunc) (Agent, error) {
	return newAgent(uri, make(chan *Message, ChannelSize), make(chan *Message, ChannelSize), run)
}

func NewAgentWithChannel(uri string, ctrl chan *Message, data chan *Message, run RunFunc) (Agent, error) {
	if ctrl == nil {
		ctrl = make(chan *Message, ChannelSize)
	}
	if data == nil {
		data = make(chan *Message, ChannelSize)
	}
	return newAgent(uri, ctrl, data, run)
}

func newAgent(uri string, ctrl chan *Message, data chan *Message, run RunFunc) (Agent, error) {
	if len(uri) == 0 {
		return nil, errors.New("error: agent URI is empty")
	}
	if ctrl == nil {
		return nil, errors.New("error: agent control channel is nil")
	}
	if run == nil {
		return nil, errors.New("error: agent RunFunc is nil")
	}
	a := new(agent)
	a.uri = uri
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
	go a.run(a)
}

// Shutdown - shutdown the agent
func (a *agent) Shutdown() {
	a.Message(NewControlMessage("", "", ShutdownEvent))
}
