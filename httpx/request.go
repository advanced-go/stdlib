package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

// ValidateRequest - validate the request given an embedded URN path
func ValidateRequest(req *http.Request, path string) (string, *core.Status) {
	if req == nil {
		return "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: Request is nil"))
	}
	reqNid, reqPath, ok := uri.UprootUrn(req.URL.Path)
	if !ok {
		return "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, path is not valid: \"%v\"", req.URL.Path)))
	}
	if reqNid != path {
		return "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, NID does not match: \"%v\" \"%v\"", req.URL.Path, path)))
	}
	return reqPath, core.StatusOK()
}
