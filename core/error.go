package core

import (
	"fmt"
	"github.com/advaced-go/stdlib/fmt2"
	"log"
	"strconv"
	"strings"
	"time"
)

//https://github.com/advanced-go/core/blob/main/runtime/env.go?line=27

const (
	TimestampName  = "timestamp"
	StatusName     = "status"
	CodeName       = "code"
	TraceName      = "trace"
	RequestIdName  = "request-id"
	ErrorsName     = "errors"
	githubHost     = "github"
	githubDotCom   = "github.com"
	githubTemplate = "https://%v/tree/main%v"
	fragmentId     = "#"
	urnSeparator   = ":"
)

const (
	markupNull   = "\"%v\":null"
	markupString = "\"%v\":\"%v\""
	markupValue  = "\"%v\":%v"
)

// Formatter - output formatting type
type Formatter func(ts time.Time, code int, status, requestId string, errs []error, trace []string) string

// SetFormatter - optional override of error formatting
func SetFormatter(fn Formatter) {
	if fn != nil {
		formatter = fn
	}
}

// Logger - log function
type Logger func(code int, status, requestId string, errs []error, trace []string)

// SetLogger - optional override of logging
func SetLogger(fn Logger) {
	if fn != nil {
		logger = fn
	}
}

var (
	formatter            = defaultFormatter
	logger               = defaultLogger
	defaultLogger Logger = func(code int, status, requestId string, errs []error, trace []string) {
		log.Default().Println(formatter(time.Now().UTC(), code, status, requestId, errs, trace))
	}
)

// ErrorHandler - error handler interface
type ErrorHandler interface {
	Handle(s *Status, requestId string) *Status
}

// Bypass - bypass error handler
type Bypass struct{}

// Handle - bypass error handler
func (h Bypass) Handle(s *Status, _ string) *Status {
	return s
}

// Output - standard output error handler
type Output struct{}

// Handle - output error handler
func (h Output) Handle(s *Status, requestId string) *Status {
	if s == nil {
		return StatusOK()
	}
	if s.OK() {
		return s
	}
	if s.Err != nil && !s.Handled {
		s.addParentLocation()
		fmt.Printf("%v", formatter(time.Now().UTC(), s.Code, HttpStatus(s.Code), requestId, []error{s.Err}, s.Trace()))
		s.Handled = true
	}
	return s
}

// Log - log error handler
type Log struct{}

// Handle - log error handler
func (h Log) Handle(s *Status, requestId string) *Status {
	if s == nil {
		return StatusOK()
	}
	if s.OK() {
		return s
	}
	if s.Err != nil && !s.Handled {
		s.addParentLocation()
		go logger(s.Code, HttpStatus(s.Code), requestId, []error{s.Err}, s.Trace())
		s.Handled = true
	}
	return s
}

/*
func handle(s *Status, requestId string, output func()) *Status {
	if s == nil {
		return StatusOK()
	}
	if s.OK() {
		return s
	}
	if s.Err != nil && !s.Handled {
		s.addParentLocation()
		output
		//go logger(s.Code, HttpStatus(s.Code), requestId, []error{s.Err}, s.Trace())
		s.Handled = true
	}
	return s
}

*/

func defaultFormatter(ts time.Time, code int, status, requestId string, errs []error, trace []string) string {
	str := strconv.Itoa(code)
	return fmt.Sprintf("{ %v, %v, %v, %v, %v, %v }\n",
		jsonMarkup(TimestampName, fmt2.FmtRFC3339Millis(ts), true),
		jsonMarkup(CodeName, str, false),
		jsonMarkup(StatusName, status, true),
		jsonMarkup(RequestIdName, requestId, true),
		formatErrors(ErrorsName, errs),
		formatTrace(TraceName, trace))
}

func formatTrace(name string, trace []string) string {
	if len(trace) == 0 {
		return fmt.Sprintf("\"%v\" : null", name)
	}
	result := fmt.Sprintf("\"%v\" : [ ", name)
	for i := len(trace) - 1; i >= 0; i-- {
		if i < len(trace)-1 {
			result += ","
		}
		result += fmt.Sprintf("\"%v\"", formatUri(trace[i]))
	}
	return result + " ]"
}

func formatErrors(name string, errs []error) string {
	if len(errs) == 0 || errs[0] == nil {
		return fmt.Sprintf("\"%v\" : null", name)
	}
	result := fmt.Sprintf("\"%v\" : [ ", name)
	for i, e := range errs {
		if i != 0 {
			result += ","
		}
		result += fmt.Sprintf("\"%v\"", e.Error())
	}
	return result + " ]"
}

func formatUri(uri string) string {
	i := strings.Index(uri, githubHost)
	if i == -1 {
		return uri
	}
	uri = strings.Replace(uri, githubHost, githubDotCom, len(githubDotCom))
	i = strings.LastIndex(uri, "/")
	if i != -1 {
		first := uri[:i]
		last := uri[i:]
		last = strings.Replace(last, urnSeparator, fragmentId, len(fragmentId))
		return fmt.Sprintf(githubTemplate, first, last)
	}
	return uri
}

// jsonMarkup - markup a name/value pair
func jsonMarkup(name, value string, stringValue bool) string {
	if len(value) == 0 {
		return fmt.Sprintf(markupNull, name)
	}
	format := markupString
	if !stringValue {
		format = markupValue
	}
	return fmt.Sprintf(format, name, value)
}
