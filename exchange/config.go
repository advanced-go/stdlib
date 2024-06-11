package exchange

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

type Config struct {
	RouteName string `json:"route"`
	Host      string `json:"host"`
	Authority string `json:"authority"`
	Duration  time.Duration
}

func New(cfg *Config, handler core.HttpExchange) *Controller {
	var prime *Resource
	var second *Resource
	if handler == nil {
		prime = NewPrimaryResource(cfg.Host, cfg.Authority, cfg.Duration, nil)
	} else {
		prime = NewPrimaryResource("", cfg.Authority, cfg.Duration, handler)
		second = NewSecondaryResource(cfg.Host, cfg.Authority, cfg.Duration, nil)
	}
	return NewController(cfg.RouteName, prime, second)
}

func RegisterControllerFromConfig(config *Config, ex core.HttpExchange) *core.Status {
	ctrl := New(config, ex)
	err := RegisterController(ctrl)
	if err != nil {
		return core.NewStatusError(core.StatusInvalidArgument, err)
	}
	return core.StatusOK()
}
