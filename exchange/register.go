package exchange

import (
	"errors"
	"fmt"
)

// RegisterController - add a controller for an egress route
func RegisterController(ctrl *Controller) error {
	if ctrl == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller is nil"))
	}
	//if ctrl.Router == nil {
	//	return errors.New(fmt.Sprintf("invalid argument: Controller router is nil"))
	//}
	if ctrl.Primary == nil {
		return errors.New(fmt.Sprintf("invalid argument: Controller rimary resource is nil"))
	}
	if len(ctrl.Primary.Authority) == 0 {
		if ctrl.Primary.Host == "" {
			return errors.New(fmt.Sprintf("invalid argument: Controller primary resource host is empty"))
		}
		return ctrlMap.register(ctrl)
	}
	return ctrlMap.registerWithAuthority(ctrl)
}
