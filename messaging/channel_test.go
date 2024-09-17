package messaging

import (
	"fmt"
	"time"
)

func ExampleNewChannel() {
	c := NewChannel()

	fmt.Printf("test: NewChannel() -> [enabled:%v]\n", c.IsEnabled())

	c.Enable()
	fmt.Printf("test: NewChannel_Enable()  -> [enabled:%v]\n", c.IsEnabled())

	c.Disable()
	fmt.Printf("test: NewChannel_Disable() -> [enabled:%v]\n", c.IsEnabled())

	c.Close()
	fmt.Printf("test: NewChannel_Close()   -> [closed:%v]\n", c.C == nil)

	//Output:
	//test: NewChannel() -> [enabled:false]
	//test: NewChannel_Enable()  -> [enabled:true]
	//test: NewChannel_Disable() -> [enabled:false]
	//test: NewChannel_Close()   -> [closed:true]

}

func ExampleNewChannel_Send() {
	c := NewChannel()
	msg := NewControlMessage("", "", StartupEvent)

	c.Enable()
	c.Send(msg)
	time.Sleep(time.Second * 2)

	msg2 := <-c.C
	fmt.Printf("test: NewChannel_Send() -> [enabled:%v] [msg:%v]\n", c.enabled, msg2)

	//Output:
	//test: NewChannel_Send() -> [enabled:true] [msg:[chan:CTRL] [from:] [to:] [event:startup]]

}
