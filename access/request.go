package access

import (
	"net/http"
)

type Request interface {
	Url() string
	Header() http.Header
	Method() string
	//RouteName() string
	//Duration() time.Duration
	//ContextTimeout(ctx context.Context) (context.Context, context.CancelFunc)
}

type request struct {
	url    string
	header http.Header
	method string
	//duration  time.Duration
	//routeName string
}

func NewRequest(method, url string, h http.Header) Request {
	r := new(request)
	r.url = url
	r.method = method
	r.header = h
	//r.routeName = routeName
	//r.duration = duration
	return r
}

func (r *request) Url() string {
	return r.url
}

func (r *request) Method() string {
	return r.method
}

/*
func (r *request) RouteName() string {
	return r.routeName
}

func (r *request) Duration() time.Duration {
	return r.duration
}


*/

func (r *request) Header() http.Header {
	return r.header
}
