package host

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/http/httptest"
)

func ExampleHost() {
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("", "http://localhost:8081/github/advanced-go/search:google?q=golang", nil)

	hostExchange[core.Output](r, req, 0, func(r *http.Request) (*http.Response, *core.Status) {
		return &http.Response{StatusCode: http.StatusOK}, core.StatusOK()
	})

	fmt.Printf("test: hostExchange() -> [resp:%v]\n", r.Result().StatusCode)

	//Output:
	//test: hostExchange() -> [resp:200]

}
