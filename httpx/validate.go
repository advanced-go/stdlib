package httpx

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

// ValidateURL - validate a URL
func ValidateURL(url *url.URL, authority string) (p *uri.Parsed, status *core.Status) {
	if url == nil {
		return &uri.Parsed{}, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	if len(authority) == 0 {
		return &uri.Parsed{}, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: authority is empty"))
	}
	if url.Path == core.AuthorityRootPath {
		return &uri.Parsed{Path: core.AuthorityPath}, core.StatusOK()
	}
	if url.RawQuery != "" {
		p = uri.Uproot(url.Path + "?" + url.RawQuery)
	} else {
		p = uri.Uproot(url.Path)
	}
	if !p.Valid {
		return p, core.NewStatusError(http.StatusBadRequest, p.Err)
	}
	if p.Authority != authority {
		return p, core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, authority does not match: \"%v\" \"%v\"", url.Path, authority)))
	}
	if len(p.Path) == 0 {
		return p, core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid URI, path only contains an authority: \"%v\"", url.Path)))
	}
	return p, core.StatusOK()
}
