package controller

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

type Config struct {
	CtrlName     string
	Host         string `json:"host"`
	Authority    string `json:"authority"`
	LivenessPath string `json:"liveness"`
	Duration     time.Duration
	Handler      core.HttpExchange
}

func Create(cfg Config) *Controller {
	var prime *Resource
	var second *Resource
	if cfg.Handler == nil {
		prime = NewPrimaryResource(cfg.Host, cfg.Authority, cfg.Duration, cfg.LivenessPath, nil)
	} else {
		prime = NewPrimaryResource("", cfg.Authority, cfg.Duration, cfg.LivenessPath, cfg.Handler)
		second = NewSecondaryResource(cfg.Host, cfg.Authority, cfg.Duration, cfg.LivenessPath, nil)
	}
	return NewController(cfg.CtrlName, prime, second)
}
