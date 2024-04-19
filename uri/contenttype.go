package uri

import "strings"

const (
	statusToken = "status"
	jsonExt     = ".json"
)

// IsJsonURL - does the URI have a .json extension
func IsJsonURL(uri string) bool {
	return strings.HasSuffix(uri, jsonExt)
}

// IsStatusURL - determine if the file name of the URL contains the text 'status'
func IsStatusURL(url string) bool {
	if len(url) == 0 {
		return false
	}
	i := strings.LastIndex(url, statusToken)
	if i == -1 {
		return false
	}
	return strings.LastIndex(url, "/") < i
}
