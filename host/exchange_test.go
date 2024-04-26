package host

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/http/httptest"
)

func appHttpExchange(r *http.Request) (*http.Response, *core.Status) {
	status := core.NewStatus(http.StatusTeapot)
	return &http.Response{StatusCode: status.Code}, status
}

func testAuthExchangeOK(r *http.Request) (*http.Response, *core.Status) {
	status := core.StatusOK()
	return &http.Response{StatusCode: status.Code}, status
}

func testAuthExchangeFail(r *http.Request) (*http.Response, *core.Status) {
	status := core.NewStatus(http.StatusUnauthorized)
	//fmt.Fprint(w, "Missing authorization header")
	return &http.Response{StatusCode: status.Code}, status
}

func testExchange(r *http.Request) (*http.Response, *core.Status) {
	req, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://www.google.com/search?q=golang", nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp == nil {
		return resp, core.NewStatus(http.StatusGatewayTimeout)
	} else {
		return resp, core.NewStatus(resp.StatusCode)
	}
}

func Example_TestExchange() {
	pattern := "github/advanced-go/example-domain/activity"
	r, _ := http.NewRequest("PUT", "http://localhost:8080/github/advanced-go/example-domain/activity:entry", nil)

	RegisterExchange(pattern, appHttpExchange)

	rec := httptest.NewRecorder()
	HttpHandler2(rec, r)

	fmt.Printf("test: HttpHandler2() -> %v\n", rec.Result().StatusCode)

	//Output:
	//test: HttpHandler2() -> 418

}
