package messaging

/*
// NewAgent - create an agent
func NewCtrlAgent(uri string, ctrlHandler Handler, public bool) (*Agent, error) {
	if len(uri) == 0 {
		return nil, errors.New("error: agent URI is empty")
	}
	if ctrlHandler == nil {
		return nil, errors.New("error: agent controller message handler is nil")
	}
	a := new(Agent)
	a.m = NewMailbox(uri, nil)
	a.m.public = public
	a.run = func(m *Mailbox) {
		DefaultRun(m, ctrlHandler)
	}

	return newDefaultAgent(uri, ctrlHandler, HostExchange, public)
}

// NewDefaultAgent - create an agent with only a control channel, registered with the HostDirectory,
// and using the default run function.
func NewDefaultAgent(uri string, ctrlHandler Handler, public bool) (*Agent, error) {
	return newDefaultAgent(uri, ctrlHandler, HostExchange, public)
}

func newDefaultAgent(uri string, ctrlHandler Handler, e *Exchange, public bool) (*Agent, error) {
	if len(uri) == 0 {
		return nil, errors.New("error: agent URI is empty")
	}
	if ctrlHandler == nil {
		return nil, errors.New("error: agent controller message handler is nil")
	}
	a := new(Agent)
	a.m = NewMailbox(uri, nil)
	a.m.public = public
	a.run = func(m *Mailbox) {
		DefaultRun(m, ctrlHandler)
	}
	return a, a.Register(e)
}

// DefaultRun - a simple run function that only handles control messages, and dispatches via a message handler
func DefaultRun(m *Mailbox, ctrlHandler Handler) {
	for {
		select {
		case msg, open := <-m.ctrl:
			if !open {
				return
			}
			switch msg.Event() {
			case ShutdownEvent:
				ctrlHandler(NewMessageWithStatus(ChannelControl, "", "", msg.Event(), core.StatusOK()))
				m.close()
				return
			default:
				ctrlHandler(msg)
			}
		default:
		}
	}
}


*/
