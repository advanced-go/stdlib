package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

type FinalizeFunc func(*http.Response)

type MatchFunc[T any] func(item *T, r *http.Request) bool
type PatchProcessFunc[PATCH any, T any] func(content *PATCH, list *[]T) *http.Response
type PostProcessFunc[POST any, T any] func(content *POST, list *[]T) *http.Response

type GetFunc[T any] func(r *http.Request, list []T, match MatchFunc[T], finalize FinalizeFunc) *http.Response
type DeleteFunc[T any] func(r *http.Request, list *[]T, match MatchFunc[T], finalize FinalizeFunc) *http.Response
type PutFunc[T any] func(r *http.Request, list *[]T, finalize FinalizeFunc) *http.Response
type PatchFunc[PATCH any, T any] func(r *http.Request, list *[]T, patch PatchProcessFunc[PATCH, T], finalize FinalizeFunc) *http.Response
type PostFunc[POST any, T any] func(r *http.Request, list *[]T, post PostProcessFunc[POST, T], finalize FinalizeFunc) *http.Response

type Resource[POST any, PATCH any, T any] struct {
	Name             string
	List             []T
	Identity         *http.Response
	MethodNotAllowed *http.Response
	Finalize         FinalizeFunc
	Match            MatchFunc[T]
	PostProcess      PostProcessFunc[POST, T]
	PatchProcess     PatchProcessFunc[PATCH, T]
}

func NewResource[POST any, PATCH any, T any](name string, match MatchFunc[T], finalize FinalizeFunc, patch PatchProcessFunc[PATCH, T], post PostProcessFunc[POST, T]) *Resource[POST, PATCH, T] {
	a := new(Resource[POST, PATCH, T])
	a.Identity = NewAuthorityResponse(name)
	a.MethodNotAllowed = NewResponse(core.NewStatus(http.StatusMethodNotAllowed), nil)
	a.Finalize = finalize
	if a.Finalize == nil {
		a.Finalize = func(resp *http.Response) {
			if resp.Header == nil {
				resp.Header = make(http.Header)
				if resp.Request != nil {
					resp.Header.Add("X-Method", resp.Request.Method)
				}
			}
		}
	}
	a.Match = match
	if a.Match == nil {
		a.Match = func(item *T, r *http.Request) bool { return false }
	}
	return a
}

func (a *Resource[POST, PATCH, T]) Do(req *http.Request) *http.Response {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Path == core.AuthorityRootPath {
			return a.Identity
		}
		return GetT[T](req, a.List, a.Match, a.Finalize)
	case http.MethodPut:
		return PutT[T](req, &a.List, a.Finalize)
	case http.MethodPatch:
		if a.PatchProcess == nil {
			return NewResponse(core.NewStatus(core.StatusInvalidArgument), nil)
		}
		return PatchT(req, &a.List, a.PatchProcess, a.Finalize)
	case http.MethodPost:
		if a.PostProcess == nil {
			return NewResponse(core.NewStatus(core.StatusInvalidArgument), nil)
		}
		return PostT(req, &a.List, a.PostProcess, a.Finalize)
	case http.MethodDelete:
		return DeleteT(req, &a.List, a.Match, a.Finalize)
	default:
		status := core.NewStatusError(http.StatusMethodNotAllowed, errors.New(fmt.Sprintf("unsupported method: %v", req.Method)))
		return NewResponse(status, status.Err)
	}
}

func FinalizeResponse(status *core.Status, r *http.Request, finalize FinalizeFunc) *http.Response {
	resp := NewResponse(status, status.Err)
	resp.Request = r
	if finalize != nil {
		finalize(resp)
	}
	return resp
}
