package core

import (
	"net/http"
)

const (
	MethodInfo = "INFO"
)

var (
	req *http.Request
)

func init() {
	req, _ = http.NewRequest(MethodInfo, "/info", nil)
}

func Authority(h HttpExchange) (authority, version string) {
	if h == nil {
		return "", ""
	}
	resp, status := h(req)
	if status.OK() {
		return resp.Header.Get(XAuthority), resp.Header.Get(XVersion)
	}
	return "", ""
}
