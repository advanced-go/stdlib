package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

type exchangeResult struct {
	Failure bool
	Resp    *http.Response
	Status  *core.Status
}

var results []exchangeResult

func onResponse(item RequestItem, resp *http.Response, status *core.Status) (proceed bool) {
	//fmt.Printf("[req:%v]\n [resp:%v]\n [status:%v]\n", resp.Request, resp, status)
	fmt.Printf("[id:%v] [status:%v]\n", item.Id, status)
	results = append(results, exchangeResult{
		Failure: false,
		Resp:    resp,
		Status:  status,
	})
	return true
}

func ExampleMultiExchange() {
	var reqs []RequestItem
	r, _ := http.NewRequest("", "https://www.google.com/search?q=golang", nil)
	reqs = append(reqs, RequestItem{Id: "1", Request: r})

	r, _ = http.NewRequest("", "https://www.search.yahoo.com/search?q=golang", nil)
	reqs = append(reqs, RequestItem{Id: "2", Request: r})

	r, _ = http.NewRequest("", "https://www.bing.com/search?q=golang", nil)
	reqs = append(reqs, RequestItem{Id: "3", Request: r})

	r, _ = http.NewRequest("", "https://www.duckduckgo.com/search?q=golang", nil)
	reqs = append(reqs, RequestItem{Id: "4", Request: r})

	MultiExchange(reqs, onResponse)
	fmt.Printf("test: MultiExchange() -> [count:%v]\n", len(results))

	//Output:
	//[id:1] [status:OK]
	//[id:3] [status:OK]
	//[id:4] [status:OK]
	//[id:2] [status:OK]
	//test: MultiExchange() -> [count:4]

}

/*
func ExampleMultiExchangeFailure() {
	var reqs []*http.Request
	r, _ := http.NewRequest("", "http://localhost:8080/search?q=golang", nil)
	reqs = append(reqs, r)

	r, _ = http.NewRequest("", "https://www.search.yahoo.com/search?q=golang", nil)
	reqs = append(reqs, r)

	r, _ = http.NewRequest("", "https://www.bing.com/search?q=golang", nil)
	reqs = append(reqs, r)

	r, _ = http.NewRequest("", "https://www.duckduckgo.com/search?q=golang", nil)
	reqs = append(reqs, r)

	results, status := MultiExchange(reqs, onResponse)
	fmt.Printf("test: ExampleMultiExchange() -> [count:%v] [%v]\n", len(results), status)

	//Output:
	//[status:Internal Error [Get "http://localhost:8080/search?q=golang": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.]]
	//[status:OK]
	//[status:OK]
	//[status:OK]
	//test: ExampleMultiExchange() -> [count:4] [Execution Error [error: request failures]]

}


*/
