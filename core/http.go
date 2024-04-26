package core

import "net/http"

const (
	HealthLivenessPath  = "/health/liveness"
	HealthReadinessPath = "/health/rediness"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)
