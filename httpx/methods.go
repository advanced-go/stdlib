package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

func FinalizeResponse(status *core.Status, r *http.Request, finalize FinalizeFunc) *http.Response {
	resp := NewResponse(status, status.Err)
	resp.Request = r
	if finalize != nil {
		finalize(resp)
	}
	return resp
}

func PutT[T any](r *http.Request, list *[]T, finalize FinalizeFunc) *http.Response {
	items, status := json.New[[]T](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize)
	}
	if len(items) == 0 {
		return FinalizeResponse(core.StatusNotFound(), r, finalize)
	}
	*list = append(*list, items...)
	return FinalizeResponse(core.StatusOK(), r, finalize)
}

func GetT[T any](r *http.Request, list []T, match MatchFunc[T], finalize FinalizeFunc) *http.Response {
	if match == nil {
		return FinalizeResponse(core.StatusBadRequest(), r, finalize)
	}
	var items []T
	for _, target := range list {
		if match(&target, r) {
			items = append(items, target)
		}
	}
	if len(items) == 0 {
		return FinalizeResponse(core.StatusNotFound(), r, finalize)
	}
	resp, status := NewJsonResponse(items, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize)
	}
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp
}

func DeleteT[T any](r *http.Request, list *[]T, match MatchFunc[T], finalize FinalizeFunc) *http.Response {
	if match == nil {
		return FinalizeResponse(core.StatusNotFound(), r, finalize)
	}
	for i, target := range *list {
		if match(&target, r) {
			*list = append((*list)[:i], (*list)[i+1:]...)
			break
		}
	}
	return FinalizeResponse(core.StatusOK(), r, finalize)
}

func PatchT[T, U any](r *http.Request, list *[]T, patch PatchProcessFunc[T, U], finalize FinalizeFunc) *http.Response {
	if patch == nil {
		FinalizeResponse(core.StatusBadRequest(), r, finalize)
	}
	content, status := json.New[U](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize)
	}
	resp := patch(list, &content)
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp

}

func PostT[T any, V any](r *http.Request, list *[]T, post PostProcessFunc[T, V], finalize FinalizeFunc) *http.Response {
	if post == nil {
		FinalizeResponse(core.StatusBadRequest(), r, finalize)
	}
	content, status := json.New[V](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize)
	}
	resp := post(list, &content)
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp
}
