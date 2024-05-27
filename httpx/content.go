package httpx

import (
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"sync"
)

type MatchFunc[T any] func(r *http.Request, item *T) bool
type PatchProcessFunc[T any, U any] func(r *http.Request, list *[]T, content *U) *core.Status
type PostProcessFunc[T any, V any] func(r *http.Request, list *[]T, content *V) *core.Status

type Mutex interface {
	Acquire() func()
}

type Lock struct {
	mu sync.Mutex
}

func (l Lock) Acquire() func() {
	l.mu.Lock()
	return func() { l.mu.Unlock() }
}

type Bypass struct{}

func (l Bypass) Acquire() func() {
	return func() {}
}

type Content[T any, U any, V any, W Mutex] interface {
	Count() int
	Empty()
	Get(r *http.Request) ([]T, *core.Status)
	Put(r *http.Request, items []T) *core.Status
	Delete(r *http.Request) *core.Status
	Patch(r *http.Request, post *U) *core.Status
	Post(r *http.Request, post *V) *core.Status
}

type ListContent[T any, U any, V any, W Mutex] struct {
	List  []T
	mutex W
	match MatchFunc[T]
	patch PatchProcessFunc[T, U]
	post  PostProcessFunc[T, V]
}

func NewListContent[T any, U any, V any, W Mutex](match MatchFunc[T], patch PatchProcessFunc[T, U], post PostProcessFunc[T, V]) Content[T, U, V, W] {
	c := new(ListContent[T, U, V, W])
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

func (c *ListContent[T, U, V, W]) Count() int {
	defer c.mutex.Acquire()
	return len(c.List)
}

func (c *ListContent[T, U, V, W]) Empty() {
	defer c.mutex.Acquire()
	c.List = nil
}

func (c *ListContent[T, U, V, W]) Get(r *http.Request) ([]T, *core.Status) {
	if r == nil {
		return nil, core.NewStatus(core.StatusInvalidArgument)
	}
	defer c.mutex.Acquire()
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

func (c *ListContent[T, U, V, W]) Put(r *http.Request, items []T) *core.Status {
	if r == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	defer c.mutex.Acquire()
	if len(items) != 0 {
		c.List = append(c.List, items...)
	}
	return core.StatusOK()
}

func (c *ListContent[T, U, V, W]) Delete(r *http.Request) *core.Status {
	if r == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	defer c.mutex.Acquire()
	count := 0
	deleted := true
	for deleted {
		deleted = false
		for i, target := range c.List {
			if c.match(r, &target) {
				c.List = append(c.List[:i], c.List[i+1:]...)
				deleted = true
				count++
				break
			}
		}
	}
	if count == 0 {
		return core.StatusNotFound()
	}
	return core.StatusOK()
}

func (c *ListContent[T, U, V, W]) Patch(r *http.Request, patch *U) *core.Status {
	if r == nil || patch == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	if c.patch == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	defer c.mutex.Acquire()
	return c.patch(r, &c.List, patch)
}

func (c *ListContent[T, U, V, W]) Post(r *http.Request, post *V) *core.Status {
	if r == nil || post == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}
	if c.post == nil {
		return core.NewStatus(http.StatusBadRequest)
	}
	defer c.mutex.Acquire()
	return c.post(r, &c.List, post)
}
