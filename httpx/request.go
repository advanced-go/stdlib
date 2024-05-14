package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	VersionPrefix = "v"
)

// ValidateRequest - validate the request URL path
func ValidateRequest(req *http.Request, authority string) (ver string, path string, status *core.Status) {
	if req == nil {
		return "", "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if len(authority) == 0 {
		return "", "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: authority is empty"))
	}
	p := uri.Uproot(req.URL.Path)
	if !p.Valid {
		return "", "", core.NewStatusError(http.StatusBadRequest, p.Err)
	}
	if p.Authority != authority {
		return "", "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, authority does not match: \"%v\" \"%v\"", req.URL.Path, authority)))
	}
	if len(p.Path) == 0 {
		return "", "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, path is empty: \"%v\"", req.URL.Path)))
	}
	return p.Version, p.Path, core.StatusOK()
}
