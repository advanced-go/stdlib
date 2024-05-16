package core

import (
	"fmt"
	"net/http"
)

const (
	HealthLivenessPath  = "health/liveness"
	HealthReadinessPath = "health/readiness"
	VersionPath         = "version"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpExchange func(r *http.Request) (*http.Response, *Status)

func VersionContent(s string) string {
	return fmt.Sprintf("{ \"version\": \"%v\" }", s)
}

func HealthContent(s string) string {
	return fmt.Sprintf("{ \"status\": \"%v\" }", s)
}
