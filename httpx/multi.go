package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"sync"
)

type OnResponse func(resp *http.Response, status *core.Status) (failure bool)

func MultiExchange(reqs []*http.Request, ex core.HttpExchange, handler OnResponse) ([]core.ExchangeResult, *core.Status) {
	cnt := len(reqs)
	if cnt == 0 {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: no requests were found to process"))
	}
	if ex == nil {
		ex = Exchange
	}
	var wg sync.WaitGroup
	failures := 0

	results := make([]core.ExchangeResult, cnt)
	for i := 0; i < cnt; i++ {
		if reqs[i] == nil {
			continue
		}
		wg.Add(1)
		go func(req *http.Request, res *core.ExchangeResult) {
			defer wg.Done()
			res.Resp, res.Status = ex(req)
			if handler != nil {
				if handler(res.Resp, res.Status) {
					failures++
				}
			}
		}(reqs[i], &results[i])
	}
	wg.Wait()
	if failures > 0 {
		return results, core.NewStatusError(core.StatusExecError, errors.New(fmt.Sprintf("error: %v requests failed", failures)))
	}
	return results, core.StatusOK()
}
