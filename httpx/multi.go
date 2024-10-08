package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"sync"
)

type RequestItem struct {
	Id      string
	Request *http.Request
}

type OnResponse func(id string, resp *http.Response, status *core.Status)

func MultiExchange(reqs []RequestItem, handler OnResponse) {
	cnt := len(reqs)
	if cnt == 0 || handler == nil {
		fmt.Printf("%v", "error: no requests were found to process, or OnResponse handler is nil")
		return
	}
	var wg sync.WaitGroup

	for i := 0; i < cnt; i++ {
		if reqs[i].Request == nil {
			continue
		}
		wg.Add(1)
		go func(item RequestItem) {
			defer wg.Done()
			resp, status := Exchange(item.Request)
			handler(item.Id, resp, status)
		}(reqs[i])
	}
	wg.Wait()
}
