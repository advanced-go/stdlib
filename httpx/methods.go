package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

type FinalizeFunc func(*http.Response)

type MatchFunc[T any] func(item *T, r *http.Request) bool
type PatchProcessFunc[T any, U any] func(list *[]T, content *U) *http.Response
type PostProcessFunc[T any, V any] func(list *[]T, content *V) *http.Response

type GetFunc[T any] func(r *http.Request, list []T, match MatchFunc[T], finalize FinalizeFunc) *http.Response
type DeleteFunc[T any] func(r *http.Request, list *[]T, match MatchFunc[T], finalize FinalizeFunc) *http.Response
type PutFunc[T any] func(r *http.Request, list *[]T, finalize FinalizeFunc) *http.Response
type PatchFunc[T any, U any] func(r *http.Request, list *[]T, patch PatchProcessFunc[T, U], finalize FinalizeFunc) *http.Response
type PostFunc[T any, V any] func(r *http.Request, list *[]T, post PostProcessFunc[T, V], finalize FinalizeFunc) *http.Response

func FinalizeResponse(status *core.Status, r *http.Request, finalize FinalizeFunc) *http.Response {
	resp := NewResponse(status, status.Err)
	resp.Request = r
	if finalize != nil {
		finalize(resp)
	}
	return resp
}

func PutT[T any](r *http.Request, list *[]T, finalize FinalizeFunc) (*http.Response, *core.Status) {
	items, status := json.New[[]T](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize), status
	}
	if len(items) == 0 {
		return FinalizeResponse(core.StatusNotFound(), r, finalize), core.StatusNotFound()
	}
	*list = append(*list, items...)
	return FinalizeResponse(core.StatusOK(), r, finalize), core.StatusOK()
}

func GetT[T any](r *http.Request, list []T, match MatchFunc[T], finalize FinalizeFunc) (*http.Response, *core.Status) {
	if match == nil {
		return FinalizeResponse(core.StatusBadRequest(), r, finalize), core.StatusBadRequest()
	}
	var items []T
	for _, target := range list {
		if match(&target, r) {
			items = append(items, target)
		}
	}
	if len(items) == 0 {
		return FinalizeResponse(core.StatusNotFound(), r, finalize), core.StatusNotFound()
	}
	resp, status := NewJsonResponse(items, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize), status
	}
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp, core.StatusOK()
}

func DeleteT[T any](r *http.Request, list *[]T, match MatchFunc[T], finalize FinalizeFunc) (*http.Response, *core.Status) {
	if match == nil {
		return FinalizeResponse(core.StatusNotFound(), r, finalize), core.StatusNotFound()
	}
	for i, target := range *list {
		if match(&target, r) {
			*list = append((*list)[:i], (*list)[i+1:]...)
			break
		}
	}
	return FinalizeResponse(core.StatusOK(), r, finalize), core.StatusOK()
}

func PatchT[T, U any](r *http.Request, list *[]T, patch PatchProcessFunc[T, U], finalize FinalizeFunc) (*http.Response, *core.Status) {
	if patch == nil {
		return FinalizeResponse(core.StatusBadRequest(), r, finalize), core.StatusBadRequest()
	}
	content, status := json.New[U](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize), status
	}
	resp := patch(list, &content)
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp, core.StatusOK()

}

func PostT[T any, V any](r *http.Request, list *[]T, post PostProcessFunc[T, V], finalize FinalizeFunc) (*http.Response, *core.Status) {
	if post == nil {
		return FinalizeResponse(core.StatusBadRequest(), r, finalize), core.StatusBadRequest()
	}
	content, status := json.New[V](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize), status
	}
	resp := post(list, &content)
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp, core.StatusOK()
}
