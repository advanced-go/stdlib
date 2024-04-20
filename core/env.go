package core

import (
	"os"
	"strings"
)

type runtimeEnv int

const (
	EnvKey            = "RUNTIME_ENV"
	debug  runtimeEnv = iota
	test
	stage
	production
)

var (
	rte = debug
)

func init() {
	setEnv()
}

func setEnv() {
	rte = debug
	s := strings.ToLower(os.Getenv(EnvKey))
	if strings.Contains(s, "prod") {
		SetProdEnvironment()
		return
	}
	if strings.Contains(s, "stage") {
		SetStageEnvironment()
		return
	}
	if strings.Contains(s, "test") {
		SetTestEnvironment()
		return
	}
}

// IsProdEnvironment - determine if production environment
func IsProdEnvironment() bool {
	return rte == production
}

// SetProdEnvironment - set production environment
func SetProdEnvironment() {
	rte = production
}

// IsTestEnvironment - determine if test environment
func IsTestEnvironment() bool {
	return rte == test
}

// SetTestEnvironment - set test environment
func SetTestEnvironment() {
	rte = test
}

// IsStageEnvironment - determine if staging environment
func IsStageEnvironment() bool {
	return rte == stage
}

// SetStageEnvironment - set staging environment
func SetStageEnvironment() {
	rte = stage
}

// IsDebugEnvironment - determine if debug environment
func IsDebugEnvironment() bool {
	return rte == debug
}

// EnvStr - string representation for the environment
func EnvStr() string {
	switch rte {
	case debug:
		return "debug"
	case test:
		return "test"
	case stage:
		return "stage"
	case production:
		return "prod"
	}
	return "unknown"
}
