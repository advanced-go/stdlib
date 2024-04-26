package core

import "net/http"

const (
	HealthLivenessPath  = "/health/liveness"
	HealthReadinessPath = "/health/readiness"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)
