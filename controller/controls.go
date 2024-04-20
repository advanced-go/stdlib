package controller

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	uri2 "github.com/advanced-go/stdlib/uri"
	"sync"
)

var (
	ctrlMap = NewControls()
)

// RegisterController - add a controller for a URI
func RegisterController(uri string, ctrl *Controller) *core.Status {
	return ctrlMap.register(uri, ctrl)
}

func Lookup(uri string) (*Controller, *core.Status) {
	nid, _, ok := uri2.UprootUrn(uri)
	if !ok {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: path is invalid: [%v]", uri)))
	}
	return ctrlMap.lookupByNID(nid)
}

// controls - key value pairs of a URI -> *Controller
type controls struct {
	m *sync.Map
}

// NewControls - create a new Controls map
func NewControls() *controls {
	p := new(controls)
	p.m = new(sync.Map)
	return p
}

// Register - add a controller
func (p *controls) register(uri string, ctrl *Controller) *core.Status {
	if len(uri) == 0 {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: path is empty"))
	}
	nid, _, ok := uri2.UprootUrn(uri)
	if !ok {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: path is invalid: [%v]", uri)))
	}
	if ctrl == nil {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: Controller is nil: [%v]", uri)))
	}
	_, ok1 := p.m.Load(nid)
	if ok1 {
		return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: Controller already exists: [%v]", uri)))
	}
	p.m.Store(nid, ctrl)
	return core.StatusOK()
}

// Lookup - get a Controller using a URI as the key
func (p *controls) lookup(uri string) (*Controller, *core.Status) {
	nid, _, ok := uri2.UprootUrn(uri)
	if !ok {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: path is invalid: [%v]", uri)))
	}
	return p.lookupByNID(nid)
}

// LookupByNID - get a Controller using an NID as a key
func (p *controls) lookupByNID(nid string) (*Controller, *core.Status) {
	v, ok := p.m.Load(nid)
	if !ok {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: Controller does not exist: [%v]", nid)))
	}
	if ctrl, ok1 := v.(*Controller); ok1 {
		return ctrl, core.StatusOK()
	}
	return nil, core.NewStatus(core.StatusInvalidContent)
}
