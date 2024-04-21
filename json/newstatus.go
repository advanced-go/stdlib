package json

import "strings"

const (
	statusToken = "status"
)

// isStatusURL - determine if the file name of the URL contains the text 'status'
func isStatusURL(url string) bool {
	if len(url) == 0 {
		return false
	}
	i := strings.LastIndex(url, statusToken)
	if i == -1 {
		return false
	}
	return strings.LastIndex(url, "/") < i
}
