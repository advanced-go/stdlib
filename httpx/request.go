package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

// ValidateURL - validate ta URL
func ValidateURL(url *url.URL, authority string) (ver string, path string, status *core.Status) {
	if url == nil {
		return "", "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	if len(authority) == 0 {
		return "", "", core.NewStatusError(core.StatusInvalidArgument, errors.New("error: authority is empty"))
	}
	if url.Path == core.AuthorityRootPath {
		return "", core.AuthorityPath, core.StatusOK()
	}
	p := uri.Uproot(url.Path)
	if !p.Valid {
		return "", "", core.NewStatusError(http.StatusBadRequest, p.Err)
	}
	if p.Authority != authority {
		return "", "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, authority does not match: \"%v\" \"%v\"", url.Path, authority)))
	}
	if len(p.Path) == 0 {
		return "", "", core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, path only contains an authority: \"%v\"", url.Path)))
	}
	return p.Version, p.Path, core.StatusOK()
}
