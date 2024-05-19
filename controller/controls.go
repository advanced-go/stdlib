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

// RegisterController - add a controller for an egress route
func RegisterController(ctrl *Controller) error {
	if ctrl == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller is nil"))
	}
	if ctrl.Router == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller router is nil"))
	}
	if ctrl.Router.Primary == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller router primary resource is nil"))
	}
	if len(ctrl.Router.Primary.Authority) == 0 {
		if ctrl.Router.Primary.Host == "" {
			return errors.New(fmt.Sprintf("invalid argument: Controller router primary resource host is empty"))
		}
		return ctrlMap.register(ctrl)
	}
	return ctrlMap.registerWithAuthority(ctrl)
}

func Lookup(req *http.Request) (ctrl *Controller, status *core.Status) {
	if req == nil || req.URL == nil {
		return nil, core.NewStatus(http.StatusNotFound)
	}

	// Try host first
	ctrl, status = ctrlMap.lookup(req.Host)
	if status.OK() {
		return
	}

	// Default to embedded authority
	p := uri2.Uproot(req.URL.Path)
	if p.Valid {
		ctrl, status = ctrlMap.lookup(p.Authority)
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

func (p *controls) register(ctrl *Controller) error {
	if ctrl == nil {
		return errors.New("invalid argument: Controller is nil")
	}
	_, ok1 := p.m.Load(ctrl.Router.Primary.Host)
	if ok1 {
		return errors.New(fmt.Sprintf("invalid argument: Controller already exists for authority: [%v]", ctrl.Router.Primary))
	}
	p.m.Store(ctrl.Router.Primary.Host, ctrl)
	return nil
}

func (p *controls) registerWithAuthority(ctrl *Controller) error {
	if ctrl == nil {
		return errors.New("invalid argument: Controller is nil")
	}
	//parsed := uri2.Uproot(ctrl.Router.Primary.Authority)
	//if !parsed.Valid {
	//	return errors.New(fmt.Sprintf("invalid argument: Controller primary authority is invalid: [%v] [%v]", ctrl.Router.Primary.Authority, parsed.Err))
	//}
	_, ok1 := p.m.Load(ctrl.Router.Primary.Authority)
	if ok1 {
		return errors.New(fmt.Sprintf("invalid argument: Controller already exists for authority : [%v]", ctrl.Router.Primary.Authority))
	}
	p.m.Store(ctrl.Router.Primary.Authority, ctrl)
	return nil
}

// Lookup - get a Controller using a URI as the key
/*
func (p *controls) lookup(uri string) (*Controller, *core.Status) {
	if uri == "" {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: uri is empty"))
	}
	parsed := uri2.Uproot(uri)
	if !parsed.Valid {
		return nil, core.NewStatusError(core.StatusInvalidArgument, parsed.Err)
	}
	return p.lookupByAuthority(parsed.Authority)
}


*/

// Lookup - get a Controller using an authority
func (p *controls) lookup(authority string) (*Controller, *core.Status) {
	if authority == "" {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: authority is empty"))
	}
	v, ok := p.m.Load(authority)
	if !ok {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: Controller does not exist: [%v]", authority)))
	}
	if ctrl, ok1 := v.(*Controller); ok1 {
		return ctrl, core.StatusOK()
	}
	return nil, core.NewStatus(core.StatusInvalidContent)
}
