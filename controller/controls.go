package controller

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	uri2 "github.com/advanced-go/stdlib/uri"
	"net/http"
	"sync"
)

var (
	ctrlMap = NewControls()
)

// RegisterController - add a controller for a host authority
func RegisterController(ctrl *Controller) error {
	return ctrlMap.register(ctrl)
}

// RegisterControllerWithAuthority - add a controller for an embedded authority
func RegisterControllerWithAuthority(authority string, ctrl *Controller) error {
	return ctrlMap.registerWithAuthority(authority, ctrl)
}

func Lookup(req *http.Request) (ctrl *Controller, status *core.Status) {
	if req == nil || req.URL == nil {
		return nil, core.NewStatus(http.StatusNotFound)
	}

	// Try host authority first
	ctrl, status = ctrlMap.lookupByAuthority(req.Host)
	if status.OK() {
		return
	}

	// Default to embedded authority
	p := uri2.Uproot(req.URL.Path)
	if p.Valid {
		ctrl, status = ctrlMap.lookupByAuthority(p.Authority)
		if status.OK() {
			return
		}
	}
	return nil, core.NewStatus(http.StatusNotFound)
}

// controls - key value pairs of an authority -> *Controller
type controls struct {
	m *sync.Map
}

// NewControls - create a new Controls map
func NewControls() *controls {
	p := new(controls)
	p.m = new(sync.Map)
	return p
}

func (p *controls) registerWithAuthority(authority string, ctrl *Controller) error {
	if len(authority) == 0 {
		return errors.New("invalid argument: authority is empty")
	}
	if ctrl == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller is nil for authority: [%v]", authority))
	}
	parsed := uri2.Uproot(authority)
	if !parsed.Valid {
		return errors.New(fmt.Sprintf("invalid argument: authority is invalid: [%v] [%v]", authority, parsed.Err))
	}
	_, ok1 := p.m.Load(parsed.Authority)
	if ok1 {
		return errors.New(fmt.Sprintf("invalid argument: Controller already exists for authority : [%v]", authority))
		//return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: controller already exists: [%v]", authority)))
	}
	p.m.Store(parsed.Authority, ctrl)
	return nil
}

func (p *controls) register(ctrl *Controller) error {
	if ctrl == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller is nil"))
	}
	if ctrl.Router == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller Router is nil"))
	}
	if ctrl.Router.primary == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller Router primary resource is nil"))
	}
	_, ok1 := p.m.Load(ctrl.Router.primary.Authority)
	if ok1 {
		return errors.New(fmt.Sprintf("invalid argument: Controller already exists for authority: [%v]", ctrl.Router.primary))
	}
	p.m.Store(ctrl.Router.primary.Authority, ctrl)
	return nil
}

// Lookup - get a Controller using a URI as the key
func (p *controls) lookup(uri string) (*Controller, *core.Status) {
	parsed := uri2.Uproot(uri)
	if !parsed.Valid {
		return nil, core.NewStatusError(core.StatusInvalidArgument, parsed.Err)
	}
	return p.lookupByAuthority(parsed.Authority)
}

// LookupByAuthority - get a Controller using an authority
func (p *controls) lookupByAuthority(authority string) (*Controller, *core.Status) {
	v, ok := p.m.Load(authority)
	if !ok {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: Controller does not exist: [%v]", authority)))
	}
	if ctrl, ok1 := v.(*Controller); ok1 {
		return ctrl, core.StatusOK()
	}
	return nil, core.NewStatus(core.StatusInvalidContent)
}
