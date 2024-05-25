package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

type Resource2[T any] struct {
	List      []T
	Authority *http.Response
	MatchFn   func(item any, r *http.Request) bool
	PatchFn   func(item any, patch *Patch)
}

func NewResource2[T any](authority string, match func(item any, r *http.Request) bool, patch func(item any, patch *Patch)) *Resource2[T] {
	r := new(Resource2[T])
	r.Authority = NewAuthorityResponse(authority)
	r.MatchFn = match
	r.PatchFn = patch
	return r
}

func (r *Resource2[T]) Count() int {
	return len(r.List)
}

func (r *Resource2[T]) Empty() {
	r.List = nil
}

func (r *Resource2[T]) Get(req *http.Request) (items []T, status *core.Status) {
	if r.MatchFn == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("MatchFunc() is nil"))
	}
	for _, target := range r.List {
		if r.MatchFn(&target, req) {
			items = append(items, target)
		}
	}
	if len(items) == 0 {
		return nil, core.StatusNotFound()
	}
	return items, core.StatusOK()
}

func (r *Resource2[T]) Put(items []T) *core.Status {
	if len(items) > 0 {
		r.List = append(r.List, items...)
	}
	return core.StatusOK()
}

func (r *Resource2[T]) Patch(req *http.Request, patch *Patch) *core.Status {
	if r.MatchFn == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("MatchFunc() is nil"))
	}
	if r.PatchFn == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("PatchFunc() is nil"))
	}
	for i, target := range r.List {
		if r.MatchFn(&target, req) {
			r.PatchFn(&r.List[i], patch)
		}
	}
	return core.StatusOK()
}

func (r *Resource2[T]) Delete(req *http.Request) *core.Status {
	if r.MatchFn == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("MatchFunc() is nil"))
	}
	for i, target := range r.List {
		if r.MatchFn(&target, req) {
			r.List = append(r.List[:i], r.List[i+1:]...)
		}
	}
	return core.StatusOK()
}

func (r *Resource2[T]) Do(req *http.Request) (*http.Response, *core.Status) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Path == core.AuthorityRootPath {
			return r.Authority, core.StatusOK()
		}
		items, status := r.Get(req)
		if !status.OK() {
			return NewResponseWithStatus(status, status.Err)
		}
		resp, status1 := NewJsonResponse(items, req.Header)
		if !status1.OK() {
			return NewResponseWithStatus(status, status.Err)
		}
		return resp, core.StatusOK()
	case http.MethodPut:
		items, status := json.New[[]T](req.Body, req.Header)
		if !status.OK() {
			return NewResponseWithStatus(status, status.Err)
		}
		if len(items) == 0 {
			return NewResponseWithStatus(core.StatusNotFound(), nil)
		}
		r.Put(items)
		return NewResponseWithStatus(core.StatusOK(), nil)
	case http.MethodPatch:
		patch, status := json.New[Patch](req.Body, req.Header)
		if !status.OK() {
			return NewResponseWithStatus(status, status.Err)
		}
		status = r.Patch(req, &patch)
		return NewResponseWithStatus(status, status.Err)
	case http.MethodDelete:
		status := r.Delete(req)
		return NewResponseWithStatus(status, status.Err)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("unsupported method: %v", req.Method)))
		return NewResponseWithStatus(status, status.Err)
	}
}
