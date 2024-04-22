package messaging

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

// Example of how to get the value of an anonymous field in a struct
//
// https://go.dev/play/p/yQULMVaQK0

func ExampleNewStatusDuration() {
	s := core.NewStatusDuration(http.StatusOK, time.Millisecond*200)

	fmt.Printf("test: NewStatusDuration() -> [status:%v] [ok:%v] [duration:%v]\n", s, s.OK(), s.Duration)

	//Output:
	//test: NewStatusDuration() -> [status:OK] [ok:true] [duration:200ms]

}
