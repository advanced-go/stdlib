package httpx

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"sync"
	"sync/atomic"
)

type OnResponse func(resp *http.Response, status *core.Status) (failure bool)

func MultiExchange(reqs []*http.Request, handler OnResponse) ([]core.ExchangeResult, *core.Status) {
	cnt := len(reqs)
	if cnt == 0 {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: no requests were found to process"))
	}
	var wg sync.WaitGroup
	failure := atomic.Bool{}

	results := make([]core.ExchangeResult, cnt)
	for i := 0; i < cnt; i++ {
		if reqs[i] == nil {
			continue
		}
		wg.Add(1)
		go func(req *http.Request, res *core.ExchangeResult) {
			defer wg.Done()
			res.Resp, res.Status = Exchange(req)
			if handler != nil {
				if handler(res.Resp, res.Status) {
					res.Failure = true
					failure.Store(true)
				}
			}
		}(reqs[i], &results[i])
	}
	wg.Wait()
	if failure.Load() {
		return results, core.NewStatusError(core.StatusExecError, errors.New("error: request failures"))
	}
	return results, core.StatusOK()
}
