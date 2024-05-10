package messaging

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// HostExchange - main exchange
var HostExchange = NewExchange()

// Exchange - exchange directory
type Exchange struct {
	m *sync.Map
}

// NewExchange - create a new exchange
func NewExchange() *Exchange {
	e := new(Exchange)
	e.m = new(sync.Map)
	return e
}

// Count - number of agents
func (d *Exchange) Count() int {
	count := 0
	d.m.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// List - a list of agent uri's
func (d *Exchange) List() []string {
	var uri []string
	d.m.Range(func(key, value any) bool {
		if str, ok := key.(string); ok {
			uri = append(uri, str)
		}
		return true
	})
	sort.Strings(uri)
	return uri
}

// Send - send a message
func (d *Exchange) Send(msg *Message) error {
	// TO DO : authenticate shutdown control message
	if msg.Event() == ShutdownEvent {
		return nil
	}
	if msg == nil {
		return errors.New(fmt.Sprintf("error: exchange.Send() failed as message is nil"))
	}
	a := d.Get(msg.To())
	if a == nil {
		return errors.New(fmt.Sprintf("error: exchange.Send() failed as the message To is empty or invalid [%v]", msg.To()))
	}
	a.Message(msg)
	return nil
}

// Register - register an agent
func (d *Exchange) Register(agent Agent) error {
	if agent == nil {
		return errors.New("error: exchange.Register() agent is nil")
	}
	//if len(m.uri) == 0 {
	//	return errors.New("error: exchange.Add() mailbox uri is empty")
	//}
	//if m.ctrl == nil {
	//	return errors.New("error: exchange.Add() mailbox command channel is nil")
	//}
	_, ok := d.m.Load(agent.Uri())
	if ok {
		return errors.New(fmt.Sprintf("error: exchange.Register() agent already exists: [%v]", agent.Uri()))
	}
	d.m.Store(agent.Uri(), agent)
	if sd, ok1 := agent.(OnShutdown); ok1 {
		sd.Add(func() {
			d.m.Delete(agent.Uri())
		})
	}
	return nil
}

func (d *Exchange) Get(uri string) Agent {
	if len(uri) == 0 {
		return nil
	}
	v, ok1 := d.m.Load(uri)
	if !ok1 {
		return nil
	}
	if a, ok2 := v.(Agent); ok2 {
		return a
	}
	return nil
}

// Shutdown - close an item's mailbox
/*
func (d *Exchange) shutdown(msg Message) error {
	// TO DO: add authentication
	return nil
}


*/

/*
func (d *exchange) shutdown(uri string) runtime.Status {
	//d.mu.RLock()
	//defer d.mu.RUnlock()
	//for _, e := range d.m {
	//	if e.ctrl != nil {
	//		e.ctrl <- Message{To: e.uri, Event: core.ShutdownEvent}
	//	}
	//}
	m, status := d.get(uri)
	if !status.OK() {
		return status
	}
	if m.data != nil {
		close(m.data)
	}
	if m.ctrl != nil {
		close(m.ctrl)
	}
	d.m.Delete(uri)
	return runtime.StatusOK()
}
*/
