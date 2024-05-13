package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"strings"
)

const (
	VersionPrefix = "v"
)

// ValidateRequest - validate the request given an embedded URN path
func ValidateRequest(req *http.Request, authority string) (ver string, path string, status *core.Status) {
	if req == nil {
		return "", "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if len(authority) == 0 {
		return "", "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: authority is empty"))
	}
	reqAuthority, reqPath, ok := uri.UprootUrn(req.URL.Path)
	if !ok {
		return "", "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, path is not valid: \"%v\"", req.URL.Path)))
	}
	if reqAuthority != authority {
		return "", "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, authority does not match: \"%v\" \"%v\"", req.URL.Path, authority)))
	}
	if strings.HasPrefix(reqPath, VersionPrefix) {
		i := strings.Index(reqPath, "/")
		if i != -1 {
			return reqPath[:i], reqPath[i+1:], core.StatusOK()
		}
	}
	return "", reqPath, core.StatusOK()
}
