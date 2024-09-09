package messagingtest

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	fmt2 "github.com/advanced-go/stdlib/fmt"
	"github.com/advanced-go/stdlib/messaging"
	"time"
)

type agent struct {
	agentId string
	ch      chan *messaging.Message
}

func NewAgent(uri string, ch chan *messaging.Message) messaging.OpsAgent {
	a := new(agent)
	a.agentId = uri
	a.ch = ch
	return a
}

func (t *agent) Uri() string                  { return t.agentId }
func (t *agent) Message(m *messaging.Message) { fmt.Printf("test: opsAgent.Message() -> %v\n", m) }
func (t *agent) Handle(status *core.Status) *core.Status {
	fmt.Printf("test: opsAgent.Handle() -> [status:%v]\n", status)
	status.Handled = true
	return status
}
func (t *agent) AddActivity(agentId string, content any) {
	fmt.Printf("test: opsAgent.AddActivity() -> %v : %v -> %v]\n", fmt2.FmtRFC3339Millis(time.Now().UTC()), agentId, content)
}

func (t *agent) Run() {}
func (t *agent) Shutdown() {
	if t.ch != nil {
		close(t.ch)
	}
}
