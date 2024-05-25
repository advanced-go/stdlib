package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

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

func PatchT[PATCH any, T any](r *http.Request, list *[]T, patch PatchProcessFunc[PATCH, T], finalize FinalizeFunc) *http.Response {
	if patch == nil {
		FinalizeResponse(core.StatusBadRequest(), r, finalize)
	}
	content, status := json.New[PATCH](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize)
	}
	resp := patch(&content, list)
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp

}

func PostT[POST any, T any](r *http.Request, list *[]T, post PostProcessFunc[POST, T], finalize FinalizeFunc) *http.Response {
	if post == nil {
		FinalizeResponse(core.StatusBadRequest(), r, finalize)
	}
	content, status := json.New[POST](r.Body, r.Header)
	if !status.OK() {
		return FinalizeResponse(status, r, finalize)
	}
	resp := post(&content, list)
	if finalize != nil {
		resp.Request = r
		finalize(resp)
	}
	return resp
}

type Resource2[T any] struct {
	List      []T
	Authority *http.Response
	MatchFn   func(item any, r *http.Request) bool
	//PatchFn
}

func (r *Resource2[T]) append(items []T) {
	r.List = append(r.List, items...)
}

func (r *Resource2[T]) remove(index int) {
	if index >= 0 && index < len(r.List) {
		r.List = append(r.List[:index], r.List[index+1:]...)
	}
}

func (r *Resource2[T]) get(req *http.Request) (items []T, status *core.Status) {
	return nil, nil
}

func (r *Resource2[T]) delete(req *http.Request) *core.Status {
	return nil
}

func (r *Resource2[T]) Do(req *http.Request) (*http.Response, *core.Status) {
	resp := &http.Response{StatusCode: http.StatusOK}
	switch req.Method {
	case http.MethodGet:
		if req.URL.Path == core.AuthorityRootPath {
			return r.Authority, core.StatusOK()
		}
		return resp, core.StatusOK()
	case http.MethodDelete:
		status := r.delete(req)
		return NewResponseWithStatus(status, status.Err)
	case http.MethodPut:
		items, status := json.New[[]T](req.Body, req.Header)
		if !status.OK() {
			return NewResponseWithStatus(status, status.Err)
		}
		if len(items) == 0 {
			return NewResponseWithStatus(core.StatusNotFound(), nil)
		}
		r.append(items)
		return NewResponseWithStatus(core.StatusOK(), nil)
	case http.MethodPatch:
		//patch, status := json.New[Patch](req.Body, req.Header)
		//if patch != nil {}
		return NewResponseWithStatus(nil, nil)
	case http.MethodPost:
		//patch, status := json.New[Patch](req.Body, req.Header)
		//if patch != nil {}
		return NewResponseWithStatus(nil, nil)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("unsupported method: %v", req.Method)))
		return NewResponseWithStatus(status, status.Err)
	}
}
