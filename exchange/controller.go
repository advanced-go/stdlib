package exchange

type Controller struct {
	RouteName string
	Primary   *Resource
	Secondary *Resource
}

func NewController(routeName string, primary, secondary *Resource) *Controller {
	c := new(Controller)
	c.RouteName = routeName
	c.Primary = primary
	c.Secondary = secondary
	return c
}
