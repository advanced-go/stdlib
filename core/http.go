package core

import "net/http"

const (
	HealthLivenessPath  = "/health/liveness"
	HealthReadinessPath = "/health/readiness"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpExchange func(r *http.Request) (*http.Response, *Status)
