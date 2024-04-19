package uri

import (
	"errors"
	"fmt"
	"strings"
)

const (
	BeginDelimiter = "{"
	EndDelimiter   = "}"
)

// Expand - expand a template string, utilizing a template variable lookup function
func Expand(t string, fn func(string) (string, error)) (string, error) {
	if fn == nil || len(t) == 0 {
		return t, nil
	}
	var buf strings.Builder
	tokens := strings.Split(t, BeginDelimiter)
	if len(tokens) == 1 {
		return t, nil
	}
	for _, s := range tokens {
		sub := strings.Split(s, EndDelimiter)
		if len(sub) > 2 {
			return "", errors.New(fmt.Sprintf("invalid argument : token has multiple end delimiters: %v", s))
		}
		// Check case where no end delimiter is found
		if len(sub) == 1 && sub[0] == s {
			buf.WriteString(s)
			continue
		}
		// Have a valid end delimiter, so lookup the variable
		t1, err := fn(sub[0])
		if err != nil {
			return "", err
		}
		buf.WriteString(t1)
		if len(sub) == 2 {
			buf.WriteString(sub[1])
		}
	}
	return buf.String(), nil
}

// TemplateToken - a template variable
func TemplateToken(s string) (string, bool) {
	if !strings.HasPrefix(s, BeginDelimiter) {
		return s, false
	}
	if !strings.HasSuffix(s, EndDelimiter) {
		return s, false
	}
	return s[1 : len(s)-1], true
}
