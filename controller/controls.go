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

func UpdatePrimaryExchange(list []core.HttpExchange) (status *core.Status) {
	if list == nil {
		return core.NewStatus(core.StatusInvalidArgument)
	}

	var ctrl *Controller
	for _, ex := range list {
		ctrl, status = LookupWithAuthority(core.Authority(ex))
		if status.OK() && ctrl.Router.Primary.Handler == nil {
			ctrl.Router.Primary.Handler = ex
		}
	}
	return status
}

func LookupWithAuthority(authority string) (ctrl *Controller, status *core.Status) {
	return ctrlMap.lookup(authority)
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

func (p *controls) remove(key string) {
	p.m.Delete(key)
}
