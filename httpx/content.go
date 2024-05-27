package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

type MatchFunc2[T any] func(r *http.Request, item *T) bool
type PatchProcessFunc2[T any, U any] func(r *http.Request, list *[]T, content *U) *core.Status
type PostProcessFunc2[T any, V any] func(r *http.Request, list *[]T, content *V) *core.Status

type Content[T any, U any, V any] interface {
	Count() int
	Empty()
	Get(r *http.Request) ([]T, *core.Status)
	Put(r *http.Request, items []T) *core.Status
	Delete(r *http.Request) *core.Status
	Patch(r *http.Request, post *U) *core.Status
	Post(r *http.Request, post *V) *core.Status
}

type ListContent[T any, U any, V any] struct {
	List  []T
	match MatchFunc2[T]
	patch PatchProcessFunc2[T, U]
	post  PostProcessFunc2[T, V]
}

func NewListContent[T any, U any, V any](match MatchFunc2[T], patch PatchProcessFunc2[T, U], post PostProcessFunc2[T, V]) Content[T, U, V] {
	c := new(ListContent[T, U, V])
	c.match = match
	c.patch = patch
	c.post = post
	if c.match == nil {
		if c.match == nil {
			c.match = func(r *http.Request, item *T) bool { return false }
		}
	}
	return c
}

func (c *ListContent[T, U, V]) Count() int {
	return len(c.List)
}

func (c *ListContent[T, U, V]) Empty() {
	c.List = nil
}

func (c *ListContent[T, U, V]) Get(r *http.Request) ([]T, *core.Status) {
	if r == nil {
		return nil, core.NewStatus(core.StatusInvalidArgument)
	}
	var items []T
	for _, target := range c.List {
		if c.match(r, &target) {
			items = append(items, target)
		}
	}
	if len(items) == 0 {
		return nil, core.StatusNotFound()
	}
	return items, core.StatusOK()
}

func (c *ListContent[T, U, V]) Put(r *http.Request, items []T) *core.Status {
	if r == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	if len(items) != 0 {
		c.List = append(c.List, items...)
	}
	return core.StatusOK()
}

func (c *ListContent[T, U, V]) Delete(r *http.Request) *core.Status {
	if r == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	deleted := true
	for deleted {
		deleted = false
		for i, target := range c.List {
			if c.match(r, &target) {
				c.List = append(c.List[:i], c.List[i+1:]...)
				deleted = true
				break
			}
		}
	}
	return core.StatusOK()
}

func (c *ListContent[T, U, V]) Patch(r *http.Request, patch *U) *core.Status {
	if r == nil || patch == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	if c.patch == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	return c.patch(r, &c.List, patch)
}

func (c *ListContent[T, U, V]) Post(r *http.Request, post *V) *core.Status {
	if r == nil || post == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	if c.post == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	return c.post(r, &c.List, post)
}
