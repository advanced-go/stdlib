package host

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
	"time"
)

func emptyRun(uri string, ctrl, data <-chan *messaging.Message, state any) {
}
func testRegister(ex *messaging.Exchange, uri string, cmd, data chan *messaging.Message) error {
	a, _ := messaging.NewAgentWithChannels(uri, cmd, data, nil, emptyRun, nil)
	ex.Register(a) //.NewMailboxWithCtrl(uri, false, cmd, data))
	return nil
}

var start time.Time

func ExampleCreateToSend() {
	uriNone := "startup/none"
	uriOne := "startup/one"
	ex := messaging.NewExchange()

	a, _ := messaging.NewAgent(uriNone, emptyRun, nil)
	err := ex.Register(a)
	if err != nil {
		fmt.Printf("test: NewAgent(%v) -> [err:%v]\n", uriNone, err)
	}

	a, _ = messaging.NewAgent(uriOne, emptyRun, nil)
	err = ex.Register(a)
	if err != nil {
		fmt.Printf("test: NewAgent(%v) -> [err:%v]\n", uriOne, err)
	}
	m := createToSend(ex, nil, nil)
	msg := m[uriNone]
	fmt.Printf("test: createToSend(nil,nil) -> [to:%v] [from:%v]\n", msg.To(), msg.From())

	//Output:
	//test: createToSend(nil,nil) -> [to:startup/none] [from:github/advanced-go/stdlib/host:Startup]

}

func ExampleStartup_Success() {
	uri1 := "github/startup/good"
	uri2 := "github/startup/bad"
	uri3 := "github/startup/depends"

	ex := messaging.NewExchange()
	start = time.Now()

	c := make(chan *messaging.Message, 16)
	testRegister(ex, uri1, c, nil)
	go startupGood(c)

	c = make(chan *messaging.Message, 16)
	testRegister(ex, uri2, c, nil)
	go startupBad(c)

	c = make(chan *messaging.Message, 16)
	testRegister(ex, uri3, c, nil)
	go startupDepends(c, nil)

	status := startup(ex, time.Second*2, nil)

	fmt.Printf("test: Startup() -> [%v]\n", status)

	//Output:
	//startup successful: [github/startup/bad] : 0s
	//startup successful: [github/startup/depends] : 0s
	//startup successful: [github/startup/good] : 0s
	//test: Startup() -> [true]

}

func ExampleStartup_Failure() {
	uri1 := "github/startup/good"
	uri2 := "github/startup/bad"
	uri3 := "github/startup/depends"
	ex := messaging.NewExchange()

	start = time.Now()

	c := make(chan *messaging.Message, 16)
	testRegister(ex, uri1, c, nil)
	go startupGood(c)

	c = make(chan *messaging.Message, 16)
	testRegister(ex, uri2, c, nil)
	go startupBad(c)

	c = make(chan *messaging.Message, 16)
	testRegister(ex, uri3, c, nil)
	go startupDepends(c, errors.New("startup failure error message"))

	status := startup(ex, time.Second*2, nil)

	fmt.Printf("test: Startup() -> [%v]\n", status)

	//Output:
	//error: startup failure [startup failure error message]
	//test: Startup() -> [false]

}

func startupGood(c chan *messaging.Message) {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
		default:
		}
	}
}

func startupBad(c chan *messaging.Message) {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			time.Sleep(time.Second + time.Millisecond*100)
			messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
		default:
		}
	}
}

func startupDepends(c chan *messaging.Message, err error) {
	for {
		select {
		case msg, open := <-c:
			if !open {
				return
			}
			if err != nil {
				time.Sleep(time.Second)
				s := core.NewStatusDuration(0, time.Since(start))
				s.Err = err
				messaging.SendReply(msg, s)
			} else {
				time.Sleep(time.Second + (time.Millisecond * 900))
				messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
			}

		default:
		}
	}
}
