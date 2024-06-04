package access

import "net/http"

type Request interface {
	Url() string
	Header() http.Header
	Method() string
}

type request struct {
	url    string
	header http.Header
	method string
}

func NewRequest(url, method string, h http.Header) Request {
	r := new(request)
	r.url = url
	r.method = method
	r.header = h
	return r
}

func (r *request) Url() string {
	return r.url
}

func (r *request) Method() string {
	return r.method
}

func (r *request) Header() http.Header {
	return r.header
}
