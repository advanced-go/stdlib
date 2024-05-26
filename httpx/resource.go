package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
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

type Resource[T any, U any, V any] struct {
	Name             string
	List             []T
	Identity         *http.Response
	MethodNotAllowed *http.Response
	Finalize         FinalizeFunc
	Match            MatchFunc[T]
	PostProcess      PostProcessFunc[T, V]
	PatchProcess     PatchProcessFunc[T, U]
}

func NewBasicResource[T any](name string, match MatchFunc[T], finalize FinalizeFunc) *Resource[T, struct{}, struct{}] {
	r := new(Resource[T, struct{}, struct{}])
	r.Identity = NewAuthorityResponse(name)
	r.MethodNotAllowed = NewResponse(core.NewStatus(http.StatusMethodNotAllowed), nil)
	r.Finalize = finalize
	if r.Finalize == nil {
		r.Finalize = defaultFinalize()
	}
	r.Match = match
	if r.Match == nil {
		r.Match = defaultMatch[T]()
	}
	return r
}

func NewResource[T any, U any, V any](name string, match MatchFunc[T], finalize FinalizeFunc, patch PatchProcessFunc[T, U], post PostProcessFunc[T, V]) *Resource[T, U, V] {
	r := new(Resource[T, U, V])
	r.Identity = NewAuthorityResponse(name)
	r.MethodNotAllowed = NewResponse(core.NewStatus(http.StatusMethodNotAllowed), nil)
	r.Finalize = finalize
	if r.Finalize == nil {
		r.Finalize = defaultFinalize()
	}
	r.Match = match
	if r.Match == nil {
		r.Match = defaultMatch[T]()
	}
	r.PatchProcess = patch
	r.PostProcess = post
	return r
}

func (r *Resource[T, U, V]) Count() int {
	return len(r.List)
}

func (r *Resource[T, U, V]) Empty() {
	r.List = nil
}

func (r *Resource[T, U, V]) Do(req *http.Request) (*http.Response, *core.Status) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Path == core.AuthorityRootPath {
			return r.Identity, core.StatusOK()
		}
		return GetT[T](req, r.List, r.Match, r.Finalize)
	case http.MethodPut:
		return PutT[T](req, &r.List, r.Finalize)
	case http.MethodPatch:
		return PatchT(req, &r.List, r.PatchProcess, r.Finalize)
	case http.MethodPost:
		return PostT(req, &r.List, r.PostProcess, r.Finalize)
	case http.MethodDelete:
		return DeleteT(req, &r.List, r.Match, r.Finalize)
	default:
		status := core.NewStatusError(http.StatusMethodNotAllowed, errors.New(fmt.Sprintf("unsupported method: %v", req.Method)))
		return NewResponse(status, status.Err), core.NewStatus(http.StatusMethodNotAllowed)
	}
}

func defaultFinalize() func(resp *http.Response) {
	return func(resp *http.Response) {
		if resp.Header == nil {
			resp.Header = make(http.Header)
			if resp.Request != nil {
				resp.Header.Add("X-Method", resp.Request.Method)
			}
		}
	}
}

func defaultMatch[T any]() func(item *T, r *http.Request) bool {
	return func(item *T, r *http.Request) bool { return false }
}
