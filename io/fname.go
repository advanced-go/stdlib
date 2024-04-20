package io

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strings"
)

const (
	CwdVariable = "[cwd]"
	statusToken = "status"
	jsonExt     = ".json"
	fileScheme  = "file"
)

var (
	basePath = ""
	win      = false
)

// init - set the base path and windows flag
func init() {
	cwd, err := os.Getwd()
	if err != nil {
		basePath = err.Error()
	}
	if os.IsPathSeparator(uint8(92)) {
		win = true
	}
	basePath = cwd
}

// FileName - return the OS correct file name from a URI
func FileName(uri any) string {
	if uri == nil {
		return "error: URL is nil"
	}
	if s, ok := uri.(string); ok {
		if len(s) == 0 {
			return "error: URL is empty"
		}
		u, _ := url.Parse(s)
		return fileName(u)
	}
	if u, ok := uri.(*url.URL); ok {
		return fileName(u)
	}
	return fmt.Sprintf("error: invalid URL type: %v", reflect.TypeOf(uri))
}

func fileName(u *url.URL) string {
	if u == nil || u.Scheme != fileScheme {
		return fmt.Sprintf("error: scheme is invalid [%v]", u.Scheme)
	}
	name := basePath
	if u.Host == CwdVariable {
		name += u.Path
	} else {
		name = u.Path[1:]
	}
	if win {
		name = strings.ReplaceAll(name, "/", "\\")
	}
	return name
}
