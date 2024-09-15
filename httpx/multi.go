package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"sync"
)

type OnResponse func(resp *http.Response, status *core.Status)

func MultiExchange(handler OnResponse, ex core.HttpExchange, reqs []*http.Request) []core.ExchangeResult {
	cnt := len(reqs)
	if cnt == 0 || ex == nil {
		return nil
	}
	var wg sync.WaitGroup

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
				handler(res.Resp, res.Status)
			}
		}(reqs[i], &results[i])
	}
	wg.Wait()
	return results
}
