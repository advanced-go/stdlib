package httpx

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func onResponse(resp *http.Response, status *core.Status) {
	//fmt.Printf("[req:%v]\n [resp:%v]\n [status:%v]\n", resp.Request, resp, status)
	fmt.Printf("[status:%v]\n", status)

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

	results := MultiExchange(onResponse, Do, reqs)
	fmt.Printf("test: ExampleMultiExchange() -> [count:%v]\n", len(results))

	//Output:
	//[status:OK]
	//[status:OK]
	//[status:Accepted]
	//[status:OK]
	//test: ExampleMultiExchange() -> [count:4]

}
