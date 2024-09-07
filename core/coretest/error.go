package coretest

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	fmt2 "github.com/advanced-go/stdlib/fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	githubHost     = "github"
	githubDotCom   = "github.com"
	githubTemplate = "https://%v/tree/main%v"
	fragmentId     = "#"
	urnSeparator   = ":"
	targetName     = "target"
)

// ErrorHandler - error handler interface
type ErrorHandler interface {
	Handle(s *core.Status, t *testing.T, target string) *core.Status
}

// Output - standard output error handler
type Output struct{}

// Handle - output error handler
func (h Output) Handle(s *core.Status, t *testing.T, target string) *core.Status {
	if s == nil {
		return core.StatusOK()
	}
	if s.OK() {
		return s
	}
	if s.Err != nil && !s.Handled {
		s.AddParentLocation()
		//fmt.Printf("%v", defaultFormatter(time.Now().UTC(), s.Code, core.HttpStatus(s.Code), s.RequestId, []error{s.Err}, s.Trace()))
		t.Errorf("%v", defaultFormatter(time.Now().UTC(), target, s.Code, core.HttpStatus(s.Code), s.RequestId, []error{s.Err}, s.Trace()))
		s.Handled = true
	}
	return s
}

func defaultFormatter(ts time.Time, target string, code int, status, requestId string, errs []error, trace []string) string {
	str := strconv.Itoa(code)
	return fmt.Sprintf("{ %v, %v %v, %v, %v, \n%v, \n%v }\n",
		fmt2.JsonMarkup(core.TimestampName, fmt2.FmtRFC3339Millis(ts), true),
		fmt2.JsonMarkup(targetName, target, true),
		fmt2.JsonMarkup(core.CodeName, str, false),
		fmt2.JsonMarkup(core.StatusName, status, true),
		fmt2.JsonMarkup(core.RequestIdName, requestId, true),
		formatErrors(core.ErrorsName, errs),
		formatTrace(core.TraceName, trace))
}

func formatErrors(name string, errs []error) string {
	if len(errs) == 0 || errs[0] == nil {
		return fmt.Sprintf("\"%v\" : null", name)
	}
	result := fmt.Sprintf("\"%v\" : [ \n", name)
	for i, e := range errs {
		if i != 0 {
			result += ",\n"
		}
		result += fmt.Sprintf("   \"%v\"", e.Error())
	}
	return result + "\n ]"
}

func formatTrace(name string, trace []string) string {
	if len(trace) == 0 {
		return fmt.Sprintf("\"%v\" : null", name)
	}
	result := fmt.Sprintf("\"%v\" : [ \n", name)
	for i := len(trace) - 1; i >= 0; i-- {
		if i < len(trace)-1 {
			result += ",\n"
		}
		result += fmt.Sprintf("   \"%v\"", formatUri(trace[i]))
	}
	return result + "\n ]"
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
