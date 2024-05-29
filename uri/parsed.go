package uri

import "net/url"

type Parsed struct {
	Valid     bool
	Authority string
	Version   string
	Path      string
	Query     string
	Err       error
}

func (p Parsed) PathURL() *url.URL {
	rawURL := p.Path
	if p.Query != "" {
		rawURL = p.Path + "?" + p.Query
	}
	u, _ := url.Parse(rawURL)
	return u
}
