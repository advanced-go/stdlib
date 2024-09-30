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

type OnResponse func(item RequestItem, resp *http.Response, status *core.Status) (proceed bool)

func MultiExchange(reqs []RequestItem, handler OnResponse) {
	cnt := len(reqs)
	if cnt == 0 || handler == nil {
		fmt.Printf("%v", "error: no requests were found to process, or OnResponse handler is nil")
		return //nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: no requests were found to process"))
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
			if !handler(item, resp, status) {
				return
			}
		}(reqs[i])
	}
	wg.Wait()
}
