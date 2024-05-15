package httpx

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

const (
	internalError           = "Internal Error"
	fileScheme              = "file"
	contextDeadlineExceeded = "context deadline exceeded"
)

var (
	Client = http.DefaultClient
)

func init() {
	t, ok := http.DefaultTransport.(*http.Transport)
	if ok {
		// Used clone instead of assignment due to presence of sync.Mutex fields
		var transport = t.Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		transport.MaxIdleConns = 200
		transport.MaxIdleConnsPerHost = 100
		Client = &http.Client{Transport: transport, Timeout: time.Second * 5}
	} else {
		Client = &http.Client{Transport: http.DefaultTransport, Timeout: time.Second * 5}
	}
}

func DeadlineExceededError(t any) bool {
	if t == nil {
		return false
	}
	if r, ok := t.(*http.Request); ok {
		return r.Context() != nil && r.Context().Err() == context.DeadlineExceeded
	}
	if e, ok := t.(error); ok {
		return e == context.DeadlineExceeded
	}
	return false
}

// Do - process an HTTP request, checking for file:// scheme
func Do(req *http.Request) (resp *http.Response, status *core.Status) {
	if req == nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument : request is nil"))
	}
	if req.URL.Scheme == fileScheme {
		resp1, status1 := readResponse(req.URL)
		if !status1.OK() {
			return resp1, status1.AddLocation()
		}
		return resp1, core.NewStatus(resp1.StatusCode)
	}
	var err error

	resp, err = Client.Do(req)
	if err != nil {
		// catch connectivity error, even with a valid URL
		if resp == nil {
			resp = serverErrorResponse()
		}
		// check for an error of deadline exceeded
		if req.Context() != nil && req.Context().Err() == context.DeadlineExceeded {
			resp.StatusCode = http.StatusGatewayTimeout
		}
		return resp, core.NewStatusError(resp.StatusCode, err)
	}
	return resp, core.NewStatus(resp.StatusCode)
}

// DoExchange - process an HTTP call utilizing a controller if configured
func DoExchange(req *http.Request) (*http.Response, *core.Status) {
	//req, _ := http.NewRequestWithContext(ctx, method, url, body)
	//if h != nil {
	//	req.Header = h
	//}
	ctrl, status := controller.Lookup(req)
	if status.OK() {
		return ctrl.Do(Do, req)
	}
	return Do(req)
}

func serverErrorResponse() *http.Response {
	resp := new(http.Response)
	resp.StatusCode = http.StatusInternalServerError
	resp.Status = internalError
	return resp
}
