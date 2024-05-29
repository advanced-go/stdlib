package uri

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	UrnScheme     = "urn"
	UrnSeparator  = ":"
	VersionPrefix = "v"
)

// Uproot - uproot an embedded uri in a URI or a URI path
func Uproot(in string) Parsed {
	if in == "" {
		return Parsed{Valid: false, Err: errors.New("error: invalid input, URI is empty")}
	}
	if strings.HasPrefix(in, UrnScheme) {
		return Parsed{Valid: true, Authority: in, Path: in}
	}
	u, err := url.Parse(in)
	if err != nil {
		return Parsed{Valid: false, Err: err}
	}
	var str []string
	if u.Path[0] == '/' {
		str = strings.Split(u.Path[1:], UrnSeparator)
	} else {
		str = strings.Split(u.Path, UrnSeparator)
	}
	switch len(str) {
	case 0:
		return Parsed{Valid: false, Err: errors.New(fmt.Sprintf("error: path has no URN separator [%v]", u.Path))}
	case 1:
		return Parsed{Valid: true, Authority: str[0], Query: u.RawQuery}
	case 2:
		p := Parsed{Valid: true, Authority: str[0], Path: str[1], Query: u.RawQuery}
		parseVersion(&p)
		return p
	default:
		return Parsed{Valid: false, Err: errors.New(fmt.Sprintf("error: path has multiple URN separators [%v]", u.Path))}
	}
}

func parseVersion(p *Parsed) {
	if p == nil {
		return
	}
	if strings.HasPrefix(p.Path, VersionPrefix) {
		i := strings.Index(p.Path, "/")
		if i != -1 {
			p.Version = p.Path[:i]
			p.Path = p.Path[i+1:]
		}
	}
}
