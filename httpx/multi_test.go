package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func onResponse(resp *http.Response, status *core.Status) (failure, proceed bool) {
	//fmt.Printf("[req:%v]\n [resp:%v]\n [status:%v]\n", resp.Request, resp, status)
	fmt.Printf("[status:%v]\n", status)
	return !status.OK(), true
}

func ExampleMultiExchange() {
	var reqs []*http.Request
	r, _ := http.NewRequest("", "https://www.google.com/search?q=golang", nil)
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
	//[status:OK]
	//[status:OK]
	//[status:OK]
	//[status:OK]
	//test: ExampleMultiExchange() -> [count:4] [OK]

}

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
